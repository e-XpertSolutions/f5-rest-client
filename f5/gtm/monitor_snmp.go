// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSNMPConfigList holds a list of MonitorSNMP configuration.
type MonitorSNMPConfigList struct {
	Items    []MonitorSNMPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorSNMPConfig holds the configuration of a single MonitorSNMP.
type MonitorSNMPConfig struct {
}

// MonitorSNMPEndpoint represents the REST resource for managing MonitorSNMP.
const MonitorSNMPEndpoint = "/monitor/snmp"

// MonitorSNMPResource provides an API to manage MonitorSNMP configurations.
type MonitorSNMPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorSNMP configurations.
func (r *MonitorSNMPResource) ListAll() (*MonitorSNMPConfigList, error) {
	var list MonitorSNMPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSNMPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSNMP configuration identified by id.
func (r *MonitorSNMPResource) Get(id string) (*MonitorSNMPConfig, error) {
	var item MonitorSNMPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSNMPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSNMP configuration.
func (r *MonitorSNMPResource) Create(item MonitorSNMPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSNMPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSNMP configuration identified by id.
func (r *MonitorSNMPResource) Edit(id string, item MonitorSNMPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSNMPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorSNMP configuration identified by id.
func (r *MonitorSNMPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSNMPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
