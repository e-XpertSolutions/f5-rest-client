// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorExternalList holds a list of MonitorExternal uration.
type MonitorExternalList struct {
	Items    []MonitorExternal `json:"items,omitempty"`
	Kind     string            `json:"kind,omitempty"`
	SelfLink string            `json:"selflink,omitempty"`
}

// MonitorExternal holds the uration of a single MonitorExternal.
type MonitorExternal struct {
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

// MonitorExternalEndpoint represents the REST resource for managing MonitorExternal.
const MonitorExternalEndpoint = "/monitor/external"

// MonitorExternalResource provides an API to manage MonitorExternal urations.
type MonitorExternalResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorExternal urations.
func (r *MonitorExternalResource) ListAll() (*MonitorExternalList, error) {
	var list MonitorExternalList
	if err := r.c.ReadQuery(BasePath+MonitorExternalEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorExternal uration identified by id.
func (r *MonitorExternalResource) Get(id string) (*MonitorExternal, error) {
	var item MonitorExternal
	if err := r.c.ReadQuery(BasePath+MonitorExternalEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorExternal uration.
func (r *MonitorExternalResource) Create(item MonitorExternal) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorExternalEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorExternal uration identified by id.
func (r *MonitorExternalResource) Edit(id string, item MonitorExternal) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorExternalEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorExternal uration identified by id.
func (r *MonitorExternalResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorExternalEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
