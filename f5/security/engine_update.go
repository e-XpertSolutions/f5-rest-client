// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// EngineUpdateConfigList holds a list of EngineUpdate configuration.
type EngineUpdateConfigList struct {
	Items    []EngineUpdateConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// EngineUpdateConfig holds the configuration of a single EngineUpdate.
type EngineUpdateConfig struct {
}

// EngineUpdateEndpoint represents the REST resource for managing EngineUpdate.
const EngineUpdateEndpoint = "/anti-fraud/engine-update"

// EngineUpdateResource provides an API to manage EngineUpdate configurations.
type EngineUpdateResource struct {
	c *f5.Client
}

// ListAll  lists all the EngineUpdate configurations.
func (r *EngineUpdateResource) ListAll() (*EngineUpdateConfigList, error) {
	var list EngineUpdateConfigList
	if err := r.c.ReadQuery(BasePath+EngineUpdateEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single EngineUpdate configuration identified by id.
func (r *EngineUpdateResource) Get(id string) (*EngineUpdateConfig, error) {
	var item EngineUpdateConfig
	if err := r.c.ReadQuery(BasePath+EngineUpdateEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new EngineUpdate configuration.
func (r *EngineUpdateResource) Create(item EngineUpdateConfig) error {
	if err := r.c.ModQuery("POST", BasePath+EngineUpdateEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a EngineUpdate configuration identified by id.
func (r *EngineUpdateResource) Edit(id string, item EngineUpdateConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+EngineUpdateEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single EngineUpdate configuration identified by id.
func (r *EngineUpdateResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+EngineUpdateEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
