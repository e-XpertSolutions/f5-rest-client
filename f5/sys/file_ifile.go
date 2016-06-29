// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileIFileConfigList holds a list of FileIFile configuration.
type FileIFileConfigList struct {
	Items    []FileIFileConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// FileIFileConfig holds the configuration of a single FileIFile.
type FileIFileConfig struct {
}

// FileIFileEndpoint represents the REST resource for managing FileIFile.
const FileIFileEndpoint = "/tm/sys/file/ifile"

// FileIFileResource provides an API to manage FileIFile configurations.
type FileIFileResource struct {
	c f5.Client
}

// ListAll  lists all the FileIFile configurations.
func (r *FileIFileResource) ListAll() (*FileIFileConfigList, error) {
	var list FileIFileConfigList
	if err := r.c.ReadQuery(BasePath+FileIFileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileIFile configuration identified by id.
func (r *FileIFileResource) Get(id string) (*FileIFileConfig, error) {
	var item FileIFileConfig
	if err := r.c.ReadQuery(BasePath+FileIFileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileIFile configuration.
func (r *FileIFileResource) Create(item FileIFileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileIFileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileIFile configuration identified by id.
func (r *FileIFileResource) Edit(id string, item FileIFileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileIFileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileIFile configuration identified by id.
func (r *FileIFileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileIFileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
