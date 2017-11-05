// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorTCPHalfList holds a list of MonitorTCPHalf configuration.
type MonitorTCPHalfList struct {
	Items    []MonitorTCPHalf `json:"items,omitempty"`
	Kind     string           `json:"kind,omitempty"`
	SelfLink string           `json:"selflink,omitempty"`
}

// MonitorTCPHalf holds the configuration of a single MonitorTCPHalf.
type MonitorTCPHalf struct {
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
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Transparent        string `json:"transparent,omitempty"`
}

// MonitorTCPHalfEndpoint represents the REST resource for managing MonitorTCPHalf.
const MonitorTCPHalfEndpoint = "/monitor/tcp-half-open"

// MonitorTCPHalfResource provides an API to manage MonitorTCPHalf configurations.
type MonitorTCPHalfResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorTCPHalf configurations.
func (r *MonitorTCPHalfResource) ListAll() (*MonitorTCPHalfList, error) {
	var list MonitorTCPHalfList
	if err := r.c.ReadQuery(BasePath+MonitorTCPHalfEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorTCPHalf configuration identified by id.
func (r *MonitorTCPHalfResource) Get(id string) (*MonitorTCPHalf, error) {
	var item MonitorTCPHalf
	if err := r.c.ReadQuery(BasePath+MonitorTCPHalfEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorTCPHalf configuration.
func (r *MonitorTCPHalfResource) Create(item MonitorTCPHalf) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorTCPHalfEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorTCPHalf configuration identified by id.
func (r *MonitorTCPHalfResource) Edit(id string, item MonitorTCPHalf) error {
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
