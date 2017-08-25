// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// CryptoCheckCertConfigList holds a list of CryptoCheckCert configuration.
type CryptoCheckCertConfigList struct {
	Items    []CryptoCheckCertConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// CryptoCheckCertConfig holds the configuration of a single CryptoCheckCert.
type CryptoCheckCertConfig struct {
}

// CryptoCheckCertEndpoint represents the REST resource for managing CryptoCheckCert.
const CryptoCheckCertEndpoint = "/crypto/check-cert"

// CryptoCheckCertResource provides an API to manage CryptoCheckCert configurations.
type CryptoCheckCertResource struct {
	c *f5.Client
}

// ListAll  lists all the CryptoCheckCert configurations.
func (r *CryptoCheckCertResource) ListAll() (*CryptoCheckCertConfigList, error) {
	var list CryptoCheckCertConfigList
	if err := r.c.ReadQuery(BasePath+CryptoCheckCertEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single CryptoCheckCert configuration identified by id.
func (r *CryptoCheckCertResource) Get(id string) (*CryptoCheckCertConfig, error) {
	var item CryptoCheckCertConfig
	if err := r.c.ReadQuery(BasePath+CryptoCheckCertEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new CryptoCheckCert configuration.
func (r *CryptoCheckCertResource) Create(item CryptoCheckCertConfig) error {
	if err := r.c.ModQuery("POST", BasePath+CryptoCheckCertEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a CryptoCheckCert configuration identified by id.
func (r *CryptoCheckCertResource) Edit(id string, item CryptoCheckCertConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoCheckCertEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single CryptoCheckCert configuration identified by id.
func (r *CryptoCheckCertResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoCheckCertEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
