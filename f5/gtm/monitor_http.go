// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorHTTPConfigList holds a list of MonitorHTTP configuration.
type MonitorHTTPConfigList struct {
	Items    []MonitorHTTPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorHTTPConfig holds the configuration of a single MonitorHTTP.
type MonitorHTTPConfig struct {
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeTimeout       int    `json:"probeTimeout"`
	Reverse            string `json:"reverse"`
	SelfLink           string `json:"selfLink"`
	Send               string `json:"send"`
	Timeout            int    `json:"timeout"`
	Transparent        string `json:"transparent"`
}

// MonitorHTTPEndpoint represents the REST resource for managing MonitorHTTP.
const MonitorHTTPEndpoint = "/monitor/http"

// MonitorHTTPResource provides an API to manage MonitorHTTP configurations.
type MonitorHTTPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorHTTP configurations.
func (r *MonitorHTTPResource) ListAll() (*MonitorHTTPConfigList, error) {
	var list MonitorHTTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorHTTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorHTTP configuration identified by id.
func (r *MonitorHTTPResource) Get(id string) (*MonitorHTTPConfig, error) {
	var item MonitorHTTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorHTTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorHTTP configuration.
func (r *MonitorHTTPResource) Create(item MonitorHTTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorHTTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorHTTP configuration identified by id.
func (r *MonitorHTTPResource) Edit(id string, item MonitorHTTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorHTTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorHTTP configuration identified by id.
func (r *MonitorHTTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorHTTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
