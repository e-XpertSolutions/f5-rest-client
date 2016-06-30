// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolAAAAMembersConfigList holds a list of PoolAAAAMembers configuration.
type PoolAAAAMembersConfigList struct {
	Items    []PoolAAAAMembersConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// PoolAAAAMembersConfig holds the configuration of a single PoolAAAAMembers.
type PoolAAAAMembersConfig struct {
}

// PoolAAAAMembersEndpoint represents the REST resource for managing PoolAAAAMembers.
const PoolAAAAMembersEndpoint = "/pool/aaaa_members"

// PoolAAAAMembersResource provides an API to manage PoolAAAAMembers configurations.
type PoolAAAAMembersResource struct {
	c f5.Client
}

// ListAll  lists all the PoolAAAAMembers configurations.
func (r *PoolAAAAMembersResource) ListAll() (*PoolAAAAMembersConfigList, error) {
	var list PoolAAAAMembersConfigList
	if err := r.c.ReadQuery(BasePath+PoolAAAAMembersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolAAAAMembers configuration identified by id.
func (r *PoolAAAAMembersResource) Get(id string) (*PoolAAAAMembersConfig, error) {
	var item PoolAAAAMembersConfig
	if err := r.c.ReadQuery(BasePath+PoolAAAAMembersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolAAAAMembers configuration.
func (r *PoolAAAAMembersResource) Create(item PoolAAAAMembersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolAAAAMembersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolAAAAMembers configuration identified by id.
func (r *PoolAAAAMembersResource) Edit(id string, item PoolAAAAMembersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolAAAAMembersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolAAAAMembers configuration identified by id.
func (r *PoolAAAAMembersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolAAAAMembersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
