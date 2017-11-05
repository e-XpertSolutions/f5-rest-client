// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorWMIList holds a list of MonitorWMI configuration.
type MonitorWMIList struct {
	Items    []MonitorWMI `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// MonitorWMI holds the configuration of a single MonitorWMI.
type MonitorWMI struct {
	Agent              string `json:"agent,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Method             string `json:"method,omitempty"`
	Metrics            string `json:"metrics,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	Post               string `json:"post,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	TmCommand          string `json:"tmCommand,omitempty"`
	URL                string `json:"url,omitempty"`
}

// MonitorWMIEndpoint represents the REST resource for managing MonitorWMI.
const MonitorWMIEndpoint = "/monitor/wmi"

// MonitorWMIResource provides an API to manage MonitorWMI configurations.
type MonitorWMIResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorWMI configurations.
func (r *MonitorWMIResource) ListAll() (*MonitorWMIList, error) {
	var list MonitorWMIList
	if err := r.c.ReadQuery(BasePath+MonitorWMIEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorWMI configuration identified by id.
func (r *MonitorWMIResource) Get(id string) (*MonitorWMI, error) {
	var item MonitorWMI
	if err := r.c.ReadQuery(BasePath+MonitorWMIEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorWMI configuration.
func (r *MonitorWMIResource) Create(item MonitorWMI) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorWMIEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorWMI configuration identified by id.
func (r *MonitorWMIResource) Edit(id string, item MonitorWMI) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorWMIEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorWMI configuration identified by id.
func (r *MonitorWMIResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorWMIEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
