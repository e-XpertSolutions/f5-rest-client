// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorRadiusAccountingList holds a list of MonitorRadiusAccounting configuration.
type MonitorRadiusAccountingList struct {
	Items    []MonitorRadiusAccounting `json:"items,omitempty"`
	Kind     string                    `json:"kind,omitempty"`
	SelfLink string                    `json:"selflink,omitempty"`
}

// MonitorRadiusAccounting holds the configuration of a single MonitorRadiusAccounting.
type MonitorRadiusAccounting struct {
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

// MonitorRadiusAccountingEndpoint represents the REST resource for managing MonitorRadiusAccounting.
const MonitorRadiusAccountingEndpoint = "/monitor/radius-accounting"

// MonitorRadiusAccountingResource provides an API to manage MonitorRadiusAccounting configurations.
type MonitorRadiusAccountingResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorRadiusAccounting configurations.
func (r *MonitorRadiusAccountingResource) ListAll() (*MonitorRadiusAccountingList, error) {
	var list MonitorRadiusAccountingList
	if err := r.c.ReadQuery(BasePath+MonitorRadiusAccountingEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorRadiusAccounting configuration identified by id.
func (r *MonitorRadiusAccountingResource) Get(id string) (*MonitorRadiusAccounting, error) {
	var item MonitorRadiusAccounting
	if err := r.c.ReadQuery(BasePath+MonitorRadiusAccountingEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorRadiusAccounting configuration.
func (r *MonitorRadiusAccountingResource) Create(item MonitorRadiusAccounting) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRadiusAccountingEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorRadiusAccounting configuration identified by id.
func (r *MonitorRadiusAccountingResource) Edit(id string, item MonitorRadiusAccounting) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRadiusAccountingEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorRadiusAccounting configuration identified by id.
func (r *MonitorRadiusAccountingResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRadiusAccountingEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
