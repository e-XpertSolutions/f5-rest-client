// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ConnectionConfigList holds a list of Connection configuration.
type ConnectionConfigList struct {
	Items    []ConnectionConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// ConnectionConfig holds the configuration of a single Connection.
type ConnectionConfig struct {
}

// ConnectionEndpoint represents the REST resource for managing Connection.
const ConnectionEndpoint = "/connection"

// ConnectionResource provides an API to manage Connection configurations.
type ConnectionResource struct {
	c *f5.Client
}

// ListAll  lists all the Connection configurations.
func (r *ConnectionResource) ListAll() (*ConnectionConfigList, error) {
	var list ConnectionConfigList
	if err := r.c.ReadQuery(BasePath+ConnectionEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Connection configuration identified by id.
func (r *ConnectionResource) Get(id string) (*ConnectionConfig, error) {
	var item ConnectionConfig
	if err := r.c.ReadQuery(BasePath+ConnectionEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Connection configuration.
func (r *ConnectionResource) Create(item ConnectionConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ConnectionEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Connection configuration identified by id.
func (r *ConnectionResource) Edit(id string, item ConnectionConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ConnectionEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Connection configuration identified by id.
func (r *ConnectionResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ConnectionEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
