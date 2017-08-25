// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SoftwareBlockDeviceHotfixConfigList holds a list of SoftwareBlockDeviceHotfix configuration.
type SoftwareBlockDeviceHotfixConfigList struct {
	Items    []SoftwareBlockDeviceHotfixConfig `json:"items"`
	Kind     string                            `json:"kind"`
	SelfLink string                            `json:"selflink"`
}

// SoftwareBlockDeviceHotfixConfig holds the configuration of a single SoftwareBlockDeviceHotfix.
type SoftwareBlockDeviceHotfixConfig struct {
}

// SoftwareBlockDeviceHotfixEndpoint represents the REST resource for managing SoftwareBlockDeviceHotfix.
const SoftwareBlockDeviceHotfixEndpoint = "/software/block-device-hotfix"

// SoftwareBlockDeviceHotfixResource provides an API to manage SoftwareBlockDeviceHotfix configurations.
type SoftwareBlockDeviceHotfixResource struct {
	c *f5.Client
}

// ListAll  lists all the SoftwareBlockDeviceHotfix configurations.
func (r *SoftwareBlockDeviceHotfixResource) ListAll() (*SoftwareBlockDeviceHotfixConfigList, error) {
	var list SoftwareBlockDeviceHotfixConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareBlockDeviceHotfixEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SoftwareBlockDeviceHotfix configuration identified by id.
func (r *SoftwareBlockDeviceHotfixResource) Get(id string) (*SoftwareBlockDeviceHotfixConfig, error) {
	var item SoftwareBlockDeviceHotfixConfig
	if err := r.c.ReadQuery(BasePath+SoftwareBlockDeviceHotfixEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SoftwareBlockDeviceHotfix configuration.
func (r *SoftwareBlockDeviceHotfixResource) Create(item SoftwareBlockDeviceHotfixConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SoftwareBlockDeviceHotfixEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SoftwareBlockDeviceHotfix configuration identified by id.
func (r *SoftwareBlockDeviceHotfixResource) Edit(id string, item SoftwareBlockDeviceHotfixConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SoftwareBlockDeviceHotfixEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SoftwareBlockDeviceHotfix configuration identified by id.
func (r *SoftwareBlockDeviceHotfixResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SoftwareBlockDeviceHotfixEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
