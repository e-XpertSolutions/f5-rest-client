// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallManagementIPRulesConfigList holds a list of FirewallManagementIPRules configuration.
type FirewallManagementIPRulesConfigList struct {
	Items    []FirewallManagementIPRulesConfig `json:"items"`
	Kind     string                            `json:"kind"`
	SelfLink string                            `json:"selflink"`
}

// FirewallManagementIPRulesConfig holds the configuration of a single FirewallManagementIPRules.
type FirewallManagementIPRulesConfig struct {
}

// FirewallManagementIPRulesEndpoint represents the REST resource for managing FirewallManagementIPRules.
const FirewallManagementIPRulesEndpoint = "/firewall/management-ip-rules"

// FirewallManagementIPRulesResource provides an API to manage FirewallManagementIPRules configurations.
type FirewallManagementIPRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallManagementIPRules configurations.
func (r *FirewallManagementIPRulesResource) ListAll() (*FirewallManagementIPRulesConfigList, error) {
	var list FirewallManagementIPRulesConfigList
	if err := r.c.ReadQuery(BasePath+FirewallManagementIPRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallManagementIPRules configuration identified by id.
func (r *FirewallManagementIPRulesResource) Get(id string) (*FirewallManagementIPRulesConfig, error) {
	var item FirewallManagementIPRulesConfig
	if err := r.c.ReadQuery(BasePath+FirewallManagementIPRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallManagementIPRules configuration.
func (r *FirewallManagementIPRulesResource) Create(item FirewallManagementIPRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallManagementIPRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallManagementIPRules configuration identified by id.
func (r *FirewallManagementIPRulesResource) Edit(id string, item FirewallManagementIPRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallManagementIPRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallManagementIPRules configuration identified by id.
func (r *FirewallManagementIPRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallManagementIPRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
