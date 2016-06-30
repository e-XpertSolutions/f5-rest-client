// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolConfigList holds a list of Pool configuration.
type PoolConfigList struct {
	Items    []PoolConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// PoolConfig holds the configuration of a single Pool.
type PoolConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// PoolEndpoint represents the REST resource for managing Pool.
const PoolEndpoint = "/pool"

// PoolResource provides an API to manage Pool configurations.
type PoolResource struct {
	c f5.Client
}

// ListAll  lists all the Pool configurations.
func (r *PoolResource) ListAll() (*PoolConfigList, error) {
	var list PoolConfigList
	if err := r.c.ReadQuery(BasePath+PoolEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Pool configuration identified by id.
func (r *PoolResource) Get(id string) (*PoolConfig, error) {
	var item PoolConfig
	if err := r.c.ReadQuery(BasePath+PoolEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Pool configuration.
func (r *PoolResource) Create(item PoolConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Pool configuration identified by id.
func (r *PoolResource) Edit(id string, item PoolConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Pool configuration identified by id.
func (r *PoolResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
