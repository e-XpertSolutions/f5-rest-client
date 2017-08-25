// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FeatureModuleConfigList holds a list of FeatureModule configuration.
type FeatureModuleConfigList struct {
	Items    []FeatureModuleConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// FeatureModuleConfig holds the configuration of a single FeatureModule.
type FeatureModuleConfig struct {
	Disabled   bool   `json:"disabled"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	SelfLink   string `json:"selfLink"`
}

// FeatureModuleEndpoint represents the REST resource for managing FeatureModule.
const FeatureModuleEndpoint = "/feature-module"

// FeatureModuleResource provides an API to manage FeatureModule configurations.
type FeatureModuleResource struct {
	c *f5.Client
}

// ListAll  lists all the FeatureModule configurations.
func (r *FeatureModuleResource) ListAll() (*FeatureModuleConfigList, error) {
	var list FeatureModuleConfigList
	if err := r.c.ReadQuery(BasePath+FeatureModuleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FeatureModule configuration identified by id.
func (r *FeatureModuleResource) Get(id string) (*FeatureModuleConfig, error) {
	var item FeatureModuleConfig
	if err := r.c.ReadQuery(BasePath+FeatureModuleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FeatureModule configuration.
func (r *FeatureModuleResource) Create(item FeatureModuleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FeatureModuleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FeatureModule configuration identified by id.
func (r *FeatureModuleResource) Edit(id string, item FeatureModuleConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FeatureModuleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FeatureModule configuration identified by id.
func (r *FeatureModuleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FeatureModuleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
