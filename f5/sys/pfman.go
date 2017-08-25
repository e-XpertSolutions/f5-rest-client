// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PFManConfigList holds a list of PFMan configuration.
type PFManConfigList struct {
	Items    []PFManConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// PFManConfig holds the configuration of a single PFMan.
type PFManConfig struct {
}

// PFManEndpoint represents the REST resource for managing PFMan.
const PFManEndpoint = "/pfman"

// PFManResource provides an API to manage PFMan configurations.
type PFManResource struct {
	c *f5.Client
}

// ListAll  lists all the PFMan configurations.
func (r *PFManResource) ListAll() (*PFManConfigList, error) {
	var list PFManConfigList
	if err := r.c.ReadQuery(BasePath+PFManEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PFMan configuration identified by id.
func (r *PFManResource) Get(id string) (*PFManConfig, error) {
	var item PFManConfig
	if err := r.c.ReadQuery(BasePath+PFManEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PFMan configuration.
func (r *PFManResource) Create(item PFManConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PFManEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PFMan configuration identified by id.
func (r *PFManResource) Edit(id string, item PFManConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PFManEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PFMan configuration identified by id.
func (r *PFManResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PFManEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
