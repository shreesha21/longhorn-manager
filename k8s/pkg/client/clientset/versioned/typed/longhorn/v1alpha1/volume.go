/*
Copyright 2018 The Kubernetes Authors.

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
	v1alpha1 "github.com/rancher/longhorn-manager/k8s/pkg/apis/longhorn/v1alpha1"
	scheme "github.com/rancher/longhorn-manager/k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumesGetter has a method to return a VolumeInterface.
// A group's client should implement this interface.
type VolumesGetter interface {
	Volumes(namespace string) VolumeInterface
}

// VolumeInterface has methods to work with Volume resources.
type VolumeInterface interface {
	Create(*v1alpha1.Volume) (*v1alpha1.Volume, error)
	Update(*v1alpha1.Volume) (*v1alpha1.Volume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Volume, error)
	List(opts v1.ListOptions) (*v1alpha1.VolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Volume, err error)
	VolumeExpansion
}

// volumes implements VolumeInterface
type volumes struct {
	client rest.Interface
	ns     string
}

// newVolumes returns a Volumes
func newVolumes(c *LonghornV1alpha1Client, namespace string) *volumes {
	return &volumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volume, and returns the corresponding volume object, and an error if there is any.
func (c *volumes) Get(name string, options v1.GetOptions) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Volumes that match those selectors.
func (c *volumes) List(opts v1.ListOptions) (result *v1alpha1.VolumeList, err error) {
	result = &v1alpha1.VolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumes.
func (c *volumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a volume and creates it.  Returns the server's representation of the volume, and an error, if there is any.
func (c *volumes) Create(volume *v1alpha1.Volume) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumes").
		Body(volume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a volume and updates it. Returns the server's representation of the volume, and an error, if there is any.
func (c *volumes) Update(volume *v1alpha1.Volume) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumes").
		Name(volume.Name).
		Body(volume).
		Do().
		Into(result)
	return
}

// Delete takes name of the volume and deletes it. Returns an error if one occurs.
func (c *volumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched volume.
func (c *volumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
