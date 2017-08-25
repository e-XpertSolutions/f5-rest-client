// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallOnDemandRuleDeployConfigList holds a list of FirewallOnDemandRuleDeploy configuration.
type FirewallOnDemandRuleDeployConfigList struct {
	Items    []FirewallOnDemandRuleDeployConfig `json:"items"`
	Kind     string                             `json:"kind"`
	SelfLink string                             `json:"selflink"`
}

// FirewallOnDemandRuleDeployConfig holds the configuration of a single FirewallOnDemandRuleDeploy.
type FirewallOnDemandRuleDeployConfig struct {
}

// FirewallOnDemandRuleDeployEndpoint represents the REST resource for managing FirewallOnDemandRuleDeploy.
const FirewallOnDemandRuleDeployEndpoint = "/firewall/on-demand-rule-deploy"

// FirewallOnDemandRuleDeployResource provides an API to manage FirewallOnDemandRuleDeploy configurations.
type FirewallOnDemandRuleDeployResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallOnDemandRuleDeploy configurations.
func (r *FirewallOnDemandRuleDeployResource) ListAll() (*FirewallOnDemandRuleDeployConfigList, error) {
	var list FirewallOnDemandRuleDeployConfigList
	if err := r.c.ReadQuery(BasePath+FirewallOnDemandRuleDeployEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallOnDemandRuleDeploy configuration identified by id.
func (r *FirewallOnDemandRuleDeployResource) Get(id string) (*FirewallOnDemandRuleDeployConfig, error) {
	var item FirewallOnDemandRuleDeployConfig
	if err := r.c.ReadQuery(BasePath+FirewallOnDemandRuleDeployEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallOnDemandRuleDeploy configuration.
func (r *FirewallOnDemandRuleDeployResource) Create(item FirewallOnDemandRuleDeployConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallOnDemandRuleDeployEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallOnDemandRuleDeploy configuration identified by id.
func (r *FirewallOnDemandRuleDeployResource) Edit(id string, item FirewallOnDemandRuleDeployConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallOnDemandRuleDeployEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallOnDemandRuleDeploy configuration identified by id.
func (r *FirewallOnDemandRuleDeployResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallOnDemandRuleDeployEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
