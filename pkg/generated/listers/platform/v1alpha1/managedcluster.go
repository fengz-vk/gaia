/*
Copyright The Gaia Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "gaia.io/gaia/pkg/apis/platform/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagedClusterLister helps list ManagedClusters.
// All objects returned here must be treated as read-only.
type ManagedClusterLister interface {
	// List lists all ManagedClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedCluster, err error)
	// ManagedClusters returns an object that can list and get ManagedClusters.
	ManagedClusters(namespace string) ManagedClusterNamespaceLister
	ManagedClusterListerExpansion
}

// managedClusterLister implements the ManagedClusterLister interface.
type managedClusterLister struct {
	indexer cache.Indexer
}

// NewManagedClusterLister returns a new ManagedClusterLister.
func NewManagedClusterLister(indexer cache.Indexer) ManagedClusterLister {
	return &managedClusterLister{indexer: indexer}
}

// List lists all ManagedClusters in the indexer.
func (s *managedClusterLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedCluster))
	})
	return ret, err
}

// ManagedClusters returns an object that can list and get ManagedClusters.
func (s *managedClusterLister) ManagedClusters(namespace string) ManagedClusterNamespaceLister {
	return managedClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagedClusterNamespaceLister helps list and get ManagedClusters.
// All objects returned here must be treated as read-only.
type ManagedClusterNamespaceLister interface {
	// List lists all ManagedClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagedCluster, err error)
	// Get retrieves the ManagedCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagedCluster, error)
	ManagedClusterNamespaceListerExpansion
}

// managedClusterNamespaceLister implements the ManagedClusterNamespaceLister
// interface.
type managedClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagedClusters in the indexer for a given namespace.
func (s managedClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagedCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagedCluster))
	})
	return ret, err
}

// Get retrieves the ManagedCluster from the indexer for a given namespace and name.
func (s managedClusterNamespaceLister) Get(name string) (*v1alpha1.ManagedCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managedcluster"), name)
	}
	return obj.(*v1alpha1.ManagedCluster), nil
}
