// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// CryptoCSRConfigList holds a list of CryptoCSR configuration.
type CryptoCSRConfigList struct {
	Items    []CryptoCSRConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// CryptoCSRConfig holds the configuration of a single CryptoCSR.
type CryptoCSRConfig struct {
}

// CryptoCSREndpoint represents the REST resource for managing CryptoCSR.
const CryptoCSREndpoint = "/crypto/csr"

// CryptoCSRResource provides an API to manage CryptoCSR configurations.
type CryptoCSRResource struct {
	c *f5.Client
}

// ListAll  lists all the CryptoCSR configurations.
func (r *CryptoCSRResource) ListAll() (*CryptoCSRConfigList, error) {
	var list CryptoCSRConfigList
	if err := r.c.ReadQuery(BasePath+CryptoCSREndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single CryptoCSR configuration identified by id.
func (r *CryptoCSRResource) Get(id string) (*CryptoCSRConfig, error) {
	var item CryptoCSRConfig
	if err := r.c.ReadQuery(BasePath+CryptoCSREndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new CryptoCSR configuration.
func (r *CryptoCSRResource) Create(item CryptoCSRConfig) error {
	if err := r.c.ModQuery("POST", BasePath+CryptoCSREndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a CryptoCSR configuration identified by id.
func (r *CryptoCSRResource) Edit(id string, item CryptoCSRConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoCSREndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single CryptoCSR configuration identified by id.
func (r *CryptoCSRResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoCSREndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
