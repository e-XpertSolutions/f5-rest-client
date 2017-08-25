// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPIntelligenseBlacklistCategoryConfigList holds a list of IPIntelligenseBlacklistCategory configuration.
type IPIntelligenseBlacklistCategoryConfigList struct {
	Items    []IPIntelligenseBlacklistCategoryConfig `json:"items"`
	Kind     string                                  `json:"kind"`
	SelfLink string                                  `json:"selflink"`
}

// IPIntelligenseBlacklistCategoryConfig holds the configuration of a single IPIntelligenseBlacklistCategory.
type IPIntelligenseBlacklistCategoryConfig struct {
	BlMatchDirection string `json:"blMatchDirection"`
	Description      string `json:"description"`
	FullPath         string `json:"fullPath"`
	Generation       int    `json:"generation"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
	Partition        string `json:"partition"`
	SelfLink         string `json:"selfLink"`
}

// IPIntelligenseBlacklistCategoryEndpoint represents the REST resource for managing IPIntelligenseBlacklistCategory.
const IPIntelligenseBlacklistCategoryEndpoint = "/ip-intelligence/blacklist-category"

// IPIntelligenseBlacklistCategoryResource provides an API to manage IPIntelligenseBlacklistCategory configurations.
type IPIntelligenseBlacklistCategoryResource struct {
	c *f5.Client
}

// ListAll  lists all the IPIntelligenseBlacklistCategory configurations.
func (r *IPIntelligenseBlacklistCategoryResource) ListAll() (*IPIntelligenseBlacklistCategoryConfigList, error) {
	var list IPIntelligenseBlacklistCategoryConfigList
	if err := r.c.ReadQuery(BasePath+IPIntelligenseBlacklistCategoryEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPIntelligenseBlacklistCategory configuration identified by id.
func (r *IPIntelligenseBlacklistCategoryResource) Get(id string) (*IPIntelligenseBlacklistCategoryConfig, error) {
	var item IPIntelligenseBlacklistCategoryConfig
	if err := r.c.ReadQuery(BasePath+IPIntelligenseBlacklistCategoryEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPIntelligenseBlacklistCategory configuration.
func (r *IPIntelligenseBlacklistCategoryResource) Create(item IPIntelligenseBlacklistCategoryConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPIntelligenseBlacklistCategoryEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPIntelligenseBlacklistCategory configuration identified by id.
func (r *IPIntelligenseBlacklistCategoryResource) Edit(id string, item IPIntelligenseBlacklistCategoryConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPIntelligenseBlacklistCategoryEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPIntelligenseBlacklistCategory configuration identified by id.
func (r *IPIntelligenseBlacklistCategoryResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPIntelligenseBlacklistCategoryEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
