// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorOracleList holds a list of MonitorOracle configuration.
type MonitorOracleList struct {
	Items    []MonitorOracle `json:"items,omitempty"`
	Kind     string          `json:"kind,omitempty"`
	SelfLink string          `json:"selflink,omitempty"`
}

// MonitorOracle holds the configuration of a single MonitorOracle.
type MonitorOracle struct {
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

// MonitorOracleEndpoint represents the REST resource for managing MonitorOracle.
const MonitorOracleEndpoint = "/monitor/oracle"

// MonitorOracleResource provides an API to manage MonitorOracle configurations.
type MonitorOracleResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorOracle configurations.
func (r *MonitorOracleResource) ListAll() (*MonitorOracleList, error) {
	var list MonitorOracleList
	if err := r.c.ReadQuery(BasePath+MonitorOracleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorOracle configuration identified by id.
func (r *MonitorOracleResource) Get(id string) (*MonitorOracle, error) {
	var item MonitorOracle
	if err := r.c.ReadQuery(BasePath+MonitorOracleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorOracle configuration.
func (r *MonitorOracleResource) Create(item MonitorOracle) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorOracleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorOracle configuration identified by id.
func (r *MonitorOracleResource) Edit(id string, item MonitorOracle) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorOracleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorOracle configuration identified by id.
func (r *MonitorOracleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorOracleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
