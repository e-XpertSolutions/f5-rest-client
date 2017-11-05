// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorFirepassList holds a list of MonitorFirepass uration.
type MonitorFirepassList struct {
	Items    []MonitorFirepass `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

// MonitorFirepass holds the uration of a single MonitorFirepass.
type MonitorFirepass struct {
	Cipherlist         string `json:"cipherlist,omitempty"`
	ConcurrencyLimit   int    `json:"concurrencyLimit,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	MaxLoadAverage     int    `json:"maxLoadAverage,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Username           string `json:"username,omitempty"`
}

// MonitorFirepassEndpoint represents the REST resource for managing MonitorFirepass.
const MonitorFirepassEndpoint = "/monitor/firepass"

// MonitorFirepassResource provides an API to manage MonitorFirepass urations.
type MonitorFirepassResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorFirepass urations.
func (r *MonitorFirepassResource) ListAll() (*MonitorFirepassList, error) {
	var list MonitorFirepassList
	if err := r.c.ReadQuery(BasePath+MonitorFirepassEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorFirepass uration identified by id.
func (r *MonitorFirepassResource) Get(id string) (*MonitorFirepass, error) {
	var item MonitorFirepass
	if err := r.c.ReadQuery(BasePath+MonitorFirepassEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorFirepass uration.
func (r *MonitorFirepassResource) Create(item MonitorFirepass) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorFirepassEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorFirepass uration identified by id.
func (r *MonitorFirepassResource) Edit(id string, item MonitorFirepass) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorFirepassEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorFirepass uration identified by id.
func (r *MonitorFirepassResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorFirepassEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
