// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSProfileDOSNetworkConfigList holds a list of DOSProfileDOSNetwork configuration.
type DOSProfileDOSNetworkConfigList struct {
	Items    []DOSProfileDOSNetworkConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// DOSProfileDOSNetworkConfig holds the configuration of a single DOSProfileDOSNetwork.
type DOSProfileDOSNetworkConfig struct {
}

// DOSProfileDOSNetworkEndpoint represents the REST resource for managing DOSProfileDOSNetwork.
const DOSProfileDOSNetworkEndpoint = "/dos/profile_dos-network"

// DOSProfileDOSNetworkResource provides an API to manage DOSProfileDOSNetwork configurations.
type DOSProfileDOSNetworkResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSProfileDOSNetwork configurations.
func (r *DOSProfileDOSNetworkResource) ListAll() (*DOSProfileDOSNetworkConfigList, error) {
	var list DOSProfileDOSNetworkConfigList
	if err := r.c.ReadQuery(BasePath+DOSProfileDOSNetworkEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSProfileDOSNetwork configuration identified by id.
func (r *DOSProfileDOSNetworkResource) Get(id string) (*DOSProfileDOSNetworkConfig, error) {
	var item DOSProfileDOSNetworkConfig
	if err := r.c.ReadQuery(BasePath+DOSProfileDOSNetworkEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSProfileDOSNetwork configuration.
func (r *DOSProfileDOSNetworkResource) Create(item DOSProfileDOSNetworkConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSProfileDOSNetworkEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSProfileDOSNetwork configuration identified by id.
func (r *DOSProfileDOSNetworkResource) Edit(id string, item DOSProfileDOSNetworkConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSProfileDOSNetworkEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSProfileDOSNetwork configuration identified by id.
func (r *DOSProfileDOSNetworkResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSProfileDOSNetworkEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
