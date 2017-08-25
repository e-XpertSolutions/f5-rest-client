// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ProvisionConfigList holds a list of Provision configuration.
type ProvisionConfigList struct {
	Items    []ProvisionConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// ProvisionConfig holds the configuration of a single Provision.
type ProvisionConfig struct {
	CPURatio    int    `json:"cpuRatio"`
	DiskRatio   int    `json:"diskRatio"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Kind        string `json:"kind"`
	Level       string `json:"level"`
	MemoryRatio int    `json:"memoryRatio"`
	Name        string `json:"name"`
	SelfLink    string `json:"selfLink"`
}

// ProvisionEndpoint represents the REST resource for managing Provision.
const ProvisionEndpoint = "/provision"

// ProvisionResource provides an API to manage Provision configurations.
type ProvisionResource struct {
	c *f5.Client
}

// ListAll  lists all the Provision configurations.
func (r *ProvisionResource) ListAll() (*ProvisionConfigList, error) {
	var list ProvisionConfigList
	if err := r.c.ReadQuery(BasePath+ProvisionEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Provision configuration identified by id.
func (r *ProvisionResource) Get(id string) (*ProvisionConfig, error) {
	var item ProvisionConfig
	if err := r.c.ReadQuery(BasePath+ProvisionEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Provision configuration.
func (r *ProvisionResource) Create(item ProvisionConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ProvisionEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Provision configuration identified by id.
func (r *ProvisionResource) Edit(id string, item ProvisionConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ProvisionEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Provision configuration identified by id.
func (r *ProvisionResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ProvisionEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
