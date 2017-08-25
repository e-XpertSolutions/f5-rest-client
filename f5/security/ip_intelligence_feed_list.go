// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPIntelligenceFeedListConfigList holds a list of IPIntelligenceFeedList configuration.
type IPIntelligenceFeedListConfigList struct {
	Items    []IPIntelligenceFeedListConfig `json:"items"`
	Kind     string                         `json:"kind"`
	SelfLink string                         `json:"selflink"`
}

// IPIntelligenceFeedListConfig holds the configuration of a single IPIntelligenceFeedList.
type IPIntelligenceFeedListConfig struct {
}

// IPIntelligenceFeedListEndpoint represents the REST resource for managing IPIntelligenceFeedList.
const IPIntelligenceFeedListEndpoint = "/ip-intelligence/feed-list"

// IPIntelligenceFeedListResource provides an API to manage IPIntelligenceFeedList configurations.
type IPIntelligenceFeedListResource struct {
	c *f5.Client
}

// ListAll  lists all the IPIntelligenceFeedList configurations.
func (r *IPIntelligenceFeedListResource) ListAll() (*IPIntelligenceFeedListConfigList, error) {
	var list IPIntelligenceFeedListConfigList
	if err := r.c.ReadQuery(BasePath+IPIntelligenceFeedListEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPIntelligenceFeedList configuration identified by id.
func (r *IPIntelligenceFeedListResource) Get(id string) (*IPIntelligenceFeedListConfig, error) {
	var item IPIntelligenceFeedListConfig
	if err := r.c.ReadQuery(BasePath+IPIntelligenceFeedListEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPIntelligenceFeedList configuration.
func (r *IPIntelligenceFeedListResource) Create(item IPIntelligenceFeedListConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPIntelligenceFeedListEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPIntelligenceFeedList configuration identified by id.
func (r *IPIntelligenceFeedListResource) Edit(id string, item IPIntelligenceFeedListConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPIntelligenceFeedListEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPIntelligenceFeedList configuration identified by id.
func (r *IPIntelligenceFeedListResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPIntelligenceFeedListEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
