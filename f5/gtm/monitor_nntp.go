// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorNNTPConfigList holds a list of MonitorNNTP configuration.
type MonitorNNTPConfigList struct {
	Items    []MonitorNNTPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorNNTPConfig holds the configuration of a single MonitorNNTP.
type MonitorNNTPConfig struct {
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

// MonitorNNTPEndpoint represents the REST resource for managing MonitorNNTP.
const MonitorNNTPEndpoint = "/monitor/nntp"

// MonitorNNTPResource provides an API to manage MonitorNNTP configurations.
type MonitorNNTPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorNNTP configurations.
func (r *MonitorNNTPResource) ListAll() (*MonitorNNTPConfigList, error) {
	var list MonitorNNTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorNNTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorNNTP configuration identified by id.
func (r *MonitorNNTPResource) Get(id string) (*MonitorNNTPConfig, error) {
	var item MonitorNNTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorNNTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorNNTP configuration.
func (r *MonitorNNTPResource) Create(item MonitorNNTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorNNTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorNNTP configuration identified by id.
func (r *MonitorNNTPResource) Edit(id string, item MonitorNNTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorNNTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorNNTP configuration identified by id.
func (r *MonitorNNTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorNNTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
