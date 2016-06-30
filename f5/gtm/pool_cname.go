// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolCNAMEConfigList holds a list of PoolCNAME configuration.
type PoolCNAMEConfigList struct {
	Items    []PoolCNAMEConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// PoolCNAMEConfig holds the configuration of a single PoolCNAME.
type PoolCNAMEConfig struct {
}

// PoolCNAMEEndpoint represents the REST resource for managing PoolCNAME.
const PoolCNAMEEndpoint = "/pool/cname"

// PoolCNAMEResource provides an API to manage PoolCNAME configurations.
type PoolCNAMEResource struct {
	c f5.Client
}

// ListAll  lists all the PoolCNAME configurations.
func (r *PoolCNAMEResource) ListAll() (*PoolCNAMEConfigList, error) {
	var list PoolCNAMEConfigList
	if err := r.c.ReadQuery(BasePath+PoolCNAMEEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolCNAME configuration identified by id.
func (r *PoolCNAMEResource) Get(id string) (*PoolCNAMEConfig, error) {
	var item PoolCNAMEConfig
	if err := r.c.ReadQuery(BasePath+PoolCNAMEEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolCNAME configuration.
func (r *PoolCNAMEResource) Create(item PoolCNAMEConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolCNAMEEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolCNAME configuration identified by id.
func (r *PoolCNAMEResource) Edit(id string, item PoolCNAMEConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolCNAMEEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolCNAME configuration identified by id.
func (r *PoolCNAMEResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolCNAMEEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
