// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSConfigList holds a list of DOS configuration.
type DOSConfigList struct {
	Items    []DOSConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// DOSConfig holds the configuration of a single DOS.
type DOSConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// DOSEndpoint represents the REST resource for managing DOS.
const DOSEndpoint = "/dos"

// DOSResource provides an API to manage DOS configurations.
type DOSResource struct {
	c *f5.Client
}

// ListAll  lists all the DOS configurations.
func (r *DOSResource) ListAll() (*DOSConfigList, error) {
	var list DOSConfigList
	if err := r.c.ReadQuery(BasePath+DOSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOS configuration identified by id.
func (r *DOSResource) Get(id string) (*DOSConfig, error) {
	var item DOSConfig
	if err := r.c.ReadQuery(BasePath+DOSEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOS configuration.
func (r *DOSResource) Create(item DOSConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOS configuration identified by id.
func (r *DOSResource) Edit(id string, item DOSConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOS configuration identified by id.
func (r *DOSResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
