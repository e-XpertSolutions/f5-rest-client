// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogConfigConfigList holds a list of LogConfig configuration.
type LogConfigConfigList struct {
	Items    []LogConfigConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// LogConfigConfig holds the configuration of a single LogConfig.
type LogConfigConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// LogConfigEndpoint represents the REST resource for managing LogConfig.
const LogConfigEndpoint = "/log-config"

// LogConfigResource provides an API to manage LogConfig configurations.
type LogConfigResource struct {
	c *f5.Client
}

// ListAll  lists all the LogConfig configurations.
func (r *LogConfigResource) ListAll() (*LogConfigConfigList, error) {
	var list LogConfigConfigList
	if err := r.c.ReadQuery(BasePath+LogConfigEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogConfig configuration identified by id.
func (r *LogConfigResource) Get(id string) (*LogConfigConfig, error) {
	var item LogConfigConfig
	if err := r.c.ReadQuery(BasePath+LogConfigEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogConfig configuration.
func (r *LogConfigResource) Create(item LogConfigConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogConfigEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogConfig configuration identified by id.
func (r *LogConfigResource) Edit(id string, item LogConfigConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogConfigEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogConfig configuration identified by id.
func (r *LogConfigResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogConfigEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
