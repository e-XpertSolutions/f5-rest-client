// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSNMPLinkConfigList holds a list of MonitorSNMPLink configuration.
type MonitorSNMPLinkConfigList struct {
	Items    []MonitorSNMPLinkConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// MonitorSNMPLinkConfig holds the configuration of a single MonitorSNMPLink.
type MonitorSNMPLinkConfig struct {
	Community          string `json:"community"`
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	Port               int    `json:"port"`
	ProbeAttempts      int    `json:"probeAttempts"`
	ProbeInterval      int    `json:"probeInterval"`
	ProbeTimeout       int    `json:"probeTimeout"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
	Version            string `json:"version"`
}

// MonitorSNMPLinkEndpoint represents the REST resource for managing MonitorSNMPLink.
const MonitorSNMPLinkEndpoint = "/monitor/snmp-link"

// MonitorSNMPLinkResource provides an API to manage MonitorSNMPLink configurations.
type MonitorSNMPLinkResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorSNMPLink configurations.
func (r *MonitorSNMPLinkResource) ListAll() (*MonitorSNMPLinkConfigList, error) {
	var list MonitorSNMPLinkConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSNMPLinkEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSNMPLink configuration identified by id.
func (r *MonitorSNMPLinkResource) Get(id string) (*MonitorSNMPLinkConfig, error) {
	var item MonitorSNMPLinkConfig
	if err := r.c.ReadQuery(BasePath+MonitorSNMPLinkEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSNMPLink configuration.
func (r *MonitorSNMPLinkResource) Create(item MonitorSNMPLinkConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSNMPLinkEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSNMPLink configuration identified by id.
func (r *MonitorSNMPLinkResource) Edit(id string, item MonitorSNMPLinkConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSNMPLinkEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorSNMPLink configuration identified by id.
func (r *MonitorSNMPLinkResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSNMPLinkEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
