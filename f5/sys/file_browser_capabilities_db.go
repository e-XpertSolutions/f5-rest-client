// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileBrowserCapabilitiesDBConfigList holds a list of FileBrowserCapabilitiesDB configuration.
type FileBrowserCapabilitiesDBConfigList struct {
	Items    []FileBrowserCapabilitiesDBConfig `json:"items"`
	Kind     string                            `json:"kind"`
	SelfLink string                            `json:"selflink"`
}

// FileBrowserCapabilitiesDBConfig holds the configuration of a single FileBrowserCapabilitiesDB.
type FileBrowserCapabilitiesDBConfig struct {
}

// FileBrowserCapabilitiesDBEndpoint represents the REST resource for managing FileBrowserCapabilitiesDB.
const FileBrowserCapabilitiesDBEndpoint = "/file/browser-capabilities-db"

// FileBrowserCapabilitiesDBResource provides an API to manage FileBrowserCapabilitiesDB configurations.
type FileBrowserCapabilitiesDBResource struct {
	c *f5.Client
}

// ListAll  lists all the FileBrowserCapabilitiesDB configurations.
func (r *FileBrowserCapabilitiesDBResource) ListAll() (*FileBrowserCapabilitiesDBConfigList, error) {
	var list FileBrowserCapabilitiesDBConfigList
	if err := r.c.ReadQuery(BasePath+FileBrowserCapabilitiesDBEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileBrowserCapabilitiesDB configuration identified by id.
func (r *FileBrowserCapabilitiesDBResource) Get(id string) (*FileBrowserCapabilitiesDBConfig, error) {
	var item FileBrowserCapabilitiesDBConfig
	if err := r.c.ReadQuery(BasePath+FileBrowserCapabilitiesDBEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileBrowserCapabilitiesDB configuration.
func (r *FileBrowserCapabilitiesDBResource) Create(item FileBrowserCapabilitiesDBConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileBrowserCapabilitiesDBEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileBrowserCapabilitiesDB configuration identified by id.
func (r *FileBrowserCapabilitiesDBResource) Edit(id string, item FileBrowserCapabilitiesDBConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileBrowserCapabilitiesDBEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileBrowserCapabilitiesDB configuration identified by id.
func (r *FileBrowserCapabilitiesDBResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileBrowserCapabilitiesDBEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
