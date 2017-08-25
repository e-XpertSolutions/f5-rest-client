// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogNetworkStorageFieldConfigList holds a list of LogNetworkStorageField configuration.
type LogNetworkStorageFieldConfigList struct {
	Items    []LogNetworkStorageFieldConfig `json:"items"`
	Kind     string                         `json:"kind"`
	SelfLink string                         `json:"selflink"`
}

// LogNetworkStorageFieldConfig holds the configuration of a single LogNetworkStorageField.
type LogNetworkStorageFieldConfig struct {
}

// LogNetworkStorageFieldEndpoint represents the REST resource for managing LogNetworkStorageField.
const LogNetworkStorageFieldEndpoint = "/log/network-storage-field"

// LogNetworkStorageFieldResource provides an API to manage LogNetworkStorageField configurations.
type LogNetworkStorageFieldResource struct {
	c *f5.Client
}

// ListAll  lists all the LogNetworkStorageField configurations.
func (r *LogNetworkStorageFieldResource) ListAll() (*LogNetworkStorageFieldConfigList, error) {
	var list LogNetworkStorageFieldConfigList
	if err := r.c.ReadQuery(BasePath+LogNetworkStorageFieldEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogNetworkStorageField configuration identified by id.
func (r *LogNetworkStorageFieldResource) Get(id string) (*LogNetworkStorageFieldConfig, error) {
	var item LogNetworkStorageFieldConfig
	if err := r.c.ReadQuery(BasePath+LogNetworkStorageFieldEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogNetworkStorageField configuration.
func (r *LogNetworkStorageFieldResource) Create(item LogNetworkStorageFieldConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogNetworkStorageFieldEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogNetworkStorageField configuration identified by id.
func (r *LogNetworkStorageFieldResource) Edit(id string, item LogNetworkStorageFieldConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogNetworkStorageFieldEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogNetworkStorageField configuration identified by id.
func (r *LogNetworkStorageFieldResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogNetworkStorageFieldEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
