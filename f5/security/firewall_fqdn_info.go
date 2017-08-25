// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallFQDNInfoConfigList holds a list of FirewallFQDNInfo configuration.
type FirewallFQDNInfoConfigList struct {
	Items    []FirewallFQDNInfoConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// FirewallFQDNInfoConfig holds the configuration of a single FirewallFQDNInfo.
type FirewallFQDNInfoConfig struct {
}

// FirewallFQDNInfoEndpoint represents the REST resource for managing FirewallFQDNInfo.
const FirewallFQDNInfoEndpoint = "/firewall/fqdn-info"

// FirewallFQDNInfoResource provides an API to manage FirewallFQDNInfo configurations.
type FirewallFQDNInfoResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallFQDNInfo configurations.
func (r *FirewallFQDNInfoResource) ListAll() (*FirewallFQDNInfoConfigList, error) {
	var list FirewallFQDNInfoConfigList
	if err := r.c.ReadQuery(BasePath+FirewallFQDNInfoEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallFQDNInfo configuration identified by id.
func (r *FirewallFQDNInfoResource) Get(id string) (*FirewallFQDNInfoConfig, error) {
	var item FirewallFQDNInfoConfig
	if err := r.c.ReadQuery(BasePath+FirewallFQDNInfoEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallFQDNInfo configuration.
func (r *FirewallFQDNInfoResource) Create(item FirewallFQDNInfoConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallFQDNInfoEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallFQDNInfo configuration identified by id.
func (r *FirewallFQDNInfoResource) Edit(id string, item FirewallFQDNInfoConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallFQDNInfoEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallFQDNInfo configuration identified by id.
func (r *FirewallFQDNInfoResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallFQDNInfoEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
