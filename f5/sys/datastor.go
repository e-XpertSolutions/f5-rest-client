// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DataStorConfigList holds a list of DataStor configuration.
type DataStorConfigList struct {
	Items    []DataStorConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// DataStorConfig holds the configuration of a single DataStor.
type DataStorConfig struct {
}

// DataStorEndpoint represents the REST resource for managing DataStor.
const DataStorEndpoint = "/datastor"

// DataStorResource provides an API to manage DataStor configurations.
type DataStorResource struct {
	c *f5.Client
}

// ListAll  lists all the DataStor configurations.
func (r *DataStorResource) ListAll() (*DataStorConfigList, error) {
	var list DataStorConfigList
	if err := r.c.ReadQuery(BasePath+DataStorEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DataStor configuration identified by id.
func (r *DataStorResource) Get(id string) (*DataStorConfig, error) {
	var item DataStorConfig
	if err := r.c.ReadQuery(BasePath+DataStorEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DataStor configuration.
func (r *DataStorResource) Create(item DataStorConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DataStorEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DataStor configuration identified by id.
func (r *DataStorResource) Edit(id string, item DataStorConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DataStorEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DataStor configuration identified by id.
func (r *DataStorResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DataStorEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
