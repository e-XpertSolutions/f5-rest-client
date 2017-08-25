// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SettingsConfigList holds a list of Settings configuration.
type SettingsConfigList struct {
	Items    []SettingsConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// SettingsConfig holds the configuration of a single Settings.
type SettingsConfig struct {
}

// SettingsEndpoint represents the REST resource for managing Settings.
const SettingsEndpoint = "/analytics/settings"

// SettingsResource provides an API to manage Settings configurations.
type SettingsResource struct {
	c *f5.Client
}

// ListAll  lists all the Settings configurations.
func (r *SettingsResource) ListAll() (*SettingsConfigList, error) {
	var list SettingsConfigList
	if err := r.c.ReadQuery(BasePath+SettingsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Settings configuration identified by id.
func (r *SettingsResource) Get(id string) (*SettingsConfig, error) {
	var item SettingsConfig
	if err := r.c.ReadQuery(BasePath+SettingsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Settings configuration.
func (r *SettingsResource) Create(item SettingsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SettingsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Settings configuration identified by id.
func (r *SettingsResource) Edit(id string, item SettingsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SettingsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Settings configuration identified by id.
func (r *SettingsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SettingsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
