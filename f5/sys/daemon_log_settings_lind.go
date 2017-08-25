// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DaemonLogSettingsLindConfigList holds a list of DaemonLogSettingsLind configuration.
type DaemonLogSettingsLindConfigList struct {
	Items    []DaemonLogSettingsLindConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// DaemonLogSettingsLindConfig holds the configuration of a single DaemonLogSettingsLind.
type DaemonLogSettingsLindConfig struct {
}

// DaemonLogSettingsLindEndpoint represents the REST resource for managing DaemonLogSettingsLind.
const DaemonLogSettingsLindEndpoint = "/daemon-log-settings/lind"

// DaemonLogSettingsLindResource provides an API to manage DaemonLogSettingsLind configurations.
type DaemonLogSettingsLindResource struct {
	c *f5.Client
}

// ListAll  lists all the DaemonLogSettingsLind configurations.
func (r *DaemonLogSettingsLindResource) ListAll() (*DaemonLogSettingsLindConfigList, error) {
	var list DaemonLogSettingsLindConfigList
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsLindEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DaemonLogSettingsLind configuration identified by id.
func (r *DaemonLogSettingsLindResource) Get(id string) (*DaemonLogSettingsLindConfig, error) {
	var item DaemonLogSettingsLindConfig
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsLindEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DaemonLogSettingsLind configuration.
func (r *DaemonLogSettingsLindResource) Create(item DaemonLogSettingsLindConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DaemonLogSettingsLindEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DaemonLogSettingsLind configuration identified by id.
func (r *DaemonLogSettingsLindResource) Edit(id string, item DaemonLogSettingsLindConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DaemonLogSettingsLindEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DaemonLogSettingsLind configuration identified by id.
func (r *DaemonLogSettingsLindResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DaemonLogSettingsLindEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
