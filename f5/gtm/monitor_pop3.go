// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorPOP3ConfigList holds a list of MonitorPOP3 configuration.
type MonitorPOP3ConfigList struct {
	Items    []MonitorPOP3Config `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorPOP3Config holds the configuration of a single MonitorPOP3.
type MonitorPOP3Config struct {
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

// MonitorPOP3Endpoint represents the REST resource for managing MonitorPOP3.
const MonitorPOP3Endpoint = "/monitor/pop3"

// MonitorPOP3Resource provides an API to manage MonitorPOP3 configurations.
type MonitorPOP3Resource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorPOP3 configurations.
func (r *MonitorPOP3Resource) ListAll() (*MonitorPOP3ConfigList, error) {
	var list MonitorPOP3ConfigList
	if err := r.c.ReadQuery(BasePath+MonitorPOP3Endpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorPOP3 configuration identified by id.
func (r *MonitorPOP3Resource) Get(id string) (*MonitorPOP3Config, error) {
	var item MonitorPOP3Config
	if err := r.c.ReadQuery(BasePath+MonitorPOP3Endpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorPOP3 configuration.
func (r *MonitorPOP3Resource) Create(item MonitorPOP3Config) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorPOP3Endpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorPOP3 configuration identified by id.
func (r *MonitorPOP3Resource) Edit(id string, item MonitorPOP3Config) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorPOP3Endpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorPOP3 configuration identified by id.
func (r *MonitorPOP3Resource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorPOP3Endpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
