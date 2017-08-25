// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallPortMisusePolicyConfigList holds a list of FirewallPortMisusePolicy configuration.
type FirewallPortMisusePolicyConfigList struct {
	Items    []FirewallPortMisusePolicyConfig `json:"items"`
	Kind     string                           `json:"kind"`
	SelfLink string                           `json:"selflink"`
}

// FirewallPortMisusePolicyConfig holds the configuration of a single FirewallPortMisusePolicy.
type FirewallPortMisusePolicyConfig struct {
}

// FirewallPortMisusePolicyEndpoint represents the REST resource for managing FirewallPortMisusePolicy.
const FirewallPortMisusePolicyEndpoint = "/firewall/port-misuse-policy"

// FirewallPortMisusePolicyResource provides an API to manage FirewallPortMisusePolicy configurations.
type FirewallPortMisusePolicyResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallPortMisusePolicy configurations.
func (r *FirewallPortMisusePolicyResource) ListAll() (*FirewallPortMisusePolicyConfigList, error) {
	var list FirewallPortMisusePolicyConfigList
	if err := r.c.ReadQuery(BasePath+FirewallPortMisusePolicyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallPortMisusePolicy configuration identified by id.
func (r *FirewallPortMisusePolicyResource) Get(id string) (*FirewallPortMisusePolicyConfig, error) {
	var item FirewallPortMisusePolicyConfig
	if err := r.c.ReadQuery(BasePath+FirewallPortMisusePolicyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallPortMisusePolicy configuration.
func (r *FirewallPortMisusePolicyResource) Create(item FirewallPortMisusePolicyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallPortMisusePolicyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallPortMisusePolicy configuration identified by id.
func (r *FirewallPortMisusePolicyResource) Edit(id string, item FirewallPortMisusePolicyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallPortMisusePolicyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallPortMisusePolicy configuration identified by id.
func (r *FirewallPortMisusePolicyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallPortMisusePolicyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
