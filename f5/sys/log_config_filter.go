// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogConfigFilterConfigList holds a list of LogConfigFilter configuration.
type LogConfigFilterConfigList struct {
	Items    []LogConfigFilterConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// LogConfigFilterConfig holds the configuration of a single LogConfigFilter.
type LogConfigFilterConfig struct {
}

// LogConfigFilterEndpoint represents the REST resource for managing LogConfigFilter.
const LogConfigFilterEndpoint = "/log-config/filter"

// LogConfigFilterResource provides an API to manage LogConfigFilter configurations.
type LogConfigFilterResource struct {
	c *f5.Client
}

// ListAll  lists all the LogConfigFilter configurations.
func (r *LogConfigFilterResource) ListAll() (*LogConfigFilterConfigList, error) {
	var list LogConfigFilterConfigList
	if err := r.c.ReadQuery(BasePath+LogConfigFilterEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogConfigFilter configuration identified by id.
func (r *LogConfigFilterResource) Get(id string) (*LogConfigFilterConfig, error) {
	var item LogConfigFilterConfig
	if err := r.c.ReadQuery(BasePath+LogConfigFilterEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogConfigFilter configuration.
func (r *LogConfigFilterResource) Create(item LogConfigFilterConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogConfigFilterEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogConfigFilter configuration identified by id.
func (r *LogConfigFilterResource) Edit(id string, item LogConfigFilterConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogConfigFilterEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogConfigFilter configuration identified by id.
func (r *LogConfigFilterResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogConfigFilterEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
