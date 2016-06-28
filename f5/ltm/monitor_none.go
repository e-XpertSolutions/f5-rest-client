// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorNoneConfigList struct {
	Items    []MonitorNoneConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorNoneConfig struct {
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
	TimeUntilUp        int    `json:"timeUntilUp"`
	Timeout            int    `json:"timeout"`
	UpInterval         int    `json:"upInterval"`
}

const MonitorNoneEndpoint = "/monitor/none"

type MonitorNoneResource struct {
	c f5.Client
}

func (r *MonitorNoneResource) ListAll() (*MonitorNoneConfigList, error) {
	var list MonitorNoneConfigList
	if err := r.c.ReadQuery(BasePath+MonitorNoneEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorNoneResource) Get(id string) (*MonitorNoneConfig, error) {
	var item MonitorNoneConfig
	if err := r.c.ReadQuery(BasePath+MonitorNoneEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorNoneResource) Create(item MonitorNoneConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorNoneEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorNoneResource) Edit(id string, item MonitorNoneConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorNoneEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorNoneResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorNoneEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
