/*
RabbitMQ Cluster Operator

Copyright 2020 VMware, Inc. All Rights Reserved.

This product is licensed to you under the Mozilla Public license, Version 2.0 (the "License").  You may not use this product except in compliance with the Mozilla Public License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	rabbitmqcomv1beta1 "github.com/rabbitmq/cluster-operator/api/v1beta1"
	versioned "github.com/rabbitmq/cluster-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/rabbitmq/cluster-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/rabbitmq/cluster-operator/pkg/generated/listers/rabbitmq.com/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FederationInformer provides access to a shared informer and lister for
// Federations.
type FederationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FederationLister
}

type federationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederationInformer constructs a new informer for Federation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederationInformer constructs a new informer for Federation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RabbitmqV1beta1().Federations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RabbitmqV1beta1().Federations(namespace).Watch(context.TODO(), options)
			},
		},
		&rabbitmqcomv1beta1.Federation{},
		resyncPeriod,
		indexers,
	)
}

func (f *federationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rabbitmqcomv1beta1.Federation{}, f.defaultInformer)
}

func (f *federationInformer) Lister() v1beta1.FederationLister {
	return v1beta1.NewFederationLister(f.Informer().GetIndexer())
}