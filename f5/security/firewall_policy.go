// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallPolicyConfigList holds a list of FirewallPolicy configuration.
type FirewallPolicyConfigList struct {
	Items    []FirewallPolicyConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// FirewallPolicyConfig holds the configuration of a single FirewallPolicy.
type FirewallPolicyConfig struct {
}

// FirewallPolicyEndpoint represents the REST resource for managing FirewallPolicy.
const FirewallPolicyEndpoint = "/firewall/policy"

// FirewallPolicyResource provides an API to manage FirewallPolicy configurations.
type FirewallPolicyResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallPolicy configurations.
func (r *FirewallPolicyResource) ListAll() (*FirewallPolicyConfigList, error) {
	var list FirewallPolicyConfigList
	if err := r.c.ReadQuery(BasePath+FirewallPolicyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallPolicy configuration identified by id.
func (r *FirewallPolicyResource) Get(id string) (*FirewallPolicyConfig, error) {
	var item FirewallPolicyConfig
	if err := r.c.ReadQuery(BasePath+FirewallPolicyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallPolicy configuration.
func (r *FirewallPolicyResource) Create(item FirewallPolicyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallPolicyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallPolicy configuration identified by id.
func (r *FirewallPolicyResource) Edit(id string, item FirewallPolicyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallPolicyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallPolicy configuration identified by id.
func (r *FirewallPolicyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallPolicyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
