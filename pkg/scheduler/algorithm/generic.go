// This file was copied from k8s.io/kubernetes/pkg/scheduler/generic_scheduler.go and modified

package algorithm

import (
	"context"
	"fmt"

	"sync"
	"sync/atomic"
	"time"

	"gonum.org/v1/gonum/mat"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	utiltrace "k8s.io/utils/trace"

	"github.com/lmxia/gaia/pkg/apis/apps/v1alpha1"
	clusterapi "github.com/lmxia/gaia/pkg/apis/platform/v1alpha1"

	"github.com/lmxia/gaia/pkg/common"
	applisterv1alpha1 "github.com/lmxia/gaia/pkg/generated/listers/apps/v1alpha1"
	schedulerapis "github.com/lmxia/gaia/pkg/scheduler/apis"
	schedulercache "github.com/lmxia/gaia/pkg/scheduler/cache"
	framework2 "github.com/lmxia/gaia/pkg/scheduler/framework"
	framework "github.com/lmxia/gaia/pkg/scheduler/framework/interfaces"
	"github.com/lmxia/gaia/pkg/scheduler/framework/runtime"
	"github.com/lmxia/gaia/pkg/scheduler/metrics"
	"github.com/lmxia/gaia/pkg/scheduler/parallelize"
)

// ErrNoClustersAvailable is used to describe the error that no clusters available to schedule subscriptions.
var ErrNoClustersAvailable = fmt.Errorf("no clusters available to schedule subscriptions")

type genericScheduler struct {
	cache                       schedulercache.Cache
	percentageOfClustersToScore int32
	nextStartClusterIndex       int
}

func (g *genericScheduler) SetRBLister(lister applisterv1alpha1.ResourceBindingLister) {
	g.cache.SetRBLister(lister)
}

func (g *genericScheduler) SetSelfClusterName(name string) {
	g.cache.SetSelfClusterName(name)
}

