// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// BlacklistPublisherBlacklistPublisherStatsConfigList holds a list of BlacklistPublisherBlacklistPublisherStats configuration.
type BlacklistPublisherBlacklistPublisherStatsConfigList struct {
	Items    []BlacklistPublisherBlacklistPublisherStatsConfig `json:"items"`
	Kind     string                                            `json:"kind"`
	SelfLink string                                            `json:"selflink"`
}

// BlacklistPublisherBlacklistPublisherStatsConfig holds the configuration of a single BlacklistPublisherBlacklistPublisherStats.
type BlacklistPublisherBlacklistPublisherStatsConfig struct {
}

// BlacklistPublisherBlacklistPublisherStatsEndpoint represents the REST resource for managing BlacklistPublisherBlacklistPublisherStats.
const BlacklistPublisherBlacklistPublisherStatsEndpoint = "/blacklist-publisher/blacklist-publisher-stats"

// BlacklistPublisherBlacklistPublisherStatsResource provides an API to manage BlacklistPublisherBlacklistPublisherStats configurations.
type BlacklistPublisherBlacklistPublisherStatsResource struct {
	c *f5.Client
}

// ListAll  lists all the BlacklistPublisherBlacklistPublisherStats configurations.
func (r *BlacklistPublisherBlacklistPublisherStatsResource) ListAll() (*BlacklistPublisherBlacklistPublisherStatsConfigList, error) {
	var list BlacklistPublisherBlacklistPublisherStatsConfigList
	if err := r.c.ReadQuery(BasePath+BlacklistPublisherBlacklistPublisherStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single BlacklistPublisherBlacklistPublisherStats configuration identified by id.
func (r *BlacklistPublisherBlacklistPublisherStatsResource) Get(id string) (*BlacklistPublisherBlacklistPublisherStatsConfig, error) {
	var item BlacklistPublisherBlacklistPublisherStatsConfig
	if err := r.c.ReadQuery(BasePath+BlacklistPublisherBlacklistPublisherStatsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new BlacklistPublisherBlacklistPublisherStats configuration.
func (r *BlacklistPublisherBlacklistPublisherStatsResource) Create(item BlacklistPublisherBlacklistPublisherStatsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+BlacklistPublisherBlacklistPublisherStatsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a BlacklistPublisherBlacklistPublisherStats configuration identified by id.
func (r *BlacklistPublisherBlacklistPublisherStatsResource) Edit(id string, item BlacklistPublisherBlacklistPublisherStatsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+BlacklistPublisherBlacklistPublisherStatsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single BlacklistPublisherBlacklistPublisherStats configuration identified by id.
func (r *BlacklistPublisherBlacklistPublisherStatsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+BlacklistPublisherBlacklistPublisherStatsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
