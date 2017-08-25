// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorMSSQLConfigList holds a list of MonitorMSSQL configuration.
type MonitorMSSQLConfigList struct {
	Items    []MonitorMSSQLConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// MonitorMSSQLConfig holds the configuration of a single MonitorMSSQL.
type MonitorMSSQLConfig struct {
	Count              string `json:"count"`
	Debug              string `json:"debug"`
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeTimeout       int    `json:"probeTimeout"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
}

// MonitorMSSQLEndpoint represents the REST resource for managing MonitorMSSQL.
const MonitorMSSQLEndpoint = "/monitor/mssql"

// MonitorMSSQLResource provides an API to manage MonitorMSSQL configurations.
type MonitorMSSQLResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorMSSQL configurations.
func (r *MonitorMSSQLResource) ListAll() (*MonitorMSSQLConfigList, error) {
	var list MonitorMSSQLConfigList
	if err := r.c.ReadQuery(BasePath+MonitorMSSQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorMSSQL configuration identified by id.
func (r *MonitorMSSQLResource) Get(id string) (*MonitorMSSQLConfig, error) {
	var item MonitorMSSQLConfig
	if err := r.c.ReadQuery(BasePath+MonitorMSSQLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorMSSQL configuration.
func (r *MonitorMSSQLResource) Create(item MonitorMSSQLConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorMSSQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorMSSQL configuration identified by id.
func (r *MonitorMSSQLResource) Edit(id string, item MonitorMSSQLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorMSSQLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorMSSQL configuration identified by id.
func (r *MonitorMSSQLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorMSSQLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
