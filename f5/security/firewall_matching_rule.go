// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallMatchingRuleConfigList holds a list of FirewallMatchingRule configuration.
type FirewallMatchingRuleConfigList struct {
	Items    []FirewallMatchingRuleConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// FirewallMatchingRuleConfig holds the configuration of a single FirewallMatchingRule.
type FirewallMatchingRuleConfig struct {
}

// FirewallMatchingRuleEndpoint represents the REST resource for managing FirewallMatchingRule.
const FirewallMatchingRuleEndpoint = "/firewall/matching-rule"

// FirewallMatchingRuleResource provides an API to manage FirewallMatchingRule configurations.
type FirewallMatchingRuleResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallMatchingRule configurations.
func (r *FirewallMatchingRuleResource) ListAll() (*FirewallMatchingRuleConfigList, error) {
	var list FirewallMatchingRuleConfigList
	if err := r.c.ReadQuery(BasePath+FirewallMatchingRuleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallMatchingRule configuration identified by id.
func (r *FirewallMatchingRuleResource) Get(id string) (*FirewallMatchingRuleConfig, error) {
	var item FirewallMatchingRuleConfig
	if err := r.c.ReadQuery(BasePath+FirewallMatchingRuleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallMatchingRule configuration.
func (r *FirewallMatchingRuleResource) Create(item FirewallMatchingRuleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallMatchingRuleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallMatchingRule configuration identified by id.
func (r *FirewallMatchingRuleResource) Edit(id string, item FirewallMatchingRuleConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallMatchingRuleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallMatchingRule configuration identified by id.
func (r *FirewallMatchingRuleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallMatchingRuleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
