// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FPGAConfigList holds a list of FPGA configuration.
type FPGAConfigList struct {
	Items    []FPGAConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// FPGAConfig holds the configuration of a single FPGA.
type FPGAConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// FPGAEndpoint represents the REST resource for managing FPGA.
const FPGAEndpoint = "/fpga"

// FPGAResource provides an API to manage FPGA configurations.
type FPGAResource struct {
	c *f5.Client
}

// ListAll  lists all the FPGA configurations.
func (r *FPGAResource) ListAll() (*FPGAConfigList, error) {
	var list FPGAConfigList
	if err := r.c.ReadQuery(BasePath+FPGAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FPGA configuration identified by id.
func (r *FPGAResource) Get(id string) (*FPGAConfig, error) {
	var item FPGAConfig
	if err := r.c.ReadQuery(BasePath+FPGAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FPGA configuration.
func (r *FPGAResource) Create(item FPGAConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FPGAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FPGA configuration identified by id.
func (r *FPGAResource) Edit(id string, item FPGAConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FPGAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FPGA configuration identified by id.
func (r *FPGAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FPGAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
