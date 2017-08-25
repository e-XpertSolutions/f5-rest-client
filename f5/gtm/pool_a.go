// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolAConfigList holds a list of PoolA configuration.
type PoolAConfigList struct {
	Items    []PoolAConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// PoolAConfig holds the configuration of a single PoolA.
type PoolAConfig struct {
}

// PoolAEndpoint represents the REST resource for managing PoolA.
const PoolAEndpoint = "/pool/a"

// PoolAResource provides an API to manage PoolA configurations.
type PoolAResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolA configurations.
func (r *PoolAResource) ListAll() (*PoolAConfigList, error) {
	var list PoolAConfigList
	if err := r.c.ReadQuery(BasePath+PoolAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolA configuration identified by id.
func (r *PoolAResource) Get(id string) (*PoolAConfig, error) {
	var item PoolAConfig
	if err := r.c.ReadQuery(BasePath+PoolAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolA configuration.
func (r *PoolAResource) Create(item PoolAConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolA configuration identified by id.
func (r *PoolAResource) Edit(id string, item PoolAConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolA configuration identified by id.
func (r *PoolAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
