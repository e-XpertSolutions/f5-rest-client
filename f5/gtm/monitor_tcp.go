// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorTCPList holds a list of MonitorTCP configuration.
type MonitorTCPList struct {
	Items    []MonitorTCP `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// MonitorTCP holds the configuration of a single MonitorTCP.
type MonitorTCP struct {
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	Reverse            string `json:"reverse,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Transparent        string `json:"transparent,omitempty"`
}

// MonitorTCPEndpoint represents the REST resource for managing MonitorTCP.
const MonitorTCPEndpoint = "/monitor/tcp"

// MonitorTCPResource provides an API to manage MonitorTCP configurations.
type MonitorTCPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorTCP configurations.
func (r *MonitorTCPResource) ListAll() (*MonitorTCPList, error) {
	var list MonitorTCPList
	if err := r.c.ReadQuery(BasePath+MonitorTCPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorTCP configuration identified by id.
func (r *MonitorTCPResource) Get(id string) (*MonitorTCP, error) {
	var item MonitorTCP
	if err := r.c.ReadQuery(BasePath+MonitorTCPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorTCP configuration.
func (r *MonitorTCPResource) Create(item MonitorTCP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorTCPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorTCP configuration identified by id.
func (r *MonitorTCPResource) Edit(id string, item MonitorTCP) error {
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