// Schedule
func (g *genericScheduler) Schedule(ctx context.Context, fwk framework.Framework, desc *v1alpha1.Description) (result ScheduleResult, err error) {
	trace := utiltrace.New("Scheduling", utiltrace.Field{Key: "namespace", Value: desc.Namespace}, utiltrace.Field{Key: "name", Value: desc.Name})
	defer trace.LogIfLong(100 * time.Millisecond)

	//1. get rbs, if has
	rbs := make([]*v1alpha1.ResourceBinding, 0)
	if desc.Namespace != common.GaiaReservedNamespace {
		rbs = g.cache.ListResourceBindings(desc, common.StatusScheduling)
	}
	//2. get backup clusters.
	if g.cache.NumClusters() == 0 {
		return result, ErrNoClustersAvailable
	}
	allClusters, _ := g.cache.ListClusters(&metav1.LabelSelector{})
	numComponent := len(desc.Spec.Component)
	// 3, means allspread, one spread, 2 spread.
	allResultGlobal := make([][]mat.Matrix, 3)
	for i := 0; i < 3; i++ {
		allResultGlobal[i] = make([]mat.Matrix, numComponent)
	}
	// has parent
	// first dim means resouce binding, second means component, third means spread level.
	allResultWithRB := make([][][]mat.Matrix, len(rbs))
	if desc.Namespace != common.GaiaReservedNamespace {
		//desc has resource binding
		for i := 0; i < len(rbs); i++ {
			allResultWithRB[i] = make([][]mat.Matrix, 3)
			for j := 0; j < 3; j++ {
				allResultWithRB[i][j] = make([]mat.Matrix, numComponent)
			}
		}
	}

	for i, comm := range desc.Spec.Component {
		localComponent := &comm
		// NO.1 pre filter
		feasibleClusters, diagnosis, err := g.findClustersThatFitComponent(ctx, fwk, localComponent)
		if err != nil {
			return result, err
		}

		if len(feasibleClusters) == 0 {
			return result, &framework.FitError{
				Description:    desc,
				NumAllClusters: g.cache.NumClusters(),
				Diagnosis:      diagnosis,
			}
		}
		// spread level info: full level, 2 level, 1 level
		spreadLevels := []int64{int64(len(feasibleClusters)), 2, 1}
		allPlan := nomalizeClusters(feasibleClusters, allClusters)
		// desc come from reserved namespace, that means no resource bindings
		if desc.Namespace == common.GaiaReservedNamespace {
			for j, spreadLevel := range spreadLevels {
				componentMat := makeComponentPlans(allPlan, int64(comm.Workload.TraitDeployment.Replicas), spreadLevel)
				allResultGlobal[j][i] = componentMat
			}
		} else {
			for j, rb := range rbs {
				replicas := getComponentClusterTotal(rb.Spec.RbApps, g.cache.GetSelfClusterName(), comm.Name)
				for k, spreadLevel := range spreadLevels {
					componentMat := makeComponentPlans(allPlan, replicas, spreadLevel)
					allResultWithRB[j][k][i] = componentMat
				}
			}
		}
	}

	/*
		add one-by-one network filter plugin.
	*/

	// NO.2 first we should spawn rbs.
	if desc.Namespace == common.GaiaReservedNamespace {
		rbs := spawnResourceBindings(allResultGlobal, allClusters, desc)
		return ScheduleResult{
			ResourceBindings: rbs,
		}, nil
	} else {
		rbIndex := 0
		rbsNew := make([]*v1alpha1.ResourceBinding, 0)
		for i, rbOld := range rbs {
			rbForrb := spawnResourceBindings(allResultWithRB[i], allClusters, desc)
			var resourceBindingApps *v1alpha1.ResourceBindingApps = nil
			for _, rbapp := range rbOld.Spec.RbApps {
				if rbapp.ClusterName == g.cache.GetSelfClusterName() {
					resourceBindingApps = rbapp
					break
				}
			}
			for _, item := range rbForrb {
				rbNew := &v1alpha1.ResourceBinding{
					ObjectMeta: metav1.ObjectMeta{
						Name: fmt.Sprintf("%s-rs-%d", desc.Name, rbIndex),
						Labels: map[string]string{
							common.GaiaDescriptionLabel: desc.Name,
						},
					},
					Spec: v1alpha1.ResourceBindingSpec{
						AppID:    desc.Name,
						ParentRB: rbOld.Name,
						RbApps:   rbOld.Spec.RbApps,
					},
				}
				rbNew.Kind = "ResourceBinding"
				rbNew.APIVersion = "platform.gaia.io/v1alpha1"
				rbIndex += 1
				resourceBindingApps.Children = item.Spec.RbApps
				rbsNew = append(rbsNew, rbNew)
			}
		}
		return ScheduleResult{
			ResourceBindings: rbsNew,
		}, err
	}
}

func nomalizeClusters(feasibleClusters []*framework2.ClusterInfo, allClusters []*clusterapi.ManagedCluster) []*framework2.ClusterInfo {
	indexCluster := make(map[string]*framework2.ClusterInfo)
	result := make([]*framework2.ClusterInfo, len(allClusters))
	for _, feasibleCluster := range feasibleClusters {
		indexCluster[feasibleCluster.Cluster.Name] = feasibleCluster
	}

	for i, cluster := range allClusters {
		if k, ok := indexCluster[cluster.Name]; ok {
			result[i] = k
		} else {
			clusterInfo := &framework2.ClusterInfo{
				Cluster: cluster,
				Total:   0,
			}
			result[i] = clusterInfo
		}
	}
	return result
}

// selectClusters takes a prioritized list of clusters and then picks a fraction of clusters
// in a reservoir sampling manner from the clusters that had the highest score.
func (g *genericScheduler) selectClusters(clusterScoreList framework.ClusterScoreList, _ *v1alpha1.Description) ([]string, error) {
	if len(clusterScoreList) == 0 {
		return nil, fmt.Errorf("empty clusterScoreList")
	}

	var selected []string
	for _, clusterScore := range clusterScoreList {
		// TODO: sampling with scores
		selected = append(selected, clusterScore.NamespacedName)
	}
	return selected, nil
}

