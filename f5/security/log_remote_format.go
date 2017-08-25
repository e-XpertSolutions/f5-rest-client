// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogRemoteFormatConfigList holds a list of LogRemoteFormat configuration.
type LogRemoteFormatConfigList struct {
	Items    []LogRemoteFormatConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// LogRemoteFormatConfig holds the configuration of a single LogRemoteFormat.
type LogRemoteFormatConfig struct {
}

// LogRemoteFormatEndpoint represents the REST resource for managing LogRemoteFormat.
const LogRemoteFormatEndpoint = "/log/remote-format"

// LogRemoteFormatResource provides an API to manage LogRemoteFormat configurations.
type LogRemoteFormatResource struct {
	c *f5.Client
}

// ListAll  lists all the LogRemoteFormat configurations.
func (r *LogRemoteFormatResource) ListAll() (*LogRemoteFormatConfigList, error) {
	var list LogRemoteFormatConfigList
	if err := r.c.ReadQuery(BasePath+LogRemoteFormatEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogRemoteFormat configuration identified by id.
func (r *LogRemoteFormatResource) Get(id string) (*LogRemoteFormatConfig, error) {
	var item LogRemoteFormatConfig
	if err := r.c.ReadQuery(BasePath+LogRemoteFormatEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogRemoteFormat configuration.
func (r *LogRemoteFormatResource) Create(item LogRemoteFormatConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogRemoteFormatEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogRemoteFormat configuration identified by id.
func (r *LogRemoteFormatResource) Edit(id string, item LogRemoteFormatConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogRemoteFormatEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogRemoteFormat configuration identified by id.
func (r *LogRemoteFormatResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogRemoteFormatEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
