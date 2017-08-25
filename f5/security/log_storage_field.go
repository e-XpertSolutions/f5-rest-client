// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogStorageFieldConfigList holds a list of LogStorageField configuration.
type LogStorageFieldConfigList struct {
	Items    []LogStorageFieldConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// LogStorageFieldConfig holds the configuration of a single LogStorageField.
type LogStorageFieldConfig struct {
	Format     string `json:"format"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	ID         int    `json:"id"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	SelfLink   string `json:"selfLink"`
}

// LogStorageFieldEndpoint represents the REST resource for managing LogStorageField.
const LogStorageFieldEndpoint = "/log/storage-field"

// LogStorageFieldResource provides an API to manage LogStorageField configurations.
type LogStorageFieldResource struct {
	c *f5.Client
}

// ListAll  lists all the LogStorageField configurations.
func (r *LogStorageFieldResource) ListAll() (*LogStorageFieldConfigList, error) {
	var list LogStorageFieldConfigList
	if err := r.c.ReadQuery(BasePath+LogStorageFieldEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogStorageField configuration identified by id.
func (r *LogStorageFieldResource) Get(id string) (*LogStorageFieldConfig, error) {
	var item LogStorageFieldConfig
	if err := r.c.ReadQuery(BasePath+LogStorageFieldEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogStorageField configuration.
func (r *LogStorageFieldResource) Create(item LogStorageFieldConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogStorageFieldEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogStorageField configuration identified by id.
func (r *LogStorageFieldResource) Edit(id string, item LogStorageFieldConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogStorageFieldEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogStorageField configuration identified by id.
func (r *LogStorageFieldResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogStorageFieldEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
