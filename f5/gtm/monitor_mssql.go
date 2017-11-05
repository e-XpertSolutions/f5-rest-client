// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorMSSQLList holds a list of MonitorMSSQL uration.
type MonitorMSSQLList struct {
	Items    []MonitorMSSQL `json:"items,omitempty"`
	Kind     string         `json:"kind,omitempty"`
	SelfLink string         `json:"selflink,omitempty"`
}

// MonitorMSSQL holds the uration of a single MonitorMSSQL.
type MonitorMSSQL struct {
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

// MonitorMSSQLEndpoint represents the REST resource for managing MonitorMSSQL.
const MonitorMSSQLEndpoint = "/monitor/mssql"

// MonitorMSSQLResource provides an API to manage MonitorMSSQL urations.
type MonitorMSSQLResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorMSSQL urations.
func (r *MonitorMSSQLResource) ListAll() (*MonitorMSSQLList, error) {
	var list MonitorMSSQLList
	if err := r.c.ReadQuery(BasePath+MonitorMSSQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorMSSQL uration identified by id.
func (r *MonitorMSSQLResource) Get(id string) (*MonitorMSSQL, error) {
	var item MonitorMSSQL
	if err := r.c.ReadQuery(BasePath+MonitorMSSQLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorMSSQL uration.
func (r *MonitorMSSQLResource) Create(item MonitorMSSQL) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorMSSQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorMSSQL uration identified by id.
func (r *MonitorMSSQLResource) Edit(id string, item MonitorMSSQL) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorMSSQLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorMSSQL uration identified by id.
func (r *MonitorMSSQLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorMSSQLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
