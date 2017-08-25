// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// RAIDConfigList holds a list of RAID configuration.
type RAIDConfigList struct {
	Items    []RAIDConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// RAIDConfig holds the configuration of a single RAID.
type RAIDConfig struct {
}

// RAIDEndpoint represents the REST resource for managing RAID.
const RAIDEndpoint = "/raid"

// RAIDResource provides an API to manage RAID configurations.
type RAIDResource struct {
	c *f5.Client
}

// ListAll  lists all the RAID configurations.
func (r *RAIDResource) ListAll() (*RAIDConfigList, error) {
	var list RAIDConfigList
	if err := r.c.ReadQuery(BasePath+RAIDEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single RAID configuration identified by id.
func (r *RAIDResource) Get(id string) (*RAIDConfig, error) {
	var item RAIDConfig
	if err := r.c.ReadQuery(BasePath+RAIDEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new RAID configuration.
func (r *RAIDResource) Create(item RAIDConfig) error {
	if err := r.c.ModQuery("POST", BasePath+RAIDEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a RAID configuration identified by id.
func (r *RAIDResource) Edit(id string, item RAIDConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+RAIDEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single RAID configuration identified by id.
func (r *RAIDResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+RAIDEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
