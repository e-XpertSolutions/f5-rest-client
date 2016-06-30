// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolSRVMembersConfigList holds a list of PoolSRVMembers configuration.
type PoolSRVMembersConfigList struct {
	Items    []PoolSRVMembersConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// PoolSRVMembersConfig holds the configuration of a single PoolSRVMembers.
type PoolSRVMembersConfig struct {
}

// PoolSRVMembersEndpoint represents the REST resource for managing PoolSRVMembers.
const PoolSRVMembersEndpoint = "/pool/srv_members"

// PoolSRVMembersResource provides an API to manage PoolSRVMembers configurations.
type PoolSRVMembersResource struct {
	c f5.Client
}

// ListAll  lists all the PoolSRVMembers configurations.
func (r *PoolSRVMembersResource) ListAll() (*PoolSRVMembersConfigList, error) {
	var list PoolSRVMembersConfigList
	if err := r.c.ReadQuery(BasePath+PoolSRVMembersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolSRVMembers configuration identified by id.
func (r *PoolSRVMembersResource) Get(id string) (*PoolSRVMembersConfig, error) {
	var item PoolSRVMembersConfig
	if err := r.c.ReadQuery(BasePath+PoolSRVMembersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolSRVMembers configuration.
func (r *PoolSRVMembersResource) Create(item PoolSRVMembersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolSRVMembersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolSRVMembers configuration identified by id.
func (r *PoolSRVMembersResource) Edit(id string, item PoolSRVMembersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolSRVMembersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolSRVMembers configuration identified by id.
func (r *PoolSRVMembersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolSRVMembersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
