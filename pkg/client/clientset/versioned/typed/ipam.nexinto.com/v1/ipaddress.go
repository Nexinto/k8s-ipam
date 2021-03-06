/*
Copyright 2018 Nexinto

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

package v1

import (
	v1 "github.com/Nexinto/k8s-ipam/pkg/apis/ipam.nexinto.com/v1"
	scheme "github.com/Nexinto/k8s-ipam/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IpAddressesGetter has a method to return a IpAddressInterface.
// A group's client should implement this interface.
type IpAddressesGetter interface {
	IpAddresses(namespace string) IpAddressInterface
}

// IpAddressInterface has methods to work with IpAddress resources.
type IpAddressInterface interface {
	Create(*v1.IpAddress) (*v1.IpAddress, error)
	Update(*v1.IpAddress) (*v1.IpAddress, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.IpAddress, error)
	List(opts meta_v1.ListOptions) (*v1.IpAddressList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.IpAddress, err error)
	IpAddressExpansion
}

// ipAddresses implements IpAddressInterface
type ipAddresses struct {
	client rest.Interface
	ns     string
}

// newIpAddresses returns a IpAddresses
func newIpAddresses(c *IpamV1Client, namespace string) *ipAddresses {
	return &ipAddresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the ipAddress, and returns the corresponding ipAddress object, and an error if there is any.
func (c *ipAddresses) Get(name string, options meta_v1.GetOptions) (result *v1.IpAddress, err error) {
	result = &v1.IpAddress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipaddresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IpAddresses that match those selectors.
func (c *ipAddresses) List(opts meta_v1.ListOptions) (result *v1.IpAddressList, err error) {
	result = &v1.IpAddressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipaddresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ipAddresses.
func (c *ipAddresses) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ipaddresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a ipAddress and creates it.  Returns the server's representation of the ipAddress, and an error, if there is any.
func (c *ipAddresses) Create(ipAddress *v1.IpAddress) (result *v1.IpAddress, err error) {
	result = &v1.IpAddress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ipaddresses").
		Body(ipAddress).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ipAddress and updates it. Returns the server's representation of the ipAddress, and an error, if there is any.
func (c *ipAddresses) Update(ipAddress *v1.IpAddress) (result *v1.IpAddress, err error) {
	result = &v1.IpAddress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ipaddresses").
		Name(ipAddress.Name).
		Body(ipAddress).
		Do().
		Into(result)
	return
}

// Delete takes name of the ipAddress and deletes it. Returns an error if one occurs.
func (c *ipAddresses) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipaddresses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ipAddresses) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipaddresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ipAddress.
func (c *ipAddresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.IpAddress, err error) {
	result = &v1.IpAddress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ipaddresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
