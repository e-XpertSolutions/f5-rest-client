// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorGTPConfigList holds a list of MonitorGTP configuration.
type MonitorGTPConfigList struct {
	Items    []MonitorGTPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// MonitorGTPConfig holds the configuration of a single MonitorGTP.
type MonitorGTPConfig struct {
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeAttempts      int    `json:"probeAttempts"`
	ProbeInterval      int    `json:"probeInterval"`
	ProbeTimeout       int    `json:"probeTimeout"`
	ProtocolVersion    int    `json:"protocolVersion"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
}

// MonitorGTPEndpoint represents the REST resource for managing MonitorGTP.
const MonitorGTPEndpoint = "/monitor/gtp"

// MonitorGTPResource provides an API to manage MonitorGTP configurations.
type MonitorGTPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorGTP configurations.
func (r *MonitorGTPResource) ListAll() (*MonitorGTPConfigList, error) {
	var list MonitorGTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorGTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorGTP configuration identified by id.
func (r *MonitorGTPResource) Get(id string) (*MonitorGTPConfig, error) {
	var item MonitorGTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorGTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorGTP configuration.
func (r *MonitorGTPResource) Create(item MonitorGTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorGTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorGTP configuration identified by id.
func (r *MonitorGTPResource) Edit(id string, item MonitorGTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorGTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorGTP configuration identified by id.
func (r *MonitorGTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorGTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
