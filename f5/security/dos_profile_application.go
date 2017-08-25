// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSProfileApplicationConfigList holds a list of DOSProfileApplication configuration.
type DOSProfileApplicationConfigList struct {
	Items    []DOSProfileApplicationConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// DOSProfileApplicationConfig holds the configuration of a single DOSProfileApplication.
type DOSProfileApplicationConfig struct {
}

// DOSProfileApplicationEndpoint represents the REST resource for managing DOSProfileApplication.
const DOSProfileApplicationEndpoint = "/dos/profile_application"

// DOSProfileApplicationResource provides an API to manage DOSProfileApplication configurations.
type DOSProfileApplicationResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSProfileApplication configurations.
func (r *DOSProfileApplicationResource) ListAll() (*DOSProfileApplicationConfigList, error) {
	var list DOSProfileApplicationConfigList
	if err := r.c.ReadQuery(BasePath+DOSProfileApplicationEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSProfileApplication configuration identified by id.
func (r *DOSProfileApplicationResource) Get(id string) (*DOSProfileApplicationConfig, error) {
	var item DOSProfileApplicationConfig
	if err := r.c.ReadQuery(BasePath+DOSProfileApplicationEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSProfileApplication configuration.
func (r *DOSProfileApplicationResource) Create(item DOSProfileApplicationConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSProfileApplicationEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSProfileApplication configuration identified by id.
func (r *DOSProfileApplicationResource) Edit(id string, item DOSProfileApplicationConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSProfileApplicationEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSProfileApplication configuration identified by id.
func (r *DOSProfileApplicationResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSProfileApplicationEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
