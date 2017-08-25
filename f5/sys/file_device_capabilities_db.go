// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileDeviceCapabilitiesDBConfigList holds a list of FileDeviceCapabilitiesDB configuration.
type FileDeviceCapabilitiesDBConfigList struct {
	Items    []FileDeviceCapabilitiesDBConfig `json:"items"`
	Kind     string                           `json:"kind"`
	SelfLink string                           `json:"selflink"`
}

// FileDeviceCapabilitiesDBConfig holds the configuration of a single FileDeviceCapabilitiesDB.
type FileDeviceCapabilitiesDBConfig struct {
}

// FileDeviceCapabilitiesDBEndpoint represents the REST resource for managing FileDeviceCapabilitiesDB.
const FileDeviceCapabilitiesDBEndpoint = "/file/device-capabilities-db"

// FileDeviceCapabilitiesDBResource provides an API to manage FileDeviceCapabilitiesDB configurations.
type FileDeviceCapabilitiesDBResource struct {
	c *f5.Client
}

// ListAll  lists all the FileDeviceCapabilitiesDB configurations.
func (r *FileDeviceCapabilitiesDBResource) ListAll() (*FileDeviceCapabilitiesDBConfigList, error) {
	var list FileDeviceCapabilitiesDBConfigList
	if err := r.c.ReadQuery(BasePath+FileDeviceCapabilitiesDBEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileDeviceCapabilitiesDB configuration identified by id.
func (r *FileDeviceCapabilitiesDBResource) Get(id string) (*FileDeviceCapabilitiesDBConfig, error) {
	var item FileDeviceCapabilitiesDBConfig
	if err := r.c.ReadQuery(BasePath+FileDeviceCapabilitiesDBEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileDeviceCapabilitiesDB configuration.
func (r *FileDeviceCapabilitiesDBResource) Create(item FileDeviceCapabilitiesDBConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileDeviceCapabilitiesDBEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileDeviceCapabilitiesDB configuration identified by id.
func (r *FileDeviceCapabilitiesDBResource) Edit(id string, item FileDeviceCapabilitiesDBConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileDeviceCapabilitiesDBEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileDeviceCapabilitiesDB configuration identified by id.
func (r *FileDeviceCapabilitiesDBResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileDeviceCapabilitiesDBEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
