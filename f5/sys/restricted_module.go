// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// RestrictedModuleConfigList holds a list of RestrictedModule configuration.
type RestrictedModuleConfigList struct {
	Items    []RestrictedModuleConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// RestrictedModuleConfig holds the configuration of a single RestrictedModule.
type RestrictedModuleConfig struct {
}

// RestrictedModuleEndpoint represents the REST resource for managing RestrictedModule.
const RestrictedModuleEndpoint = "/restricted-module"

// RestrictedModuleResource provides an API to manage RestrictedModule configurations.
type RestrictedModuleResource struct {
	c *f5.Client
}

// ListAll  lists all the RestrictedModule configurations.
func (r *RestrictedModuleResource) ListAll() (*RestrictedModuleConfigList, error) {
	var list RestrictedModuleConfigList
	if err := r.c.ReadQuery(BasePath+RestrictedModuleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single RestrictedModule configuration identified by id.
func (r *RestrictedModuleResource) Get(id string) (*RestrictedModuleConfig, error) {
	var item RestrictedModuleConfig
	if err := r.c.ReadQuery(BasePath+RestrictedModuleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new RestrictedModule configuration.
func (r *RestrictedModuleResource) Create(item RestrictedModuleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+RestrictedModuleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a RestrictedModule configuration identified by id.
func (r *RestrictedModuleResource) Edit(id string, item RestrictedModuleConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+RestrictedModuleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single RestrictedModule configuration identified by id.
func (r *RestrictedModuleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+RestrictedModuleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
