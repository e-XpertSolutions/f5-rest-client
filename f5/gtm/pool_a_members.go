// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolAMembersConfigList holds a list of PoolAMembers configuration.
type PoolAMembersConfigList struct {
	Items    []PoolAMembersConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// PoolAMembersConfig holds the configuration of a single PoolAMembers.
type PoolAMembersConfig struct {
}

// PoolAMembersEndpoint represents the REST resource for managing PoolAMembers.
const PoolAMembersEndpoint = "/pool/a_members"

// PoolAMembersResource provides an API to manage PoolAMembers configurations.
type PoolAMembersResource struct {
	c f5.Client
}

// ListAll  lists all the PoolAMembers configurations.
func (r *PoolAMembersResource) ListAll() (*PoolAMembersConfigList, error) {
	var list PoolAMembersConfigList
	if err := r.c.ReadQuery(BasePath+PoolAMembersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolAMembers configuration identified by id.
func (r *PoolAMembersResource) Get(id string) (*PoolAMembersConfig, error) {
	var item PoolAMembersConfig
	if err := r.c.ReadQuery(BasePath+PoolAMembersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolAMembers configuration.
func (r *PoolAMembersResource) Create(item PoolAMembersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolAMembersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolAMembers configuration identified by id.
func (r *PoolAMembersResource) Edit(id string, item PoolAMembersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolAMembersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolAMembers configuration identified by id.
func (r *PoolAMembersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolAMembersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
