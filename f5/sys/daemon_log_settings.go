// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DaemonLogSettingsConfigList holds a list of DaemonLogSettings configuration.
type DaemonLogSettingsConfigList struct {
	Items    []DaemonLogSettingsConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

// DaemonLogSettingsConfig holds the configuration of a single DaemonLogSettings.
type DaemonLogSettingsConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// DaemonLogSettingsEndpoint represents the REST resource for managing DaemonLogSettings.
const DaemonLogSettingsEndpoint = "/daemon-log-settings"

// DaemonLogSettingsResource provides an API to manage DaemonLogSettings configurations.
type DaemonLogSettingsResource struct {
	c *f5.Client
}

// ListAll  lists all the DaemonLogSettings configurations.
func (r *DaemonLogSettingsResource) ListAll() (*DaemonLogSettingsConfigList, error) {
	var list DaemonLogSettingsConfigList
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DaemonLogSettings configuration identified by id.
func (r *DaemonLogSettingsResource) Get(id string) (*DaemonLogSettingsConfig, error) {
	var item DaemonLogSettingsConfig
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DaemonLogSettings configuration.
func (r *DaemonLogSettingsResource) Create(item DaemonLogSettingsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DaemonLogSettingsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DaemonLogSettings configuration identified by id.
func (r *DaemonLogSettingsResource) Edit(id string, item DaemonLogSettingsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DaemonLogSettingsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DaemonLogSettings configuration identified by id.
func (r *DaemonLogSettingsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DaemonLogSettingsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
