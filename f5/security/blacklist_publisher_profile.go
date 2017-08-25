// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// BlacklistPublisherProfileConfigList holds a list of BlacklistPublisherProfile configuration.
type BlacklistPublisherProfileConfigList struct {
	Items    []BlacklistPublisherProfileConfig `json:"items"`
	Kind     string                            `json:"kind"`
	SelfLink string                            `json:"selflink"`
}

// BlacklistPublisherProfileConfig holds the configuration of a single BlacklistPublisherProfile.
type BlacklistPublisherProfileConfig struct {
}

// BlacklistPublisherProfileEndpoint represents the REST resource for managing BlacklistPublisherProfile.
const BlacklistPublisherProfileEndpoint = "/blacklist-publisher/profile"

// BlacklistPublisherProfileResource provides an API to manage BlacklistPublisherProfile configurations.
type BlacklistPublisherProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the BlacklistPublisherProfile configurations.
func (r *BlacklistPublisherProfileResource) ListAll() (*BlacklistPublisherProfileConfigList, error) {
	var list BlacklistPublisherProfileConfigList
	if err := r.c.ReadQuery(BasePath+BlacklistPublisherProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single BlacklistPublisherProfile configuration identified by id.
func (r *BlacklistPublisherProfileResource) Get(id string) (*BlacklistPublisherProfileConfig, error) {
	var item BlacklistPublisherProfileConfig
	if err := r.c.ReadQuery(BasePath+BlacklistPublisherProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new BlacklistPublisherProfile configuration.
func (r *BlacklistPublisherProfileResource) Create(item BlacklistPublisherProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+BlacklistPublisherProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a BlacklistPublisherProfile configuration identified by id.
func (r *BlacklistPublisherProfileResource) Edit(id string, item BlacklistPublisherProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+BlacklistPublisherProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single BlacklistPublisherProfile configuration identified by id.
func (r *BlacklistPublisherProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+BlacklistPublisherProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
