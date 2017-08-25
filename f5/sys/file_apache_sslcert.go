// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileApacheSSLCertConfigList holds a list of FileApacheSSLCert configuration.
type FileApacheSSLCertConfigList struct {
	Items    []FileApacheSSLCertConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

// FileApacheSSLCertConfig holds the configuration of a single FileApacheSSLCert.
type FileApacheSSLCertConfig struct {
}

// FileApacheSSLCertEndpoint represents the REST resource for managing FileApacheSSLCert.
const FileApacheSSLCertEndpoint = "/file/apache-ssl-cert"

// FileApacheSSLCertResource provides an API to manage FileApacheSSLCert configurations.
type FileApacheSSLCertResource struct {
	c *f5.Client
}

// ListAll  lists all the FileApacheSSLCert configurations.
func (r *FileApacheSSLCertResource) ListAll() (*FileApacheSSLCertConfigList, error) {
	var list FileApacheSSLCertConfigList
	if err := r.c.ReadQuery(BasePath+FileApacheSSLCertEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileApacheSSLCert configuration identified by id.
func (r *FileApacheSSLCertResource) Get(id string) (*FileApacheSSLCertConfig, error) {
	var item FileApacheSSLCertConfig
	if err := r.c.ReadQuery(BasePath+FileApacheSSLCertEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileApacheSSLCert configuration.
func (r *FileApacheSSLCertResource) Create(item FileApacheSSLCertConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileApacheSSLCertEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileApacheSSLCert configuration identified by id.
func (r *FileApacheSSLCertResource) Edit(id string, item FileApacheSSLCertConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileApacheSSLCertEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileApacheSSLCert configuration identified by id.
func (r *FileApacheSSLCertResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileApacheSSLCertEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
