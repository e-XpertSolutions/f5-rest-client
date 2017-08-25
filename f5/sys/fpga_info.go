// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FPGAInfoConfigList holds a list of FPGAInfo configuration.
type FPGAInfoConfigList struct {
	Items    []FPGAInfoConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// FPGAInfoConfig holds the configuration of a single FPGAInfo.
type FPGAInfoConfig struct {
}

// FPGAInfoEndpoint represents the REST resource for managing FPGAInfo.
const FPGAInfoEndpoint = "/fpga/info"

// FPGAInfoResource provides an API to manage FPGAInfo configurations.
type FPGAInfoResource struct {
	c *f5.Client
}

// ListAll  lists all the FPGAInfo configurations.
func (r *FPGAInfoResource) ListAll() (*FPGAInfoConfigList, error) {
	var list FPGAInfoConfigList
	if err := r.c.ReadQuery(BasePath+FPGAInfoEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FPGAInfo configuration identified by id.
func (r *FPGAInfoResource) Get(id string) (*FPGAInfoConfig, error) {
	var item FPGAInfoConfig
	if err := r.c.ReadQuery(BasePath+FPGAInfoEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FPGAInfo configuration.
func (r *FPGAInfoResource) Create(item FPGAInfoConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FPGAInfoEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FPGAInfo configuration identified by id.
func (r *FPGAInfoResource) Edit(id string, item FPGAInfoConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FPGAInfoEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FPGAInfo configuration identified by id.
func (r *FPGAInfoResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FPGAInfoEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
