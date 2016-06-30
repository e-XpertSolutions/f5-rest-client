// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorUDPConfigList holds a list of MonitorUDP configuration.
type MonitorUDPConfigList struct {
	Items    []MonitorUDPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// MonitorUDPConfig holds the configuration of a single MonitorUDP.
type MonitorUDPConfig struct {
	Debug              string `json:"debug"`
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeAttempts      int    `json:"probeAttempts"`
	ProbeInterval      int    `json:"probeInterval"`
	ProbeTimeout       int    `json:"probeTimeout"`
	Reverse            string `json:"reverse"`
	SelfLink           string `json:"selfLink"`
	Send               string `json:"send"`
	Timeout            int    `json:"timeout"`
	Transparent        string `json:"transparent"`
}

// MonitorUDPEndpoint represents the REST resource for managing MonitorUDP.
const MonitorUDPEndpoint = "/monitor/udp"

// MonitorUDPResource provides an API to manage MonitorUDP configurations.
type MonitorUDPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorUDP configurations.
func (r *MonitorUDPResource) ListAll() (*MonitorUDPConfigList, error) {
	var list MonitorUDPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorUDPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorUDP configuration identified by id.
func (r *MonitorUDPResource) Get(id string) (*MonitorUDPConfig, error) {
	var item MonitorUDPConfig
	if err := r.c.ReadQuery(BasePath+MonitorUDPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorUDP configuration.
func (r *MonitorUDPResource) Create(item MonitorUDPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorUDPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorUDP configuration identified by id.
func (r *MonitorUDPResource) Edit(id string, item MonitorUDPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorUDPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorUDP configuration identified by id.
func (r *MonitorUDPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorUDPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
