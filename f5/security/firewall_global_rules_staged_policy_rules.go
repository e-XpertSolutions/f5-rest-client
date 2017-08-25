// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallGlobalRulesStagedPolicyRulesConfigList holds a list of FirewallGlobalRulesStagedPolicyRules configuration.
type FirewallGlobalRulesStagedPolicyRulesConfigList struct {
	Items    []FirewallGlobalRulesStagedPolicyRulesConfig `json:"items"`
	Kind     string                                       `json:"kind"`
	SelfLink string                                       `json:"selflink"`
}

// FirewallGlobalRulesStagedPolicyRulesConfig holds the configuration of a single FirewallGlobalRulesStagedPolicyRules.
type FirewallGlobalRulesStagedPolicyRulesConfig struct {
}

// FirewallGlobalRulesStagedPolicyRulesEndpoint represents the REST resource for managing FirewallGlobalRulesStagedPolicyRules.
const FirewallGlobalRulesStagedPolicyRulesEndpoint = "/firewall/global-rules_staged-policy-rules"

// FirewallGlobalRulesStagedPolicyRulesResource provides an API to manage FirewallGlobalRulesStagedPolicyRules configurations.
type FirewallGlobalRulesStagedPolicyRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallGlobalRulesStagedPolicyRules configurations.
func (r *FirewallGlobalRulesStagedPolicyRulesResource) ListAll() (*FirewallGlobalRulesStagedPolicyRulesConfigList, error) {
	var list FirewallGlobalRulesStagedPolicyRulesConfigList
	if err := r.c.ReadQuery(BasePath+FirewallGlobalRulesStagedPolicyRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallGlobalRulesStagedPolicyRules configuration identified by id.
func (r *FirewallGlobalRulesStagedPolicyRulesResource) Get(id string) (*FirewallGlobalRulesStagedPolicyRulesConfig, error) {
	var item FirewallGlobalRulesStagedPolicyRulesConfig
	if err := r.c.ReadQuery(BasePath+FirewallGlobalRulesStagedPolicyRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallGlobalRulesStagedPolicyRules configuration.
func (r *FirewallGlobalRulesStagedPolicyRulesResource) Create(item FirewallGlobalRulesStagedPolicyRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallGlobalRulesStagedPolicyRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallGlobalRulesStagedPolicyRules configuration identified by id.
func (r *FirewallGlobalRulesStagedPolicyRulesResource) Edit(id string, item FirewallGlobalRulesStagedPolicyRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallGlobalRulesStagedPolicyRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallGlobalRulesStagedPolicyRules configuration identified by id.
func (r *FirewallGlobalRulesStagedPolicyRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallGlobalRulesStagedPolicyRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
