// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolNAPTRMembersConfigList holds a list of PoolNAPTRMembers configuration.
type PoolNAPTRMembersConfigList struct {
	Items    []PoolNAPTRMembersConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// PoolNAPTRMembersConfig holds the configuration of a single PoolNAPTRMembers.
type PoolNAPTRMembersConfig struct {
}

// PoolNAPTRMembersEndpoint represents the REST resource for managing PoolNAPTRMembers.
const PoolNAPTRMembersEndpoint = "/pool/naptr_members"

// PoolNAPTRMembersResource provides an API to manage PoolNAPTRMembers configurations.
type PoolNAPTRMembersResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolNAPTRMembers configurations.
func (r *PoolNAPTRMembersResource) ListAll() (*PoolNAPTRMembersConfigList, error) {
	var list PoolNAPTRMembersConfigList
	if err := r.c.ReadQuery(BasePath+PoolNAPTRMembersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolNAPTRMembers configuration identified by id.
func (r *PoolNAPTRMembersResource) Get(id string) (*PoolNAPTRMembersConfig, error) {
	var item PoolNAPTRMembersConfig
	if err := r.c.ReadQuery(BasePath+PoolNAPTRMembersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolNAPTRMembers configuration.
func (r *PoolNAPTRMembersResource) Create(item PoolNAPTRMembersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolNAPTRMembersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolNAPTRMembers configuration identified by id.
func (r *PoolNAPTRMembersResource) Edit(id string, item PoolNAPTRMembersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolNAPTRMembersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolNAPTRMembers configuration identified by id.
func (r *PoolNAPTRMembersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolNAPTRMembersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
