// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallConfigList holds a list of Firewall configuration.
type FirewallConfigList struct {
	Items    []FirewallConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// FirewallConfig holds the configuration of a single Firewall.
type FirewallConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// FirewallEndpoint represents the REST resource for managing Firewall.
const FirewallEndpoint = "/firewall"

// FirewallResource provides an API to manage Firewall configurations.
type FirewallResource struct {
	c *f5.Client
}

// ListAll  lists all the Firewall configurations.
func (r *FirewallResource) ListAll() (*FirewallConfigList, error) {
	var list FirewallConfigList
	if err := r.c.ReadQuery(BasePath+FirewallEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Firewall configuration identified by id.
func (r *FirewallResource) Get(id string) (*FirewallConfig, error) {
	var item FirewallConfig
	if err := r.c.ReadQuery(BasePath+FirewallEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Firewall configuration.
func (r *FirewallResource) Create(item FirewallConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Firewall configuration identified by id.
func (r *FirewallResource) Edit(id string, item FirewallConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Firewall configuration identified by id.
func (r *FirewallResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
