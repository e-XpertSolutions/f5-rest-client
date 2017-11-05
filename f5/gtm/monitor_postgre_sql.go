// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorPostgreSQLList holds a list of MonitorPostgreSQL configuration.
type MonitorPostgreSQLList struct {
	Items    []MonitorPostgreSQL `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

// MonitorPostgreSQL holds the configuration of a single MonitorPostgreSQL.
type MonitorPostgreSQL struct {
	Count              string `json:"count,omitempty"`
	Debug              string `json:"debug,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// MonitorPostgreSQLEndpoint represents the REST resource for managing MonitorPostgreSQL.
const MonitorPostgreSQLEndpoint = "/monitor/postgresql"

// MonitorPostgreSQLResource provides an API to manage MonitorPostgreSQL configurations.
type MonitorPostgreSQLResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorPostgreSQL configurations.
func (r *MonitorPostgreSQLResource) ListAll() (*MonitorPostgreSQLList, error) {
	var list MonitorPostgreSQLList
	if err := r.c.ReadQuery(BasePath+MonitorPostgreSQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorPostgreSQL configuration identified by id.
func (r *MonitorPostgreSQLResource) Get(id string) (*MonitorPostgreSQL, error) {
	var item MonitorPostgreSQL
	if err := r.c.ReadQuery(BasePath+MonitorPostgreSQLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorPostgreSQL configuration.
func (r *MonitorPostgreSQLResource) Create(item MonitorPostgreSQL) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorPostgreSQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorPostgreSQL configuration identified by id.
func (r *MonitorPostgreSQLResource) Edit(id string, item MonitorPostgreSQL) error {
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
