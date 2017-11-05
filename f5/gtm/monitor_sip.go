// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSIPList holds a list of MonitorSIP configuration.
type MonitorSIPList struct {
	Items    []MonitorSIP `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// MonitorSIP holds the configuration of a single MonitorSIP.
type MonitorSIP struct {
	Compatibility      string `json:"compatibility,omitempty"`
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

// MonitorSIPEndpoint represents the REST resource for managing MonitorSIP.
const MonitorSIPEndpoint = "/monitor/sip"

// MonitorSIPResource provides an API to manage MonitorSIP configurations.
type MonitorSIPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorSIP configurations.
func (r *MonitorSIPResource) ListAll() (*MonitorSIPList, error) {
	var list MonitorSIPList
	if err := r.c.ReadQuery(BasePath+MonitorSIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSIP configuration identified by id.
func (r *MonitorSIPResource) Get(id string) (*MonitorSIP, error) {
	var item MonitorSIP
	if err := r.c.ReadQuery(BasePath+MonitorSIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSIP configuration.
func (r *MonitorSIPResource) Create(item MonitorSIP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSIP configuration identified by id.
func (r *MonitorSIPResource) Edit(id string, item MonitorSIP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorSIP configuration identified by id.
func (r *MonitorSIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