// Filters the clusters to find the ones that fit the subscription based on the framework filter plugins.
func (g *genericScheduler) findClustersThatFitComponent(ctx context.Context, fwk framework.Framework, comm *v1alpha1.Components) ([]*framework2.ClusterInfo, framework.Diagnosis, error) {
	diagnosis := framework.Diagnosis{
		ClusterToStatusMap:   make(framework.ClusterToStatusMap),
		UnschedulablePlugins: sets.NewString(),
	}

	var allClusters []*clusterapi.ManagedCluster
	// 1. Geo labels
	mergedSelector := &metav1.LabelSelector{}

	if comm.SchedulePolicy.Geolocation != nil {
		mergedSelector.MatchLabels = comm.SchedulePolicy.Geolocation.MatchLabels
		mergedSelector.MatchExpressions = comm.SchedulePolicy.Geolocation.MatchExpressions
	}
	// 2. Net labels
	if comm.SchedulePolicy.Netenvironment != nil {
		if mergedSelector.MatchLabels != nil {
			for key := range comm.SchedulePolicy.Netenvironment.MatchLabels {
				if val, ok := mergedSelector.MatchLabels[key]; !ok {
					mergedSelector.MatchLabels[key] = val
				}
			}
		}  else {
			mergedSelector.MatchLabels = comm.SchedulePolicy.Netenvironment.MatchLabels
		}

		if mergedSelector.MatchExpressions != nil {
			mergedSelector.MatchExpressions = append(mergedSelector.MatchExpressions,
				comm.SchedulePolicy.Netenvironment.MatchExpressions...)
		} else {
			mergedSelector.MatchExpressions = comm.SchedulePolicy.Netenvironment.MatchExpressions
		}
	}

	// 3. specific sn
	if comm.SchedulePolicy.SpecificResource != nil {
		if mergedSelector.MatchLabels != nil {
			for key := range comm.SchedulePolicy.SpecificResource.MatchLabels {
				if val, ok := mergedSelector.MatchLabels[key]; !ok {
					mergedSelector.MatchLabels[key] = val
				}
			}
		}  else {
			mergedSelector.MatchLabels = comm.SchedulePolicy.SpecificResource.MatchLabels
		}
		if mergedSelector.MatchExpressions != nil {
			mergedSelector.MatchExpressions = append(mergedSelector.MatchExpressions,
				comm.SchedulePolicy.SpecificResource.MatchExpressions...)
		} else {
			mergedSelector.MatchExpressions = comm.SchedulePolicy.SpecificResource.MatchExpressions
		}
	}

	// 4. provider
	if comm.SchedulePolicy.Provider != nil {
		if mergedSelector.MatchLabels != nil {
			for key := range comm.SchedulePolicy.Provider.MatchLabels {
				if val, ok := mergedSelector.MatchLabels[key]; !ok {
					mergedSelector.MatchLabels[key] = val
				}
			}
		} else {
			mergedSelector.MatchLabels = comm.SchedulePolicy.Provider.MatchLabels
		}

		if mergedSelector.MatchExpressions != nil {
			mergedSelector.MatchExpressions = append(mergedSelector.MatchExpressions,
				comm.SchedulePolicy.Provider.MatchExpressions...)
		} else {
			mergedSelector.MatchExpressions = comm.SchedulePolicy.Provider.MatchExpressions
		}
	}

	clusters, err := g.cache.ListClusters(mergedSelector)

	if err != nil {
		return nil, diagnosis, err
	}
	allClusters = append(allClusters, clusters...)

	allClusters = normalizedClusters(allClusters)
	// Return immediately if no clusters match the cluster affinity.
	if len(allClusters) == 0 {
		return nil, diagnosis, nil
	}

	// Run "prefilter" plugins. ALL PASS FOR NOW. we don't know how to use it.
	s := fwk.RunPreFilterPlugins(ctx, comm)
	if !s.IsSuccess() {
		if !s.IsUnschedulable() {
			return nil, diagnosis, s.AsError()
		}
		// All clusters will have the same status. Some non trivial refactoring is
		// needed to avoid this copy.
		for _, n := range allClusters {
			diagnosis.ClusterToStatusMap[klog.KObj(n).String()] = s
		}
		// Status satisfying IsUnschedulable() gets injected into diagnosis.UnschedulablePlugins.
		diagnosis.UnschedulablePlugins.Insert(s.FailedPlugin())
		return nil, diagnosis, nil
	}

	feasibleClusters, err := g.findClustersThatPassFilters(ctx, fwk, comm, diagnosis, allClusters)
	if err != nil {
		return nil, diagnosis, err
	}

	// aggregate all container resource
	non0CPU, non0MEM, _ := calculateResource(comm.Module)
	result, _ := scheduleWorkload(non0CPU, non0MEM, feasibleClusters)

	return result, diagnosis, nil
}

