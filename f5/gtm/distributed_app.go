// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DistributedAppConfigList holds a list of DistributedApp configuration.
type DistributedAppConfigList struct {
	Items    []DistributedAppConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// DistributedAppConfig holds the configuration of a single DistributedApp.
type DistributedAppConfig struct {
}

// DistributedAppEndpoint represents the REST resource for managing DistributedApp.
const DistributedAppEndpoint = "/distributed-app"

// DistributedAppResource provides an API to manage DistributedApp configurations.
type DistributedAppResource struct {
	c f5.Client
}

// ListAll  lists all the DistributedApp configurations.
func (r *DistributedAppResource) ListAll() (*DistributedAppConfigList, error) {
	var list DistributedAppConfigList
	if err := r.c.ReadQuery(BasePath+DistributedAppEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DistributedApp configuration identified by id.
func (r *DistributedAppResource) Get(id string) (*DistributedAppConfig, error) {
	var item DistributedAppConfig
	if err := r.c.ReadQuery(BasePath+DistributedAppEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DistributedApp configuration.
func (r *DistributedAppResource) Create(item DistributedAppConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DistributedAppEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DistributedApp configuration identified by id.
func (r *DistributedAppResource) Edit(id string, item DistributedAppConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DistributedAppEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DistributedApp configuration identified by id.
func (r *DistributedAppResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DistributedAppEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
