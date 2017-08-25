// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SSHProfileRulesConfigList holds a list of SSHProfileRules configuration.
type SSHProfileRulesConfigList struct {
	Items    []SSHProfileRulesConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// SSHProfileRulesConfig holds the configuration of a single SSHProfileRules.
type SSHProfileRulesConfig struct {
}

// SSHProfileRulesEndpoint represents the REST resource for managing SSHProfileRules.
const SSHProfileRulesEndpoint = "/ssh/profile_rules"

// SSHProfileRulesResource provides an API to manage SSHProfileRules configurations.
type SSHProfileRulesResource struct {
	c *f5.Client
}

// ListAll  lists all the SSHProfileRules configurations.
func (r *SSHProfileRulesResource) ListAll() (*SSHProfileRulesConfigList, error) {
	var list SSHProfileRulesConfigList
	if err := r.c.ReadQuery(BasePath+SSHProfileRulesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SSHProfileRules configuration identified by id.
func (r *SSHProfileRulesResource) Get(id string) (*SSHProfileRulesConfig, error) {
	var item SSHProfileRulesConfig
	if err := r.c.ReadQuery(BasePath+SSHProfileRulesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SSHProfileRules configuration.
func (r *SSHProfileRulesResource) Create(item SSHProfileRulesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SSHProfileRulesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SSHProfileRules configuration identified by id.
func (r *SSHProfileRulesResource) Edit(id string, item SSHProfileRulesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SSHProfileRulesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SSHProfileRules configuration identified by id.
func (r *SSHProfileRulesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SSHProfileRulesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
