// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorRealServerConfigList holds a list of MonitorRealServer configuration.
type MonitorRealServerConfigList struct {
	Items    []MonitorRealServerConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

// MonitorRealServerConfig holds the configuration of a single MonitorRealServer.
type MonitorRealServerConfig struct {
	Agent              string `json:"agent"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Method             string `json:"method"`
	Metrics            string `json:"metrics"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeTimeout       int    `json:"probeTimeout"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
	TmCommand          string `json:"tmCommand"`
}

// MonitorRealServerEndpoint represents the REST resource for managing MonitorRealServer.
const MonitorRealServerEndpoint = "/monitor/real-server"

// MonitorRealServerResource provides an API to manage MonitorRealServer configurations.
type MonitorRealServerResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorRealServer configurations.
func (r *MonitorRealServerResource) ListAll() (*MonitorRealServerConfigList, error) {
	var list MonitorRealServerConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRealServerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorRealServer configuration identified by id.
func (r *MonitorRealServerResource) Get(id string) (*MonitorRealServerConfig, error) {
	var item MonitorRealServerConfig
	if err := r.c.ReadQuery(BasePath+MonitorRealServerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorRealServer configuration.
func (r *MonitorRealServerResource) Create(item MonitorRealServerConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRealServerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorRealServer configuration identified by id.
func (r *MonitorRealServerResource) Edit(id string, item MonitorRealServerConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRealServerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorRealServer configuration identified by id.
func (r *MonitorRealServerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRealServerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
