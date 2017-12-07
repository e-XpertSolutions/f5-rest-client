// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcmp

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// VirtualDiskConfigList holds a list of VirtualDisk configuration.
type VirtualDiskConfigList struct {
	Items    []VirtualDiskConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

// VirtualDiskConfig holds the configuration of a single VirtualDisk.
type VirtualDiskConfig struct {
	Name            string `json:"name,omitempty"`
	Kind            string `json:"kind,omitempty"`
	FullPath        string `json:"fullPath,omitempty"`
	Generation      int    `json:"generation,omitempty"`
	SelfLink        string `json:"selfLink,omitempty"`
	OperatingSystem string `json:"operatingSystem,omitempty"`
}

// VirtualDiskEndpoint represents the REST resource for managing VirtualDisk.
const VirtualDiskEndpoint = "/virtual-disk"

// VirtualDiskResource provides an API to manage VirtualDisk configurations.
type VirtualDiskResource struct {
	c *f5.Client
}

// ListAll  lists all the VirtualDisk configurations.
func (r *VirtualDiskResource) ListAll() (*VirtualDiskConfigList, error) {
	var list VirtualDiskConfigList
	if err := r.c.ReadQuery(BasePath+VirtualDiskEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single VirtualDisk configuration identified by id.
func (r *VirtualDiskResource) Get(id string) (*VirtualDiskConfig, error) {
	var item VirtualDiskConfig
	if err := r.c.ReadQuery(BasePath+VirtualDiskEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new VirtualDisk configuration.
func (r *VirtualDiskResource) Create(item VirtualDiskConfig) error {
	if err := r.c.ModQuery("POST", BasePath+VirtualDiskEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a VirtualDisk configuration identified by id.
func (r *VirtualDiskResource) Edit(id string, item VirtualDiskConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+VirtualDiskEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single VirtualDisk configuration identified by id.
func (r *VirtualDiskResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+VirtualDiskEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
