// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// HAGroupConfigList holds a list of HAGroup configuration.
type HAGroupConfigList struct {
	Items    []HAGroupConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// HAGroupConfig holds the configuration of a single HAGroup.
type HAGroupConfig struct {
}

// HAGroupEndpoint represents the REST resource for managing HAGroup.
const HAGroupEndpoint = "/ha-group"

// HAGroupResource provides an API to manage HAGroup configurations.
type HAGroupResource struct {
	c *f5.Client
}

// ListAll  lists all the HAGroup configurations.
func (r *HAGroupResource) ListAll() (*HAGroupConfigList, error) {
	var list HAGroupConfigList
	if err := r.c.ReadQuery(BasePath+HAGroupEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single HAGroup configuration identified by id.
func (r *HAGroupResource) Get(id string) (*HAGroupConfig, error) {
	var item HAGroupConfig
	if err := r.c.ReadQuery(BasePath+HAGroupEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new HAGroup configuration.
func (r *HAGroupResource) Create(item HAGroupConfig) error {
	if err := r.c.ModQuery("POST", BasePath+HAGroupEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a HAGroup configuration identified by id.
func (r *HAGroupResource) Edit(id string, item HAGroupConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+HAGroupEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single HAGroup configuration identified by id.
func (r *HAGroupResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+HAGroupEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
