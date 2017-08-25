// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DaemonLogSettingsTMMConfigList holds a list of DaemonLogSettingsTMM configuration.
type DaemonLogSettingsTMMConfigList struct {
	Items    []DaemonLogSettingsTMMConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// DaemonLogSettingsTMMConfig holds the configuration of a single DaemonLogSettingsTMM.
type DaemonLogSettingsTMMConfig struct {
}

// DaemonLogSettingsTMMEndpoint represents the REST resource for managing DaemonLogSettingsTMM.
const DaemonLogSettingsTMMEndpoint = "/daemon-log-settings/tmm"

// DaemonLogSettingsTMMResource provides an API to manage DaemonLogSettingsTMM configurations.
type DaemonLogSettingsTMMResource struct {
	c *f5.Client
}

// ListAll  lists all the DaemonLogSettingsTMM configurations.
func (r *DaemonLogSettingsTMMResource) ListAll() (*DaemonLogSettingsTMMConfigList, error) {
	var list DaemonLogSettingsTMMConfigList
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsTMMEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DaemonLogSettingsTMM configuration identified by id.
func (r *DaemonLogSettingsTMMResource) Get(id string) (*DaemonLogSettingsTMMConfig, error) {
	var item DaemonLogSettingsTMMConfig
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsTMMEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DaemonLogSettingsTMM configuration.
func (r *DaemonLogSettingsTMMResource) Create(item DaemonLogSettingsTMMConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DaemonLogSettingsTMMEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DaemonLogSettingsTMM configuration identified by id.
func (r *DaemonLogSettingsTMMResource) Edit(id string, item DaemonLogSettingsTMMConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DaemonLogSettingsTMMEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DaemonLogSettingsTMM configuration identified by id.
func (r *DaemonLogSettingsTMMResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DaemonLogSettingsTMMEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
