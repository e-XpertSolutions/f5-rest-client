// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSMTPConfigList holds a list of MonitorSMTP configuration.
type MonitorSMTPConfigList struct {
	Items    []MonitorSMTPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorSMTPConfig holds the configuration of a single MonitorSMTP.
type MonitorSMTPConfig struct {
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

// MonitorSMTPEndpoint represents the REST resource for managing MonitorSMTP.
const MonitorSMTPEndpoint = "/monitor/smtp"

// MonitorSMTPResource provides an API to manage MonitorSMTP configurations.
type MonitorSMTPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorSMTP configurations.
func (r *MonitorSMTPResource) ListAll() (*MonitorSMTPConfigList, error) {
	var list MonitorSMTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSMTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSMTP configuration identified by id.
func (r *MonitorSMTPResource) Get(id string) (*MonitorSMTPConfig, error) {
	var item MonitorSMTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSMTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSMTP configuration.
func (r *MonitorSMTPResource) Create(item MonitorSMTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSMTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSMTP configuration identified by id.
func (r *MonitorSMTPResource) Edit(id string, item MonitorSMTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSMTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorSMTP configuration identified by id.
func (r *MonitorSMTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSMTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
