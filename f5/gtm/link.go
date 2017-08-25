// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LinkConfigList holds a list of Link configuration.
type LinkConfigList struct {
	Items    []LinkConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// LinkConfig holds the configuration of a single Link.
type LinkConfig struct {
}

// LinkEndpoint represents the REST resource for managing Link.
const LinkEndpoint = "/link"

// LinkResource provides an API to manage Link configurations.
type LinkResource struct {
	c *f5.Client
}

// ListAll  lists all the Link configurations.
func (r *LinkResource) ListAll() (*LinkConfigList, error) {
	var list LinkConfigList
	if err := r.c.ReadQuery(BasePath+LinkEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Link configuration identified by id.
func (r *LinkResource) Get(id string) (*LinkConfig, error) {
	var item LinkConfig
	if err := r.c.ReadQuery(BasePath+LinkEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Link configuration.
func (r *LinkResource) Create(item LinkConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LinkEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Link configuration identified by id.
func (r *LinkResource) Edit(id string, item LinkConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LinkEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Link configuration identified by id.
func (r *LinkResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LinkEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
