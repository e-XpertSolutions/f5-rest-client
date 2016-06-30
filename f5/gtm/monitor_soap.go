// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorSOAPConfigList holds a list of MonitorSOAP configuration.
type MonitorSOAPConfigList struct {
	Items    []MonitorSOAPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorSOAPConfig holds the configuration of a single MonitorSOAP.
type MonitorSOAPConfig struct {
	Debug              string `json:"debug"`
	Destination        string `json:"destination"`
	ExpectFault        string `json:"expectFault"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeTimeout       int    `json:"probeTimeout"`
	Protocol           string `json:"protocol"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
}

// MonitorSOAPEndpoint represents the REST resource for managing MonitorSOAP.
const MonitorSOAPEndpoint = "/monitor/soap"

// MonitorSOAPResource provides an API to manage MonitorSOAP configurations.
type MonitorSOAPResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorSOAP configurations.
func (r *MonitorSOAPResource) ListAll() (*MonitorSOAPConfigList, error) {
	var list MonitorSOAPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSOAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorSOAP configuration identified by id.
func (r *MonitorSOAPResource) Get(id string) (*MonitorSOAPConfig, error) {
	var item MonitorSOAPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSOAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorSOAP configuration.
func (r *MonitorSOAPResource) Create(item MonitorSOAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSOAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorSOAP configuration identified by id.
func (r *MonitorSOAPResource) Edit(id string, item MonitorSOAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSOAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorSOAP configuration identified by id.
func (r *MonitorSOAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSOAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
