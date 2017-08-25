// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileSystemSSLKeyConfigList holds a list of FileSystemSSLKey configuration.
type FileSystemSSLKeyConfigList struct {
	Items    []FileSystemSSLKeyConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// FileSystemSSLKeyConfig holds the configuration of a single FileSystemSSLKey.
type FileSystemSSLKeyConfig struct {
}

// FileSystemSSLKeyEndpoint represents the REST resource for managing FileSystemSSLKey.
const FileSystemSSLKeyEndpoint = "/file/system-ssl-key"

// FileSystemSSLKeyResource provides an API to manage FileSystemSSLKey configurations.
type FileSystemSSLKeyResource struct {
	c *f5.Client
}

// ListAll  lists all the FileSystemSSLKey configurations.
func (r *FileSystemSSLKeyResource) ListAll() (*FileSystemSSLKeyConfigList, error) {
	var list FileSystemSSLKeyConfigList
	if err := r.c.ReadQuery(BasePath+FileSystemSSLKeyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSystemSSLKey configuration identified by id.
func (r *FileSystemSSLKeyResource) Get(id string) (*FileSystemSSLKeyConfig, error) {
	var item FileSystemSSLKeyConfig
	if err := r.c.ReadQuery(BasePath+FileSystemSSLKeyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSystemSSLKey configuration.
func (r *FileSystemSSLKeyResource) Create(item FileSystemSSLKeyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileSystemSSLKeyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileSystemSSLKey configuration identified by id.
func (r *FileSystemSSLKeyResource) Edit(id string, item FileSystemSSLKeyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileSystemSSLKeyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileSystemSSLKey configuration identified by id.
func (r *FileSystemSSLKeyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSystemSSLKeyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
