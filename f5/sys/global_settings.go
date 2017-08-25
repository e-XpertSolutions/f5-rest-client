// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GlobalSettingsConfigList holds a list of GlobalSettings configuration.
type GlobalSettingsConfigList struct {
	Items    []GlobalSettingsConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// GlobalSettingsConfig holds the configuration of a single GlobalSettings.
type GlobalSettingsConfig struct {
}

// GlobalSettingsEndpoint represents the REST resource for managing GlobalSettings.
const GlobalSettingsEndpoint = "/global-settings"

// GlobalSettingsResource provides an API to manage GlobalSettings configurations.
type GlobalSettingsResource struct {
	c *f5.Client
}

// ListAll  lists all the GlobalSettings configurations.
func (r *GlobalSettingsResource) ListAll() (*GlobalSettingsConfigList, error) {
	var list GlobalSettingsConfigList
	if err := r.c.ReadQuery(BasePath+GlobalSettingsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single GlobalSettings configuration identified by id.
func (r *GlobalSettingsResource) Get(id string) (*GlobalSettingsConfig, error) {
	var item GlobalSettingsConfig
	if err := r.c.ReadQuery(BasePath+GlobalSettingsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new GlobalSettings configuration.
func (r *GlobalSettingsResource) Create(item GlobalSettingsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+GlobalSettingsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a GlobalSettings configuration identified by id.
func (r *GlobalSettingsResource) Edit(id string, item GlobalSettingsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+GlobalSettingsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single GlobalSettings configuration identified by id.
func (r *GlobalSettingsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+GlobalSettingsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
