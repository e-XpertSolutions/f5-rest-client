// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GlobalSettingsLoadBalancingConfigList holds a list of GlobalSettingsLoadBalancing configuration.
type GlobalSettingsLoadBalancingConfigList struct {
	Items    []GlobalSettingsLoadBalancingConfig `json:"items"`
	Kind     string                              `json:"kind"`
	SelfLink string                              `json:"selflink"`
}

// GlobalSettingsLoadBalancingConfig holds the configuration of a single GlobalSettingsLoadBalancing.
type GlobalSettingsLoadBalancingConfig struct {
}

// GlobalSettingsLoadBalancingEndpoint represents the REST resource for managing GlobalSettingsLoadBalancing.
const GlobalSettingsLoadBalancingEndpoint = "/global-settings/load-balancing"

// GlobalSettingsLoadBalancingResource provides an API to manage GlobalSettingsLoadBalancing configurations.
type GlobalSettingsLoadBalancingResource struct {
	c f5.Client
}

// ListAll  lists all the GlobalSettingsLoadBalancing configurations.
func (r *GlobalSettingsLoadBalancingResource) ListAll() (*GlobalSettingsLoadBalancingConfigList, error) {
	var list GlobalSettingsLoadBalancingConfigList
	if err := r.c.ReadQuery(BasePath+GlobalSettingsLoadBalancingEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single GlobalSettingsLoadBalancing configuration identified by id.
func (r *GlobalSettingsLoadBalancingResource) Get(id string) (*GlobalSettingsLoadBalancingConfig, error) {
	var item GlobalSettingsLoadBalancingConfig
	if err := r.c.ReadQuery(BasePath+GlobalSettingsLoadBalancingEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new GlobalSettingsLoadBalancing configuration.
func (r *GlobalSettingsLoadBalancingResource) Create(item GlobalSettingsLoadBalancingConfig) error {
	if err := r.c.ModQuery("POST", BasePath+GlobalSettingsLoadBalancingEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a GlobalSettingsLoadBalancing configuration identified by id.
func (r *GlobalSettingsLoadBalancingResource) Edit(id string, item GlobalSettingsLoadBalancingConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+GlobalSettingsLoadBalancingEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single GlobalSettingsLoadBalancing configuration identified by id.
func (r *GlobalSettingsLoadBalancingResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+GlobalSettingsLoadBalancingEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
