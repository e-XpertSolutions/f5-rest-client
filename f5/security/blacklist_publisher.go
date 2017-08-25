// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// BlacklistPublisherConfigList holds a list of BlacklistPublisher configuration.
type BlacklistPublisherConfigList struct {
	Items    []BlacklistPublisherConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// BlacklistPublisherConfig holds the configuration of a single BlacklistPublisher.
type BlacklistPublisherConfig struct {
}

// BlacklistPublisherEndpoint represents the REST resource for managing BlacklistPublisher.
const BlacklistPublisherEndpoint = "/blacklist-publisher"

// BlacklistPublisherResource provides an API to manage BlacklistPublisher configurations.
type BlacklistPublisherResource struct {
	c *f5.Client
}

// ListAll  lists all the BlacklistPublisher configurations.
func (r *BlacklistPublisherResource) ListAll() (*BlacklistPublisherConfigList, error) {
	var list BlacklistPublisherConfigList
	if err := r.c.ReadQuery(BasePath+BlacklistPublisherEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single BlacklistPublisher configuration identified by id.
func (r *BlacklistPublisherResource) Get(id string) (*BlacklistPublisherConfig, error) {
	var item BlacklistPublisherConfig
	if err := r.c.ReadQuery(BasePath+BlacklistPublisherEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new BlacklistPublisher configuration.
func (r *BlacklistPublisherResource) Create(item BlacklistPublisherConfig) error {
	if err := r.c.ModQuery("POST", BasePath+BlacklistPublisherEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a BlacklistPublisher configuration identified by id.
func (r *BlacklistPublisherResource) Edit(id string, item BlacklistPublisherConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+BlacklistPublisherEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single BlacklistPublisher configuration identified by id.
func (r *BlacklistPublisherResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+BlacklistPublisherEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
