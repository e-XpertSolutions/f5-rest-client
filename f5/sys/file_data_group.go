// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileDataGroupConfigList holds a list of FileDataGroup configuration.
type FileDataGroupConfigList struct {
	Items    []FileDataGroupConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// FileDataGroupConfig holds the configuration of a single FileDataGroup.
type FileDataGroupConfig struct {
}

// FileDataGroupEndpoint represents the REST resource for managing FileDataGroup.
const FileDataGroupEndpoint = "/file/data-group"

// FileDataGroupResource provides an API to manage FileDataGroup configurations.
type FileDataGroupResource struct {
	c *f5.Client
}

// ListAll  lists all the FileDataGroup configurations.
func (r *FileDataGroupResource) ListAll() (*FileDataGroupConfigList, error) {
	var list FileDataGroupConfigList
	if err := r.c.ReadQuery(BasePath+FileDataGroupEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileDataGroup configuration identified by id.
func (r *FileDataGroupResource) Get(id string) (*FileDataGroupConfig, error) {
	var item FileDataGroupConfig
	if err := r.c.ReadQuery(BasePath+FileDataGroupEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileDataGroup configuration.
func (r *FileDataGroupResource) Create(item FileDataGroupConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileDataGroupEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileDataGroup configuration identified by id.
func (r *FileDataGroupResource) Edit(id string, item FileDataGroupConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileDataGroupEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileDataGroup configuration identified by id.
func (r *FileDataGroupResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileDataGroupEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
