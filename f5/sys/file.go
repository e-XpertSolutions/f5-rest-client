// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileConfigList holds a list of File configuration.
type FileConfigList struct {
	Items    []FileConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// FileConfig holds the configuration of a single File.
type FileConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// FileEndpoint represents the REST resource for managing File.
const FileEndpoint = "/tm/sys/file"

// FileResource provides an API to manage File configurations.
type FileResource struct {
	c f5.Client
}

// ListAll  lists all the File configurations.
func (r *FileResource) ListAll() (*FileConfigList, error) {
	var list FileConfigList
	if err := r.c.ReadQuery(BasePath+FileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single File configuration identified by id.
func (r *FileResource) Get(id string) (*FileConfig, error) {
	var item FileConfig
	if err := r.c.ReadQuery(BasePath+FileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new File configuration.
func (r *FileResource) Create(item FileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a File configuration identified by id.
func (r *FileResource) Edit(id string, item FileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single File configuration identified by id.
func (r *FileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
