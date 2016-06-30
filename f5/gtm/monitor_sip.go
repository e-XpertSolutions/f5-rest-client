// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSIPConfigList holds a list of MonitorSIP configuration.
type MonitorSIPConfigList struct {
	Items    []MonitorSIPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// MonitorSIPConfig holds the configuration of a single MonitorSIP.
type MonitorSIPConfig struct {
}

// MonitorSIPEndpoint represents the REST resource for managing MonitorSIP.
const MonitorSIPEndpoint = "/monitor/sip"

// MonitorSIPResource provides an API to manage MonitorSIP configurations.
type MonitorSIPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorSIP configurations.
func (r *MonitorSIPResource) ListAll() (*MonitorSIPConfigList, error) {
	var list MonitorSIPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSIP configuration identified by id.
func (r *MonitorSIPResource) Get(id string) (*MonitorSIPConfig, error) {
	var item MonitorSIPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSIP configuration.
func (r *MonitorSIPResource) Create(item MonitorSIPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSIP configuration identified by id.
func (r *MonitorSIPResource) Edit(id string, item MonitorSIPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorSIP configuration identified by id.
func (r *MonitorSIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
