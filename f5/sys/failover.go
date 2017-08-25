// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FailoverConfigList holds a list of Failover configuration.
type FailoverConfigList struct {
	Items    []FailoverConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// FailoverConfig holds the configuration of a single Failover.
type FailoverConfig struct {
}

// FailoverEndpoint represents the REST resource for managing Failover.
const FailoverEndpoint = "/failover"

// FailoverResource provides an API to manage Failover configurations.
type FailoverResource struct {
	c *f5.Client
}

// ListAll  lists all the Failover configurations.
func (r *FailoverResource) ListAll() (*FailoverConfigList, error) {
	var list FailoverConfigList
	if err := r.c.ReadQuery(BasePath+FailoverEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Failover configuration identified by id.
func (r *FailoverResource) Get(id string) (*FailoverConfig, error) {
	var item FailoverConfig
	if err := r.c.ReadQuery(BasePath+FailoverEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Failover configuration.
func (r *FailoverResource) Create(item FailoverConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FailoverEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Failover configuration identified by id.
func (r *FailoverResource) Edit(id string, item FailoverConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FailoverEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Failover configuration identified by id.
func (r *FailoverResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FailoverEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
