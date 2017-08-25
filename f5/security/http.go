// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// HTTPConfigList holds a list of HTTP configuration.
type HTTPConfigList struct {
	Items    []HTTPConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// HTTPConfig holds the configuration of a single HTTP.
type HTTPConfig struct {
}

// HTTPEndpoint represents the REST resource for managing HTTP.
const HTTPEndpoint = "/http"

// HTTPResource provides an API to manage HTTP configurations.
type HTTPResource struct {
	c *f5.Client
}

// ListAll  lists all the HTTP configurations.
func (r *HTTPResource) ListAll() (*HTTPConfigList, error) {
	var list HTTPConfigList
	if err := r.c.ReadQuery(BasePath+HTTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single HTTP configuration identified by id.
func (r *HTTPResource) Get(id string) (*HTTPConfig, error) {
	var item HTTPConfig
	if err := r.c.ReadQuery(BasePath+HTTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new HTTP configuration.
func (r *HTTPResource) Create(item HTTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+HTTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a HTTP configuration identified by id.
func (r *HTTPResource) Edit(id string, item HTTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+HTTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single HTTP configuration identified by id.
func (r *HTTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+HTTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
