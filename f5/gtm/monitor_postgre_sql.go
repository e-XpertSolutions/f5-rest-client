// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorPostgreSQLConfigList holds a list of MonitorPostgreSQL configuration.
type MonitorPostgreSQLConfigList struct {
	Items    []MonitorPostgreSQLConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

// MonitorPostgreSQLConfig holds the configuration of a single MonitorPostgreSQL.
type MonitorPostgreSQLConfig struct {
}

// MonitorPostgreSQLEndpoint represents the REST resource for managing MonitorPostgreSQL.
const MonitorPostgreSQLEndpoint = "/monitor/postgresql"

// MonitorPostgreSQLResource provides an API to manage MonitorPostgreSQL configurations.
type MonitorPostgreSQLResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorPostgreSQL configurations.
func (r *MonitorPostgreSQLResource) ListAll() (*MonitorPostgreSQLConfigList, error) {
	var list MonitorPostgreSQLConfigList
	if err := r.c.ReadQuery(BasePath+MonitorPostgreSQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorPostgreSQL configuration identified by id.
func (r *MonitorPostgreSQLResource) Get(id string) (*MonitorPostgreSQLConfig, error) {
	var item MonitorPostgreSQLConfig
	if err := r.c.ReadQuery(BasePath+MonitorPostgreSQLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorPostgreSQL configuration.
func (r *MonitorPostgreSQLResource) Create(item MonitorPostgreSQLConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorPostgreSQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorPostgreSQL configuration identified by id.
func (r *MonitorPostgreSQLResource) Edit(id string, item MonitorPostgreSQLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorPostgreSQLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorPostgreSQL configuration identified by id.
func (r *MonitorPostgreSQLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorPostgreSQLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
