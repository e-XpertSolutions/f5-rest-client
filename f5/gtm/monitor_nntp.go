// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorNNTPList holds a list of MonitorNNTP uration.
type MonitorNNTPList struct {
	Items    []MonitorNNTP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorNNTP holds the uration of a single MonitorNNTP.
type MonitorNNTP struct {
	Debug              string `json:"debug,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// MonitorNNTPEndpoint represents the REST resource for managing MonitorNNTP.
const MonitorNNTPEndpoint = "/monitor/nntp"

// MonitorNNTPResource provides an API to manage MonitorNNTP urations.
type MonitorNNTPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorNNTP urations.
func (r *MonitorNNTPResource) ListAll() (*MonitorNNTPList, error) {
	var list MonitorNNTPList
	if err := r.c.ReadQuery(BasePath+MonitorNNTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorNNTP uration identified by id.
func (r *MonitorNNTPResource) Get(id string) (*MonitorNNTP, error) {
	var item MonitorNNTP
	if err := r.c.ReadQuery(BasePath+MonitorNNTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorNNTP uration.
func (r *MonitorNNTPResource) Create(item MonitorNNTP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorNNTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorNNTP uration identified by id.
func (r *MonitorNNTPResource) Edit(id string, item MonitorNNTP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorNNTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorNNTP uration identified by id.
func (r *MonitorNNTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorNNTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
