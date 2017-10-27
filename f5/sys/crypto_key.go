// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// CryptoKeyConfigList holds a list of CryptoKey configuration.
type CryptoKeyConfigList struct {
	Items    []CryptoKeyConfig `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

// CryptoKeyConfig holds the configuration of a single CryptoKey.
type CryptoKeyConfig struct {
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	KeySize      string `json:"keySize,omitempty"`
	KeyType      string `json:"keyType,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	SecurityType string `json:"securityType,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
}

// CryptoKeyEndpoint represents the REST resource for managing CryptoKey.
const CryptoKeyEndpoint = "/crypto/key"

// CryptoKeyResource provides an API to manage CryptoKey configurations.
type CryptoKeyResource struct {
	c *f5.Client
}

// ListAll  lists all the CryptoKey configurations.
func (r *CryptoKeyResource) ListAll() (*CryptoKeyConfigList, error) {
	var list CryptoKeyConfigList
	if err := r.c.ReadQuery(BasePath+CryptoKeyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single CryptoKey configuration identified by id.
func (r *CryptoKeyResource) Get(id string) (*CryptoKeyConfig, error) {
	var item CryptoKeyConfig
	if err := r.c.ReadQuery(BasePath+CryptoKeyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new CryptoKey configuration.
func (r *CryptoKeyResource) Create(item CryptoKeyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+CryptoKeyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a CryptoKey configuration identified by id.
func (r *CryptoKeyResource) Edit(id string, item CryptoKeyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoKeyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single CryptoKey configuration identified by id.
func (r *CryptoKeyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoKeyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
