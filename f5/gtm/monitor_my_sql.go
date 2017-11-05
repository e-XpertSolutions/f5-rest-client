// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorMySQLList holds a list of MonitorMySQL uration.
type MonitorMySQLList struct {
	Items    []MonitorMySQL `json:"items,omitempty"`
	Kind     string         `json:"kind,omitempty"`
	SelfLink string         `json:"selflink,omitempty"`
}

// MonitorMySQL holds the uration of a single MonitorMySQL.
type MonitorMySQL struct {
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

// MonitorMySQLEndpoint represents the REST resource for managing MonitorMySQL.
const MonitorMySQLEndpoint = "/monitor/mysql"

// MonitorMySQLResource provides an API to manage MonitorMySQL urations.
type MonitorMySQLResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorMySQL urations.
func (r *MonitorMySQLResource) ListAll() (*MonitorMySQLList, error) {
	var list MonitorMySQLList
	if err := r.c.ReadQuery(BasePath+MonitorMySQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorMySQL uration identified by id.
func (r *MonitorMySQLResource) Get(id string) (*MonitorMySQL, error) {
	var item MonitorMySQL
	if err := r.c.ReadQuery(BasePath+MonitorMySQLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorMySQL uration.
func (r *MonitorMySQLResource) Create(item MonitorMySQL) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorMySQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorMySQL uration identified by id.
func (r *MonitorMySQLResource) Edit(id string, item MonitorMySQL) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorMySQLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorMySQL uration identified by id.
func (r *MonitorMySQLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorMySQLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
