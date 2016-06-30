// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorTCPHalfConfigList holds a list of MonitorTCPHalf configuration.
type MonitorTCPHalfConfigList struct {
	Items    []MonitorTCPHalfConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// MonitorTCPHalfConfig holds the configuration of a single MonitorTCPHalf.
type MonitorTCPHalfConfig struct {
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

// MonitorTCPHalfEndpoint represents the REST resource for managing MonitorTCPHalf.
const MonitorTCPHalfEndpoint = "/monitor/tcp-half-open"

// MonitorTCPHalfResource provides an API to manage MonitorTCPHalf configurations.
type MonitorTCPHalfResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorTCPHalf configurations.
func (r *MonitorTCPHalfResource) ListAll() (*MonitorTCPHalfConfigList, error) {
	var list MonitorTCPHalfConfigList
	if err := r.c.ReadQuery(BasePath+MonitorTCPHalfEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorTCPHalf configuration identified by id.
func (r *MonitorTCPHalfResource) Get(id string) (*MonitorTCPHalfConfig, error) {
	var item MonitorTCPHalfConfig
	if err := r.c.ReadQuery(BasePath+MonitorTCPHalfEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorTCPHalf configuration.
func (r *MonitorTCPHalfResource) Create(item MonitorTCPHalfConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorTCPHalfEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorTCPHalf configuration identified by id.
func (r *MonitorTCPHalfResource) Edit(id string, item MonitorTCPHalfConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorTCPHalfEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorTCPHalf configuration identified by id.
func (r *MonitorTCPHalfResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorTCPHalfEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
