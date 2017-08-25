// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallGlobalRulesConfigList holds a list of FirewallGlobalRules configuration.
type FirewallGlobalRulesConfigList struct {
	Items    []FirewallGlobalRulesConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// FirewallGlobalRulesConfig holds the configuration of a single FirewallGlobalRules.
type FirewallGlobalRulesConfig struct {
}

// FirewallGlobalRulesEndpoint represents the REST resource for managing FirewallGlobalRules.
const FirewallGlobalRulesEndpoint = "/firewall/global-rules"

// FirewallGlobalRulesResource provides an API to manage FirewallGlobalRules configurations.
type FirewallGlobalRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallGlobalRules configurations.
func (r *FirewallGlobalRulesResource) ListAll() (*FirewallGlobalRulesConfigList, error) {
	var list FirewallGlobalRulesConfigList
	if err := r.c.ReadQuery(BasePath+FirewallGlobalRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallGlobalRules configuration identified by id.
func (r *FirewallGlobalRulesResource) Get(id string) (*FirewallGlobalRulesConfig, error) {
	var item FirewallGlobalRulesConfig
	if err := r.c.ReadQuery(BasePath+FirewallGlobalRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallGlobalRules configuration.
func (r *FirewallGlobalRulesResource) Create(item FirewallGlobalRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallGlobalRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallGlobalRules configuration identified by id.
func (r *FirewallGlobalRulesResource) Edit(id string, item FirewallGlobalRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallGlobalRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallGlobalRules configuration identified by id.
func (r *FirewallGlobalRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallGlobalRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
