// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ServerConfigList holds a list of Server configuration.
type ServerConfigList struct {
	Items    []ServerConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// ServerConfig holds the configuration of a single Server.
type ServerConfig struct {
}

// ServerEndpoint represents the REST resource for managing Server.
const ServerEndpoint = "/service"

// ServerResource provides an API to manage Server configurations.
type ServerResource struct {
	c *f5.Client
}

// ListAll  lists all the Server configurations.
func (r *ServerResource) ListAll() (*ServerConfigList, error) {
	var list ServerConfigList
	if err := r.c.ReadQuery(BasePath+ServerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Server configuration identified by id.
func (r *ServerResource) Get(id string) (*ServerConfig, error) {
	var item ServerConfig
	if err := r.c.ReadQuery(BasePath+ServerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Server configuration.
func (r *ServerResource) Create(item ServerConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ServerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Server configuration identified by id.
func (r *ServerResource) Edit(id string, item ServerConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ServerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Server configuration identified by id.
func (r *ServerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ServerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
