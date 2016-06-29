// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileSSLCRLConfigList holds a list of FileSSLCRL configuration.
type FileSSLCRLConfigList struct {
	Items    []FileSSLCRLConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// FileSSLCRLConfig holds the configuration of a single FileSSLCRL.
type FileSSLCRLConfig struct {
}

// FileSSLCRLEndpoint represents the REST resource for managing FileSSLCRL.
const FileSSLCRLEndpoint = "/tm/sys/file/ssl-crl"

// FileSSLCRLResource provides an API to manage FileSSLCRL configurations.
type FileSSLCRLResource struct {
	c f5.Client
}

// ListAll  lists all the FileSSLCRL configurations.
func (r *FileSSLCRLResource) ListAll() (*FileSSLCRLConfigList, error) {
	var list FileSSLCRLConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLCRLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSSLCRL configuration identified by id.
func (r *FileSSLCRLResource) Get(id string) (*FileSSLCRLConfig, error) {
	var item FileSSLCRLConfig
	if err := r.c.ReadQuery(BasePath+FileSSLCRLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSSLCRL configuration.
func (r *FileSSLCRLResource) Create(item FileSSLCRLConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileSSLCRLEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileSSLCRL configuration identified by id.
func (r *FileSSLCRLResource) Edit(id string, item FileSSLCRLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileSSLCRLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileSSLCRL configuration identified by id.
func (r *FileSSLCRLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSSLCRLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
