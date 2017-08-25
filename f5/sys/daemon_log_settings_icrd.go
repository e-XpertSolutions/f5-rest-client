// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DaemonLogSettingsICRDConfigList holds a list of DaemonLogSettingsICRD configuration.
type DaemonLogSettingsICRDConfigList struct {
	Items    []DaemonLogSettingsICRDConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// DaemonLogSettingsICRDConfig holds the configuration of a single DaemonLogSettingsICRD.
type DaemonLogSettingsICRDConfig struct {
}

// DaemonLogSettingsICRDEndpoint represents the REST resource for managing DaemonLogSettingsICRD.
const DaemonLogSettingsICRDEndpoint = "/daemon-log-settings/icrd"

// DaemonLogSettingsICRDResource provides an API to manage DaemonLogSettingsICRD configurations.
type DaemonLogSettingsICRDResource struct {
	c *f5.Client
}

// ListAll  lists all the DaemonLogSettingsICRD configurations.
func (r *DaemonLogSettingsICRDResource) ListAll() (*DaemonLogSettingsICRDConfigList, error) {
	var list DaemonLogSettingsICRDConfigList
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsICRDEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DaemonLogSettingsICRD configuration identified by id.
func (r *DaemonLogSettingsICRDResource) Get(id string) (*DaemonLogSettingsICRDConfig, error) {
	var item DaemonLogSettingsICRDConfig
	if err := r.c.ReadQuery(BasePath+DaemonLogSettingsICRDEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DaemonLogSettingsICRD configuration.
func (r *DaemonLogSettingsICRDResource) Create(item DaemonLogSettingsICRDConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DaemonLogSettingsICRDEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DaemonLogSettingsICRD configuration identified by id.
func (r *DaemonLogSettingsICRDResource) Edit(id string, item DaemonLogSettingsICRDConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DaemonLogSettingsICRDEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DaemonLogSettingsICRD configuration identified by id.
func (r *DaemonLogSettingsICRDResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DaemonLogSettingsICRDEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
