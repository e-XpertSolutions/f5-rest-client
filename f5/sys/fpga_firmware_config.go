// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FPGAFirmwareConfigConfigList holds a list of FPGAFirmwareConfig configuration.
type FPGAFirmwareConfigConfigList struct {
	Items    []FPGAFirmwareConfigConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// FPGAFirmwareConfigConfig holds the configuration of a single FPGAFirmwareConfig.
type FPGAFirmwareConfigConfig struct {
}

// FPGAFirmwareConfigEndpoint represents the REST resource for managing FPGAFirmwareConfig.
const FPGAFirmwareConfigEndpoint = "/fpga/firmware-config"

// FPGAFirmwareConfigResource provides an API to manage FPGAFirmwareConfig configurations.
type FPGAFirmwareConfigResource struct {
	c *f5.Client
}

// ListAll  lists all the FPGAFirmwareConfig configurations.
func (r *FPGAFirmwareConfigResource) ListAll() (*FPGAFirmwareConfigConfigList, error) {
	var list FPGAFirmwareConfigConfigList
	if err := r.c.ReadQuery(BasePath+FPGAFirmwareConfigEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FPGAFirmwareConfig configuration identified by id.
func (r *FPGAFirmwareConfigResource) Get(id string) (*FPGAFirmwareConfigConfig, error) {
	var item FPGAFirmwareConfigConfig
	if err := r.c.ReadQuery(BasePath+FPGAFirmwareConfigEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FPGAFirmwareConfig configuration.
func (r *FPGAFirmwareConfigResource) Create(item FPGAFirmwareConfigConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FPGAFirmwareConfigEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FPGAFirmwareConfig configuration identified by id.
func (r *FPGAFirmwareConfigResource) Edit(id string, item FPGAFirmwareConfigConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FPGAFirmwareConfigEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FPGAFirmwareConfig configuration identified by id.
func (r *FPGAFirmwareConfigResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FPGAFirmwareConfigEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
