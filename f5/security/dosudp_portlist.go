// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSUDPPortlistConfigList holds a list of DOSUDPPortlist configuration.
type DOSUDPPortlistConfigList struct {
	Items    []DOSUDPPortlistConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// DOSUDPPortlistConfig holds the configuration of a single DOSUDPPortlist.
type DOSUDPPortlistConfig struct {
}

// DOSUDPPortlistEndpoint represents the REST resource for managing DOSUDPPortlist.
const DOSUDPPortlistEndpoint = "/dos/udp-portlist"

// DOSUDPPortlistResource provides an API to manage DOSUDPPortlist configurations.
type DOSUDPPortlistResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSUDPPortlist configurations.
func (r *DOSUDPPortlistResource) ListAll() (*DOSUDPPortlistConfigList, error) {
	var list DOSUDPPortlistConfigList
	if err := r.c.ReadQuery(BasePath+DOSUDPPortlistEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSUDPPortlist configuration identified by id.
func (r *DOSUDPPortlistResource) Get(id string) (*DOSUDPPortlistConfig, error) {
	var item DOSUDPPortlistConfig
	if err := r.c.ReadQuery(BasePath+DOSUDPPortlistEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSUDPPortlist configuration.
func (r *DOSUDPPortlistResource) Create(item DOSUDPPortlistConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSUDPPortlistEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSUDPPortlist configuration identified by id.
func (r *DOSUDPPortlistResource) Edit(id string, item DOSUDPPortlistConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSUDPPortlistEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSUDPPortlist configuration identified by id.
func (r *DOSUDPPortlistResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSUDPPortlistEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
