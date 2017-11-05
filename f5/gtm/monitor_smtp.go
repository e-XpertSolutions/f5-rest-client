// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSMTPList holds a list of MonitorSMTP configuration.
type MonitorSMTPList struct {
	Items    []MonitorSMTP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorSMTP holds the configuration of a single MonitorSMTP.
type MonitorSMTP struct {
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

// MonitorSMTPEndpoint represents the REST resource for managing MonitorSMTP.
const MonitorSMTPEndpoint = "/monitor/smtp"

// MonitorSMTPResource provides an API to manage MonitorSMTP configurations.
type MonitorSMTPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorSMTP configurations.
func (r *MonitorSMTPResource) ListAll() (*MonitorSMTPList, error) {
	var list MonitorSMTPList
	if err := r.c.ReadQuery(BasePath+MonitorSMTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSMTP configuration identified by id.
func (r *MonitorSMTPResource) Get(id string) (*MonitorSMTP, error) {
	var item MonitorSMTP
	if err := r.c.ReadQuery(BasePath+MonitorSMTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSMTP configuration.
func (r *MonitorSMTPResource) Create(item MonitorSMTP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSMTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSMTP configuration identified by id.
func (r *MonitorSMTPResource) Edit(id string, item MonitorSMTP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSMTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorSMTP configuration identified by id.
func (r *MonitorSMTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSMTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
