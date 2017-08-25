// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallUserGroupEntityConfigList holds a list of FirewallUserGroupEntity configuration.
type FirewallUserGroupEntityConfigList struct {
	Items    []FirewallUserGroupEntityConfig `json:"items"`
	Kind     string                          `json:"kind"`
	SelfLink string                          `json:"selflink"`
}

// FirewallUserGroupEntityConfig holds the configuration of a single FirewallUserGroupEntity.
type FirewallUserGroupEntityConfig struct {
}

// FirewallUserGroupEntityEndpoint represents the REST resource for managing FirewallUserGroupEntity.
const FirewallUserGroupEntityEndpoint = "/firewall/user-group-entity"

// FirewallUserGroupEntityResource provides an API to manage FirewallUserGroupEntity configurations.
type FirewallUserGroupEntityResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallUserGroupEntity configurations.
func (r *FirewallUserGroupEntityResource) ListAll() (*FirewallUserGroupEntityConfigList, error) {
	var list FirewallUserGroupEntityConfigList
	if err := r.c.ReadQuery(BasePath+FirewallUserGroupEntityEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallUserGroupEntity configuration identified by id.
func (r *FirewallUserGroupEntityResource) Get(id string) (*FirewallUserGroupEntityConfig, error) {
	var item FirewallUserGroupEntityConfig
	if err := r.c.ReadQuery(BasePath+FirewallUserGroupEntityEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallUserGroupEntity configuration.
func (r *FirewallUserGroupEntityResource) Create(item FirewallUserGroupEntityConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallUserGroupEntityEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallUserGroupEntity configuration identified by id.
func (r *FirewallUserGroupEntityResource) Edit(id string, item FirewallUserGroupEntityConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallUserGroupEntityEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallUserGroupEntity configuration identified by id.
func (r *FirewallUserGroupEntityResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallUserGroupEntityEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
