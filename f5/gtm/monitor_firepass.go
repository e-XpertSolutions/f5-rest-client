// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorFirepassConfigList holds a list of MonitorFirepass configuration.
type MonitorFirepassConfigList struct {
	Items    []MonitorFirepassConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// MonitorFirepassConfig holds the configuration of a single MonitorFirepass.
type MonitorFirepassConfig struct {
}

// MonitorFirepassEndpoint represents the REST resource for managing MonitorFirepass.
const MonitorFirepassEndpoint = "/monitor/firepass"

// MonitorFirepassResource provides an API to manage MonitorFirepass configurations.
type MonitorFirepassResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorFirepass configurations.
func (r *MonitorFirepassResource) ListAll() (*MonitorFirepassConfigList, error) {
	var list MonitorFirepassConfigList
	if err := r.c.ReadQuery(BasePath+MonitorFirepassEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorFirepass configuration identified by id.
func (r *MonitorFirepassResource) Get(id string) (*MonitorFirepassConfig, error) {
	var item MonitorFirepassConfig
	if err := r.c.ReadQuery(BasePath+MonitorFirepassEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorFirepass configuration.
func (r *MonitorFirepassResource) Create(item MonitorFirepassConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorFirepassEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorFirepass configuration identified by id.
func (r *MonitorFirepassResource) Edit(id string, item MonitorFirepassConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorFirepassEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorFirepass configuration identified by id.
func (r *MonitorFirepassResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorFirepassEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
