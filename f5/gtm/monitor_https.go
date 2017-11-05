// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorHTTPSList holds a list of MonitorHTTPS uration.
type MonitorHTTPSList struct {
	Items    []MonitorHTTPS `json:"items,omitempty"`
	Kind     string         `json:"kind,omitempty"`
	SelfLink string         `json:"selflink,omitempty"`
}

// MonitorHTTPS holds the uration of a single MonitorHTTPS.
type MonitorHTTPS struct {
	Cipherlist         string `json:"cipherlist,omitempty"`
	Compatibility      string `json:"compatibility,omitempty"`
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

// MonitorHTTPSEndpoint represents the REST resource for managing MonitorHTTPS.
const MonitorHTTPSEndpoint = "/monitor/https"

// MonitorHTTPSResource provides an API to manage MonitorHTTPS urations.
type MonitorHTTPSResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorHTTPS urations.
func (r *MonitorHTTPSResource) ListAll() (*MonitorHTTPSList, error) {
	var list MonitorHTTPSList
	if err := r.c.ReadQuery(BasePath+MonitorHTTPSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorHTTPS uration identified by id.
func (r *MonitorHTTPSResource) Get(id string) (*MonitorHTTPS, error) {
	var item MonitorHTTPS
	if err := r.c.ReadQuery(BasePath+MonitorHTTPSEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorHTTPS uration.
func (r *MonitorHTTPSResource) Create(item MonitorHTTPS) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorHTTPSEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorHTTPS uration identified by id.
func (r *MonitorHTTPSResource) Edit(id string, item MonitorHTTPS) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorHTTPSEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorHTTPS uration identified by id.
func (r *MonitorHTTPSResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorHTTPSEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
