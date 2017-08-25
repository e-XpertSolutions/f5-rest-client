// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PFManDeviceConfigList holds a list of PFManDevice configuration.
type PFManDeviceConfigList struct {
	Items    []PFManDeviceConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// PFManDeviceConfig holds the configuration of a single PFManDevice.
type PFManDeviceConfig struct {
}

// PFManDeviceEndpoint represents the REST resource for managing PFManDevice.
const PFManDeviceEndpoint = "/pfman/device"

// PFManDeviceResource provides an API to manage PFManDevice configurations.
type PFManDeviceResource struct {
	c *f5.Client
}

// ListAll  lists all the PFManDevice configurations.
func (r *PFManDeviceResource) ListAll() (*PFManDeviceConfigList, error) {
	var list PFManDeviceConfigList
	if err := r.c.ReadQuery(BasePath+PFManDeviceEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PFManDevice configuration identified by id.
func (r *PFManDeviceResource) Get(id string) (*PFManDeviceConfig, error) {
	var item PFManDeviceConfig
	if err := r.c.ReadQuery(BasePath+PFManDeviceEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PFManDevice configuration.
func (r *PFManDeviceResource) Create(item PFManDeviceConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PFManDeviceEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PFManDevice configuration identified by id.
func (r *PFManDeviceResource) Edit(id string, item PFManDeviceConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PFManDeviceEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PFManDevice configuration identified by id.
func (r *PFManDeviceResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PFManDeviceEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
