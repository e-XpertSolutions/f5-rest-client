// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogProfileNetworkConfigList holds a list of LogProfileNetwork configuration.
type LogProfileNetworkConfigList struct {
	Items    []LogProfileNetworkConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

// LogProfileNetworkConfig holds the configuration of a single LogProfileNetwork.
type LogProfileNetworkConfig struct {
}

// LogProfileNetworkEndpoint represents the REST resource for managing LogProfileNetwork.
const LogProfileNetworkEndpoint = "/log/profile_network"

// LogProfileNetworkResource provides an API to manage LogProfileNetwork configurations.
type LogProfileNetworkResource struct {
	c *f5.Client
}

// ListAll  lists all the LogProfileNetwork configurations.
func (r *LogProfileNetworkResource) ListAll() (*LogProfileNetworkConfigList, error) {
	var list LogProfileNetworkConfigList
	if err := r.c.ReadQuery(BasePath+LogProfileNetworkEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogProfileNetwork configuration identified by id.
func (r *LogProfileNetworkResource) Get(id string) (*LogProfileNetworkConfig, error) {
	var item LogProfileNetworkConfig
	if err := r.c.ReadQuery(BasePath+LogProfileNetworkEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogProfileNetwork configuration.
func (r *LogProfileNetworkResource) Create(item LogProfileNetworkConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogProfileNetworkEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogProfileNetwork configuration identified by id.
func (r *LogProfileNetworkResource) Edit(id string, item LogProfileNetworkConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogProfileNetworkEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogProfileNetwork configuration identified by id.
func (r *LogProfileNetworkResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogProfileNetworkEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
