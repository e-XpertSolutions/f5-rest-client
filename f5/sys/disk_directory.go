// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DiskDirectoryConfigList holds a list of DiskDirectory configuration.
type DiskDirectoryConfigList struct {
	Items    []DiskDirectoryConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// DiskDirectoryConfig holds the configuration of a single DiskDirectory.
type DiskDirectoryConfig struct {
}

// DiskDirectoryEndpoint represents the REST resource for managing DiskDirectory.
const DiskDirectoryEndpoint = "/disk/directory"

// DiskDirectoryResource provides an API to manage DiskDirectory configurations.
type DiskDirectoryResource struct {
	c *f5.Client
}

// ListAll  lists all the DiskDirectory configurations.
func (r *DiskDirectoryResource) ListAll() (*DiskDirectoryConfigList, error) {
	var list DiskDirectoryConfigList
	if err := r.c.ReadQuery(BasePath+DiskDirectoryEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DiskDirectory configuration identified by id.
func (r *DiskDirectoryResource) Get(id string) (*DiskDirectoryConfig, error) {
	var item DiskDirectoryConfig
	if err := r.c.ReadQuery(BasePath+DiskDirectoryEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DiskDirectory configuration.
func (r *DiskDirectoryResource) Create(item DiskDirectoryConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DiskDirectoryEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DiskDirectory configuration identified by id.
func (r *DiskDirectoryResource) Edit(id string, item DiskDirectoryConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DiskDirectoryEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DiskDirectory configuration identified by id.
func (r *DiskDirectoryResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DiskDirectoryEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
