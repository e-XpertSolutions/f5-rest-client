// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogProtocolSIPStorageFieldConfigList holds a list of LogProtocolSIPStorageField configuration.
type LogProtocolSIPStorageFieldConfigList struct {
	Items    []LogProtocolSIPStorageFieldConfig `json:"items"`
	Kind     string                             `json:"kind"`
	SelfLink string                             `json:"selflink"`
}

// LogProtocolSIPStorageFieldConfig holds the configuration of a single LogProtocolSIPStorageField.
type LogProtocolSIPStorageFieldConfig struct {
}

// LogProtocolSIPStorageFieldEndpoint represents the REST resource for managing LogProtocolSIPStorageField.
const LogProtocolSIPStorageFieldEndpoint = "/log/protocol-sip-storage-field"

// LogProtocolSIPStorageFieldResource provides an API to manage LogProtocolSIPStorageField configurations.
type LogProtocolSIPStorageFieldResource struct {
	c *f5.Client
}

// ListAll  lists all the LogProtocolSIPStorageField configurations.
func (r *LogProtocolSIPStorageFieldResource) ListAll() (*LogProtocolSIPStorageFieldConfigList, error) {
	var list LogProtocolSIPStorageFieldConfigList
	if err := r.c.ReadQuery(BasePath+LogProtocolSIPStorageFieldEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogProtocolSIPStorageField configuration identified by id.
func (r *LogProtocolSIPStorageFieldResource) Get(id string) (*LogProtocolSIPStorageFieldConfig, error) {
	var item LogProtocolSIPStorageFieldConfig
	if err := r.c.ReadQuery(BasePath+LogProtocolSIPStorageFieldEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogProtocolSIPStorageField configuration.
func (r *LogProtocolSIPStorageFieldResource) Create(item LogProtocolSIPStorageFieldConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogProtocolSIPStorageFieldEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogProtocolSIPStorageField configuration identified by id.
func (r *LogProtocolSIPStorageFieldResource) Edit(id string, item LogProtocolSIPStorageFieldConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogProtocolSIPStorageFieldEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogProtocolSIPStorageField configuration identified by id.
func (r *LogProtocolSIPStorageFieldResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogProtocolSIPStorageFieldEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
