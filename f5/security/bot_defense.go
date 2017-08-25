// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// BotDefenseConfigList holds a list of BotDefense configuration.
type BotDefenseConfigList struct {
	Items    []BotDefenseConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// BotDefenseConfig holds the configuration of a single BotDefense.
type BotDefenseConfig struct {
}

// BotDefenseEndpoint represents the REST resource for managing BotDefense.
const BotDefenseEndpoint = "/bot-defense"

// BotDefenseResource provides an API to manage BotDefense configurations.
type BotDefenseResource struct {
	c *f5.Client
}

// ListAll  lists all the BotDefense configurations.
func (r *BotDefenseResource) ListAll() (*BotDefenseConfigList, error) {
	var list BotDefenseConfigList
	if err := r.c.ReadQuery(BasePath+BotDefenseEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single BotDefense configuration identified by id.
func (r *BotDefenseResource) Get(id string) (*BotDefenseConfig, error) {
	var item BotDefenseConfig
	if err := r.c.ReadQuery(BasePath+BotDefenseEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new BotDefense configuration.
func (r *BotDefenseResource) Create(item BotDefenseConfig) error {
	if err := r.c.ModQuery("POST", BasePath+BotDefenseEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a BotDefense configuration identified by id.
func (r *BotDefenseResource) Edit(id string, item BotDefenseConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+BotDefenseEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single BotDefense configuration identified by id.
func (r *BotDefenseResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+BotDefenseEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
