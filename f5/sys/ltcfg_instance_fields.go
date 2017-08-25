// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LTCFGInstanceFieldsConfigList holds a list of LTCFGInstanceFields configuration.
type LTCFGInstanceFieldsConfigList struct {
	Items    []LTCFGInstanceFieldsConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// LTCFGInstanceFieldsConfig holds the configuration of a single LTCFGInstanceFields.
type LTCFGInstanceFieldsConfig struct {
}

// LTCFGInstanceFieldsEndpoint represents the REST resource for managing LTCFGInstanceFields.
const LTCFGInstanceFieldsEndpoint = "/ltcfg-instance_fields"

// LTCFGInstanceFieldsResource provides an API to manage LTCFGInstanceFields configurations.
type LTCFGInstanceFieldsResource struct {
	c *f5.Client
}

// ListAll  lists all the LTCFGInstanceFields configurations.
func (r *LTCFGInstanceFieldsResource) ListAll() (*LTCFGInstanceFieldsConfigList, error) {
	var list LTCFGInstanceFieldsConfigList
	if err := r.c.ReadQuery(BasePath+LTCFGInstanceFieldsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LTCFGInstanceFields configuration identified by id.
func (r *LTCFGInstanceFieldsResource) Get(id string) (*LTCFGInstanceFieldsConfig, error) {
	var item LTCFGInstanceFieldsConfig
	if err := r.c.ReadQuery(BasePath+LTCFGInstanceFieldsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LTCFGInstanceFields configuration.
func (r *LTCFGInstanceFieldsResource) Create(item LTCFGInstanceFieldsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LTCFGInstanceFieldsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LTCFGInstanceFields configuration identified by id.
func (r *LTCFGInstanceFieldsResource) Edit(id string, item LTCFGInstanceFieldsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LTCFGInstanceFieldsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LTCFGInstanceFields configuration identified by id.
func (r *LTCFGInstanceFieldsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LTCFGInstanceFieldsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
