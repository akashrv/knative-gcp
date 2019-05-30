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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/GoogleCloudPlatform/cloud-run-events/pkg/apis/events/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/cloud-run-events/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PubSubSourcesGetter has a method to return a PubSubSourceInterface.
// A group's client should implement this interface.
type PubSubSourcesGetter interface {
	PubSubSources(namespace string) PubSubSourceInterface
}

// PubSubSourceInterface has methods to work with PubSubSource resources.
type PubSubSourceInterface interface {
	Create(*v1alpha1.PubSubSource) (*v1alpha1.PubSubSource, error)
	Update(*v1alpha1.PubSubSource) (*v1alpha1.PubSubSource, error)
	UpdateStatus(*v1alpha1.PubSubSource) (*v1alpha1.PubSubSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PubSubSource, error)
	List(opts v1.ListOptions) (*v1alpha1.PubSubSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubSubSource, err error)
	PubSubSourceExpansion
}

// pubSubSources implements PubSubSourceInterface
type pubSubSources struct {
	client rest.Interface
	ns     string
}

// newPubSubSources returns a PubSubSources
func newPubSubSources(c *EventsV1alpha1Client, namespace string) *pubSubSources {
	return &pubSubSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pubSubSource, and returns the corresponding pubSubSource object, and an error if there is any.
func (c *pubSubSources) Get(name string, options v1.GetOptions) (result *v1alpha1.PubSubSource, err error) {
	result = &v1alpha1.PubSubSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pubsubsources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PubSubSources that match those selectors.
func (c *pubSubSources) List(opts v1.ListOptions) (result *v1alpha1.PubSubSourceList, err error) {
	result = &v1alpha1.PubSubSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pubsubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pubSubSources.
func (c *pubSubSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pubsubsources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a pubSubSource and creates it.  Returns the server's representation of the pubSubSource, and an error, if there is any.
func (c *pubSubSources) Create(pubSubSource *v1alpha1.PubSubSource) (result *v1alpha1.PubSubSource, err error) {
	result = &v1alpha1.PubSubSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pubsubsources").
		Body(pubSubSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pubSubSource and updates it. Returns the server's representation of the pubSubSource, and an error, if there is any.
func (c *pubSubSources) Update(pubSubSource *v1alpha1.PubSubSource) (result *v1alpha1.PubSubSource, err error) {
	result = &v1alpha1.PubSubSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pubsubsources").
		Name(pubSubSource.Name).
		Body(pubSubSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pubSubSources) UpdateStatus(pubSubSource *v1alpha1.PubSubSource) (result *v1alpha1.PubSubSource, err error) {
	result = &v1alpha1.PubSubSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pubsubsources").
		Name(pubSubSource.Name).
		SubResource("status").
		Body(pubSubSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the pubSubSource and deletes it. Returns an error if one occurs.
func (c *pubSubSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pubsubsources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pubSubSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pubsubsources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pubSubSource.
func (c *pubSubSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PubSubSource, err error) {
	result = &v1alpha1.PubSubSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pubsubsources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
