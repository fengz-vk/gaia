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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	platformv1alpha1 "github.com/lmxia/gaia/pkg/apis/platform/v1alpha1"
	versioned "github.com/lmxia/gaia/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/lmxia/gaia/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/lmxia/gaia/pkg/generated/listers/platform/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResourceBindingInformer provides access to a shared informer and lister for
// ResourceBindings.
type ResourceBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ResourceBindingLister
}

type resourceBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewResourceBindingInformer constructs a new informer for ResourceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceBindingInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredResourceBindingInformer constructs a new informer for ResourceBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PlatformV1alpha1().ResourceBindings().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PlatformV1alpha1().ResourceBindings().Watch(context.TODO(), options)
			},
		},
		&platformv1alpha1.ResourceBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceBindingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&platformv1alpha1.ResourceBinding{}, f.defaultInformer)
}

func (f *resourceBindingInformer) Lister() v1alpha1.ResourceBindingLister {
	return v1alpha1.NewResourceBindingLister(f.Informer().GetIndexer())
}