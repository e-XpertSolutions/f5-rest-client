// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// TopologyConfigList holds a list of Topology configuration.
type TopologyConfigList struct {
	Items    []TopologyConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// TopologyConfig holds the configuration of a single Topology.
type TopologyConfig struct {
}

// TopologyEndpoint represents the REST resource for managing Topology.
const TopologyEndpoint = "/topology"

// TopologyResource provides an API to manage Topology configurations.
type TopologyResource struct {
	c *f5.Client
}

// ListAll  lists all the Topology configurations.
func (r *TopologyResource) ListAll() (*TopologyConfigList, error) {
	var list TopologyConfigList
	if err := r.c.ReadQuery(BasePath+TopologyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Topology configuration identified by id.
func (r *TopologyResource) Get(id string) (*TopologyConfig, error) {
	var item TopologyConfig
	if err := r.c.ReadQuery(BasePath+TopologyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Topology configuration.
func (r *TopologyResource) Create(item TopologyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+TopologyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Topology configuration identified by id.
func (r *TopologyResource) Edit(id string, item TopologyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+TopologyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Topology configuration identified by id.
func (r *TopologyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+TopologyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
