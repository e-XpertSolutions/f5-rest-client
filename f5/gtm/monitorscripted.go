// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorscriptedConfigList holds a list of Monitorscripted configuration.
type MonitorscriptedConfigList struct {
	Items    []MonitorscriptedConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// MonitorscriptedConfig holds the configuration of a single Monitorscripted.
type MonitorscriptedConfig struct {
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

// MonitorscriptedEndpoint represents the REST resource for managing Monitorscripted.
const MonitorscriptedEndpoint = "/monitor/scripted"

// MonitorscriptedResource provides an API to manage Monitorscripted configurations.
type MonitorscriptedResource struct {
	c f5.Client
}

// ListAll  lists all the Monitorscripted configurations.
func (r *MonitorscriptedResource) ListAll() (*MonitorscriptedConfigList, error) {
	var list MonitorscriptedConfigList
	if err := r.c.ReadQuery(BasePath+MonitorscriptedEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Monitorscripted configuration identified by id.
func (r *MonitorscriptedResource) Get(id string) (*MonitorscriptedConfig, error) {
	var item MonitorscriptedConfig
	if err := r.c.ReadQuery(BasePath+MonitorscriptedEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Monitorscripted configuration.
func (r *MonitorscriptedResource) Create(item MonitorscriptedConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorscriptedEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Monitorscripted configuration identified by id.
func (r *MonitorscriptedResource) Edit(id string, item MonitorscriptedConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorscriptedEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Monitorscripted configuration identified by id.
func (r *MonitorscriptedResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorscriptedEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
