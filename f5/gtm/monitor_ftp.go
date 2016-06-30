// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorFTPConfigList holds a list of MonitorFTP configuration.
type MonitorFTPConfigList struct {
	Items    []MonitorFTPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// MonitorFTPConfig holds the configuration of a single MonitorFTP.
type MonitorFTPConfig struct {
	Debug              string `json:"debug"`
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Mode               string `json:"mode"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeTimeout       int    `json:"probeTimeout"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
}

// MonitorFTPEndpoint represents the REST resource for managing MonitorFTP.
const MonitorFTPEndpoint = "/monitor/ftp"

// MonitorFTPResource provides an API to manage MonitorFTP configurations.
type MonitorFTPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorFTP configurations.
func (r *MonitorFTPResource) ListAll() (*MonitorFTPConfigList, error) {
	var list MonitorFTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorFTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorFTP configuration identified by id.
func (r *MonitorFTPResource) Get(id string) (*MonitorFTPConfig, error) {
	var item MonitorFTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorFTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorFTP configuration.
func (r *MonitorFTPResource) Create(item MonitorFTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorFTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorFTP configuration identified by id.
func (r *MonitorFTPResource) Edit(id string, item MonitorFTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorFTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorFTP configuration identified by id.
func (r *MonitorFTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorFTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
