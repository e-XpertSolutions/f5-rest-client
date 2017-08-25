// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SoftwareBlockDeviceImageConfigList holds a list of SoftwareBlockDeviceImage configuration.
type SoftwareBlockDeviceImageConfigList struct {
	Items    []SoftwareBlockDeviceImageConfig `json:"items"`
	Kind     string                           `json:"kind"`
	SelfLink string                           `json:"selflink"`
}

// SoftwareBlockDeviceImageConfig holds the configuration of a single SoftwareBlockDeviceImage.
type SoftwareBlockDeviceImageConfig struct {
}

// SoftwareBlockDeviceImageEndpoint represents the REST resource for managing SoftwareBlockDeviceImage.
const SoftwareBlockDeviceImageEndpoint = "/software/block-device-image"

// SoftwareBlockDeviceImageResource provides an API to manage SoftwareBlockDeviceImage configurations.
type SoftwareBlockDeviceImageResource struct {
	c *f5.Client
}

// ListAll  lists all the SoftwareBlockDeviceImage configurations.
func (r *SoftwareBlockDeviceImageResource) ListAll() (*SoftwareBlockDeviceImageConfigList, error) {
	var list SoftwareBlockDeviceImageConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareBlockDeviceImageEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SoftwareBlockDeviceImage configuration identified by id.
func (r *SoftwareBlockDeviceImageResource) Get(id string) (*SoftwareBlockDeviceImageConfig, error) {
	var item SoftwareBlockDeviceImageConfig
	if err := r.c.ReadQuery(BasePath+SoftwareBlockDeviceImageEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SoftwareBlockDeviceImage configuration.
func (r *SoftwareBlockDeviceImageResource) Create(item SoftwareBlockDeviceImageConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SoftwareBlockDeviceImageEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SoftwareBlockDeviceImage configuration identified by id.
func (r *SoftwareBlockDeviceImageResource) Edit(id string, item SoftwareBlockDeviceImageConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SoftwareBlockDeviceImageEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SoftwareBlockDeviceImage configuration identified by id.
func (r *SoftwareBlockDeviceImageResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SoftwareBlockDeviceImageEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
