// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorConfigList holds a list of Monitor configuration.
type MonitorConfigList struct {
	Items    []MonitorConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// MonitorConfig holds the configuration of a single Monitor.
type MonitorConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// MonitorEndpoint represents the REST resource for managing Monitor.
const MonitorEndpoint = "/monitor"

// MonitorResource provides an API to manage Monitor configurations.
type MonitorResource struct {
	c *f5.Client
}

// ListAll  lists all the Monitor configurations.
func (r *MonitorResource) ListAll() (*MonitorConfigList, error) {
	var list MonitorConfigList
	if err := r.c.ReadQuery(BasePath+MonitorEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Monitor configuration identified by id.
func (r *MonitorResource) Get(id string) (*MonitorConfig, error) {
	var item MonitorConfig
	if err := r.c.ReadQuery(BasePath+MonitorEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Monitor configuration.
func (r *MonitorResource) Create(item MonitorConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Monitor configuration identified by id.
func (r *MonitorResource) Edit(id string, item MonitorConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Monitor configuration identified by id.
func (r *MonitorResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
