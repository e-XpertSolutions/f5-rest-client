// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogConfigList holds a list of Log configuration.
type LogConfigList struct {
	Items    []LogConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// LogConfig holds the configuration of a single Log.
type LogConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// LogEndpoint represents the REST resource for managing Log.
const LogEndpoint = "/log"

// LogResource provides an API to manage Log configurations.
type LogResource struct {
	c *f5.Client
}

// ListAll  lists all the Log configurations.
func (r *LogResource) ListAll() (*LogConfigList, error) {
	var list LogConfigList
	if err := r.c.ReadQuery(BasePath+LogEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Log configuration identified by id.
func (r *LogResource) Get(id string) (*LogConfig, error) {
	var item LogConfig
	if err := r.c.ReadQuery(BasePath+LogEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Log configuration.
func (r *LogResource) Create(item LogConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Log configuration identified by id.
func (r *LogResource) Edit(id string, item LogConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Log configuration identified by id.
func (r *LogResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
