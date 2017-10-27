// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DiskLogicalDiskConfigList holds a list of DiskLogicalDisk configuration.
type DiskLogicalDiskConfigList struct {
	Items    []DiskLogicalDiskConfig `json:"items,omitempty"`
	Kind     string                  `json:"kind,omitempty"`
	SelfLink string                  `json:"selflink,omitempty"`
}

// DiskLogicalDiskConfig holds the configuration of a single DiskLogicalDisk.
type DiskLogicalDiskConfig struct {
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Mode       string `json:"mode,omitempty"`
	Name       string `json:"name,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	Size       int    `json:"size,omitempty"`
	VgFree     int    `json:"vgFree,omitempty"`
	VgInUse    int    `json:"vgInUse,omitempty"`
	VgReserved int    `json:"vgReserved,omitempty"`
}

// DiskLogicalDiskEndpoint represents the REST resource for managing DiskLogicalDisk.
const DiskLogicalDiskEndpoint = "/disk/logical-disk"

// DiskLogicalDiskResource provides an API to manage DiskLogicalDisk configurations.
type DiskLogicalDiskResource struct {
	c *f5.Client
}

// ListAll  lists all the DiskLogicalDisk configurations.
func (r *DiskLogicalDiskResource) ListAll() (*DiskLogicalDiskConfigList, error) {
	var list DiskLogicalDiskConfigList
	if err := r.c.ReadQuery(BasePath+DiskLogicalDiskEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DiskLogicalDisk configuration identified by id.
func (r *DiskLogicalDiskResource) Get(id string) (*DiskLogicalDiskConfig, error) {
	var item DiskLogicalDiskConfig
	if err := r.c.ReadQuery(BasePath+DiskLogicalDiskEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DiskLogicalDisk configuration.
func (r *DiskLogicalDiskResource) Create(item DiskLogicalDiskConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DiskLogicalDiskEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DiskLogicalDisk configuration identified by id.
func (r *DiskLogicalDiskResource) Edit(id string, item DiskLogicalDiskConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DiskLogicalDiskEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DiskLogicalDisk configuration identified by id.
func (r *DiskLogicalDiskResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DiskLogicalDiskEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
