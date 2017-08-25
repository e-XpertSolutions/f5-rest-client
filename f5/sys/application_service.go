// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ApplicationServiceConfigList holds a list of ApplicationService configuration.
type ApplicationServiceConfigList struct {
	Items    []ApplicationServiceConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// ApplicationServiceConfig holds the configuration of a single ApplicationService.
type ApplicationServiceConfig struct {
}

// ApplicationServiceEndpoint represents the REST resource for managing ApplicationService.
const ApplicationServiceEndpoint = "/application/service"

// ApplicationServiceResource provides an API to manage ApplicationService configurations.
type ApplicationServiceResource struct {
	c *f5.Client
}

// ListAll  lists all the ApplicationService configurations.
func (r *ApplicationServiceResource) ListAll() (*ApplicationServiceConfigList, error) {
	var list ApplicationServiceConfigList
	if err := r.c.ReadQuery(BasePath+ApplicationServiceEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ApplicationService configuration identified by id.
func (r *ApplicationServiceResource) Get(id string) (*ApplicationServiceConfig, error) {
	var item ApplicationServiceConfig
	if err := r.c.ReadQuery(BasePath+ApplicationServiceEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ApplicationService configuration.
func (r *ApplicationServiceResource) Create(item ApplicationServiceConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ApplicationServiceEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ApplicationService configuration identified by id.
func (r *ApplicationServiceResource) Edit(id string, item ApplicationServiceConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ApplicationServiceEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ApplicationService configuration identified by id.
func (r *ApplicationServiceResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ApplicationServiceEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
