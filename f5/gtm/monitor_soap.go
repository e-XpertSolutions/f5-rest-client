// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSOAPList holds a list of MonitorSOAP configuration.
type MonitorSOAPList struct {
	Items    []MonitorSOAP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorSOAP holds the configuration of a single MonitorSOAP.
type MonitorSOAP struct {
	Debug              string `json:"debug,omitempty"`
	Destination        string `json:"destination,omitempty"`
	ExpectFault        string `json:"expectFault,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// MonitorSOAPEndpoint represents the REST resource for managing MonitorSOAP.
const MonitorSOAPEndpoint = "/monitor/soap"

// MonitorSOAPResource provides an API to manage MonitorSOAP configurations.
type MonitorSOAPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorSOAP configurations.
func (r *MonitorSOAPResource) ListAll() (*MonitorSOAPList, error) {
	var list MonitorSOAPList
	if err := r.c.ReadQuery(BasePath+MonitorSOAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSOAP configuration identified by id.
func (r *MonitorSOAPResource) Get(id string) (*MonitorSOAP, error) {
	var item MonitorSOAP
	if err := r.c.ReadQuery(BasePath+MonitorSOAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSOAP configuration.
func (r *MonitorSOAPResource) Create(item MonitorSOAP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSOAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSOAP configuration identified by id.
func (r *MonitorSOAPResource) Edit(id string, item MonitorSOAP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSOAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorSOAP configuration identified by id.
func (r *MonitorSOAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSOAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
