// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ManagementDHCPConfigList holds a list of ManagementDHCP configuration.
type ManagementDHCPConfigList struct {
	Items    []ManagementDHCPConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// ManagementDHCPConfig holds the configuration of a single ManagementDHCP.
type ManagementDHCPConfig struct {
	FullPath       string   `json:"fullPath"`
	Generation     int      `json:"generation"`
	Kind           string   `json:"kind"`
	Name           string   `json:"name"`
	Partition      string   `json:"partition"`
	RequestOptions []string `json:"requestOptions"`
	SelfLink       string   `json:"selfLink"`
}

// ManagementDHCPEndpoint represents the REST resource for managing ManagementDHCP.
const ManagementDHCPEndpoint = "/management-dhcp"

// ManagementDHCPResource provides an API to manage ManagementDHCP configurations.
type ManagementDHCPResource struct {
	c *f5.Client
}

// ListAll  lists all the ManagementDHCP configurations.
func (r *ManagementDHCPResource) ListAll() (*ManagementDHCPConfigList, error) {
	var list ManagementDHCPConfigList
	if err := r.c.ReadQuery(BasePath+ManagementDHCPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ManagementDHCP configuration identified by id.
func (r *ManagementDHCPResource) Get(id string) (*ManagementDHCPConfig, error) {
	var item ManagementDHCPConfig
	if err := r.c.ReadQuery(BasePath+ManagementDHCPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ManagementDHCP configuration.
func (r *ManagementDHCPResource) Create(item ManagementDHCPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ManagementDHCPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ManagementDHCP configuration identified by id.
func (r *ManagementDHCPResource) Edit(id string, item ManagementDHCPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ManagementDHCPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ManagementDHCP configuration identified by id.
func (r *ManagementDHCPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ManagementDHCPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
