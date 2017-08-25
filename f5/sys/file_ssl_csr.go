// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileSSLCSRConfigList holds a list of FileSSLCSR configuration.
type FileSSLCSRConfigList struct {
	Items    []FileSSLCSRConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// FileSSLCSRConfig holds the configuration of a single FileSSLCSR.
type FileSSLCSRConfig struct {
}

// FileSSLCSREndpoint represents the REST resource for managing FileSSLCSR.
const FileSSLCSREndpoint = "/file/ssl-csr"

// FileSSLCSRResource provides an API to manage FileSSLCSR configurations.
type FileSSLCSRResource struct {
	c *f5.Client
}

// ListAll  lists all the FileSSLCSR configurations.
func (r *FileSSLCSRResource) ListAll() (*FileSSLCSRConfigList, error) {
	var list FileSSLCSRConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLCSREndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSSLCSR configuration identified by id.
func (r *FileSSLCSRResource) Get(id string) (*FileSSLCSRConfig, error) {
	var item FileSSLCSRConfig
	if err := r.c.ReadQuery(BasePath+FileSSLCSREndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSSLCSR configuration.
func (r *FileSSLCSRResource) Create(item FileSSLCSRConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileSSLCSREndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileSSLCSR configuration identified by id.
func (r *FileSSLCSRResource) Edit(id string, item FileSSLCSRConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileSSLCSREndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileSSLCSR configuration identified by id.
func (r *FileSSLCSRResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSSLCSREndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
