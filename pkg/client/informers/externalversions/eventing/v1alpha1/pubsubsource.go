/*
Copyright 2019 Google LLC

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
	time "time"

	eventingv1alpha1 "github.com/GoogleCloudPlatform/cloud-run-events/pkg/apis/events/v1alpha1"
	versioned "github.com/GoogleCloudPlatform/cloud-run-events/pkg/client/clientset/versioned"
	internalinterfaces "github.com/GoogleCloudPlatform/cloud-run-events/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/GoogleCloudPlatform/cloud-run-events/pkg/client/listers/eventing/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PubSubSourceInformer provides access to a shared informer and lister for
// PubSubSources.
type PubSubSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PubSubSourceLister
}

type pubSubSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPubSubSourceInformer constructs a new informer for PubSubSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPubSubSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPubSubSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPubSubSourceInformer constructs a new informer for PubSubSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPubSubSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventsV1alpha1().PubSubSources(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventsV1alpha1().PubSubSources(namespace).Watch(options)
			},
		},
		&eventingv1alpha1.PubSubSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *pubSubSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPubSubSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pubSubSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventingv1alpha1.PubSubSource{}, f.defaultInformer)
}

func (f *pubSubSourceInformer) Lister() v1alpha1.PubSubSourceLister {
	return v1alpha1.NewPubSubSourceLister(f.Informer().GetIndexer())
}
