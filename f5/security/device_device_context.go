// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DeviceDeviceContextConfigList holds a list of DeviceDeviceContext configuration.
type DeviceDeviceContextConfigList struct {
	Items    []DeviceDeviceContextConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// DeviceDeviceContextConfig holds the configuration of a single DeviceDeviceContext.
type DeviceDeviceContextConfig struct {
}

// DeviceDeviceContextEndpoint represents the REST resource for managing DeviceDeviceContext.
const DeviceDeviceContextEndpoint = "/device/device-context"

// DeviceDeviceContextResource provides an API to manage DeviceDeviceContext configurations.
type DeviceDeviceContextResource struct {
	c *f5.Client
}

// ListAll  lists all the DeviceDeviceContext configurations.
func (r *DeviceDeviceContextResource) ListAll() (*DeviceDeviceContextConfigList, error) {
	var list DeviceDeviceContextConfigList
	if err := r.c.ReadQuery(BasePath+DeviceDeviceContextEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DeviceDeviceContext configuration identified by id.
func (r *DeviceDeviceContextResource) Get(id string) (*DeviceDeviceContextConfig, error) {
	var item DeviceDeviceContextConfig
	if err := r.c.ReadQuery(BasePath+DeviceDeviceContextEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DeviceDeviceContext configuration.
func (r *DeviceDeviceContextResource) Create(item DeviceDeviceContextConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DeviceDeviceContextEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DeviceDeviceContext configuration identified by id.
func (r *DeviceDeviceContextResource) Edit(id string, item DeviceDeviceContextConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DeviceDeviceContextEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DeviceDeviceContext configuration identified by id.
func (r *DeviceDeviceContextResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DeviceDeviceContextEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
