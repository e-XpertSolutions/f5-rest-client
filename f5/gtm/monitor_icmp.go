// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorICMPList holds a list of MonitorICMP uration.
type MonitorICMPList struct {
	Items    []MonitorICMP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorICMP holds the uration of a single MonitorICMP.
type MonitorICMP struct {
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

// MonitorICMPEndpoint represents the REST resource for managing MonitorICMP.
const MonitorICMPEndpoint = "/monitor/gateway-icmp"

// MonitorICMPResource provides an API to manage MonitorICMP urations.
type MonitorICMPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorICMP urations.
func (r *MonitorICMPResource) ListAll() (*MonitorICMPList, error) {
	var list MonitorICMPList
	if err := r.c.ReadQuery(BasePath+MonitorICMPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorICMP uration identified by id.
func (r *MonitorICMPResource) Get(id string) (*MonitorICMP, error) {
	var item MonitorICMP
	if err := r.c.ReadQuery(BasePath+MonitorICMPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorICMP uration.
func (r *MonitorICMPResource) Create(item MonitorICMP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorICMPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorICMP uration identified by id.
func (r *MonitorICMPResource) Edit(id string, item MonitorICMP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorICMPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorICMP uration identified by id.
func (r *MonitorICMPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorICMPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
