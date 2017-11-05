// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorRealServerList holds a list of MonitorRealServer configuration.
type MonitorRealServerList struct {
	Items    []MonitorRealServer `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

// MonitorRealServer holds the configuration of a single MonitorRealServer.
type MonitorRealServer struct {
	Agent              string `json:"agent,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Method             string `json:"method,omitempty"`
	Metrics            string `json:"metrics,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
	TmCommand          string `json:"tmCommand,omitempty"`
}

// MonitorRealServerEndpoint represents the REST resource for managing MonitorRealServer.
const MonitorRealServerEndpoint = "/monitor/real-server"

// MonitorRealServerResource provides an API to manage MonitorRealServer configurations.
type MonitorRealServerResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorRealServer configurations.
func (r *MonitorRealServerResource) ListAll() (*MonitorRealServerList, error) {
	var list MonitorRealServerList
	if err := r.c.ReadQuery(BasePath+MonitorRealServerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorRealServer configuration identified by id.
func (r *MonitorRealServerResource) Get(id string) (*MonitorRealServer, error) {
	var item MonitorRealServer
	if err := r.c.ReadQuery(BasePath+MonitorRealServerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorRealServer configuration.
func (r *MonitorRealServerResource) Create(item MonitorRealServer) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRealServerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorRealServer configuration identified by id.
func (r *MonitorRealServerResource) Edit(id string, item MonitorRealServer) error {
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
