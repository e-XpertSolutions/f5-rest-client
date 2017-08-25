// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// CryptoConfigList holds a list of Crypto configuration.
type CryptoConfigList struct {
	Items    []CryptoConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// CryptoConfig holds the configuration of a single Crypto.
type CryptoConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// CryptoEndpoint represents the REST resource for managing Crypto.
const CryptoEndpoint = "/crypto"

// CryptoResource provides an API to manage Crypto configurations.
type CryptoResource struct {
	c *f5.Client
}

// ListAll  lists all the Crypto configurations.
func (r *CryptoResource) ListAll() (*CryptoConfigList, error) {
	var list CryptoConfigList
	if err := r.c.ReadQuery(BasePath+CryptoEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Crypto configuration identified by id.
func (r *CryptoResource) Get(id string) (*CryptoConfig, error) {
	var item CryptoConfig
	if err := r.c.ReadQuery(BasePath+CryptoEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Crypto configuration.
func (r *CryptoResource) Create(item CryptoConfig) error {
	if err := r.c.ModQuery("POST", BasePath+CryptoEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Crypto configuration identified by id.
func (r *CryptoResource) Edit(id string, item CryptoConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Crypto configuration identified by id.
func (r *CryptoResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
