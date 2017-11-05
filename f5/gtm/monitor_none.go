// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorNoneList holds a list of MonitorNone uration.
type MonitorNoneList struct {
	Items    []MonitorNone `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorNone holds the uration of a single MonitorNone.
type MonitorNone struct {
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
	TimeUntilUp        int    `json:"timeUntilUp,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	UpInterval         int    `json:"upInterval,omitempty"`
}

// MonitorNoneEndpoint represents the REST resource for managing MonitorNone.
const MonitorNoneEndpoint = "/monitor/none"

// MonitorNoneResource provides an API to manage MonitorNone urations.
type MonitorNoneResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorNone urations.
func (r *MonitorNoneResource) ListAll() (*MonitorNoneList, error) {
	var list MonitorNoneList
	if err := r.c.ReadQuery(BasePath+MonitorNoneEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorNone uration identified by id.
func (r *MonitorNoneResource) Get(id string) (*MonitorNone, error) {
	var item MonitorNone
	if err := r.c.ReadQuery(BasePath+MonitorNoneEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorNone uration.
func (r *MonitorNoneResource) Create(item MonitorNone) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorNoneEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorNone uration identified by id.
func (r *MonitorNoneResource) Edit(id string, item MonitorNone) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorNoneEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorNone uration identified by id.
func (r *MonitorNoneResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorNoneEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
