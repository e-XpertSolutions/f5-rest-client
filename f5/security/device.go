// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DeviceConfigList holds a list of Device configuration.
type DeviceConfigList struct {
	Items    []DeviceConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// DeviceConfig holds the configuration of a single Device.
type DeviceConfig struct {
}

// DeviceEndpoint represents the REST resource for managing Device.
const DeviceEndpoint = "/device"

// DeviceResource provides an API to manage Device configurations.
type DeviceResource struct {
	c *f5.Client
}

// ListAll  lists all the Device configurations.
func (r *DeviceResource) ListAll() (*DeviceConfigList, error) {
	var list DeviceConfigList
	if err := r.c.ReadQuery(BasePath+DeviceEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Device configuration identified by id.
func (r *DeviceResource) Get(id string) (*DeviceConfig, error) {
	var item DeviceConfig
	if err := r.c.ReadQuery(BasePath+DeviceEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Device configuration.
func (r *DeviceResource) Create(item DeviceConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DeviceEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Device configuration identified by id.
func (r *DeviceResource) Edit(id string, item DeviceConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DeviceEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Device configuration identified by id.
func (r *DeviceResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DeviceEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
