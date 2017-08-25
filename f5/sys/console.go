// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ConsoleConfigList holds a list of Console configuration.
type ConsoleConfigList struct {
	Items    []ConsoleConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// ConsoleConfig holds the configuration of a single Console.
type ConsoleConfig struct {
}

// ConsoleEndpoint represents the REST resource for managing Console.
const ConsoleEndpoint = "/console"

// ConsoleResource provides an API to manage Console configurations.
type ConsoleResource struct {
	c *f5.Client
}

// ListAll  lists all the Console configurations.
func (r *ConsoleResource) ListAll() (*ConsoleConfigList, error) {
	var list ConsoleConfigList
	if err := r.c.ReadQuery(BasePath+ConsoleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Console configuration identified by id.
func (r *ConsoleResource) Get(id string) (*ConsoleConfig, error) {
	var item ConsoleConfig
	if err := r.c.ReadQuery(BasePath+ConsoleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Console configuration.
func (r *ConsoleResource) Create(item ConsoleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ConsoleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Console configuration identified by id.
func (r *ConsoleResource) Edit(id string, item ConsoleConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ConsoleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Console configuration identified by id.
func (r *ConsoleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ConsoleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
