// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcmp

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// VCMPConfigList holds a list of VCMP configuration.
type VCMPConfigList struct {
	Items    []VCMPConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// VCMPConfig holds the configuration of a single VCMP.
type VCMPConfig struct {
}

// VCMPEndpoint represents the REST resource for managing VCMP.
const VCMPEndpoint = ""

// VCMPResource provides an API to manage VCMP configurations.
type VCMPResource struct {
	c *f5.Client
}

// ListAll  lists all the VCMP configurations.
func (r *VCMPResource) ListAll() (*VCMPConfigList, error) {
	var list VCMPConfigList
	if err := r.c.ReadQuery(BasePath+VCMPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single VCMP configuration identified by id.
func (r *VCMPResource) Get(id string) (*VCMPConfig, error) {
	var item VCMPConfig
	if err := r.c.ReadQuery(BasePath+VCMPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new VCMP configuration.
func (r *VCMPResource) Create(item VCMPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+VCMPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a VCMP configuration identified by id.
func (r *VCMPResource) Edit(id string, item VCMPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+VCMPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single VCMP configuration identified by id.
func (r *VCMPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+VCMPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
