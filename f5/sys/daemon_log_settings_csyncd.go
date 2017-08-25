// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DaemonLogSettingsCsyncdConfigList holds a list of DaemonLogSettingsCsyncd configuration.
type DaemonLogSettingsCsyncdConfigList struct {
	Items    []DaemonLogSettingsCsyncdConfig `json:"items"`
	Kind     string                          `json:"kind"`
	SelfLink string                          `json:"selflink"`
}

// DaemonLogSettingsCsyncdConfig holds the configuration of a single DaemonLogSettingsCsyncd.
type DaemonLogSettingsCsyncdConfig struct {
}

// DaemonLogSettingsCsyncdEndpoint represents the REST resource for managing DaemonLogSettingsCsyncd.
const DaemonLogSettingsCsyncdEndpoint = "/daemon-log-settings/csyncd"

// DaemonLogSettingsCsyncdResource provides an API to manage DaemonLogSettingsCsyncd configurations.
type DaemonLogSettingsCsyncdResource struct {
	c *f5.Client
}

// ListAll  lists all the DaemonLogSettingsCsyncd configurations.
func (r *DaemonLogSettingsCsyncdResource) ListAll() (*DaemonLogSettingsCsyncdConfigList, error) {
	var list DaemonLogSettingsCsyncdConfigList
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsCsyncdEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DaemonLogSettingsCsyncd configuration identified by id.
func (r *DaemonLogSettingsCsyncdResource) Get(id string) (*DaemonLogSettingsCsyncdConfig, error) {
	var item DaemonLogSettingsCsyncdConfig
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsCsyncdEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DaemonLogSettingsCsyncd configuration.
func (r *DaemonLogSettingsCsyncdResource) Create(item DaemonLogSettingsCsyncdConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DaemonLogSettingsCsyncdEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DaemonLogSettingsCsyncd configuration identified by id.
func (r *DaemonLogSettingsCsyncdResource) Edit(id string, item DaemonLogSettingsCsyncdConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DaemonLogSettingsCsyncdEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DaemonLogSettingsCsyncd configuration identified by id.
func (r *DaemonLogSettingsCsyncdResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DaemonLogSettingsCsyncdEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
