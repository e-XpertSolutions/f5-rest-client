// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AOMConfigList holds a list of AOM configuration.
type AOMConfigList struct {
	Items    []AOMConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// AOMConfig holds the configuration of a single AOM.
type AOMConfig struct {
}

// AOMEndpoint represents the REST resource for managing AOM.
const AOMEndpoint = "/aom"

// AOMResource provides an API to manage AOM configurations.
type AOMResource struct {
	c *f5.Client
}

// ListAll  lists all the AOM configurations.
func (r *AOMResource) ListAll() (*AOMConfigList, error) {
	var list AOMConfigList
	if err := r.c.ReadQuery(BasePath+AOMEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AOM configuration identified by id.
func (r *AOMResource) Get(id string) (*AOMConfig, error) {
	var item AOMConfig
	if err := r.c.ReadQuery(BasePath+AOMEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AOM configuration.
func (r *AOMResource) Create(item AOMConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AOMEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AOM configuration identified by id.
func (r *AOMResource) Edit(id string, item AOMConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AOMEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AOM configuration identified by id.
func (r *AOMResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AOMEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
