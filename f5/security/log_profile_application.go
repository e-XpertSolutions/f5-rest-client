// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogProfileApplicationConfigList holds a list of LogProfileApplication configuration.
type LogProfileApplicationConfigList struct {
	Items    []LogProfileApplicationConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// LogProfileApplicationConfig holds the configuration of a single LogProfileApplication.
type LogProfileApplicationConfig struct {
}

// LogProfileApplicationEndpoint represents the REST resource for managing LogProfileApplication.
const LogProfileApplicationEndpoint = "/log/profile_application"

// LogProfileApplicationResource provides an API to manage LogProfileApplication configurations.
type LogProfileApplicationResource struct {
	c *f5.Client
}

// ListAll  lists all the LogProfileApplication configurations.
func (r *LogProfileApplicationResource) ListAll() (*LogProfileApplicationConfigList, error) {
	var list LogProfileApplicationConfigList
	if err := r.c.ReadQuery(BasePath+LogProfileApplicationEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogProfileApplication configuration identified by id.
func (r *LogProfileApplicationResource) Get(id string) (*LogProfileApplicationConfig, error) {
	var item LogProfileApplicationConfig
	if err := r.c.ReadQuery(BasePath+LogProfileApplicationEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogProfileApplication configuration.
func (r *LogProfileApplicationResource) Create(item LogProfileApplicationConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogProfileApplicationEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogProfileApplication configuration identified by id.
func (r *LogProfileApplicationResource) Edit(id string, item LogProfileApplicationConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogProfileApplicationEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogProfileApplication configuration identified by id.
func (r *LogProfileApplicationResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogProfileApplicationEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
