// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallUserListConfigList holds a list of FirewallUserList configuration.
type FirewallUserListConfigList struct {
	Items    []FirewallUserListConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// FirewallUserListConfig holds the configuration of a single FirewallUserList.
type FirewallUserListConfig struct {
}

// FirewallUserListEndpoint represents the REST resource for managing FirewallUserList.
const FirewallUserListEndpoint = "/firewall/user-list"

// FirewallUserListResource provides an API to manage FirewallUserList configurations.
type FirewallUserListResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallUserList configurations.
func (r *FirewallUserListResource) ListAll() (*FirewallUserListConfigList, error) {
	var list FirewallUserListConfigList
	if err := r.c.ReadQuery(BasePath+FirewallUserListEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallUserList configuration identified by id.
func (r *FirewallUserListResource) Get(id string) (*FirewallUserListConfig, error) {
	var item FirewallUserListConfig
	if err := r.c.ReadQuery(BasePath+FirewallUserListEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallUserList configuration.
func (r *FirewallUserListResource) Create(item FirewallUserListConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallUserListEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallUserList configuration identified by id.
func (r *FirewallUserListResource) Edit(id string, item FirewallUserListConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallUserListEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallUserList configuration identified by id.
func (r *FirewallUserListResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallUserListEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
