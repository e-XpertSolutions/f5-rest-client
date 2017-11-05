// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorUDPList holds a list of MonitorUDP configuration.
type MonitorUDPList struct {
	Items    []MonitorUDP `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// MonitorUDP holds the configuration of a single MonitorUDP.
type MonitorUDP struct {
	Debug              string `json:"debug,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeAttempts      int    `json:"probeAttempts,omitempty"`
	ProbeInterval      int    `json:"probeInterval,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	Reverse            string `json:"reverse,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Send               string `json:"send,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Transparent        string `json:"transparent,omitempty"`
}

// MonitorUDPEndpoint represents the REST resource for managing MonitorUDP.
const MonitorUDPEndpoint = "/monitor/udp"

// MonitorUDPResource provides an API to manage MonitorUDP configurations.
type MonitorUDPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorUDP configurations.
func (r *MonitorUDPResource) ListAll() (*MonitorUDPList, error) {
	var list MonitorUDPList
	if err := r.c.ReadQuery(BasePath+MonitorUDPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorUDP configuration identified by id.
func (r *MonitorUDPResource) Get(id string) (*MonitorUDP, error) {
	var item MonitorUDP
	if err := r.c.ReadQuery(BasePath+MonitorUDPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorUDP configuration.
func (r *MonitorUDPResource) Create(item MonitorUDP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorUDPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorUDP configuration identified by id.
func (r *MonitorUDPResource) Edit(id string, item MonitorUDP) error {
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
