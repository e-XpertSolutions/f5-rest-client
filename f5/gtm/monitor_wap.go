// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorWAPList holds a list of MonitorWAP configuration.
type MonitorWAPList struct {
	Items    []MonitorWAP `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// MonitorWAP holds the configuration of a single MonitorWAP.
type MonitorWAP struct {
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

// MonitorWAPEndpoint represents the REST resource for managing MonitorWAP.
const MonitorWAPEndpoint = "/monitor/wap"

// MonitorWAPResource provides an API to manage MonitorWAP configurations.
type MonitorWAPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorWAP configurations.
func (r *MonitorWAPResource) ListAll() (*MonitorWAPList, error) {
	var list MonitorWAPList
	if err := r.c.ReadQuery(BasePath+MonitorWAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorWAP configuration identified by id.
func (r *MonitorWAPResource) Get(id string) (*MonitorWAP, error) {
	var item MonitorWAP
	if err := r.c.ReadQuery(BasePath+MonitorWAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorWAP configuration.
func (r *MonitorWAPResource) Create(item MonitorWAP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorWAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorWAP configuration identified by id.
func (r *MonitorWAPResource) Edit(id string, item MonitorWAP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorWAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorWAP configuration identified by id.
func (r *MonitorWAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorWAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
