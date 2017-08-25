// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LTCFGClassConfigList holds a list of LTCFGClass configuration.
type LTCFGClassConfigList struct {
	Items    []LTCFGClassConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// LTCFGClassConfig holds the configuration of a single LTCFGClass.
type LTCFGClassConfig struct {
}

// LTCFGClassEndpoint represents the REST resource for managing LTCFGClass.
const LTCFGClassEndpoint = "/ltcfg-class"

// LTCFGClassResource provides an API to manage LTCFGClass configurations.
type LTCFGClassResource struct {
	c *f5.Client
}

// ListAll  lists all the LTCFGClass configurations.
func (r *LTCFGClassResource) ListAll() (*LTCFGClassConfigList, error) {
	var list LTCFGClassConfigList
	if err := r.c.ReadQuery(BasePath+LTCFGClassEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LTCFGClass configuration identified by id.
func (r *LTCFGClassResource) Get(id string) (*LTCFGClassConfig, error) {
	var item LTCFGClassConfig
	if err := r.c.ReadQuery(BasePath+LTCFGClassEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LTCFGClass configuration.
func (r *LTCFGClassResource) Create(item LTCFGClassConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LTCFGClassEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LTCFGClass configuration identified by id.
func (r *LTCFGClassResource) Edit(id string, item LTCFGClassConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LTCFGClassEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LTCFGClass configuration identified by id.
func (r *LTCFGClassResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LTCFGClassEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
