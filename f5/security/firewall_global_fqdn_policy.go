// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallGlobalFQDNPolicyConfigList holds a list of FirewallGlobalFQDNPolicy configuration.
type FirewallGlobalFQDNPolicyConfigList struct {
	Items    []FirewallGlobalFQDNPolicyConfig `json:"items"`
	Kind     string                           `json:"kind"`
	SelfLink string                           `json:"selflink"`
}

// FirewallGlobalFQDNPolicyConfig holds the configuration of a single FirewallGlobalFQDNPolicy.
type FirewallGlobalFQDNPolicyConfig struct {
}

// FirewallGlobalFQDNPolicyEndpoint represents the REST resource for managing FirewallGlobalFQDNPolicy.
const FirewallGlobalFQDNPolicyEndpoint = "/firewall/global-fqdn-policy"

// FirewallGlobalFQDNPolicyResource provides an API to manage FirewallGlobalFQDNPolicy configurations.
type FirewallGlobalFQDNPolicyResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallGlobalFQDNPolicy configurations.
func (r *FirewallGlobalFQDNPolicyResource) ListAll() (*FirewallGlobalFQDNPolicyConfigList, error) {
	var list FirewallGlobalFQDNPolicyConfigList
	if err := r.c.ReadQuery(BasePath+FirewallGlobalFQDNPolicyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallGlobalFQDNPolicy configuration identified by id.
func (r *FirewallGlobalFQDNPolicyResource) Get(id string) (*FirewallGlobalFQDNPolicyConfig, error) {
	var item FirewallGlobalFQDNPolicyConfig
	if err := r.c.ReadQuery(BasePath+FirewallGlobalFQDNPolicyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallGlobalFQDNPolicy configuration.
func (r *FirewallGlobalFQDNPolicyResource) Create(item FirewallGlobalFQDNPolicyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallGlobalFQDNPolicyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallGlobalFQDNPolicy configuration identified by id.
func (r *FirewallGlobalFQDNPolicyResource) Edit(id string, item FirewallGlobalFQDNPolicyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallGlobalFQDNPolicyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallGlobalFQDNPolicy configuration identified by id.
func (r *FirewallGlobalFQDNPolicyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallGlobalFQDNPolicyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
