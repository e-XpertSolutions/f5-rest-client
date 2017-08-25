// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FixConnectionConfigList holds a list of FixConnection configuration.
type FixConnectionConfigList struct {
	Items    []FixConnectionConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// FixConnectionConfig holds the configuration of a single FixConnection.
type FixConnectionConfig struct {
}

// FixConnectionEndpoint represents the REST resource for managing FixConnection.
const FixConnectionEndpoint = "/fix-connection"

// FixConnectionResource provides an API to manage FixConnection configurations.
type FixConnectionResource struct {
	c *f5.Client
}

// ListAll  lists all the FixConnection configurations.
func (r *FixConnectionResource) ListAll() (*FixConnectionConfigList, error) {
	var list FixConnectionConfigList
	if err := r.c.ReadQuery(BasePath+FixConnectionEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FixConnection configuration identified by id.
func (r *FixConnectionResource) Get(id string) (*FixConnectionConfig, error) {
	var item FixConnectionConfig
	if err := r.c.ReadQuery(BasePath+FixConnectionEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FixConnection configuration.
func (r *FixConnectionResource) Create(item FixConnectionConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FixConnectionEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FixConnection configuration identified by id.
func (r *FixConnectionResource) Edit(id string, item FixConnectionConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FixConnectionEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FixConnection configuration identified by id.
func (r *FixConnectionResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FixConnectionEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
