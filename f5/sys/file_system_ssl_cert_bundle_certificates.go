// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileSystemSSLCertBundleCertificatesConfigList holds a list of FileSystemSSLCertBundleCertificates configuration.
type FileSystemSSLCertBundleCertificatesConfigList struct {
	Items    []FileSystemSSLCertBundleCertificatesConfig `json:"items"`
	Kind     string                                      `json:"kind"`
	SelfLink string                                      `json:"selflink"`
}

// FileSystemSSLCertBundleCertificatesConfig holds the configuration of a single FileSystemSSLCertBundleCertificates.
type FileSystemSSLCertBundleCertificatesConfig struct {
}

// FileSystemSSLCertBundleCertificatesEndpoint represents the REST resource for managing FileSystemSSLCertBundleCertificates.
const FileSystemSSLCertBundleCertificatesEndpoint = "/file/system-ssl-cert_bundle-certificates"

// FileSystemSSLCertBundleCertificatesResource provides an API to manage FileSystemSSLCertBundleCertificates configurations.
type FileSystemSSLCertBundleCertificatesResource struct {
	c *f5.Client
}

// ListAll  lists all the FileSystemSSLCertBundleCertificates configurations.
func (r *FileSystemSSLCertBundleCertificatesResource) ListAll() (*FileSystemSSLCertBundleCertificatesConfigList, error) {
	var list FileSystemSSLCertBundleCertificatesConfigList
	if err := r.c.ReadQuery(BasePath+FileSystemSSLCertBundleCertificatesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSystemSSLCertBundleCertificates configuration identified by id.
func (r *FileSystemSSLCertBundleCertificatesResource) Get(id string) (*FileSystemSSLCertBundleCertificatesConfig, error) {
	var item FileSystemSSLCertBundleCertificatesConfig
	if err := r.c.ReadQuery(BasePath+FileSystemSSLCertBundleCertificatesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSystemSSLCertBundleCertificates configuration.
func (r *FileSystemSSLCertBundleCertificatesResource) Create(item FileSystemSSLCertBundleCertificatesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileSystemSSLCertBundleCertificatesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileSystemSSLCertBundleCertificates configuration identified by id.
func (r *FileSystemSSLCertBundleCertificatesResource) Edit(id string, item FileSystemSSLCertBundleCertificatesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileSystemSSLCertBundleCertificatesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileSystemSSLCertBundleCertificates configuration identified by id.
func (r *FileSystemSSLCertBundleCertificatesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSystemSSLCertBundleCertificatesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
