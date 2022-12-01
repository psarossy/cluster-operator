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

// SchemaReplicationInformer provides access to a shared informer and lister for
// SchemaReplications.
type SchemaReplicationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.SchemaReplicationLister
}

type schemaReplicationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSchemaReplicationInformer constructs a new informer for SchemaReplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSchemaReplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSchemaReplicationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSchemaReplicationInformer constructs a new informer for SchemaReplication type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSchemaReplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RabbitmqV1beta1().SchemaReplications(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RabbitmqV1beta1().SchemaReplications(namespace).Watch(context.TODO(), options)
			},
		},
		&rabbitmqcomv1beta1.SchemaReplication{},
		resyncPeriod,
		indexers,
	)
}

func (f *schemaReplicationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSchemaReplicationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *schemaReplicationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rabbitmqcomv1beta1.SchemaReplication{}, f.defaultInformer)
}

func (f *schemaReplicationInformer) Lister() v1beta1.SchemaReplicationLister {
	return v1beta1.NewSchemaReplicationLister(f.Informer().GetIndexer())
}
