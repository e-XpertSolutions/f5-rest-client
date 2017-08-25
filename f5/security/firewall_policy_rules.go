// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallPolicyRulesConfigList holds a list of FirewallPolicyRules configuration.
type FirewallPolicyRulesConfigList struct {
	Items    []FirewallPolicyRulesConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// FirewallPolicyRulesConfig holds the configuration of a single FirewallPolicyRules.
type FirewallPolicyRulesConfig struct {
}

// FirewallPolicyRulesEndpoint represents the REST resource for managing FirewallPolicyRules.
const FirewallPolicyRulesEndpoint = "/firewall/policy_rules"

// FirewallPolicyRulesResource provides an API to manage FirewallPolicyRules configurations.
type FirewallPolicyRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallPolicyRules configurations.
func (r *FirewallPolicyRulesResource) ListAll() (*FirewallPolicyRulesConfigList, error) {
	var list FirewallPolicyRulesConfigList
	if err := r.c.ReadQuery(BasePath+FirewallPolicyRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallPolicyRules configuration identified by id.
func (r *FirewallPolicyRulesResource) Get(id string) (*FirewallPolicyRulesConfig, error) {
	var item FirewallPolicyRulesConfig
	if err := r.c.ReadQuery(BasePath+FirewallPolicyRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallPolicyRules configuration.
func (r *FirewallPolicyRulesResource) Create(item FirewallPolicyRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallPolicyRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallPolicyRules configuration identified by id.
func (r *FirewallPolicyRulesResource) Edit(id string, item FirewallPolicyRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallPolicyRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallPolicyRules configuration identified by id.
func (r *FirewallPolicyRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallPolicyRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
