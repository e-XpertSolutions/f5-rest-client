// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileSSLCertBundleCertificatesConfigList holds a list of FileSSLCertBundleCertificates configuration.
type FileSSLCertBundleCertificatesConfigList struct {
	Items    []FileSSLCertBundleCertificatesConfig `json:"items"`
	Kind     string                                `json:"kind"`
	SelfLink string                                `json:"selflink"`
}

// FileSSLCertBundleCertificatesConfig holds the configuration of a single FileSSLCertBundleCertificates.
type FileSSLCertBundleCertificatesConfig struct {
}

// FileSSLCertBundleCertificatesEndpoint represents the REST resource for managing FileSSLCertBundleCertificates.
const FileSSLCertBundleCertificatesEndpoint = "/file/ssl-cert_bundle-certificates"

// FileSSLCertBundleCertificatesResource provides an API to manage FileSSLCertBundleCertificates configurations.
type FileSSLCertBundleCertificatesResource struct {
	c *f5.Client
}

// ListAll  lists all the FileSSLCertBundleCertificates configurations.
func (r *FileSSLCertBundleCertificatesResource) ListAll() (*FileSSLCertBundleCertificatesConfigList, error) {
	var list FileSSLCertBundleCertificatesConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLCertBundleCertificatesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSSLCertBundleCertificates configuration identified by id.
func (r *FileSSLCertBundleCertificatesResource) Get(id string) (*FileSSLCertBundleCertificatesConfig, error) {
	var item FileSSLCertBundleCertificatesConfig
	if err := r.c.ReadQuery(BasePath+FileSSLCertBundleCertificatesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSSLCertBundleCertificates configuration.
func (r *FileSSLCertBundleCertificatesResource) Create(item FileSSLCertBundleCertificatesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileSSLCertBundleCertificatesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileSSLCertBundleCertificates configuration identified by id.
func (r *FileSSLCertBundleCertificatesResource) Edit(id string, item FileSSLCertBundleCertificatesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileSSLCertBundleCertificatesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileSSLCertBundleCertificates configuration identified by id.
func (r *FileSSLCertBundleCertificatesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSSLCertBundleCertificatesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
