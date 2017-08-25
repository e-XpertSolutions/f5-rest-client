// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallUserDomainConfigList holds a list of FirewallUserDomain configuration.
type FirewallUserDomainConfigList struct {
	Items    []FirewallUserDomainConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// FirewallUserDomainConfig holds the configuration of a single FirewallUserDomain.
type FirewallUserDomainConfig struct {
}

// FirewallUserDomainEndpoint represents the REST resource for managing FirewallUserDomain.
const FirewallUserDomainEndpoint = "/firewall/user-domain"

// FirewallUserDomainResource provides an API to manage FirewallUserDomain configurations.
type FirewallUserDomainResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallUserDomain configurations.
func (r *FirewallUserDomainResource) ListAll() (*FirewallUserDomainConfigList, error) {
	var list FirewallUserDomainConfigList
	if err := r.c.ReadQuery(BasePath+FirewallUserDomainEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallUserDomain configuration identified by id.
func (r *FirewallUserDomainResource) Get(id string) (*FirewallUserDomainConfig, error) {
	var item FirewallUserDomainConfig
	if err := r.c.ReadQuery(BasePath+FirewallUserDomainEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallUserDomain configuration.
func (r *FirewallUserDomainResource) Create(item FirewallUserDomainConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallUserDomainEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallUserDomain configuration identified by id.
func (r *FirewallUserDomainResource) Edit(id string, item FirewallUserDomainConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallUserDomainEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallUserDomain configuration identified by id.
func (r *FirewallUserDomainResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallUserDomainEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
