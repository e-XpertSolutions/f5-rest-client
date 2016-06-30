// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// RegionConfigList holds a list of Region configuration.
type RegionConfigList struct {
	Items    []RegionConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// RegionConfig holds the configuration of a single Region.
type RegionConfig struct {
}

// RegionEndpoint represents the REST resource for managing Region.
const RegionEndpoint = "/region"

// RegionResource provides an API to manage Region configurations.
type RegionResource struct {
	c f5.Client
}

// ListAll  lists all the Region configurations.
func (r *RegionResource) ListAll() (*RegionConfigList, error) {
	var list RegionConfigList
	if err := r.c.ReadQuery(BasePath+RegionEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Region configuration identified by id.
func (r *RegionResource) Get(id string) (*RegionConfig, error) {
	var item RegionConfig
	if err := r.c.ReadQuery(BasePath+RegionEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Region configuration.
func (r *RegionResource) Create(item RegionConfig) error {
	if err := r.c.ModQuery("POST", BasePath+RegionEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Region configuration identified by id.
func (r *RegionResource) Edit(id string, item RegionConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+RegionEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Region configuration identified by id.
func (r *RegionResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+RegionEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
