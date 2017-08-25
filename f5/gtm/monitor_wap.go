// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorWAPConfigList holds a list of MonitorWAP configuration.
type MonitorWAPConfigList struct {
	Items    []MonitorWAPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// MonitorWAPConfig holds the configuration of a single MonitorWAP.
type MonitorWAPConfig struct {
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

// MonitorWAPEndpoint represents the REST resource for managing MonitorWAP.
const MonitorWAPEndpoint = "/monitor/wap"

// MonitorWAPResource provides an API to manage MonitorWAP configurations.
type MonitorWAPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorWAP configurations.
func (r *MonitorWAPResource) ListAll() (*MonitorWAPConfigList, error) {
	var list MonitorWAPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorWAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorWAP configuration identified by id.
func (r *MonitorWAPResource) Get(id string) (*MonitorWAPConfig, error) {
	var item MonitorWAPConfig
	if err := r.c.ReadQuery(BasePath+MonitorWAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorWAP configuration.
func (r *MonitorWAPResource) Create(item MonitorWAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorWAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorWAP configuration identified by id.
func (r *MonitorWAPResource) Edit(id string, item MonitorWAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorWAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorWAP configuration identified by id.
func (r *MonitorWAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorWAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
