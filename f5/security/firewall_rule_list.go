// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallRuleListConfigList holds a list of FirewallRuleList configuration.
type FirewallRuleListConfigList struct {
	Items    []FirewallRuleListConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// FirewallRuleListConfig holds the configuration of a single FirewallRuleList.
type FirewallRuleListConfig struct {
	FullPath       string `json:"fullPath"`
	Generation     int    `json:"generation"`
	Kind           string `json:"kind"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	RulesReference struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"rulesReference"`
	SelfLink string `json:"selfLink"`
}

// FirewallRuleListEndpoint represents the REST resource for managing FirewallRuleList.
const FirewallRuleListEndpoint = "/firewall/rule-list"

// FirewallRuleListResource provides an API to manage FirewallRuleList configurations.
type FirewallRuleListResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallRuleList configurations.
func (r *FirewallRuleListResource) ListAll() (*FirewallRuleListConfigList, error) {
	var list FirewallRuleListConfigList
	if err := r.c.ReadQuery(BasePath+FirewallRuleListEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallRuleList configuration identified by id.
func (r *FirewallRuleListResource) Get(id string) (*FirewallRuleListConfig, error) {
	var item FirewallRuleListConfig
	if err := r.c.ReadQuery(BasePath+FirewallRuleListEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallRuleList configuration.
func (r *FirewallRuleListResource) Create(item FirewallRuleListConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallRuleListEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallRuleList configuration identified by id.
func (r *FirewallRuleListResource) Edit(id string, item FirewallRuleListConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallRuleListEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallRuleList configuration identified by id.
func (r *FirewallRuleListResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallRuleListEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
