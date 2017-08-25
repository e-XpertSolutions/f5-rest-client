// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ManagementOVSDBConfigList holds a list of ManagementOVSDB configuration.
type ManagementOVSDBConfigList struct {
	Items    []ManagementOVSDBConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// ManagementOVSDBConfig holds the configuration of a single ManagementOVSDB.
type ManagementOVSDBConfig struct {
}

// ManagementOVSDBEndpoint represents the REST resource for managing ManagementOVSDB.
const ManagementOVSDBEndpoint = "/management-ovsdb"

// ManagementOVSDBResource provides an API to manage ManagementOVSDB configurations.
type ManagementOVSDBResource struct {
	c *f5.Client
}

// ListAll  lists all the ManagementOVSDB configurations.
func (r *ManagementOVSDBResource) ListAll() (*ManagementOVSDBConfigList, error) {
	var list ManagementOVSDBConfigList
	if err := r.c.ReadQuery(BasePath+ManagementOVSDBEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ManagementOVSDB configuration identified by id.
func (r *ManagementOVSDBResource) Get(id string) (*ManagementOVSDBConfig, error) {
	var item ManagementOVSDBConfig
	if err := r.c.ReadQuery(BasePath+ManagementOVSDBEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ManagementOVSDB configuration.
func (r *ManagementOVSDBResource) Create(item ManagementOVSDBConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ManagementOVSDBEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ManagementOVSDB configuration identified by id.
func (r *ManagementOVSDBResource) Edit(id string, item ManagementOVSDBConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ManagementOVSDBEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ManagementOVSDB configuration identified by id.
func (r *ManagementOVSDBResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ManagementOVSDBEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
