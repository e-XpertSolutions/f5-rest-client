// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// RegionList holds a list of Region configuration.
type RegionList struct {
	Items    []Region `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

// Region holds the configuration of a single Region.
type Region struct {
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Kind          string `json:"kind,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	RegionMembers []struct {
		Name string `json:"name,omitempty"`
	} `json:"regionMembers,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// RegionEndpoint represents the REST resource for managing Region.
const RegionEndpoint = "/region"

// RegionResource provides an API to manage Region configurations.
type RegionResource struct {
	c *f5.Client
}

// ListAll  lists all the Region configurations.
func (r *RegionResource) ListAll() (*RegionList, error) {
	var list RegionList
	if err := r.c.ReadQuery(BasePath+RegionEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Region configuration identified by id.
func (r *RegionResource) Get(id string) (*Region, error) {
	var item Region
	if err := r.c.ReadQuery(BasePath+RegionEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Region configuration.
func (r *RegionResource) Create(item Region) error {
	if err := r.c.ModQuery("POST", BasePath+RegionEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Region configuration identified by id.
func (r *RegionResource) Edit(id string, item Region) error {
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
