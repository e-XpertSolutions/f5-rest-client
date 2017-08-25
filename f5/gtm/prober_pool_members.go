// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ProberPoolMembersConfigList holds a list of ProberPoolMembers configuration.
type ProberPoolMembersConfigList struct {
	Items    []ProberPoolMembersConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

// ProberPoolMembersConfig holds the configuration of a single ProberPoolMembers.
type ProberPoolMembersConfig struct {
}

// ProberPoolMembersEndpoint represents the REST resource for managing ProberPoolMembers.
const ProberPoolMembersEndpoint = "/prober-pool_members"

// ProberPoolMembersResource provides an API to manage ProberPoolMembers configurations.
type ProberPoolMembersResource struct {
	c *f5.Client
}

// ListAll  lists all the ProberPoolMembers configurations.
func (r *ProberPoolMembersResource) ListAll() (*ProberPoolMembersConfigList, error) {
	var list ProberPoolMembersConfigList
	if err := r.c.ReadQuery(BasePath+ProberPoolMembersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ProberPoolMembers configuration identified by id.
func (r *ProberPoolMembersResource) Get(id string) (*ProberPoolMembersConfig, error) {
	var item ProberPoolMembersConfig
	if err := r.c.ReadQuery(BasePath+ProberPoolMembersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ProberPoolMembers configuration.
func (r *ProberPoolMembersResource) Create(item ProberPoolMembersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ProberPoolMembersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ProberPoolMembers configuration identified by id.
func (r *ProberPoolMembersResource) Edit(id string, item ProberPoolMembersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ProberPoolMembersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ProberPoolMembers configuration identified by id.
func (r *ProberPoolMembersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ProberPoolMembersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
