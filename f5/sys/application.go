// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ApplicationConfigList holds a list of Application configuration.
type ApplicationConfigList struct {
	Items    []ApplicationConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// ApplicationConfig holds the configuration of a single Application.
type ApplicationConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// ApplicationEndpoint represents the REST resource for managing Application.
const ApplicationEndpoint = "/application"

// ApplicationResource provides an API to manage Application configurations.
type ApplicationResource struct {
	c *f5.Client
}

// ListAll  lists all the Application configurations.
func (r *ApplicationResource) ListAll() (*ApplicationConfigList, error) {
	var list ApplicationConfigList
	if err := r.c.ReadQuery(BasePath+ApplicationEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Application configuration identified by id.
func (r *ApplicationResource) Get(id string) (*ApplicationConfig, error) {
	var item ApplicationConfig
	if err := r.c.ReadQuery(BasePath+ApplicationEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Application configuration.
func (r *ApplicationResource) Create(item ApplicationConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ApplicationEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Application configuration identified by id.
func (r *ApplicationResource) Edit(id string, item ApplicationConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ApplicationEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Application configuration identified by id.
func (r *ApplicationResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ApplicationEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
