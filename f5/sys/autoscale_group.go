// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AutoscaleGroupConfigList holds a list of AutoscaleGroup configuration.
type AutoscaleGroupConfigList struct {
	Items    []AutoscaleGroupConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// AutoscaleGroupConfig holds the configuration of a single AutoscaleGroup.
type AutoscaleGroupConfig struct {
}

// AutoscaleGroupEndpoint represents the REST resource for managing AutoscaleGroup.
const AutoscaleGroupEndpoint = "/autoscale-group"

// AutoscaleGroupResource provides an API to manage AutoscaleGroup configurations.
type AutoscaleGroupResource struct {
	c *f5.Client
}

// ListAll  lists all the AutoscaleGroup configurations.
func (r *AutoscaleGroupResource) ListAll() (*AutoscaleGroupConfigList, error) {
	var list AutoscaleGroupConfigList
	if err := r.c.ReadQuery(BasePath+AutoscaleGroupEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AutoscaleGroup configuration identified by id.
func (r *AutoscaleGroupResource) Get(id string) (*AutoscaleGroupConfig, error) {
	var item AutoscaleGroupConfig
	if err := r.c.ReadQuery(BasePath+AutoscaleGroupEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AutoscaleGroup configuration.
func (r *AutoscaleGroupResource) Create(item AutoscaleGroupConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AutoscaleGroupEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AutoscaleGroup configuration identified by id.
func (r *AutoscaleGroupResource) Edit(id string, item AutoscaleGroupConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AutoscaleGroupEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AutoscaleGroup configuration identified by id.
func (r *AutoscaleGroupResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AutoscaleGroupEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
