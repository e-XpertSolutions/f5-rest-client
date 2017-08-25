// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ManagementIPConfigList holds a list of ManagementIP configuration.
type ManagementIPConfigList struct {
	Items    []ManagementIPConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// ManagementIPConfig holds the configuration of a single ManagementIP.
type ManagementIPConfig struct {
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	SelfLink   string `json:"selfLink"`
}

// ManagementIPEndpoint represents the REST resource for managing ManagementIP.
const ManagementIPEndpoint = "/management-ip"

// ManagementIPResource provides an API to manage ManagementIP configurations.
type ManagementIPResource struct {
	c *f5.Client
}

// ListAll  lists all the ManagementIP configurations.
func (r *ManagementIPResource) ListAll() (*ManagementIPConfigList, error) {
	var list ManagementIPConfigList
	if err := r.c.ReadQuery(BasePath+ManagementIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ManagementIP configuration identified by id.
func (r *ManagementIPResource) Get(id string) (*ManagementIPConfig, error) {
	var item ManagementIPConfig
	if err := r.c.ReadQuery(BasePath+ManagementIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ManagementIP configuration.
func (r *ManagementIPResource) Create(item ManagementIPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ManagementIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ManagementIP configuration identified by id.
func (r *ManagementIPResource) Edit(id string, item ManagementIPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ManagementIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ManagementIP configuration identified by id.
func (r *ManagementIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ManagementIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
