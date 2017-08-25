// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DaemonHAConfigList holds a list of DaemonHA configuration.
type DaemonHAConfigList struct {
	Items    []DaemonHAConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// DaemonHAConfig holds the configuration of a single DaemonHA.
type DaemonHAConfig struct {
	FullPath         string `json:"fullPath"`
	Generation       int    `json:"generation"`
	Heartbeat        string `json:"heartbeat"`
	HeartbeatAction  string `json:"heartbeatAction"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
	NotRunningAction string `json:"notRunningAction"`
	Running          string `json:"running"`
	RunningTimeout   int    `json:"runningTimeout"`
	SelfLink         string `json:"selfLink"`
}

// DaemonHAEndpoint represents the REST resource for managing DaemonHA.
const DaemonHAEndpoint = "/daemon-ha"

// DaemonHAResource provides an API to manage DaemonHA configurations.
type DaemonHAResource struct {
	c *f5.Client
}

// ListAll  lists all the DaemonHA configurations.
func (r *DaemonHAResource) ListAll() (*DaemonHAConfigList, error) {
	var list DaemonHAConfigList
	if err := r.c.ReadQuery(BasePath+DaemonHAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DaemonHA configuration identified by id.
func (r *DaemonHAResource) Get(id string) (*DaemonHAConfig, error) {
	var item DaemonHAConfig
	if err := r.c.ReadQuery(BasePath+DaemonHAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DaemonHA configuration.
func (r *DaemonHAResource) Create(item DaemonHAConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DaemonHAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DaemonHA configuration identified by id.
func (r *DaemonHAResource) Edit(id string, item DaemonHAConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DaemonHAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DaemonHA configuration identified by id.
func (r *DaemonHAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DaemonHAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
