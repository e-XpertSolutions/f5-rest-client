// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorHTTPSConfigList holds a list of MonitorHTTPS configuration.
type MonitorHTTPSConfigList struct {
	Items    []MonitorHTTPSConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// MonitorHTTPSConfig holds the configuration of a single MonitorHTTPS.
type MonitorHTTPSConfig struct {
	Cipherlist         string `json:"cipherlist"`
	Compatibility      string `json:"compatibility"`
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

// MonitorHTTPSEndpoint represents the REST resource for managing MonitorHTTPS.
const MonitorHTTPSEndpoint = "/monitor/https"

// MonitorHTTPSResource provides an API to manage MonitorHTTPS configurations.
type MonitorHTTPSResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorHTTPS configurations.
func (r *MonitorHTTPSResource) ListAll() (*MonitorHTTPSConfigList, error) {
	var list MonitorHTTPSConfigList
	if err := r.c.ReadQuery(BasePath+MonitorHTTPSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorHTTPS configuration identified by id.
func (r *MonitorHTTPSResource) Get(id string) (*MonitorHTTPSConfig, error) {
	var item MonitorHTTPSConfig
	if err := r.c.ReadQuery(BasePath+MonitorHTTPSEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorHTTPS configuration.
func (r *MonitorHTTPSResource) Create(item MonitorHTTPSConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorHTTPSEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorHTTPS configuration identified by id.
func (r *MonitorHTTPSResource) Edit(id string, item MonitorHTTPSConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorHTTPSEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorHTTPS configuration identified by id.
func (r *MonitorHTTPSResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorHTTPSEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
