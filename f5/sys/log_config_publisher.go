// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogConfigPublisherConfigList holds a list of LogConfigPublisher configuration.
type LogConfigPublisherConfigList struct {
	Items    []LogConfigPublisherConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// LogConfigPublisherConfig holds the configuration of a single LogConfigPublisher.
type LogConfigPublisherConfig struct {
	Destinations []struct {
		Name      string `json:"name"`
		Partition string `json:"partition"`
	} `json:"destinations"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	SelfLink   string `json:"selfLink"`
}

// LogConfigPublisherEndpoint represents the REST resource for managing LogConfigPublisher.
const LogConfigPublisherEndpoint = "/log-config/publisher"

// LogConfigPublisherResource provides an API to manage LogConfigPublisher configurations.
type LogConfigPublisherResource struct {
	c *f5.Client
}

// ListAll  lists all the LogConfigPublisher configurations.
func (r *LogConfigPublisherResource) ListAll() (*LogConfigPublisherConfigList, error) {
	var list LogConfigPublisherConfigList
	if err := r.c.ReadQuery(BasePath+LogConfigPublisherEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogConfigPublisher configuration identified by id.
func (r *LogConfigPublisherResource) Get(id string) (*LogConfigPublisherConfig, error) {
	var item LogConfigPublisherConfig
	if err := r.c.ReadQuery(BasePath+LogConfigPublisherEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogConfigPublisher configuration.
func (r *LogConfigPublisherResource) Create(item LogConfigPublisherConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogConfigPublisherEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogConfigPublisher configuration identified by id.
func (r *LogConfigPublisherResource) Edit(id string, item LogConfigPublisherConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogConfigPublisherEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogConfigPublisher configuration identified by id.
func (r *LogConfigPublisherResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogConfigPublisherEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
