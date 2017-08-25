// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ICallIStatsTriggerConfigList holds a list of ICallIStatsTrigger configuration.
type ICallIStatsTriggerConfigList struct {
	Items    []ICallIStatsTriggerConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// ICallIStatsTriggerConfig holds the configuration of a single ICallIStatsTrigger.
type ICallIStatsTriggerConfig struct {
}

// ICallIStatsTriggerEndpoint represents the REST resource for managing ICallIStatsTrigger.
const ICallIStatsTriggerEndpoint = "/icall/istats-trigger"

// ICallIStatsTriggerResource provides an API to manage ICallIStatsTrigger configurations.
type ICallIStatsTriggerResource struct {
	c *f5.Client
}

// ListAll  lists all the ICallIStatsTrigger configurations.
func (r *ICallIStatsTriggerResource) ListAll() (*ICallIStatsTriggerConfigList, error) {
	var list ICallIStatsTriggerConfigList
	if err := r.c.ReadQuery(BasePath+ICallIStatsTriggerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ICallIStatsTrigger configuration identified by id.
func (r *ICallIStatsTriggerResource) Get(id string) (*ICallIStatsTriggerConfig, error) {
	var item ICallIStatsTriggerConfig
	if err := r.c.ReadQuery(BasePath+ICallIStatsTriggerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ICallIStatsTrigger configuration.
func (r *ICallIStatsTriggerResource) Create(item ICallIStatsTriggerConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ICallIStatsTriggerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ICallIStatsTrigger configuration identified by id.
func (r *ICallIStatsTriggerResource) Edit(id string, item ICallIStatsTriggerConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ICallIStatsTriggerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ICallIStatsTrigger configuration identified by id.
func (r *ICallIStatsTriggerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ICallIStatsTriggerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
