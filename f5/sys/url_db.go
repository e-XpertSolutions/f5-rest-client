// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// URLDBConfigList holds a list of URLDB configuration.
type URLDBConfigList struct {
	Items    []URLDBConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// URLDBConfig holds the configuration of a single URLDB.
type URLDBConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// URLDBEndpoint represents the REST resource for managing URLDB.
const URLDBEndpoint = "/url-db"

// URLDBResource provides an API to manage URLDB configurations.
type URLDBResource struct {
	c *f5.Client
}

// ListAll  lists all the URLDB configurations.
func (r *URLDBResource) ListAll() (*URLDBConfigList, error) {
	var list URLDBConfigList
	if err := r.c.ReadQuery(BasePath+URLDBEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single URLDB configuration identified by id.
func (r *URLDBResource) Get(id string) (*URLDBConfig, error) {
	var item URLDBConfig
	if err := r.c.ReadQuery(BasePath+URLDBEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new URLDB configuration.
func (r *URLDBResource) Create(item URLDBConfig) error {
	if err := r.c.ModQuery("POST", BasePath+URLDBEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a URLDB configuration identified by id.
func (r *URLDBResource) Edit(id string, item URLDBConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+URLDBEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single URLDB configuration identified by id.
func (r *URLDBResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+URLDBEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
