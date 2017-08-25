// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolMXMembersConfigList holds a list of PoolMXMembers configuration.
type PoolMXMembersConfigList struct {
	Items    []PoolMXMembersConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// PoolMXMembersConfig holds the configuration of a single PoolMXMembers.
type PoolMXMembersConfig struct {
}

// PoolMXMembersEndpoint represents the REST resource for managing PoolMXMembers.
const PoolMXMembersEndpoint = "/pool/mx_members"

// PoolMXMembersResource provides an API to manage PoolMXMembers configurations.
type PoolMXMembersResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolMXMembers configurations.
func (r *PoolMXMembersResource) ListAll() (*PoolMXMembersConfigList, error) {
	var list PoolMXMembersConfigList
	if err := r.c.ReadQuery(BasePath+PoolMXMembersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolMXMembers configuration identified by id.
func (r *PoolMXMembersResource) Get(id string) (*PoolMXMembersConfig, error) {
	var item PoolMXMembersConfig
	if err := r.c.ReadQuery(BasePath+PoolMXMembersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolMXMembers configuration.
func (r *PoolMXMembersResource) Create(item PoolMXMembersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolMXMembersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolMXMembers configuration identified by id.
func (r *PoolMXMembersResource) Edit(id string, item PoolMXMembersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolMXMembersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolMXMembers configuration identified by id.
func (r *PoolMXMembersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolMXMembersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
