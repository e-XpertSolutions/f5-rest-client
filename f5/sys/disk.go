// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DiskConfigList holds a list of Disk configuration.
type DiskConfigList struct {
	Items    []DiskConfig `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// DiskConfig holds the configuration of a single Disk.
type DiskConfig struct {
	Reference struct {
		Link string `json:"link,omitempty"`
	} `json:"reference,omitempty"`
}

// DiskEndpoint represents the REST resource for managing Disk.
const DiskEndpoint = "/disk"

// DiskResource provides an API to manage Disk configurations.
type DiskResource struct {
	c *f5.Client
}

// ListAll  lists all the Disk configurations.
func (r *DiskResource) ListAll() (*DiskConfigList, error) {
	var list DiskConfigList
	if err := r.c.ReadQuery(BasePath+DiskEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Disk configuration identified by id.
func (r *DiskResource) Get(id string) (*DiskConfig, error) {
	var item DiskConfig
	if err := r.c.ReadQuery(BasePath+DiskEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Disk configuration.
func (r *DiskResource) Create(item DiskConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DiskEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Disk configuration identified by id.
func (r *DiskResource) Edit(id string, item DiskConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DiskEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Disk configuration identified by id.
func (r *DiskResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DiskEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
