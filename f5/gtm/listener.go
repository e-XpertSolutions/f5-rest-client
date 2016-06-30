// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ListenerConfigList holds a list of Listener configuration.
type ListenerConfigList struct {
	Items    []ListenerConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// ListenerConfig holds the configuration of a single Listener.
type ListenerConfig struct {
}

// ListenerEndpoint represents the REST resource for managing Listener.
const ListenerEndpoint = "/listener"

// ListenerResource provides an API to manage Listener configurations.
type ListenerResource struct {
	c f5.Client
}

// ListAll  lists all the Listener configurations.
func (r *ListenerResource) ListAll() (*ListenerConfigList, error) {
	var list ListenerConfigList
	if err := r.c.ReadQuery(BasePath+ListenerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Listener configuration identified by id.
func (r *ListenerResource) Get(id string) (*ListenerConfig, error) {
	var item ListenerConfig
	if err := r.c.ReadQuery(BasePath+ListenerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Listener configuration.
func (r *ListenerResource) Create(item ListenerConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ListenerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Listener configuration identified by id.
func (r *ListenerResource) Edit(id string, item ListenerConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ListenerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Listener configuration identified by id.
func (r *ListenerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ListenerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
