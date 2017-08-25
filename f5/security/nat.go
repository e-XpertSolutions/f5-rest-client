// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// NATConfigList holds a list of NAT configuration.
type NATConfigList struct {
	Items    []NATConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// NATConfig holds the configuration of a single NAT.
type NATConfig struct {
}

// NATEndpoint represents the REST resource for managing NAT.
const NATEndpoint = "/nat"

// NATResource provides an API to manage NAT configurations.
type NATResource struct {
	c *f5.Client
}

// ListAll  lists all the NAT configurations.
func (r *NATResource) ListAll() (*NATConfigList, error) {
	var list NATConfigList
	if err := r.c.ReadQuery(BasePath+NATEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single NAT configuration identified by id.
func (r *NATResource) Get(id string) (*NATConfig, error) {
	var item NATConfig
	if err := r.c.ReadQuery(BasePath+NATEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new NAT configuration.
func (r *NATResource) Create(item NATConfig) error {
	if err := r.c.ModQuery("POST", BasePath+NATEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a NAT configuration identified by id.
func (r *NATResource) Edit(id string, item NATConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+NATEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single NAT configuration identified by id.
func (r *NATResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+NATEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
