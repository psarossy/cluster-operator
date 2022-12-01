/*
RabbitMQ Cluster Operator

Copyright 2020 VMware, Inc. All Rights Reserved.

This product is licensed to you under the Mozilla Public license, Version 2.0 (the "License").  You may not use this product except in compliance with the Mozilla Public License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/rabbitmq/cluster-operator/api/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePermissions implements PermissionInterface
type FakePermissions struct {
	Fake *FakeRabbitmqV1beta1
	ns   string
}

var permissionsResource = schema.GroupVersionResource{Group: "rabbitmq.com", Version: "v1beta1", Resource: "permissions"}

var permissionsKind = schema.GroupVersionKind{Group: "rabbitmq.com", Version: "v1beta1", Kind: "Permission"}

// Get takes name of the permission, and returns the corresponding permission object, and an error if there is any.
func (c *FakePermissions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Permission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(permissionsResource, c.ns, name), &v1beta1.Permission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Permission), err
}

// List takes label and field selectors, and returns the list of Permissions that match those selectors.
func (c *FakePermissions) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.PermissionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(permissionsResource, permissionsKind, c.ns, opts), &v1beta1.PermissionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.PermissionList{ListMeta: obj.(*v1beta1.PermissionList).ListMeta}
	for _, item := range obj.(*v1beta1.PermissionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested permissions.
func (c *FakePermissions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(permissionsResource, c.ns, opts))

}

// Create takes the representation of a permission and creates it.  Returns the server's representation of the permission, and an error, if there is any.
func (c *FakePermissions) Create(ctx context.Context, permission *v1beta1.Permission, opts v1.CreateOptions) (result *v1beta1.Permission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(permissionsResource, c.ns, permission), &v1beta1.Permission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Permission), err
}

// Update takes the representation of a permission and updates it. Returns the server's representation of the permission, and an error, if there is any.
func (c *FakePermissions) Update(ctx context.Context, permission *v1beta1.Permission, opts v1.UpdateOptions) (result *v1beta1.Permission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(permissionsResource, c.ns, permission), &v1beta1.Permission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Permission), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePermissions) UpdateStatus(ctx context.Context, permission *v1beta1.Permission, opts v1.UpdateOptions) (*v1beta1.Permission, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(permissionsResource, "status", c.ns, permission), &v1beta1.Permission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Permission), err
}

// Delete takes name of the permission and deletes it. Returns an error if one occurs.
func (c *FakePermissions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(permissionsResource, c.ns, name, opts), &v1beta1.Permission{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePermissions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(permissionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.PermissionList{})
	return err
}

// Patch applies the patch and returns the patched permission.
func (c *FakePermissions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Permission, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(permissionsResource, c.ns, name, pt, data, subresources...), &v1beta1.Permission{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Permission), err
}