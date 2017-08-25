// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DaemonLogSettingsMCPDConfigList holds a list of DaemonLogSettingsMCPD configuration.
type DaemonLogSettingsMCPDConfigList struct {
	Items    []DaemonLogSettingsMCPDConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// DaemonLogSettingsMCPDConfig holds the configuration of a single DaemonLogSettingsMCPD.
type DaemonLogSettingsMCPDConfig struct {
}

// DaemonLogSettingsMCPDEndpoint represents the REST resource for managing DaemonLogSettingsMCPD.
const DaemonLogSettingsMCPDEndpoint = "/daemon-log-settings/mcpd"

// DaemonLogSettingsMCPDResource provides an API to manage DaemonLogSettingsMCPD configurations.
type DaemonLogSettingsMCPDResource struct {
	c *f5.Client
}

// ListAll  lists all the DaemonLogSettingsMCPD configurations.
func (r *DaemonLogSettingsMCPDResource) ListAll() (*DaemonLogSettingsMCPDConfigList, error) {
	var list DaemonLogSettingsMCPDConfigList
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsMCPDEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DaemonLogSettingsMCPD configuration identified by id.
func (r *DaemonLogSettingsMCPDResource) Get(id string) (*DaemonLogSettingsMCPDConfig, error) {
	var item DaemonLogSettingsMCPDConfig
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsMCPDEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DaemonLogSettingsMCPD configuration.
func (r *DaemonLogSettingsMCPDResource) Create(item DaemonLogSettingsMCPDConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DaemonLogSettingsMCPDEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DaemonLogSettingsMCPD configuration identified by id.
func (r *DaemonLogSettingsMCPDResource) Edit(id string, item DaemonLogSettingsMCPDConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DaemonLogSettingsMCPDEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DaemonLogSettingsMCPD configuration identified by id.
func (r *DaemonLogSettingsMCPDResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DaemonLogSettingsMCPDEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
