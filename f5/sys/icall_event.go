// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ICallEventConfigList holds a list of ICallEvent configuration.
type ICallEventConfigList struct {
	Items    []ICallEventConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// ICallEventConfig holds the configuration of a single ICallEvent.
type ICallEventConfig struct {
}

// ICallEventEndpoint represents the REST resource for managing ICallEvent.
const ICallEventEndpoint = "/icall/event"

// ICallEventResource provides an API to manage ICallEvent configurations.
type ICallEventResource struct {
	c *f5.Client
}

// ListAll  lists all the ICallEvent configurations.
func (r *ICallEventResource) ListAll() (*ICallEventConfigList, error) {
	var list ICallEventConfigList
	if err := r.c.ReadQuery(BasePath+ICallEventEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ICallEvent configuration identified by id.
func (r *ICallEventResource) Get(id string) (*ICallEventConfig, error) {
	var item ICallEventConfig
	if err := r.c.ReadQuery(BasePath+ICallEventEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ICallEvent configuration.
func (r *ICallEventResource) Create(item ICallEventConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ICallEventEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ICallEvent configuration identified by id.
func (r *ICallEventResource) Edit(id string, item ICallEventConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ICallEventEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ICallEvent configuration identified by id.
func (r *ICallEventResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ICallEventEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
