// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallGlobalRulesActiveConfigList holds a list of FirewallGlobalRulesActive configuration.
type FirewallGlobalRulesActiveConfigList struct {
	Items    []FirewallGlobalRulesActiveConfig `json:"items"`
	Kind     string                            `json:"kind"`
	SelfLink string                            `json:"selflink"`
}

// FirewallGlobalRulesActiveConfig holds the configuration of a single FirewallGlobalRulesActive.
type FirewallGlobalRulesActiveConfig struct {
}

// FirewallGlobalRulesActiveEndpoint represents the REST resource for managing FirewallGlobalRulesActive.
const FirewallGlobalRulesActiveEndpoint = "/firewall/global-rules_active"

// FirewallGlobalRulesActiveResource provides an API to manage FirewallGlobalRulesActive configurations.
type FirewallGlobalRulesActiveResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallGlobalRulesActive configurations.
func (r *FirewallGlobalRulesActiveResource) ListAll() (*FirewallGlobalRulesActiveConfigList, error) {
	var list FirewallGlobalRulesActiveConfigList
	if err := r.c.ReadQuery(BasePath+FirewallGlobalRulesActiveEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallGlobalRulesActive configuration identified by id.
func (r *FirewallGlobalRulesActiveResource) Get(id string) (*FirewallGlobalRulesActiveConfig, error) {
	var item FirewallGlobalRulesActiveConfig
	if err := r.c.ReadQuery(BasePath+FirewallGlobalRulesActiveEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallGlobalRulesActive configuration.
func (r *FirewallGlobalRulesActiveResource) Create(item FirewallGlobalRulesActiveConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallGlobalRulesActiveEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallGlobalRulesActive configuration identified by id.
func (r *FirewallGlobalRulesActiveResource) Edit(id string, item FirewallGlobalRulesActiveConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallGlobalRulesActiveEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallGlobalRulesActive configuration identified by id.
func (r *FirewallGlobalRulesActiveResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallGlobalRulesActiveEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
