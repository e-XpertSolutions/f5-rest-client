// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ClockConfigList holds a list of Clock configuration.
type ClockConfigList struct {
	Items    []ClockConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// ClockConfig holds the configuration of a single Clock.
type ClockConfig struct {
}

// ClockEndpoint represents the REST resource for managing Clock.
const ClockEndpoint = "/clock"

// ClockResource provides an API to manage Clock configurations.
type ClockResource struct {
	c *f5.Client
}

// ListAll  lists all the Clock configurations.
func (r *ClockResource) ListAll() (*ClockConfigList, error) {
	var list ClockConfigList
	if err := r.c.ReadQuery(BasePath+ClockEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Clock configuration identified by id.
func (r *ClockResource) Get(id string) (*ClockConfig, error) {
	var item ClockConfig
	if err := r.c.ReadQuery(BasePath+ClockEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Clock configuration.
func (r *ClockResource) Create(item ClockConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ClockEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Clock configuration identified by id.
func (r *ClockResource) Edit(id string, item ClockConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ClockEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Clock configuration identified by id.
func (r *ClockResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ClockEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
