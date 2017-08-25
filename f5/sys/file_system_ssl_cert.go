// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileSystemSSLCertConfigList holds a list of FileSystemSSLCert configuration.
type FileSystemSSLCertConfigList struct {
	Items    []FileSystemSSLCertConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

// FileSystemSSLCertConfig holds the configuration of a single FileSystemSSLCert.
type FileSystemSSLCertConfig struct {
}

// FileSystemSSLCertEndpoint represents the REST resource for managing FileSystemSSLCert.
const FileSystemSSLCertEndpoint = "/file/system-ssl-cert"

// FileSystemSSLCertResource provides an API to manage FileSystemSSLCert configurations.
type FileSystemSSLCertResource struct {
	c *f5.Client
}

// ListAll  lists all the FileSystemSSLCert configurations.
func (r *FileSystemSSLCertResource) ListAll() (*FileSystemSSLCertConfigList, error) {
	var list FileSystemSSLCertConfigList
	if err := r.c.ReadQuery(BasePath+FileSystemSSLCertEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSystemSSLCert configuration identified by id.
func (r *FileSystemSSLCertResource) Get(id string) (*FileSystemSSLCertConfig, error) {
	var item FileSystemSSLCertConfig
	if err := r.c.ReadQuery(BasePath+FileSystemSSLCertEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSystemSSLCert configuration.
func (r *FileSystemSSLCertResource) Create(item FileSystemSSLCertConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileSystemSSLCertEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileSystemSSLCert configuration identified by id.
func (r *FileSystemSSLCertResource) Edit(id string, item FileSystemSSLCertConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileSystemSSLCertEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileSystemSSLCert configuration identified by id.
func (r *FileSystemSSLCertResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSystemSSLCertEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
