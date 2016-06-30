// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorWMIConfigList holds a list of MonitorWMI configuration.
type MonitorWMIConfigList struct {
	Items    []MonitorWMIConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// MonitorWMIConfig holds the configuration of a single MonitorWMI.
type MonitorWMIConfig struct {
	Agent              string `json:"agent"`
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Method             string `json:"method"`
	Metrics            string `json:"metrics"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	Post               string `json:"post"`
	ProbeTimeout       int    `json:"probeTimeout"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
	TmCommand          string `json:"tmCommand"`
	URL                string `json:"url"`
}

// MonitorWMIEndpoint represents the REST resource for managing MonitorWMI.
const MonitorWMIEndpoint = "/monitor/wmi"

// MonitorWMIResource provides an API to manage MonitorWMI configurations.
type MonitorWMIResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorWMI configurations.
func (r *MonitorWMIResource) ListAll() (*MonitorWMIConfigList, error) {
	var list MonitorWMIConfigList
	if err := r.c.ReadQuery(BasePath+MonitorWMIEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorWMI configuration identified by id.
func (r *MonitorWMIResource) Get(id string) (*MonitorWMIConfig, error) {
	var item MonitorWMIConfig
	if err := r.c.ReadQuery(BasePath+MonitorWMIEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorWMI configuration.
func (r *MonitorWMIResource) Create(item MonitorWMIConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorWMIEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorWMI configuration identified by id.
func (r *MonitorWMIResource) Edit(id string, item MonitorWMIConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorWMIEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorWMI configuration identified by id.
func (r *MonitorWMIResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorWMIEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
