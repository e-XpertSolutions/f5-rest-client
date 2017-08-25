// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DatasyncConfigList holds a list of Datasync configuration.
type DatasyncConfigList struct {
	Items    []DatasyncConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// DatasyncConfig holds the configuration of a single Datasync.
type DatasyncConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// DatasyncEndpoint represents the REST resource for managing Datasync.
const DatasyncEndpoint = "/datasync"

// DatasyncResource provides an API to manage Datasync configurations.
type DatasyncResource struct {
	c *f5.Client
}

// ListAll  lists all the Datasync configurations.
func (r *DatasyncResource) ListAll() (*DatasyncConfigList, error) {
	var list DatasyncConfigList
	if err := r.c.ReadQuery(BasePath+DatasyncEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Datasync configuration identified by id.
func (r *DatasyncResource) Get(id string) (*DatasyncConfig, error) {
	var item DatasyncConfig
	if err := r.c.ReadQuery(BasePath+DatasyncEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Datasync configuration.
func (r *DatasyncResource) Create(item DatasyncConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DatasyncEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Datasync configuration identified by id.
func (r *DatasyncResource) Edit(id string, item DatasyncConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DatasyncEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Datasync configuration identified by id.
func (r *DatasyncResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DatasyncEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
