// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPFixConfigList holds a list of IPFix configuration.
type IPFixConfigList struct {
	Items    []IPFixConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// IPFixConfig holds the configuration of a single IPFix.
type IPFixConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// IPFixEndpoint represents the REST resource for managing IPFix.
const IPFixEndpoint = "/ipfix"

// IPFixResource provides an API to manage IPFix configurations.
type IPFixResource struct {
	c *f5.Client
}

// ListAll  lists all the IPFix configurations.
func (r *IPFixResource) ListAll() (*IPFixConfigList, error) {
	var list IPFixConfigList
	if err := r.c.ReadQuery(BasePath+IPFixEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPFix configuration identified by id.
func (r *IPFixResource) Get(id string) (*IPFixConfig, error) {
	var item IPFixConfig
	if err := r.c.ReadQuery(BasePath+IPFixEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPFix configuration.
func (r *IPFixResource) Create(item IPFixConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPFixEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPFix configuration identified by id.
func (r *IPFixResource) Edit(id string, item IPFixConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPFixEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPFix configuration identified by id.
func (r *IPFixResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPFixEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
