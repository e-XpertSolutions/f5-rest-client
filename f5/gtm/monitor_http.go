// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorHTTPList holds a list of MonitorHTTP uration.
type MonitorHTTPList struct {
	Items    []MonitorHTTP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorHTTP holds the uration of a single MonitorHTTP.
type MonitorHTTP struct {
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	Reverse            string `json:"reverse,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Send               string `json:"send,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	Transparent        string `json:"transparent,omitempty"`
}

// MonitorHTTPEndpoint represents the REST resource for managing MonitorHTTP.
const MonitorHTTPEndpoint = "/monitor/http"

// MonitorHTTPResource provides an API to manage MonitorHTTP urations.
type MonitorHTTPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorHTTP urations.
func (r *MonitorHTTPResource) ListAll() (*MonitorHTTPList, error) {
	var list MonitorHTTPList
	if err := r.c.ReadQuery(BasePath+MonitorHTTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorHTTP uration identified by id.
func (r *MonitorHTTPResource) Get(id string) (*MonitorHTTP, error) {
	var item MonitorHTTP
	if err := r.c.ReadQuery(BasePath+MonitorHTTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorHTTP uration.
func (r *MonitorHTTPResource) Create(item MonitorHTTP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorHTTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorHTTP uration identified by id.
func (r *MonitorHTTPResource) Edit(id string, item MonitorHTTP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorHTTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorHTTP uration identified by id.
func (r *MonitorHTTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorHTTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
