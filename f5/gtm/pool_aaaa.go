// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolAAAAConfigList holds a list of PoolAAAA configuration.
type PoolAAAAConfigList struct {
	Items    []PoolAAAAConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// PoolAAAAConfig holds the configuration of a single PoolAAAA.
type PoolAAAAConfig struct {
}

// PoolAAAAEndpoint represents the REST resource for managing PoolAAAA.
const PoolAAAAEndpoint = "/pool/aaaa"

// PoolAAAAResource provides an API to manage PoolAAAA configurations.
type PoolAAAAResource struct {
	c f5.Client
}

// ListAll  lists all the PoolAAAA configurations.
func (r *PoolAAAAResource) ListAll() (*PoolAAAAConfigList, error) {
	var list PoolAAAAConfigList
	if err := r.c.ReadQuery(BasePath+PoolAAAAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolAAAA configuration identified by id.
func (r *PoolAAAAResource) Get(id string) (*PoolAAAAConfig, error) {
	var item PoolAAAAConfig
	if err := r.c.ReadQuery(BasePath+PoolAAAAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolAAAA configuration.
func (r *PoolAAAAResource) Create(item PoolAAAAConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolAAAAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolAAAA configuration identified by id.
func (r *PoolAAAAResource) Edit(id string, item PoolAAAAConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolAAAAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolAAAA configuration identified by id.
func (r *PoolAAAAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolAAAAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
