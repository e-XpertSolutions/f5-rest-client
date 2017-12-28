// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// SoftwareHotfixConfigList holds a list of SoftwareHotfix configuration.
type SoftwareHotfixConfigList struct {
	Items    []SoftwareHotfixConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// SoftwareHotfixConfig holds the configuration of a single SoftwareHotfix.
type SoftwareHotfixConfig struct {
	Build      string `json:"build"`
	Checksum   string `json:"checksum"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	ID         string `json:"id"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Product    string `json:"product"`
	SelfLink   string `json:"selfLink"`
	Title      string `json:"title"`
	Verified   string `json:"verified"`
	Version    string `json:"version"`
}

// SoftwareHotfixEndpoint represents the REST resource for managing SoftwareHotfix.
const SoftwareHotfixEndpoint = "/software/hotfix"

// SoftwareHotfixResource provides an API to manage SoftwareHotfix configurations.
type SoftwareHotfixResource struct {
	c *f5.Client
}

// ListAll  lists all the SoftwareHotfix configurations.
func (r *SoftwareHotfixResource) ListAll() (*SoftwareHotfixConfigList, error) {
	var list SoftwareHotfixConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareHotfixEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SoftwareHotfix configuration identified by id.
func (r *SoftwareHotfixResource) Get(id string) (*SoftwareHotfixConfig, error) {
	var item SoftwareHotfixConfig
	if err := r.c.ReadQuery(BasePath+SoftwareHotfixEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SoftwareHotfix configuration.
func (r *SoftwareHotfixResource) Create(item SoftwareHotfixConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SoftwareHotfixEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SoftwareHotfix configuration identified by id.
func (r *SoftwareHotfixResource) Edit(id string, item SoftwareHotfixConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SoftwareHotfixEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SoftwareHotfix configuration identified by id.
func (r *SoftwareHotfixResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SoftwareHotfixEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
