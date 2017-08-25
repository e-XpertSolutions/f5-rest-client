// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DatacenterConfigList holds a list of Datacenter configuration.
type DatacenterConfigList struct {
	Items    []DatacenterConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// DatacenterConfig holds the configuration of a single Datacenter.
type DatacenterConfig struct {
}

// DatacenterEndpoint represents the REST resource for managing Datacenter.
const DatacenterEndpoint = "/datacenter"

// DatacenterResource provides an API to manage Datacenter configurations.
type DatacenterResource struct {
	c *f5.Client
}

// ListAll  lists all the Datacenter configurations.
func (r *DatacenterResource) ListAll() (*DatacenterConfigList, error) {
	var list DatacenterConfigList
	if err := r.c.ReadQuery(BasePath+DatacenterEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Datacenter configuration identified by id.
func (r *DatacenterResource) Get(id string) (*DatacenterConfig, error) {
	var item DatacenterConfig
	if err := r.c.ReadQuery(BasePath+DatacenterEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Datacenter configuration.
func (r *DatacenterResource) Create(item DatacenterConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DatacenterEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Datacenter configuration identified by id.
func (r *DatacenterResource) Edit(id string, item DatacenterConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DatacenterEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Datacenter configuration identified by id.
func (r *DatacenterResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DatacenterEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
