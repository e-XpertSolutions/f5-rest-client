// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallFQDNEntityConfigList holds a list of FirewallFQDNEntity configuration.
type FirewallFQDNEntityConfigList struct {
	Items    []FirewallFQDNEntityConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// FirewallFQDNEntityConfig holds the configuration of a single FirewallFQDNEntity.
type FirewallFQDNEntityConfig struct {
}

// FirewallFQDNEntityEndpoint represents the REST resource for managing FirewallFQDNEntity.
const FirewallFQDNEntityEndpoint = "/firewall/fqdn-entity"

// FirewallFQDNEntityResource provides an API to manage FirewallFQDNEntity configurations.
type FirewallFQDNEntityResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallFQDNEntity configurations.
func (r *FirewallFQDNEntityResource) ListAll() (*FirewallFQDNEntityConfigList, error) {
	var list FirewallFQDNEntityConfigList
	if err := r.c.ReadQuery(BasePath+FirewallFQDNEntityEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallFQDNEntity configuration identified by id.
func (r *FirewallFQDNEntityResource) Get(id string) (*FirewallFQDNEntityConfig, error) {
	var item FirewallFQDNEntityConfig
	if err := r.c.ReadQuery(BasePath+FirewallFQDNEntityEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallFQDNEntity configuration.
func (r *FirewallFQDNEntityResource) Create(item FirewallFQDNEntityConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallFQDNEntityEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallFQDNEntity configuration identified by id.
func (r *FirewallFQDNEntityResource) Edit(id string, item FirewallFQDNEntityConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallFQDNEntityEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallFQDNEntity configuration identified by id.
func (r *FirewallFQDNEntityResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallFQDNEntityEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
