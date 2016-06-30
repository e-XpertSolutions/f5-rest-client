// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorExternalConfigList holds a list of MonitorExternal configuration.
type MonitorExternalConfigList struct {
	Items    []MonitorExternalConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// MonitorExternalConfig holds the configuration of a single MonitorExternal.
type MonitorExternalConfig struct {
}

// MonitorExternalEndpoint represents the REST resource for managing MonitorExternal.
const MonitorExternalEndpoint = "/monitor/external"

// MonitorExternalResource provides an API to manage MonitorExternal configurations.
type MonitorExternalResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorExternal configurations.
func (r *MonitorExternalResource) ListAll() (*MonitorExternalConfigList, error) {
	var list MonitorExternalConfigList
	if err := r.c.ReadQuery(BasePath+MonitorExternalEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorExternal configuration identified by id.
func (r *MonitorExternalResource) Get(id string) (*MonitorExternalConfig, error) {
	var item MonitorExternalConfig
	if err := r.c.ReadQuery(BasePath+MonitorExternalEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorExternal configuration.
func (r *MonitorExternalResource) Create(item MonitorExternalConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorExternalEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorExternal configuration identified by id.
func (r *MonitorExternalResource) Edit(id string, item MonitorExternalConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorExternalEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorExternal configuration identified by id.
func (r *MonitorExternalResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorExternalEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
