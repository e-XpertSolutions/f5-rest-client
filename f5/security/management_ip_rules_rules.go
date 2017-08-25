// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ManagementIPRulesRulesConfigList holds a list of ManagementIPRulesRules configuration.
type ManagementIPRulesRulesConfigList struct {
	Items    []ManagementIPRulesRulesConfig `json:"items"`
	Kind     string                         `json:"kind"`
	SelfLink string                         `json:"selflink"`
}

// ManagementIPRulesRulesConfig holds the configuration of a single ManagementIPRulesRules.
type ManagementIPRulesRulesConfig struct {
}

// ManagementIPRulesRulesEndpoint represents the REST resource for managing ManagementIPRulesRules.
const ManagementIPRulesRulesEndpoint = "/firewall/management-ip-rules_rules"

// ManagementIPRulesRulesResource provides an API to manage ManagementIPRulesRules configurations.
type ManagementIPRulesRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the ManagementIPRulesRules configurations.
func (r *ManagementIPRulesRulesResource) ListAll() (*ManagementIPRulesRulesConfigList, error) {
	var list ManagementIPRulesRulesConfigList
	if err := r.c.ReadQuery(BasePath+ManagementIPRulesRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ManagementIPRulesRules configuration identified by id.
func (r *ManagementIPRulesRulesResource) Get(id string) (*ManagementIPRulesRulesConfig, error) {
	var item ManagementIPRulesRulesConfig
	if err := r.c.ReadQuery(BasePath+ManagementIPRulesRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ManagementIPRulesRules configuration.
func (r *ManagementIPRulesRulesResource) Create(item ManagementIPRulesRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ManagementIPRulesRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ManagementIPRulesRules configuration identified by id.
func (r *ManagementIPRulesRulesResource) Edit(id string, item ManagementIPRulesRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ManagementIPRulesRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ManagementIPRulesRules configuration identified by id.
func (r *ManagementIPRulesRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ManagementIPRulesRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
