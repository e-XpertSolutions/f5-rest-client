// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallOnDemandCompilationConfigList holds a list of FirewallOnDemandCompilation configuration.
type FirewallOnDemandCompilationConfigList struct {
	Items    []FirewallOnDemandCompilationConfig `json:"items"`
	Kind     string                              `json:"kind"`
	SelfLink string                              `json:"selflink"`
}

// FirewallOnDemandCompilationConfig holds the configuration of a single FirewallOnDemandCompilation.
type FirewallOnDemandCompilationConfig struct {
}

// FirewallOnDemandCompilationEndpoint represents the REST resource for managing FirewallOnDemandCompilation.
const FirewallOnDemandCompilationEndpoint = "/firewall/on-demand-compilation"

// FirewallOnDemandCompilationResource provides an API to manage FirewallOnDemandCompilation configurations.
type FirewallOnDemandCompilationResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallOnDemandCompilation configurations.
func (r *FirewallOnDemandCompilationResource) ListAll() (*FirewallOnDemandCompilationConfigList, error) {
	var list FirewallOnDemandCompilationConfigList
	if err := r.c.ReadQuery(BasePath+FirewallOnDemandCompilationEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallOnDemandCompilation configuration identified by id.
func (r *FirewallOnDemandCompilationResource) Get(id string) (*FirewallOnDemandCompilationConfig, error) {
	var item FirewallOnDemandCompilationConfig
	if err := r.c.ReadQuery(BasePath+FirewallOnDemandCompilationEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallOnDemandCompilation configuration.
func (r *FirewallOnDemandCompilationResource) Create(item FirewallOnDemandCompilationConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallOnDemandCompilationEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallOnDemandCompilation configuration identified by id.
func (r *FirewallOnDemandCompilationResource) Edit(id string, item FirewallOnDemandCompilationConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallOnDemandCompilationEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallOnDemandCompilation configuration identified by id.
func (r *FirewallOnDemandCompilationResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallOnDemandCompilationEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
