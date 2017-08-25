// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AlertConfigList holds a list of Alert configuration.
type AlertConfigList struct {
	Items    []AlertConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// AlertConfig holds the configuration of a single Alert.
type AlertConfig struct {
}

// AlertEndpoint represents the REST resource for managing Alert.
const AlertEndpoint = "/alert"

// AlertResource provides an API to manage Alert configurations.
type AlertResource struct {
	c *f5.Client
}

// ListAll  lists all the Alert configurations.
func (r *AlertResource) ListAll() (*AlertConfigList, error) {
	var list AlertConfigList
	if err := r.c.ReadQuery(BasePath+AlertEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Alert configuration identified by id.
func (r *AlertResource) Get(id string) (*AlertConfig, error) {
	var item AlertConfig
	if err := r.c.ReadQuery(BasePath+AlertEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Alert configuration.
func (r *AlertResource) Create(item AlertConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AlertEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Alert configuration identified by id.
func (r *AlertResource) Edit(id string, item AlertConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AlertEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Alert configuration identified by id.
func (r *AlertResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AlertEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
