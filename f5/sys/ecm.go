// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ECMConfigList holds a list of ECM configuration.
type ECMConfigList struct {
	Items    []ECMConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// ECMConfig holds the configuration of a single ECM.
type ECMConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// ECMEndpoint represents the REST resource for managing ECM.
const ECMEndpoint = "/ecm"

// ECMResource provides an API to manage ECM configurations.
type ECMResource struct {
	c *f5.Client
}

// ListAll  lists all the ECM configurations.
func (r *ECMResource) ListAll() (*ECMConfigList, error) {
	var list ECMConfigList
	if err := r.c.ReadQuery(BasePath+ECMEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ECM configuration identified by id.
func (r *ECMResource) Get(id string) (*ECMConfig, error) {
	var item ECMConfig
	if err := r.c.ReadQuery(BasePath+ECMEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ECM configuration.
func (r *ECMResource) Create(item ECMConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ECMEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ECM configuration identified by id.
func (r *ECMResource) Edit(id string, item ECMConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ECMEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ECM configuration identified by id.
func (r *ECMResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ECMEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
