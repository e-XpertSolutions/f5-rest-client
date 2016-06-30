// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorICMPConfigList holds a list of MonitorICMP configuration.
type MonitorICMPConfigList struct {
	Items    []MonitorICMPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorICMPConfig holds the configuration of a single MonitorICMP.
type MonitorICMPConfig struct {
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
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
	Transparent        string `json:"transparent"`
}

// MonitorICMPEndpoint represents the REST resource for managing MonitorICMP.
const MonitorICMPEndpoint = "/monitor/gateway-icmp"

// MonitorICMPResource provides an API to manage MonitorICMP configurations.
type MonitorICMPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorICMP configurations.
func (r *MonitorICMPResource) ListAll() (*MonitorICMPConfigList, error) {
	var list MonitorICMPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorICMPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorICMP configuration identified by id.
func (r *MonitorICMPResource) Get(id string) (*MonitorICMPConfig, error) {
	var item MonitorICMPConfig
	if err := r.c.ReadQuery(BasePath+MonitorICMPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorICMP configuration.
func (r *MonitorICMPResource) Create(item MonitorICMPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorICMPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorICMP configuration identified by id.
func (r *MonitorICMPResource) Edit(id string, item MonitorICMPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorICMPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorICMP configuration identified by id.
func (r *MonitorICMPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorICMPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
