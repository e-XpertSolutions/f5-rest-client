// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorFTPList holds a list of MonitorFTP uration.
type MonitorFTPList struct {
	Items    []MonitorFTP `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// MonitorFTP holds the uration of a single MonitorFTP.
type MonitorFTP struct {
	Debug              string `json:"debug,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Mode               string `json:"mode,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// MonitorFTPEndpoint represents the REST resource for managing MonitorFTP.
const MonitorFTPEndpoint = "/monitor/ftp"

// MonitorFTPResource provides an API to manage MonitorFTP urations.
type MonitorFTPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorFTP urations.
func (r *MonitorFTPResource) ListAll() (*MonitorFTPList, error) {
	var list MonitorFTPList
	if err := r.c.ReadQuery(BasePath+MonitorFTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorFTP uration identified by id.
func (r *MonitorFTPResource) Get(id string) (*MonitorFTP, error) {
	var item MonitorFTP
	if err := r.c.ReadQuery(BasePath+MonitorFTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorFTP uration.
func (r *MonitorFTPResource) Create(item MonitorFTP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorFTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorFTP uration identified by id.
func (r *MonitorFTPResource) Edit(id string, item MonitorFTP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorFTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorFTP uration identified by id.
func (r *MonitorFTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorFTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
