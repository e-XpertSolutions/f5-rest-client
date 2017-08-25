// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// NATPolicyConfigList holds a list of NATPolicy configuration.
type NATPolicyConfigList struct {
	Items    []NATPolicyConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// NATPolicyConfig holds the configuration of a single NATPolicy.
type NATPolicyConfig struct {
}

// NATPolicyEndpoint represents the REST resource for managing NATPolicy.
const NATPolicyEndpoint = "/nat/policy"

// NATPolicyResource provides an API to manage NATPolicy configurations.
type NATPolicyResource struct {
	c *f5.Client
}

// ListAll  lists all the NATPolicy configurations.
func (r *NATPolicyResource) ListAll() (*NATPolicyConfigList, error) {
	var list NATPolicyConfigList
	if err := r.c.ReadQuery(BasePath+NATPolicyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single NATPolicy configuration identified by id.
func (r *NATPolicyResource) Get(id string) (*NATPolicyConfig, error) {
	var item NATPolicyConfig
	if err := r.c.ReadQuery(BasePath+NATPolicyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new NATPolicy configuration.
func (r *NATPolicyResource) Create(item NATPolicyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+NATPolicyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a NATPolicy configuration identified by id.
func (r *NATPolicyResource) Edit(id string, item NATPolicyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+NATPolicyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single NATPolicy configuration identified by id.
func (r *NATPolicyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+NATPolicyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
