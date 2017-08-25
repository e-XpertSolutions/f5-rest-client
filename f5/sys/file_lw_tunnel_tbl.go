// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileLWTunnelTBLConfigList holds a list of FileLWTunnelTBL configuration.
type FileLWTunnelTBLConfigList struct {
	Items    []FileLWTunnelTBLConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// FileLWTunnelTBLConfig holds the configuration of a single FileLWTunnelTBL.
type FileLWTunnelTBLConfig struct {
}

// FileLWTunnelTBLEndpoint represents the REST resource for managing FileLWTunnelTBL.
const FileLWTunnelTBLEndpoint = "/file/lwtunneltbl"

// FileLWTunnelTBLResource provides an API to manage FileLWTunnelTBL configurations.
type FileLWTunnelTBLResource struct {
	c *f5.Client
}

// ListAll  lists all the FileLWTunnelTBL configurations.
func (r *FileLWTunnelTBLResource) ListAll() (*FileLWTunnelTBLConfigList, error) {
	var list FileLWTunnelTBLConfigList
	if err := r.c.ReadQuery(BasePath+FileLWTunnelTBLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileLWTunnelTBL configuration identified by id.
func (r *FileLWTunnelTBLResource) Get(id string) (*FileLWTunnelTBLConfig, error) {
	var item FileLWTunnelTBLConfig
	if err := r.c.ReadQuery(BasePath+FileLWTunnelTBLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileLWTunnelTBL configuration.
func (r *FileLWTunnelTBLResource) Create(item FileLWTunnelTBLConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileLWTunnelTBLEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileLWTunnelTBL configuration identified by id.
func (r *FileLWTunnelTBLResource) Edit(id string, item FileLWTunnelTBLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileLWTunnelTBLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileLWTunnelTBL configuration identified by id.
func (r *FileLWTunnelTBLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileLWTunnelTBLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
