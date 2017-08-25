// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DaemonLogSettingsClusterdConfigList holds a list of DaemonLogSettingsClusterd configuration.
type DaemonLogSettingsClusterdConfigList struct {
	Items    []DaemonLogSettingsClusterdConfig `json:"items"`
	Kind     string                            `json:"kind"`
	SelfLink string                            `json:"selflink"`
}

// DaemonLogSettingsClusterdConfig holds the configuration of a single DaemonLogSettingsClusterd.
type DaemonLogSettingsClusterdConfig struct {
}

// DaemonLogSettingsClusterdEndpoint represents the REST resource for managing DaemonLogSettingsClusterd.
const DaemonLogSettingsClusterdEndpoint = "/daemon-log-settings/clusterd"

// DaemonLogSettingsClusterdResource provides an API to manage DaemonLogSettingsClusterd configurations.
type DaemonLogSettingsClusterdResource struct {
	c *f5.Client
}

// ListAll  lists all the DaemonLogSettingsClusterd configurations.
func (r *DaemonLogSettingsClusterdResource) ListAll() (*DaemonLogSettingsClusterdConfigList, error) {
	var list DaemonLogSettingsClusterdConfigList
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsClusterdEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DaemonLogSettingsClusterd configuration identified by id.
func (r *DaemonLogSettingsClusterdResource) Get(id string) (*DaemonLogSettingsClusterdConfig, error) {
	var item DaemonLogSettingsClusterdConfig
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsClusterdEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DaemonLogSettingsClusterd configuration.
func (r *DaemonLogSettingsClusterdResource) Create(item DaemonLogSettingsClusterdConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DaemonLogSettingsClusterdEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DaemonLogSettingsClusterd configuration identified by id.
func (r *DaemonLogSettingsClusterdResource) Edit(id string, item DaemonLogSettingsClusterdConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DaemonLogSettingsClusterdEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DaemonLogSettingsClusterd configuration identified by id.
func (r *DaemonLogSettingsClusterdResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DaemonLogSettingsClusterdEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
