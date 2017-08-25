// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallRuleListRulesConfigList holds a list of FirewallRuleListRules configuration.
type FirewallRuleListRulesConfigList struct {
	Items    []FirewallRuleListRulesConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// FirewallRuleListRulesConfig holds the configuration of a single FirewallRuleListRules.
type FirewallRuleListRulesConfig struct {
}

// FirewallRuleListRulesEndpoint represents the REST resource for managing FirewallRuleListRules.
const FirewallRuleListRulesEndpoint = "/firewall/rule-list_rules"

// FirewallRuleListRulesResource provides an API to manage FirewallRuleListRules configurations.
type FirewallRuleListRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallRuleListRules configurations.
func (r *FirewallRuleListRulesResource) ListAll() (*FirewallRuleListRulesConfigList, error) {
	var list FirewallRuleListRulesConfigList
	if err := r.c.ReadQuery(BasePath+FirewallRuleListRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallRuleListRules configuration identified by id.
func (r *FirewallRuleListRulesResource) Get(id string) (*FirewallRuleListRulesConfig, error) {
	var item FirewallRuleListRulesConfig
	if err := r.c.ReadQuery(BasePath+FirewallRuleListRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallRuleListRules configuration.
func (r *FirewallRuleListRulesResource) Create(item FirewallRuleListRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallRuleListRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallRuleListRules configuration identified by id.
func (r *FirewallRuleListRulesResource) Edit(id string, item FirewallRuleListRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallRuleListRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallRuleListRules configuration identified by id.
func (r *FirewallRuleListRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallRuleListRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
