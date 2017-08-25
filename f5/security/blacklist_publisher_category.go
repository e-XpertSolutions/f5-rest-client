// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// BlacklistPublisherCategoryConfigList holds a list of BlacklistPublisherCategory configuration.
type BlacklistPublisherCategoryConfigList struct {
	Items    []BlacklistPublisherCategoryConfig `json:"items"`
	Kind     string                             `json:"kind"`
	SelfLink string                             `json:"selflink"`
}

// BlacklistPublisherCategoryConfig holds the configuration of a single BlacklistPublisherCategory.
type BlacklistPublisherCategoryConfig struct {
}

// BlacklistPublisherCategoryEndpoint represents the REST resource for managing BlacklistPublisherCategory.
const BlacklistPublisherCategoryEndpoint = "/blacklist-publisher/category"

// BlacklistPublisherCategoryResource provides an API to manage BlacklistPublisherCategory configurations.
type BlacklistPublisherCategoryResource struct {
	c *f5.Client
}

// ListAll  lists all the BlacklistPublisherCategory configurations.
func (r *BlacklistPublisherCategoryResource) ListAll() (*BlacklistPublisherCategoryConfigList, error) {
	var list BlacklistPublisherCategoryConfigList
	if err := r.c.ReadQuery(BasePath+BlacklistPublisherCategoryEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single BlacklistPublisherCategory configuration identified by id.
func (r *BlacklistPublisherCategoryResource) Get(id string) (*BlacklistPublisherCategoryConfig, error) {
	var item BlacklistPublisherCategoryConfig
	if err := r.c.ReadQuery(BasePath+BlacklistPublisherCategoryEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new BlacklistPublisherCategory configuration.
func (r *BlacklistPublisherCategoryResource) Create(item BlacklistPublisherCategoryConfig) error {
	if err := r.c.ModQuery("POST", BasePath+BlacklistPublisherCategoryEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a BlacklistPublisherCategory configuration identified by id.
func (r *BlacklistPublisherCategoryResource) Edit(id string, item BlacklistPublisherCategoryConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+BlacklistPublisherCategoryEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single BlacklistPublisherCategory configuration identified by id.
func (r *BlacklistPublisherCategoryResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+BlacklistPublisherCategoryEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
