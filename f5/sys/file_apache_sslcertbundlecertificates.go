// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileApacheSSLCertBundleCertificatesConfigList holds a list of FileApacheSSLCertBundleCertificates configuration.
type FileApacheSSLCertBundleCertificatesConfigList struct {
	Items    []FileApacheSSLCertBundleCertificatesConfig `json:"items"`
	Kind     string                                      `json:"kind"`
	SelfLink string                                      `json:"selflink"`
}

// FileApacheSSLCertBundleCertificatesConfig holds the configuration of a single FileApacheSSLCertBundleCertificates.
type FileApacheSSLCertBundleCertificatesConfig struct {
}

// FileApacheSSLCertBundleCertificatesEndpoint represents the REST resource for managing FileApacheSSLCertBundleCertificates.
const FileApacheSSLCertBundleCertificatesEndpoint = "/file/apache-ssl-cert_bundle-certificates"

// FileApacheSSLCertBundleCertificatesResource provides an API to manage FileApacheSSLCertBundleCertificates configurations.
type FileApacheSSLCertBundleCertificatesResource struct {
	c *f5.Client
}

// ListAll  lists all the FileApacheSSLCertBundleCertificates configurations.
func (r *FileApacheSSLCertBundleCertificatesResource) ListAll() (*FileApacheSSLCertBundleCertificatesConfigList, error) {
	var list FileApacheSSLCertBundleCertificatesConfigList
	if err := r.c.ReadQuery(BasePath+FileApacheSSLCertBundleCertificatesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileApacheSSLCertBundleCertificates configuration identified by id.
func (r *FileApacheSSLCertBundleCertificatesResource) Get(id string) (*FileApacheSSLCertBundleCertificatesConfig, error) {
	var item FileApacheSSLCertBundleCertificatesConfig
	if err := r.c.ReadQuery(BasePath+FileApacheSSLCertBundleCertificatesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileApacheSSLCertBundleCertificates configuration.
func (r *FileApacheSSLCertBundleCertificatesResource) Create(item FileApacheSSLCertBundleCertificatesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileApacheSSLCertBundleCertificatesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileApacheSSLCertBundleCertificates configuration identified by id.
func (r *FileApacheSSLCertBundleCertificatesResource) Edit(id string, item FileApacheSSLCertBundleCertificatesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileApacheSSLCertBundleCertificatesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileApacheSSLCertBundleCertificates configuration identified by id.
func (r *FileApacheSSLCertBundleCertificatesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileApacheSSLCertBundleCertificatesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
