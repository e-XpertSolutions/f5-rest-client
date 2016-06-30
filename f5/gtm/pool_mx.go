// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolMXConfigList holds a list of PoolMX configuration.
type PoolMXConfigList struct {
	Items    []PoolMXConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// PoolMXConfig holds the configuration of a single PoolMX.
type PoolMXConfig struct {
}

// PoolMXEndpoint represents the REST resource for managing PoolMX.
const PoolMXEndpoint = "/pool/mx"

// PoolMXResource provides an API to manage PoolMX configurations.
type PoolMXResource struct {
	c f5.Client
}

// ListAll  lists all the PoolMX configurations.
func (r *PoolMXResource) ListAll() (*PoolMXConfigList, error) {
	var list PoolMXConfigList
	if err := r.c.ReadQuery(BasePath+PoolMXEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolMX configuration identified by id.
func (r *PoolMXResource) Get(id string) (*PoolMXConfig, error) {
	var item PoolMXConfig
	if err := r.c.ReadQuery(BasePath+PoolMXEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolMX configuration.
func (r *PoolMXResource) Create(item PoolMXConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolMXEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolMX configuration identified by id.
func (r *PoolMXResource) Edit(id string, item PoolMXConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolMXEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolMX configuration identified by id.
func (r *PoolMXResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolMXEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
