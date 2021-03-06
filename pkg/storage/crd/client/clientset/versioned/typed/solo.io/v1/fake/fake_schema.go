/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	solo_io_v1 "github.com/solo-io/sqoop/pkg/storage/crd/solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSchemas implements SchemaInterface
type FakeSchemas struct {
	Fake *FakeSqoopV1
	ns   string
}

var schemasResource = schema.GroupVersionResource{Group: "sqoop.solo.io", Version: "v1", Resource: "schemas"}

var schemasKind = schema.GroupVersionKind{Group: "sqoop.solo.io", Version: "v1", Kind: "Schema"}

// Get takes name of the schema, and returns the corresponding schema object, and an error if there is any.
func (c *FakeSchemas) Get(name string, options v1.GetOptions) (result *solo_io_v1.Schema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(schemasResource, c.ns, name), &solo_io_v1.Schema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Schema), err
}

// List takes label and field selectors, and returns the list of Schemas that match those selectors.
func (c *FakeSchemas) List(opts v1.ListOptions) (result *solo_io_v1.SchemaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(schemasResource, schemasKind, c.ns, opts), &solo_io_v1.SchemaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &solo_io_v1.SchemaList{}
	for _, item := range obj.(*solo_io_v1.SchemaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested schemas.
func (c *FakeSchemas) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(schemasResource, c.ns, opts))

}

// Create takes the representation of a schema and creates it.  Returns the server's representation of the schema, and an error, if there is any.
func (c *FakeSchemas) Create(schema *solo_io_v1.Schema) (result *solo_io_v1.Schema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(schemasResource, c.ns, schema), &solo_io_v1.Schema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Schema), err
}

// Update takes the representation of a schema and updates it. Returns the server's representation of the schema, and an error, if there is any.
func (c *FakeSchemas) Update(schema *solo_io_v1.Schema) (result *solo_io_v1.Schema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(schemasResource, c.ns, schema), &solo_io_v1.Schema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Schema), err
}

// Delete takes name of the schema and deletes it. Returns an error if one occurs.
func (c *FakeSchemas) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(schemasResource, c.ns, name), &solo_io_v1.Schema{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSchemas) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(schemasResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &solo_io_v1.SchemaList{})
	return err
}

// Patch applies the patch and returns the patched schema.
func (c *FakeSchemas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *solo_io_v1.Schema, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(schemasResource, c.ns, name, data, subresources...), &solo_io_v1.Schema{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Schema), err
}
