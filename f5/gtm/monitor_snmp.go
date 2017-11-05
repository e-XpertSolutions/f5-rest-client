// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSNMPList holds a list of MonitorSNMP configuration.
type MonitorSNMPList struct {
	Items    []MonitorSNMP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorSNMP holds the configuration of a single MonitorSNMP.
type MonitorSNMP struct {
	Community          string `json:"community,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	Port               int    `json:"port,omitempty"`
	ProbeAttempts      int    `json:"probeAttempts,omitempty"`
	ProbeInterval      int    `json:"probeInterval,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Version            string `json:"version,omitempty"`
}

// MonitorSNMPEndpoint represents the REST resource for managing MonitorSNMP.
const MonitorSNMPEndpoint = "/monitor/snmp"

// MonitorSNMPResource provides an API to manage MonitorSNMP configurations.
type MonitorSNMPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorSNMP configurations.
func (r *MonitorSNMPResource) ListAll() (*MonitorSNMPList, error) {
	var list MonitorSNMPList
	if err := r.c.ReadQuery(BasePath+MonitorSNMPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSNMP configuration identified by id.
func (r *MonitorSNMPResource) Get(id string) (*MonitorSNMP, error) {
	var item MonitorSNMP
	if err := r.c.ReadQuery(BasePath+MonitorSNMPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSNMP configuration.
func (r *MonitorSNMPResource) Create(item MonitorSNMP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSNMPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSNMP configuration identified by id.
func (r *MonitorSNMPResource) Edit(id string, item MonitorSNMP) error {
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
