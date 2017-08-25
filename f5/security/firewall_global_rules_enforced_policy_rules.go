// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallGlobalRulesEnforcedPolicyRulesConfigList holds a list of FirewallGlobalRulesEnforcedPolicyRules configuration.
type FirewallGlobalRulesEnforcedPolicyRulesConfigList struct {
	Items    []FirewallGlobalRulesEnforcedPolicyRulesConfig `json:"items"`
	Kind     string                                         `json:"kind"`
	SelfLink string                                         `json:"selflink"`
}

// FirewallGlobalRulesEnforcedPolicyRulesConfig holds the configuration of a single FirewallGlobalRulesEnforcedPolicyRules.
type FirewallGlobalRulesEnforcedPolicyRulesConfig struct {
}

// FirewallGlobalRulesEnforcedPolicyRulesEndpoint represents the REST resource for managing FirewallGlobalRulesEnforcedPolicyRules.
const FirewallGlobalRulesEnforcedPolicyRulesEndpoint = "/firewall/global-rules_enforced-policy-rules"

// FirewallGlobalRulesEnforcedPolicyRulesResource provides an API to manage FirewallGlobalRulesEnforcedPolicyRules configurations.
type FirewallGlobalRulesEnforcedPolicyRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallGlobalRulesEnforcedPolicyRules configurations.
func (r *FirewallGlobalRulesEnforcedPolicyRulesResource) ListAll() (*FirewallGlobalRulesEnforcedPolicyRulesConfigList, error) {
	var list FirewallGlobalRulesEnforcedPolicyRulesConfigList
	if err := r.c.ReadQuery(BasePath+FirewallGlobalRulesEnforcedPolicyRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallGlobalRulesEnforcedPolicyRules configuration identified by id.
func (r *FirewallGlobalRulesEnforcedPolicyRulesResource) Get(id string) (*FirewallGlobalRulesEnforcedPolicyRulesConfig, error) {
	var item FirewallGlobalRulesEnforcedPolicyRulesConfig
	if err := r.c.ReadQuery(BasePath+FirewallGlobalRulesEnforcedPolicyRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallGlobalRulesEnforcedPolicyRules configuration.
func (r *FirewallGlobalRulesEnforcedPolicyRulesResource) Create(item FirewallGlobalRulesEnforcedPolicyRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallGlobalRulesEnforcedPolicyRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallGlobalRulesEnforcedPolicyRules configuration identified by id.
func (r *FirewallGlobalRulesEnforcedPolicyRulesResource) Edit(id string, item FirewallGlobalRulesEnforcedPolicyRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallGlobalRulesEnforcedPolicyRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallGlobalRulesEnforcedPolicyRules configuration identified by id.
func (r *FirewallGlobalRulesEnforcedPolicyRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallGlobalRulesEnforcedPolicyRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
