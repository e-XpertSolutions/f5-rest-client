// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// StateMirroringConfigList holds a list of StateMirroring configuration.
type StateMirroringConfigList struct {
	Items    []StateMirroringConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// StateMirroringConfig holds the configuration of a single StateMirroring.
type StateMirroringConfig struct {
}

// StateMirroringEndpoint represents the REST resource for managing StateMirroring.
const StateMirroringEndpoint = "/state-mirroring"

// StateMirroringResource provides an API to manage StateMirroring configurations.
type StateMirroringResource struct {
	c *f5.Client
}

// ListAll  lists all the StateMirroring configurations.
func (r *StateMirroringResource) ListAll() (*StateMirroringConfigList, error) {
	var list StateMirroringConfigList
	if err := r.c.ReadQuery(BasePath+StateMirroringEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single StateMirroring configuration identified by id.
func (r *StateMirroringResource) Get(id string) (*StateMirroringConfig, error) {
	var item StateMirroringConfig
	if err := r.c.ReadQuery(BasePath+StateMirroringEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new StateMirroring configuration.
func (r *StateMirroringResource) Create(item StateMirroringConfig) error {
	if err := r.c.ModQuery("POST", BasePath+StateMirroringEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a StateMirroring configuration identified by id.
func (r *StateMirroringResource) Edit(id string, item StateMirroringConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+StateMirroringEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single StateMirroring configuration identified by id.
func (r *StateMirroringResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+StateMirroringEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
