// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorRadiusConfigList holds a list of MonitorRadius configuration.
type MonitorRadiusConfigList struct {
	Items    []MonitorRadiusConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// MonitorRadiusConfig holds the configuration of a single MonitorRadius.
type MonitorRadiusConfig struct {
	Debug              string `json:"debug"`
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeTimeout       int    `json:"probeTimeout"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
}

// MonitorRadiusEndpoint represents the REST resource for managing MonitorRadius.
const MonitorRadiusEndpoint = "/monitor/radius"

// MonitorRadiusResource provides an API to manage MonitorRadius configurations.
type MonitorRadiusResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorRadius configurations.
func (r *MonitorRadiusResource) ListAll() (*MonitorRadiusConfigList, error) {
	var list MonitorRadiusConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRadiusEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorRadius configuration identified by id.
func (r *MonitorRadiusResource) Get(id string) (*MonitorRadiusConfig, error) {
	var item MonitorRadiusConfig
	if err := r.c.ReadQuery(BasePath+MonitorRadiusEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorRadius configuration.
func (r *MonitorRadiusResource) Create(item MonitorRadiusConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRadiusEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorRadius configuration identified by id.
func (r *MonitorRadiusResource) Edit(id string, item MonitorRadiusConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRadiusEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorRadius configuration identified by id.
func (r *MonitorRadiusResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRadiusEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
