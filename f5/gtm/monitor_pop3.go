// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorPOP3List holds a list of MonitorPOP3 configuration.
type MonitorPOP3List struct {
	Items    []MonitorPOP3 `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorPOP3 holds the configuration of a single MonitorPOP3.
type MonitorPOP3 struct {
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

// MonitorPOP3Endpoint represents the REST resource for managing MonitorPOP3.
const MonitorPOP3Endpoint = "/monitor/pop3"

// MonitorPOP3Resource provides an API to manage MonitorPOP3 configurations.
type MonitorPOP3Resource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorPOP3 configurations.
func (r *MonitorPOP3Resource) ListAll() (*MonitorPOP3List, error) {
	var list MonitorPOP3List
	if err := r.c.ReadQuery(BasePath+MonitorPOP3Endpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorPOP3 configuration identified by id.
func (r *MonitorPOP3Resource) Get(id string) (*MonitorPOP3, error) {
	var item MonitorPOP3
	if err := r.c.ReadQuery(BasePath+MonitorPOP3Endpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorPOP3 configuration.
func (r *MonitorPOP3Resource) Create(item MonitorPOP3) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorPOP3Endpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorPOP3 configuration identified by id.
func (r *MonitorPOP3Resource) Edit(id string, item MonitorPOP3) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorPOP3Endpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorPOP3 configuration identified by id.
func (r *MonitorPOP3Resource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorPOP3Endpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
