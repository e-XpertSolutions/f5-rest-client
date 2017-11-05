// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorBigIPList holds a list of MonitorBigIP uration.
type MonitorBigIPList struct {
	Items    []MonitorBigIP `json:"items,omitempty"`
	Kind     string         `json:"kind,omitempty"`
	SelfLink string         `json:"selflink,omitempty"`
}

// MonitorBigIP holds the uration of a single MonitorBigIP.
type MonitorBigIP struct {
	AggregateDynamicRatios string `json:"aggregateDynamicRatios,omitempty"`
	Destination            string `json:"destination,omitempty"`
	FullPath               string `json:"fullPath,omitempty"`
	Generation             int    `json:"generation,omitempty"`
	IgnoreDownResponse     string `json:"ignoreDownResponse,omitempty"`
	Interval               int    `json:"interval,omitempty"`
	Kind                   string `json:"kind,omitempty"`
	Name                   string `json:"name,omitempty"`
	Partition              string `json:"partition,omitempty"`
	SelfLink               string `json:"selfLink,omitempty"`
	Timeout                int    `json:"timeout,omitempty"`
}

// MonitorBigIPEndpoint represents the REST resource for managing MonitorBigIP.
const MonitorBigIPEndpoint = "/monitor/bigip"

// MonitorBigIPResource provides an API to manage MonitorBigIP urations.
type MonitorBigIPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorBigIP urations.
func (r *MonitorBigIPResource) ListAll() (*MonitorBigIPList, error) {
	var list MonitorBigIPList
	if err := r.c.ReadQuery(BasePath+MonitorBigIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorBigIP uration identified by id.
func (r *MonitorBigIPResource) Get(id string) (*MonitorBigIP, error) {
	var item MonitorBigIP
	if err := r.c.ReadQuery(BasePath+MonitorBigIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorBigIP uration.
func (r *MonitorBigIPResource) Create(item MonitorBigIP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorBigIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorBigIP uration identified by id.
func (r *MonitorBigIPResource) Edit(id string, item MonitorBigIP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorBigIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorBigIP uration identified by id.
func (r *MonitorBigIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorBigIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
