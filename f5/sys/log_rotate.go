// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogRotateConfigList holds a list of LogRotate configuration.
type LogRotateConfigList struct {
	Items    []LogRotateConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// LogRotateConfig holds the configuration of a single LogRotate.
type LogRotateConfig struct {
}

// LogRotateEndpoint represents the REST resource for managing LogRotate.
const LogRotateEndpoint = "/log-rotate"

// LogRotateResource provides an API to manage LogRotate configurations.
type LogRotateResource struct {
	c *f5.Client
}

// ListAll  lists all the LogRotate configurations.
func (r *LogRotateResource) ListAll() (*LogRotateConfigList, error) {
	var list LogRotateConfigList
	if err := r.c.ReadQuery(BasePath+LogRotateEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogRotate configuration identified by id.
func (r *LogRotateResource) Get(id string) (*LogRotateConfig, error) {
	var item LogRotateConfig
	if err := r.c.ReadQuery(BasePath+LogRotateEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogRotate configuration.
func (r *LogRotateResource) Create(item LogRotateConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogRotateEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogRotate configuration identified by id.
func (r *LogRotateResource) Edit(id string, item LogRotateConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogRotateEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogRotate configuration identified by id.
func (r *LogRotateResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogRotateEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
