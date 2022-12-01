/*
RabbitMQ Cluster Operator

Copyright 2020 VMware, Inc. All Rights Reserved.

This product is licensed to you under the Mozilla Public license, Version 2.0 (the "License").  You may not use this product except in compliance with the Mozilla Public License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"net/http"

	v1beta1 "github.com/rabbitmq/cluster-operator/api/v1beta1"
	"github.com/rabbitmq/cluster-operator/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type RabbitmqV1beta1Interface interface {
	RESTClient() rest.Interface
	BindingsGetter
	ExchangesGetter
	FederationsGetter
	PermissionsGetter
	PoliciesGetter
	QueuesGetter
	SchemaReplicationsGetter
	ShovelsGetter
	UsersGetter
	VhostsGetter
}

// RabbitmqV1beta1Client is used to interact with features provided by the rabbitmq.com group.
type RabbitmqV1beta1Client struct {
	restClient rest.Interface
}

func (c *RabbitmqV1beta1Client) Bindings(namespace string) BindingInterface {
	return newBindings(c, namespace)
}

func (c *RabbitmqV1beta1Client) Exchanges(namespace string) ExchangeInterface {
	return newExchanges(c, namespace)
}

func (c *RabbitmqV1beta1Client) Federations(namespace string) FederationInterface {
	return newFederations(c, namespace)
}

func (c *RabbitmqV1beta1Client) Permissions(namespace string) PermissionInterface {
	return newPermissions(c, namespace)
}

func (c *RabbitmqV1beta1Client) Policies(namespace string) PolicyInterface {
	return newPolicies(c, namespace)
}

func (c *RabbitmqV1beta1Client) Queues(namespace string) QueueInterface {
	return newQueues(c, namespace)
}

func (c *RabbitmqV1beta1Client) SchemaReplications(namespace string) SchemaReplicationInterface {
	return newSchemaReplications(c, namespace)
}

func (c *RabbitmqV1beta1Client) Shovels(namespace string) ShovelInterface {
	return newShovels(c, namespace)
}

func (c *RabbitmqV1beta1Client) Users(namespace string) UserInterface {
	return newUsers(c, namespace)
}

func (c *RabbitmqV1beta1Client) Vhosts(namespace string) VhostInterface {
	return newVhosts(c, namespace)
}

// NewForConfig creates a new RabbitmqV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*RabbitmqV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new RabbitmqV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*RabbitmqV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &RabbitmqV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new RabbitmqV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *RabbitmqV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new RabbitmqV1beta1Client for the given RESTClient.
func New(c rest.Interface) *RabbitmqV1beta1Client {
	return &RabbitmqV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *RabbitmqV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}