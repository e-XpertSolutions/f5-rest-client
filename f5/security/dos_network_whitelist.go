// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSNetworkWhitelistConfigList holds a list of DOSNetworkWhitelist configuration.
type DOSNetworkWhitelistConfigList struct {
	Items    []DOSNetworkWhitelistConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// DOSNetworkWhitelistConfig holds the configuration of a single DOSNetworkWhitelist.
type DOSNetworkWhitelistConfig struct {
}

// DOSNetworkWhitelistEndpoint represents the REST resource for managing DOSNetworkWhitelist.
const DOSNetworkWhitelistEndpoint = "/dos/network-whitelist"

// DOSNetworkWhitelistResource provides an API to manage DOSNetworkWhitelist configurations.
type DOSNetworkWhitelistResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSNetworkWhitelist configurations.
func (r *DOSNetworkWhitelistResource) ListAll() (*DOSNetworkWhitelistConfigList, error) {
	var list DOSNetworkWhitelistConfigList
	if err := r.c.ReadQuery(BasePath+DOSNetworkWhitelistEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSNetworkWhitelist configuration identified by id.
func (r *DOSNetworkWhitelistResource) Get(id string) (*DOSNetworkWhitelistConfig, error) {
	var item DOSNetworkWhitelistConfig
	if err := r.c.ReadQuery(BasePath+DOSNetworkWhitelistEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSNetworkWhitelist configuration.
func (r *DOSNetworkWhitelistResource) Create(item DOSNetworkWhitelistConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSNetworkWhitelistEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSNetworkWhitelist configuration identified by id.
func (r *DOSNetworkWhitelistResource) Edit(id string, item DOSNetworkWhitelistConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSNetworkWhitelistEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSNetworkWhitelist configuration identified by id.
func (r *DOSNetworkWhitelistResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSNetworkWhitelistEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
