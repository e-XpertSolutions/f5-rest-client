// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// RuleList holds a list of Rule configuration.
type RuleList struct {
	Items    []Rule `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// Rule holds the configuration of a single Rule.
type Rule struct {
	APIAnonymous string `json:"apiAnonymous,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
}

// RuleEndpoint represents the REST resource for managing Rule.
const RuleEndpoint = "/rule"

// RuleResource provides an API to manage Rule configurations.
type RuleResource struct {
	c *f5.Client
}

// ListAll  lists all the Rule configurations.
func (r *RuleResource) ListAll() (*RuleList, error) {
	var list RuleList
	if err := r.c.ReadQuery(BasePath+RuleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Rule configuration identified by id.
func (r *RuleResource) Get(id string) (*Rule, error) {
	var item Rule
	if err := r.c.ReadQuery(BasePath+RuleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Rule configuration.
func (r *RuleResource) Create(item Rule) error {
	if err := r.c.ModQuery("POST", BasePath+RuleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Rule configuration identified by id.
func (r *RuleResource) Edit(id string, item Rule) error {
	if err := r.c.ModQuery("PUT", BasePath+RuleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Rule configuration identified by id.
func (r *RuleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+RuleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
