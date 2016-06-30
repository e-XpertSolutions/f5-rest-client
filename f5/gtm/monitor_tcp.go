// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorTCPConfigList holds a list of MonitorTCP configuration.
type MonitorTCPConfigList struct {
	Items    []MonitorTCPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// MonitorTCPConfig holds the configuration of a single MonitorTCP.
type MonitorTCPConfig struct {
}

// MonitorTCPEndpoint represents the REST resource for managing MonitorTCP.
const MonitorTCPEndpoint = "/monitor/tcp"

// MonitorTCPResource provides an API to manage MonitorTCP configurations.
type MonitorTCPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorTCP configurations.
func (r *MonitorTCPResource) ListAll() (*MonitorTCPConfigList, error) {
	var list MonitorTCPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorTCPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorTCP configuration identified by id.
func (r *MonitorTCPResource) Get(id string) (*MonitorTCPConfig, error) {
	var item MonitorTCPConfig
	if err := r.c.ReadQuery(BasePath+MonitorTCPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorTCP configuration.
func (r *MonitorTCPResource) Create(item MonitorTCPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorTCPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorTCP configuration identified by id.
func (r *MonitorTCPResource) Edit(id string, item MonitorTCPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorTCPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorTCP configuration identified by id.
func (r *MonitorTCPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorTCPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
