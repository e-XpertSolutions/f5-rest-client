// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSDeviceConfigConfigList holds a list of DOSDeviceConfig configuration.
type DOSDeviceConfigConfigList struct {
	Items    []DOSDeviceConfigConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// DOSDeviceConfigConfig holds the configuration of a single DOSDeviceConfig.
type DOSDeviceConfigConfig struct {
}

// DOSDeviceConfigEndpoint represents the REST resource for managing DOSDeviceConfig.
const DOSDeviceConfigEndpoint = "/dos/device-config"

// DOSDeviceConfigResource provides an API to manage DOSDeviceConfig configurations.
type DOSDeviceConfigResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSDeviceConfig configurations.
func (r *DOSDeviceConfigResource) ListAll() (*DOSDeviceConfigConfigList, error) {
	var list DOSDeviceConfigConfigList
	if err := r.c.ReadQuery(BasePath+DOSDeviceConfigEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSDeviceConfig configuration identified by id.
func (r *DOSDeviceConfigResource) Get(id string) (*DOSDeviceConfigConfig, error) {
	var item DOSDeviceConfigConfig
	if err := r.c.ReadQuery(BasePath+DOSDeviceConfigEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSDeviceConfig configuration.
func (r *DOSDeviceConfigResource) Create(item DOSDeviceConfigConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSDeviceConfigEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSDeviceConfig configuration identified by id.
func (r *DOSDeviceConfigResource) Edit(id string, item DOSDeviceConfigConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSDeviceConfigEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSDeviceConfig configuration identified by id.
func (r *DOSDeviceConfigResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSDeviceConfigEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
