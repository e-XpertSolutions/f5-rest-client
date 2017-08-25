// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ServerVirtualServersConfigList holds a list of ServerVirtualServers configuration.
type ServerVirtualServersConfigList struct {
	Items    []ServerVirtualServersConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// ServerVirtualServersConfig holds the configuration of a single ServerVirtualServers.
type ServerVirtualServersConfig struct {
}

// ServerVirtualServersEndpoint represents the REST resource for managing ServerVirtualServers.
const ServerVirtualServersEndpoint = "/server_virtual-servers"

// ServerVirtualServersResource provides an API to manage ServerVirtualServers configurations.
type ServerVirtualServersResource struct {
	c *f5.Client
}

// ListAll  lists all the ServerVirtualServers configurations.
func (r *ServerVirtualServersResource) ListAll() (*ServerVirtualServersConfigList, error) {
	var list ServerVirtualServersConfigList
	if err := r.c.ReadQuery(BasePath+ServerVirtualServersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ServerVirtualServers configuration identified by id.
func (r *ServerVirtualServersResource) Get(id string) (*ServerVirtualServersConfig, error) {
	var item ServerVirtualServersConfig
	if err := r.c.ReadQuery(BasePath+ServerVirtualServersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ServerVirtualServers configuration.
func (r *ServerVirtualServersResource) Create(item ServerVirtualServersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ServerVirtualServersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ServerVirtualServers configuration identified by id.
func (r *ServerVirtualServersResource) Edit(id string, item ServerVirtualServersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ServerVirtualServersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ServerVirtualServers configuration identified by id.
func (r *ServerVirtualServersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ServerVirtualServersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
