// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorscriptedList holds a list of Monitorscripted uration.
type MonitorscriptedList struct {
	Items    []Monitorscripted `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

// Monitorscripted holds the uration of a single Monitorscripted.
type Monitorscripted struct {
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

// MonitorscriptedEndpoint represents the REST resource for managing Monitorscripted.
const MonitorscriptedEndpoint = "/monitor/scripted"

// MonitorscriptedResource provides an API to manage Monitorscripted urations.
type MonitorscriptedResource struct {
	c *f5.Client
}

// ListAll  lists all the Monitorscripted urations.
func (r *MonitorscriptedResource) ListAll() (*MonitorscriptedList, error) {
	var list MonitorscriptedList
	if err := r.c.ReadQuery(BasePath+MonitorscriptedEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Monitorscripted uration identified by id.
func (r *MonitorscriptedResource) Get(id string) (*Monitorscripted, error) {
	var item Monitorscripted
	if err := r.c.ReadQuery(BasePath+MonitorscriptedEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Monitorscripted uration.
func (r *MonitorscriptedResource) Create(item Monitorscripted) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorscriptedEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Monitorscripted uration identified by id.
func (r *MonitorscriptedResource) Edit(id string, item Monitorscripted) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorscriptedEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Monitorscripted uration identified by id.
func (r *MonitorscriptedResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorscriptedEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
