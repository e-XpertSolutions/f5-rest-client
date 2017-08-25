// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// CryptoClientConfigList holds a list of CryptoClient configuration.
type CryptoClientConfigList struct {
	Items    []CryptoClientConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// CryptoClientConfig holds the configuration of a single CryptoClient.
type CryptoClientConfig struct {
}

// CryptoClientEndpoint represents the REST resource for managing CryptoClient.
const CryptoClientEndpoint = "/crypto/client"

// CryptoClientResource provides an API to manage CryptoClient configurations.
type CryptoClientResource struct {
	c *f5.Client
}

// ListAll  lists all the CryptoClient configurations.
func (r *CryptoClientResource) ListAll() (*CryptoClientConfigList, error) {
	var list CryptoClientConfigList
	if err := r.c.ReadQuery(BasePath+CryptoClientEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single CryptoClient configuration identified by id.
func (r *CryptoClientResource) Get(id string) (*CryptoClientConfig, error) {
	var item CryptoClientConfig
	if err := r.c.ReadQuery(BasePath+CryptoClientEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new CryptoClient configuration.
func (r *CryptoClientResource) Create(item CryptoClientConfig) error {
	if err := r.c.ModQuery("POST", BasePath+CryptoClientEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a CryptoClient configuration identified by id.
func (r *CryptoClientResource) Edit(id string, item CryptoClientConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoClientEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single CryptoClient configuration identified by id.
func (r *CryptoClientResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoClientEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
