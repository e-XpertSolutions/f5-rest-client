// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ECMConfigConfigList holds a list of ECMConfig configuration.
type ECMConfigConfigList struct {
	Items    []ECMConfigConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// ECMConfigConfig holds the configuration of a single ECMConfig.
type ECMConfigConfig struct {
}

// ECMConfigEndpoint represents the REST resource for managing ECMConfig.
const ECMConfigEndpoint = "/ecm/config"

// ECMConfigResource provides an API to manage ECMConfig configurations.
type ECMConfigResource struct {
	c *f5.Client
}

// ListAll  lists all the ECMConfig configurations.
func (r *ECMConfigResource) ListAll() (*ECMConfigConfigList, error) {
	var list ECMConfigConfigList
	if err := r.c.ReadQuery(BasePath+ECMConfigEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ECMConfig configuration identified by id.
func (r *ECMConfigResource) Get(id string) (*ECMConfigConfig, error) {
	var item ECMConfigConfig
	if err := r.c.ReadQuery(BasePath+ECMConfigEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ECMConfig configuration.
func (r *ECMConfigResource) Create(item ECMConfigConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ECMConfigEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ECMConfig configuration identified by id.
func (r *ECMConfigResource) Edit(id string, item ECMConfigConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ECMConfigEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ECMConfig configuration identified by id.
func (r *ECMConfigResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ECMConfigEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
