// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolCNAMEMembersConfigList holds a list of PoolCNAMEMembers configuration.
type PoolCNAMEMembersConfigList struct {
	Items    []PoolCNAMEMembersConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// PoolCNAMEMembersConfig holds the configuration of a single PoolCNAMEMembers.
type PoolCNAMEMembersConfig struct {
}

// PoolCNAMEMembersEndpoint represents the REST resource for managing PoolCNAMEMembers.
const PoolCNAMEMembersEndpoint = "/pool/cname_members"

// PoolCNAMEMembersResource provides an API to manage PoolCNAMEMembers configurations.
type PoolCNAMEMembersResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolCNAMEMembers configurations.
func (r *PoolCNAMEMembersResource) ListAll() (*PoolCNAMEMembersConfigList, error) {
	var list PoolCNAMEMembersConfigList
	if err := r.c.ReadQuery(BasePath+PoolCNAMEMembersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolCNAMEMembers configuration identified by id.
func (r *PoolCNAMEMembersResource) Get(id string) (*PoolCNAMEMembersConfig, error) {
	var item PoolCNAMEMembersConfig
	if err := r.c.ReadQuery(BasePath+PoolCNAMEMembersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PoolCNAMEMembers configuration.
func (r *PoolCNAMEMembersResource) Create(item PoolCNAMEMembersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PoolCNAMEMembersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolCNAMEMembers configuration identified by id.
func (r *PoolCNAMEMembersResource) Edit(id string, item PoolCNAMEMembersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolCNAMEMembersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolCNAMEMembers configuration identified by id.
func (r *PoolCNAMEMembersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolCNAMEMembersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
