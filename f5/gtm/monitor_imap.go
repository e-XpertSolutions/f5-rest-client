// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorIMAPList holds a list of MonitorIMAP uration.
type MonitorIMAPList struct {
	Items    []MonitorIMAP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorIMAP holds the uration of a single MonitorIMAP.
type MonitorIMAP struct {
	Debug              string `json:"debug,omitempty"`
	Destination        string `json:"destination,omitempty"`
	Folder             string `json:"folder,omitempty"`
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

// MonitorIMAPEndpoint represents the REST resource for managing MonitorIMAP.
const MonitorIMAPEndpoint = "/monitor/imap"

// MonitorIMAPResource provides an API to manage MonitorIMAP urations.
type MonitorIMAPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorIMAP urations.
func (r *MonitorIMAPResource) ListAll() (*MonitorIMAPList, error) {
	var list MonitorIMAPList
	if err := r.c.ReadQuery(BasePath+MonitorIMAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorIMAP uration identified by id.
func (r *MonitorIMAPResource) Get(id string) (*MonitorIMAP, error) {
	var item MonitorIMAP
	if err := r.c.ReadQuery(BasePath+MonitorIMAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorIMAP uration.
func (r *MonitorIMAPResource) Create(item MonitorIMAP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorIMAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorIMAP uration identified by id.
func (r *MonitorIMAPResource) Edit(id string, item MonitorIMAP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorIMAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorIMAP uration identified by id.
func (r *MonitorIMAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorIMAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
