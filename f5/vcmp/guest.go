// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcmp

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GuestConfigList holds a list of Guest configuration.
type GuestConfigList struct {
	Items    []GuestConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// GuestConfig holds the configuration of a single Guest.
type GuestConfig struct {
}

// GuestEndpoint represents the REST resource for managing Guest.
const GuestEndpoint = "/guest"

// GuestResource provides an API to manage Guest configurations.
type GuestResource struct {
	c f5.Client
}

// ListAll  lists all the Guest configurations.
func (r *GuestResource) ListAll() (*GuestConfigList, error) {
	var list GuestConfigList
	if err := r.c.ReadQuery(BasePath+GuestEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Guest configuration identified by id.
func (r *GuestResource) Get(id string) (*GuestConfig, error) {
	var item GuestConfig
	if err := r.c.ReadQuery(BasePath+GuestEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Guest configuration.
func (r *GuestResource) Create(item GuestConfig) error {
	if err := r.c.ModQuery("POST", BasePath+GuestEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Guest configuration identified by id.
func (r *GuestResource) Edit(id string, item GuestConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+GuestEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Guest configuration identified by id.
func (r *GuestResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+GuestEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
