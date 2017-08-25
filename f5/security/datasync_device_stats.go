// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DatasyncDeviceStatsConfigList holds a list of DatasyncDeviceStats configuration.
type DatasyncDeviceStatsConfigList struct {
	Items    []DatasyncDeviceStatsConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// DatasyncDeviceStatsConfig holds the configuration of a single DatasyncDeviceStats.
type DatasyncDeviceStatsConfig struct {
}

// DatasyncDeviceStatsEndpoint represents the REST resource for managing DatasyncDeviceStats.
const DatasyncDeviceStatsEndpoint = "/datasync/device-stats"

// DatasyncDeviceStatsResource provides an API to manage DatasyncDeviceStats configurations.
type DatasyncDeviceStatsResource struct {
	c *f5.Client
}

// ListAll  lists all the DatasyncDeviceStats configurations.
func (r *DatasyncDeviceStatsResource) ListAll() (*DatasyncDeviceStatsConfigList, error) {
	var list DatasyncDeviceStatsConfigList
	if err := r.c.ReadQuery(BasePath+DatasyncDeviceStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DatasyncDeviceStats configuration identified by id.
func (r *DatasyncDeviceStatsResource) Get(id string) (*DatasyncDeviceStatsConfig, error) {
	var item DatasyncDeviceStatsConfig
	if err := r.c.ReadQuery(BasePath+DatasyncDeviceStatsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DatasyncDeviceStats configuration.
func (r *DatasyncDeviceStatsResource) Create(item DatasyncDeviceStatsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DatasyncDeviceStatsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DatasyncDeviceStats configuration identified by id.
func (r *DatasyncDeviceStatsResource) Edit(id string, item DatasyncDeviceStatsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DatasyncDeviceStatsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DatasyncDeviceStats configuration identified by id.
func (r *DatasyncDeviceStatsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DatasyncDeviceStatsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
