// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorIMAPConfigList holds a list of MonitorIMAP configuration.
type MonitorIMAPConfigList struct {
	Items    []MonitorIMAPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorIMAPConfig holds the configuration of a single MonitorIMAP.
type MonitorIMAPConfig struct {
	Debug              string `json:"debug"`
	Destination        string `json:"destination"`
	Folder             string `json:"folder"`
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

// MonitorIMAPEndpoint represents the REST resource for managing MonitorIMAP.
const MonitorIMAPEndpoint = "/monitor/imap"

// MonitorIMAPResource provides an API to manage MonitorIMAP configurations.
type MonitorIMAPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorIMAP configurations.
func (r *MonitorIMAPResource) ListAll() (*MonitorIMAPConfigList, error) {
	var list MonitorIMAPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorIMAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorIMAP configuration identified by id.
func (r *MonitorIMAPResource) Get(id string) (*MonitorIMAPConfig, error) {
	var item MonitorIMAPConfig
	if err := r.c.ReadQuery(BasePath+MonitorIMAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorIMAP configuration.
func (r *MonitorIMAPResource) Create(item MonitorIMAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorIMAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorIMAP configuration identified by id.
func (r *MonitorIMAPResource) Edit(id string, item MonitorIMAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorIMAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorIMAP configuration identified by id.
func (r *MonitorIMAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorIMAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
