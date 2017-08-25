// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SysConfigList holds a list of Sys configuration.
type SysConfigList struct {
	Items    []SysConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// SysConfig holds the configuration of a single Sys.
type SysConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// SysEndpoint represents the REST resource for managing Sys.
const SysEndpoint = "/tm/sys"

// SysResource provides an API to manage Sys configurations.
type SysResource struct {
	c *f5.Client
}

// ListAll  lists all the Sys configurations.
func (r *SysResource) ListAll() (*SysConfigList, error) {
	var list SysConfigList
	if err := r.c.ReadQuery(BasePath+SysEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Sys configuration identified by id.
func (r *SysResource) Get(id string) (*SysConfig, error) {
	var item SysConfig
	if err := r.c.ReadQuery(BasePath+SysEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Sys configuration.
func (r *SysResource) Create(item SysConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SysEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Sys configuration identified by id.
func (r *SysResource) Edit(id string, item SysConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SysEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Sys configuration identified by id.
func (r *SysResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SysEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
