// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorGTPList holds a list of MonitorGTP uration.
type MonitorGTPList struct {
	Items    []MonitorGTP `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// MonitorGTP holds the uration of a single MonitorGTP.
type MonitorGTP struct {
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
	ProtocolVersion    int    `json:"protocolVersion,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// MonitorGTPEndpoint represents the REST resource for managing MonitorGTP.
const MonitorGTPEndpoint = "/monitor/gtp"

// MonitorGTPResource provides an API to manage MonitorGTP urations.
type MonitorGTPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorGTP urations.
func (r *MonitorGTPResource) ListAll() (*MonitorGTPList, error) {
	var list MonitorGTPList
	if err := r.c.ReadQuery(BasePath+MonitorGTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorGTP uration identified by id.
func (r *MonitorGTPResource) Get(id string) (*MonitorGTP, error) {
	var item MonitorGTP
	if err := r.c.ReadQuery(BasePath+MonitorGTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorGTP uration.
func (r *MonitorGTPResource) Create(item MonitorGTP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorGTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorGTP uration identified by id.
func (r *MonitorGTPResource) Edit(id string, item MonitorGTP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorGTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorGTP uration identified by id.
func (r *MonitorGTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorGTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
