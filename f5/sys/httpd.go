// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// HTTPDConfigList holds a list of HTTPD configuration.
type HTTPDConfigList struct {
	Items    []HTTPDConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// HTTPDConfig holds the configuration of a single HTTPD.
type HTTPDConfig struct {
}

// HTTPDEndpoint represents the REST resource for managing HTTPD.
const HTTPDEndpoint = "/httpd"

// HTTPDResource provides an API to manage HTTPD configurations.
type HTTPDResource struct {
	c *f5.Client
}

// ListAll  lists all the HTTPD configurations.
func (r *HTTPDResource) ListAll() (*HTTPDConfigList, error) {
	var list HTTPDConfigList
	if err := r.c.ReadQuery(BasePath+HTTPDEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single HTTPD configuration identified by id.
func (r *HTTPDResource) Get(id string) (*HTTPDConfig, error) {
	var item HTTPDConfig
	if err := r.c.ReadQuery(BasePath+HTTPDEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new HTTPD configuration.
func (r *HTTPDResource) Create(item HTTPDConfig) error {
	if err := r.c.ModQuery("POST", BasePath+HTTPDEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a HTTPD configuration identified by id.
func (r *HTTPDResource) Edit(id string, item HTTPDConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+HTTPDEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single HTTPD configuration identified by id.
func (r *HTTPDResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+HTTPDEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
