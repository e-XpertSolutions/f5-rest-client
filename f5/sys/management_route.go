// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ManagementRouteConfigList holds a list of ManagementRoute configuration.
type ManagementRouteConfigList struct {
	Items    []ManagementRouteConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// ManagementRouteConfig holds the configuration of a single ManagementRoute.
type ManagementRouteConfig struct {
	FullPath   string `json:"fullPath"`
	Gateway    string `json:"gateway"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Mtu        int    `json:"mtu"`
	Name       string `json:"name"`
	Network    string `json:"network"`
	Partition  string `json:"partition"`
	SelfLink   string `json:"selfLink"`
}

// ManagementRouteEndpoint represents the REST resource for managing ManagementRoute.
const ManagementRouteEndpoint = "/management-route"

// ManagementRouteResource provides an API to manage ManagementRoute configurations.
type ManagementRouteResource struct {
	c *f5.Client
}

// ListAll  lists all the ManagementRoute configurations.
func (r *ManagementRouteResource) ListAll() (*ManagementRouteConfigList, error) {
	var list ManagementRouteConfigList
	if err := r.c.ReadQuery(BasePath+ManagementRouteEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ManagementRoute configuration identified by id.
func (r *ManagementRouteResource) Get(id string) (*ManagementRouteConfig, error) {
	var item ManagementRouteConfig
	if err := r.c.ReadQuery(BasePath+ManagementRouteEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ManagementRoute configuration.
func (r *ManagementRouteResource) Create(item ManagementRouteConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ManagementRouteEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ManagementRoute configuration identified by id.
func (r *ManagementRouteResource) Edit(id string, item ManagementRouteConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ManagementRouteEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ManagementRoute configuration identified by id.
func (r *ManagementRouteResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ManagementRouteEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
