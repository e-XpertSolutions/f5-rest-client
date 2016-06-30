// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorBigIPConfigList holds a list of MonitorBigIP configuration.
type MonitorBigIPConfigList struct {
	Items    []MonitorBigIPConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// MonitorBigIPConfig holds the configuration of a single MonitorBigIP.
type MonitorBigIPConfig struct {
	AggregateDynamicRatios string `json:"aggregateDynamicRatios"`
	Destination            string `json:"destination"`
	FullPath               string `json:"fullPath"`
	Generation             int    `json:"generation"`
	IgnoreDownResponse     string `json:"ignoreDownResponse"`
	Interval               int    `json:"interval"`
	Kind                   string `json:"kind"`
	Name                   string `json:"name"`
	Partition              string `json:"partition"`
	SelfLink               string `json:"selfLink"`
	Timeout                int    `json:"timeout"`
}

// MonitorBigIPEndpoint represents the REST resource for managing MonitorBigIP.
const MonitorBigIPEndpoint = "/monitor/bigip"

// MonitorBigIPResource provides an API to manage MonitorBigIP configurations.
type MonitorBigIPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorBigIP configurations.
func (r *MonitorBigIPResource) ListAll() (*MonitorBigIPConfigList, error) {
	var list MonitorBigIPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorBigIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorBigIP configuration identified by id.
func (r *MonitorBigIPResource) Get(id string) (*MonitorBigIPConfig, error) {
	var item MonitorBigIPConfig
	if err := r.c.ReadQuery(BasePath+MonitorBigIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorBigIP configuration.
func (r *MonitorBigIPResource) Create(item MonitorBigIPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorBigIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorBigIP configuration identified by id.
func (r *MonitorBigIPResource) Edit(id string, item MonitorBigIPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorBigIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorBigIP configuration identified by id.
func (r *MonitorBigIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorBigIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
