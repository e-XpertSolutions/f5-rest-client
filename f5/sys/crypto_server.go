// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// CryptoServerConfigList holds a list of CryptoServer configuration.
type CryptoServerConfigList struct {
	Items    []CryptoServerConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// CryptoServerConfig holds the configuration of a single CryptoServer.
type CryptoServerConfig struct {
}

// CryptoServerEndpoint represents the REST resource for managing CryptoServer.
const CryptoServerEndpoint = "/crypto/server"

// CryptoServerResource provides an API to manage CryptoServer configurations.
type CryptoServerResource struct {
	c *f5.Client
}

// ListAll  lists all the CryptoServer configurations.
func (r *CryptoServerResource) ListAll() (*CryptoServerConfigList, error) {
	var list CryptoServerConfigList
	if err := r.c.ReadQuery(BasePath+CryptoServerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single CryptoServer configuration identified by id.
func (r *CryptoServerResource) Get(id string) (*CryptoServerConfig, error) {
	var item CryptoServerConfig
	if err := r.c.ReadQuery(BasePath+CryptoServerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new CryptoServer configuration.
func (r *CryptoServerResource) Create(item CryptoServerConfig) error {
	if err := r.c.ModQuery("POST", BasePath+CryptoServerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a CryptoServer configuration identified by id.
func (r *CryptoServerResource) Edit(id string, item CryptoServerConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoServerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single CryptoServer configuration identified by id.
func (r *CryptoServerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoServerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
