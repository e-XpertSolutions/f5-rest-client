// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolNAPTRConfigList holds a list of PoolNAPTR configuration.
type PoolNAPTRConfigList struct {
	Items    []PoolNAPTRConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// PoolNAPTRConfig holds the configuration of a single PoolNAPTR.
type PoolNAPTRConfig struct {
}

// PoolNAPTREndpoint represents the REST resource for managing PoolNAPTR.
const PoolNAPTREndpoint = "/pool/naptr"

// PoolNAPTRResource provides an API to manage PoolNAPTR configurations.
type PoolNAPTRResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolNAPTR configurations.
func (r *PoolNAPTRResource) ListAll() (*PoolNAPTRConfigList, error) {
	var list PoolNAPTRConfigList
	if err := r.c.ReadQuery(BasePath+PoolNAPTREndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolNAPTR configuration identified by id.
func (r *PoolNAPTRResource) Get(id string) (*PoolNAPTRConfig, error) {
	var item PoolNAPTRConfig
	if err := r.c.ReadQuery(BasePath+PoolNAPTREndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolNAPTR configuration.
func (r *PoolNAPTRResource) Create(item PoolNAPTRConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolNAPTREndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolNAPTR configuration identified by id.
func (r *PoolNAPTRResource) Edit(id string, item PoolNAPTRConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolNAPTREndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolNAPTR configuration identified by id.
func (r *PoolNAPTRResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolNAPTREndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
