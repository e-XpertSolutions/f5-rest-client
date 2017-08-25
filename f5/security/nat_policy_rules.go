// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// NATPolicyRulesConfigList holds a list of NATPolicyRules configuration.
type NATPolicyRulesConfigList struct {
	Items    []NATPolicyRulesConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// NATPolicyRulesConfig holds the configuration of a single NATPolicyRules.
type NATPolicyRulesConfig struct {
}

// NATPolicyRulesEndpoint represents the REST resource for managing NATPolicyRules.
const NATPolicyRulesEndpoint = "/nat/policy_rules"

// NATPolicyRulesResource provides an API to manage NATPolicyRules configurations.
type NATPolicyRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the NATPolicyRules configurations.
func (r *NATPolicyRulesResource) ListAll() (*NATPolicyRulesConfigList, error) {
	var list NATPolicyRulesConfigList
	if err := r.c.ReadQuery(BasePath+NATPolicyRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single NATPolicyRules configuration identified by id.
func (r *NATPolicyRulesResource) Get(id string) (*NATPolicyRulesConfig, error) {
	var item NATPolicyRulesConfig
	if err := r.c.ReadQuery(BasePath+NATPolicyRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new NATPolicyRules configuration.
func (r *NATPolicyRulesResource) Create(item NATPolicyRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+NATPolicyRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a NATPolicyRules configuration identified by id.
func (r *NATPolicyRulesResource) Edit(id string, item NATPolicyRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+NATPolicyRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single NATPolicyRules configuration identified by id.
func (r *NATPolicyRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+NATPolicyRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
