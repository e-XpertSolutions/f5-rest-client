// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// HTTPProfileConfigList holds a list of HTTPProfile configuration.
type HTTPProfileConfigList struct {
	Items    []HTTPProfileConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// HTTPProfileConfig holds the configuration of a single HTTPProfile.
type HTTPProfileConfig struct {
}

// HTTPProfileEndpoint represents the REST resource for managing HTTPProfile.
const HTTPProfileEndpoint = "/http/profile"

// HTTPProfileResource provides an API to manage HTTPProfile configurations.
type HTTPProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the HTTPProfile configurations.
func (r *HTTPProfileResource) ListAll() (*HTTPProfileConfigList, error) {
	var list HTTPProfileConfigList
	if err := r.c.ReadQuery(BasePath+HTTPProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single HTTPProfile configuration identified by id.
func (r *HTTPProfileResource) Get(id string) (*HTTPProfileConfig, error) {
	var item HTTPProfileConfig
	if err := r.c.ReadQuery(BasePath+HTTPProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new HTTPProfile configuration.
func (r *HTTPProfileResource) Create(item HTTPProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+HTTPProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a HTTPProfile configuration identified by id.
func (r *HTTPProfileResource) Edit(id string, item HTTPProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+HTTPProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single HTTPProfile configuration identified by id.
func (r *HTTPProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+HTTPProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
