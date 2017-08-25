// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// CryptoPKCS12ConfigList holds a list of CryptoPKCS12 configuration.
type CryptoPKCS12ConfigList struct {
	Items    []CryptoPKCS12Config `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// CryptoPKCS12Config holds the configuration of a single CryptoPKCS12.
type CryptoPKCS12Config struct {
}

// CryptoPKCS12Endpoint represents the REST resource for managing CryptoPKCS12.
const CryptoPKCS12Endpoint = "/crypto/pkcs12"

// CryptoPKCS12Resource provides an API to manage CryptoPKCS12 configurations.
type CryptoPKCS12Resource struct {
	c *f5.Client
}

// ListAll  lists all the CryptoPKCS12 configurations.
func (r *CryptoPKCS12Resource) ListAll() (*CryptoPKCS12ConfigList, error) {
	var list CryptoPKCS12ConfigList
	if err := r.c.ReadQuery(BasePath+CryptoPKCS12Endpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single CryptoPKCS12 configuration identified by id.
func (r *CryptoPKCS12Resource) Get(id string) (*CryptoPKCS12Config, error) {
	var item CryptoPKCS12Config
	if err := r.c.ReadQuery(BasePath+CryptoPKCS12Endpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new CryptoPKCS12 configuration.
func (r *CryptoPKCS12Resource) Create(item CryptoPKCS12Config) error {
	if err := r.c.ModQuery("POST", BasePath+CryptoPKCS12Endpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a CryptoPKCS12 configuration identified by id.
func (r *CryptoPKCS12Resource) Edit(id string, item CryptoPKCS12Config) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoPKCS12Endpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single CryptoPKCS12 configuration identified by id.
func (r *CryptoPKCS12Resource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoPKCS12Endpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
