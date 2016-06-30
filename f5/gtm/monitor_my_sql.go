// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorMySQLConfigList holds a list of MonitorMySQL configuration.
type MonitorMySQLConfigList struct {
	Items    []MonitorMySQLConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// MonitorMySQLConfig holds the configuration of a single MonitorMySQL.
type MonitorMySQLConfig struct {
}

// MonitorMySQLEndpoint represents the REST resource for managing MonitorMySQL.
const MonitorMySQLEndpoint = "/monitor/mysql"

// MonitorMySQLResource provides an API to manage MonitorMySQL configurations.
type MonitorMySQLResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorMySQL configurations.
func (r *MonitorMySQLResource) ListAll() (*MonitorMySQLConfigList, error) {
	var list MonitorMySQLConfigList
	if err := r.c.ReadQuery(BasePath+MonitorMySQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorMySQL configuration identified by id.
func (r *MonitorMySQLResource) Get(id string) (*MonitorMySQLConfig, error) {
	var item MonitorMySQLConfig
	if err := r.c.ReadQuery(BasePath+MonitorMySQLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorMySQL configuration.
func (r *MonitorMySQLResource) Create(item MonitorMySQLConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorMySQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorMySQL configuration identified by id.
func (r *MonitorMySQLResource) Edit(id string, item MonitorMySQLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorMySQLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorMySQL configuration identified by id.
func (r *MonitorMySQLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorMySQLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
