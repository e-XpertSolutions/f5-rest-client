// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// HTTPMandatoryHeaderConfigList holds a list of HTTPMandatoryHeader configuration.
type HTTPMandatoryHeaderConfigList struct {
	Items    []HTTPMandatoryHeaderConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// HTTPMandatoryHeaderConfig holds the configuration of a single HTTPMandatoryHeader.
type HTTPMandatoryHeaderConfig struct {
}

// HTTPMandatoryHeaderEndpoint represents the REST resource for managing HTTPMandatoryHeader.
const HTTPMandatoryHeaderEndpoint = "/http/mandatory-header"

// HTTPMandatoryHeaderResource provides an API to manage HTTPMandatoryHeader configurations.
type HTTPMandatoryHeaderResource struct {
	c *f5.Client
}

// ListAll  lists all the HTTPMandatoryHeader configurations.
func (r *HTTPMandatoryHeaderResource) ListAll() (*HTTPMandatoryHeaderConfigList, error) {
	var list HTTPMandatoryHeaderConfigList
	if err := r.c.ReadQuery(BasePath+HTTPMandatoryHeaderEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single HTTPMandatoryHeader configuration identified by id.
func (r *HTTPMandatoryHeaderResource) Get(id string) (*HTTPMandatoryHeaderConfig, error) {
	var item HTTPMandatoryHeaderConfig
	if err := r.c.ReadQuery(BasePath+HTTPMandatoryHeaderEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new HTTPMandatoryHeader configuration.
func (r *HTTPMandatoryHeaderResource) Create(item HTTPMandatoryHeaderConfig) error {
	if err := r.c.ModQuery("POST", BasePath+HTTPMandatoryHeaderEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a HTTPMandatoryHeader configuration identified by id.
func (r *HTTPMandatoryHeaderResource) Edit(id string, item HTTPMandatoryHeaderConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+HTTPMandatoryHeaderEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single HTTPMandatoryHeader configuration identified by id.
func (r *HTTPMandatoryHeaderResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+HTTPMandatoryHeaderEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
