// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// HTTPFileTypeConfigList holds a list of HTTPFileType configuration.
type HTTPFileTypeConfigList struct {
	Items    []HTTPFileTypeConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// HTTPFileTypeConfig holds the configuration of a single HTTPFileType.
type HTTPFileTypeConfig struct {
}

// HTTPFileTypeEndpoint represents the REST resource for managing HTTPFileType.
const HTTPFileTypeEndpoint = "/http/file-type"

// HTTPFileTypeResource provides an API to manage HTTPFileType configurations.
type HTTPFileTypeResource struct {
	c *f5.Client
}

// ListAll  lists all the HTTPFileType configurations.
func (r *HTTPFileTypeResource) ListAll() (*HTTPFileTypeConfigList, error) {
	var list HTTPFileTypeConfigList
	if err := r.c.ReadQuery(BasePath+HTTPFileTypeEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single HTTPFileType configuration identified by id.
func (r *HTTPFileTypeResource) Get(id string) (*HTTPFileTypeConfig, error) {
	var item HTTPFileTypeConfig
	if err := r.c.ReadQuery(BasePath+HTTPFileTypeEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new HTTPFileType configuration.
func (r *HTTPFileTypeResource) Create(item HTTPFileTypeConfig) error {
	if err := r.c.ModQuery("POST", BasePath+HTTPFileTypeEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a HTTPFileType configuration identified by id.
func (r *HTTPFileTypeResource) Edit(id string, item HTTPFileTypeConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+HTTPFileTypeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single HTTPFileType configuration identified by id.
func (r *HTTPFileTypeResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+HTTPFileTypeEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
