// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PersistConfigList holds a list of Persist configuration.
type PersistConfigList struct {
	Items    []PersistConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// PersistConfig holds the configuration of a single Persist.
type PersistConfig struct {
}

// PersistEndpoint represents the REST resource for managing Persist.
const PersistEndpoint = "/persist"

// PersistResource provides an API to manage Persist configurations.
type PersistResource struct {
	c f5.Client
}

// ListAll  lists all the Persist configurations.
func (r *PersistResource) ListAll() (*PersistConfigList, error) {
	var list PersistConfigList
	if err := r.c.ReadQuery(BasePath+PersistEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Persist configuration identified by id.
func (r *PersistResource) Get(id string) (*PersistConfig, error) {
	var item PersistConfig
	if err := r.c.ReadQuery(BasePath+PersistEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Persist configuration.
func (r *PersistResource) Create(item PersistConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PersistEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Persist configuration identified by id.
func (r *PersistResource) Edit(id string, item PersistConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PersistEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Persist configuration identified by id.
func (r *PersistResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PersistEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