// findClustersThatPassFilters finds the clusters that fit the filter plugins.
func (g *genericScheduler) findClustersThatPassFilters(ctx context.Context, fwk framework.Framework,
	com *v1alpha1.Components, diagnosis framework.Diagnosis,
	clusters []*clusterapi.ManagedCluster) ([]*clusterapi.ManagedCluster, error) {
	if !fwk.HasFilterPlugins() {
		return clusters, nil
	}

	errCh := parallelize.NewErrorChannel()
	var statusesLock sync.Mutex
	var feasibleClustersLen int32
	feasibleClusters := make([]*clusterapi.ManagedCluster, len(clusters))

	ctx, cancel := context.WithCancel(ctx)
	checkCluster := func(i int) {
		cluster := clusters[i]

		status := fwk.RunFilterPlugins(ctx, com, cluster).Merge()
		if status.Code() == framework.Error {
			errCh.SendErrorWithCancel(status.AsError(), cancel)
			return
		}
		if status.IsSuccess() {
			length := atomic.AddInt32(&feasibleClustersLen, 1)
			feasibleClusters[length-1] = cluster
		} else {
			statusesLock.Lock()
			diagnosis.ClusterToStatusMap[klog.KObj(cluster).String()] = status
			diagnosis.UnschedulablePlugins.Insert(status.FailedPlugin())
			statusesLock.Unlock()
		}
	}

	beginCheckCluster := time.Now()
	statusCode := framework.Success
	defer func() {
		// We record Filter extension point latency here instead of in framework.go because framework.RunFilterPlugins
		// function is called for each cluster, whereas we want to have an overall latency for all clusters per scheduling cycle.
		metrics.FrameworkExtensionPointDuration.WithLabelValues(runtime.Filter, statusCode.String(), fwk.ProfileName()).Observe(metrics.SinceInSeconds(beginCheckCluster))
	}()

	// Stops searching for more clusters once the configured number of feasible clusters
	// are found.
	fwk.Parallelizer().Until(ctx, len(clusters), checkCluster)

	if err := errCh.ReceiveError(); err != nil {
		statusCode = framework.Error
		return nil, err
	}
	for i, cluster := range feasibleClusters {
		if cluster == nil {
			return feasibleClusters[:i-1], nil
		}
	}
	return feasibleClusters, nil
}

// NewGenericScheduler creates a genericScheduler object.
func NewGenericScheduler(cache schedulercache.Cache) ScheduleAlgorithm {
	return &genericScheduler{
		cache:                       cache,
		percentageOfClustersToScore: schedulerapis.DefaultPercentageOfClustersToScore,
	}
}

// normalizedClusters will remove duplicate clusters. Deleting clusters will be removed as well.
func normalizedClusters(clusters []*clusterapi.ManagedCluster) []*clusterapi.ManagedCluster {
	allKeys := make(map[string]bool)
	var uniqueClusters []*clusterapi.ManagedCluster
	for _, cluster := range clusters {
		if _, ok := allKeys[klog.KObj(cluster).String()]; !ok {
			if cluster.DeletionTimestamp != nil {
				continue
			}
			uniqueClusters = append(uniqueClusters, cluster)
			allKeys[klog.KObj(cluster).String()] = true
		}
	}
	return uniqueClusters
}