// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DatasyncUpdateFileConfigList holds a list of DatasyncUpdateFile configuration.
type DatasyncUpdateFileConfigList struct {
	Items    []DatasyncUpdateFileConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// DatasyncUpdateFileConfig holds the configuration of a single DatasyncUpdateFile.
type DatasyncUpdateFileConfig struct {
}

// DatasyncUpdateFileEndpoint represents the REST resource for managing DatasyncUpdateFile.
const DatasyncUpdateFileEndpoint = "/datasync/update-file"

// DatasyncUpdateFileResource provides an API to manage DatasyncUpdateFile configurations.
type DatasyncUpdateFileResource struct {
	c *f5.Client
}

// ListAll  lists all the DatasyncUpdateFile configurations.
func (r *DatasyncUpdateFileResource) ListAll() (*DatasyncUpdateFileConfigList, error) {
	var list DatasyncUpdateFileConfigList
	if err := r.c.ReadQuery(BasePath+DatasyncUpdateFileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DatasyncUpdateFile configuration identified by id.
func (r *DatasyncUpdateFileResource) Get(id string) (*DatasyncUpdateFileConfig, error) {
	var item DatasyncUpdateFileConfig
	if err := r.c.ReadQuery(BasePath+DatasyncUpdateFileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DatasyncUpdateFile configuration.
func (r *DatasyncUpdateFileResource) Create(item DatasyncUpdateFileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DatasyncUpdateFileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DatasyncUpdateFile configuration identified by id.
func (r *DatasyncUpdateFileResource) Edit(id string, item DatasyncUpdateFileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DatasyncUpdateFileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DatasyncUpdateFile configuration identified by id.
func (r *DatasyncUpdateFileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DatasyncUpdateFileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
