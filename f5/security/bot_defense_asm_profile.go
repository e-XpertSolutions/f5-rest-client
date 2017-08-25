// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// BotDefenseASMProfileConfigList holds a list of BotDefenseASMProfile configuration.
type BotDefenseASMProfileConfigList struct {
	Items    []BotDefenseASMProfileConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// BotDefenseASMProfileConfig holds the configuration of a single BotDefenseASMProfile.
type BotDefenseASMProfileConfig struct {
}

// BotDefenseASMProfileEndpoint represents the REST resource for managing BotDefenseASMProfile.
const BotDefenseASMProfileEndpoint = "/bot-defense/asm-profile"

// BotDefenseASMProfileResource provides an API to manage BotDefenseASMProfile configurations.
type BotDefenseASMProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the BotDefenseASMProfile configurations.
func (r *BotDefenseASMProfileResource) ListAll() (*BotDefenseASMProfileConfigList, error) {
	var list BotDefenseASMProfileConfigList
	if err := r.c.ReadQuery(BasePath+BotDefenseASMProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single BotDefenseASMProfile configuration identified by id.
func (r *BotDefenseASMProfileResource) Get(id string) (*BotDefenseASMProfileConfig, error) {
	var item BotDefenseASMProfileConfig
	if err := r.c.ReadQuery(BasePath+BotDefenseASMProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new BotDefenseASMProfile configuration.
func (r *BotDefenseASMProfileResource) Create(item BotDefenseASMProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+BotDefenseASMProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a BotDefenseASMProfile configuration identified by id.
func (r *BotDefenseASMProfileResource) Edit(id string, item BotDefenseASMProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+BotDefenseASMProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single BotDefenseASMProfile configuration identified by id.
func (r *BotDefenseASMProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+BotDefenseASMProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
