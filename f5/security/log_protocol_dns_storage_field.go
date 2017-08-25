// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogProtocolDNSStorageFieldConfigList holds a list of LogProtocolDNSStorageField configuration.
type LogProtocolDNSStorageFieldConfigList struct {
	Items    []LogProtocolDNSStorageFieldConfig `json:"items"`
	Kind     string                             `json:"kind"`
	SelfLink string                             `json:"selflink"`
}

// LogProtocolDNSStorageFieldConfig holds the configuration of a single LogProtocolDNSStorageField.
type LogProtocolDNSStorageFieldConfig struct {
}

// LogProtocolDNSStorageFieldEndpoint represents the REST resource for managing LogProtocolDNSStorageField.
const LogProtocolDNSStorageFieldEndpoint = "/log/protocol-dns-storage-field"

// LogProtocolDNSStorageFieldResource provides an API to manage LogProtocolDNSStorageField configurations.
type LogProtocolDNSStorageFieldResource struct {
	c *f5.Client
}

// ListAll  lists all the LogProtocolDNSStorageField configurations.
func (r *LogProtocolDNSStorageFieldResource) ListAll() (*LogProtocolDNSStorageFieldConfigList, error) {
	var list LogProtocolDNSStorageFieldConfigList
	if err := r.c.ReadQuery(BasePath+LogProtocolDNSStorageFieldEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogProtocolDNSStorageField configuration identified by id.
func (r *LogProtocolDNSStorageFieldResource) Get(id string) (*LogProtocolDNSStorageFieldConfig, error) {
	var item LogProtocolDNSStorageFieldConfig
	if err := r.c.ReadQuery(BasePath+LogProtocolDNSStorageFieldEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogProtocolDNSStorageField configuration.
func (r *LogProtocolDNSStorageFieldResource) Create(item LogProtocolDNSStorageFieldConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogProtocolDNSStorageFieldEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogProtocolDNSStorageField configuration identified by id.
func (r *LogProtocolDNSStorageFieldResource) Edit(id string, item LogProtocolDNSStorageFieldConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogProtocolDNSStorageFieldEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogProtocolDNSStorageField configuration identified by id.
func (r *LogProtocolDNSStorageFieldResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogProtocolDNSStorageFieldEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
