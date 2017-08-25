// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LTCFGInstanceConfigList holds a list of LTCFGInstance configuration.
type LTCFGInstanceConfigList struct {
	Items    []LTCFGInstanceConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// LTCFGInstanceConfig holds the configuration of a single LTCFGInstance.
type LTCFGInstanceConfig struct {
}

// LTCFGInstanceEndpoint represents the REST resource for managing LTCFGInstance.
const LTCFGInstanceEndpoint = "/ltcfg-instance"

// LTCFGInstanceResource provides an API to manage LTCFGInstance configurations.
type LTCFGInstanceResource struct {
	c *f5.Client
}

// ListAll  lists all the LTCFGInstance configurations.
func (r *LTCFGInstanceResource) ListAll() (*LTCFGInstanceConfigList, error) {
	var list LTCFGInstanceConfigList
	if err := r.c.ReadQuery(BasePath+LTCFGInstanceEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LTCFGInstance configuration identified by id.
func (r *LTCFGInstanceResource) Get(id string) (*LTCFGInstanceConfig, error) {
	var item LTCFGInstanceConfig
	if err := r.c.ReadQuery(BasePath+LTCFGInstanceEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LTCFGInstance configuration.
func (r *LTCFGInstanceResource) Create(item LTCFGInstanceConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LTCFGInstanceEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LTCFGInstance configuration identified by id.
func (r *LTCFGInstanceResource) Edit(id string, item LTCFGInstanceConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LTCFGInstanceEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LTCFGInstance configuration identified by id.
func (r *LTCFGInstanceResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LTCFGInstanceEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
