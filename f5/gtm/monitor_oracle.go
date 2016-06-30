// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorOracleConfigList holds a list of MonitorOracle configuration.
type MonitorOracleConfigList struct {
	Items    []MonitorOracleConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// MonitorOracleConfig holds the configuration of a single MonitorOracle.
type MonitorOracleConfig struct {
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

// MonitorOracleEndpoint represents the REST resource for managing MonitorOracle.
const MonitorOracleEndpoint = "/monitor/oracle"

// MonitorOracleResource provides an API to manage MonitorOracle configurations.
type MonitorOracleResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorOracle configurations.
func (r *MonitorOracleResource) ListAll() (*MonitorOracleConfigList, error) {
	var list MonitorOracleConfigList
	if err := r.c.ReadQuery(BasePath+MonitorOracleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorOracle configuration identified by id.
func (r *MonitorOracleResource) Get(id string) (*MonitorOracleConfig, error) {
	var item MonitorOracleConfig
	if err := r.c.ReadQuery(BasePath+MonitorOracleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorOracle configuration.
func (r *MonitorOracleResource) Create(item MonitorOracleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorOracleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorOracle configuration identified by id.
func (r *MonitorOracleResource) Edit(id string, item MonitorOracleConfig) error {
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
