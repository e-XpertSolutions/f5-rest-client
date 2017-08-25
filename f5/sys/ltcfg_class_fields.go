// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LTCFGClassFieldsConfigList holds a list of LTCFGClassFields configuration.
type LTCFGClassFieldsConfigList struct {
	Items    []LTCFGClassFieldsConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// LTCFGClassFieldsConfig holds the configuration of a single LTCFGClassFields.
type LTCFGClassFieldsConfig struct {
}

// LTCFGClassFieldsEndpoint represents the REST resource for managing LTCFGClassFields.
const LTCFGClassFieldsEndpoint = "/ltcfg-class_fields"

// LTCFGClassFieldsResource provides an API to manage LTCFGClassFields configurations.
type LTCFGClassFieldsResource struct {
	c *f5.Client
}

// ListAll  lists all the LTCFGClassFields configurations.
func (r *LTCFGClassFieldsResource) ListAll() (*LTCFGClassFieldsConfigList, error) {
	var list LTCFGClassFieldsConfigList
	if err := r.c.ReadQuery(BasePath+LTCFGClassFieldsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LTCFGClassFields configuration identified by id.
func (r *LTCFGClassFieldsResource) Get(id string) (*LTCFGClassFieldsConfig, error) {
	var item LTCFGClassFieldsConfig
	if err := r.c.ReadQuery(BasePath+LTCFGClassFieldsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LTCFGClassFields configuration.
func (r *LTCFGClassFieldsResource) Create(item LTCFGClassFieldsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LTCFGClassFieldsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LTCFGClassFields configuration identified by id.
func (r *LTCFGClassFieldsResource) Edit(id string, item LTCFGClassFieldsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LTCFGClassFieldsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LTCFGClassFields configuration identified by id.
func (r *LTCFGClassFieldsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LTCFGClassFieldsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
