// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorSMBConfigList struct {
	Items    []MonitorSMBConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorSMBConfig struct {
	Debug        string `json:"debug"`
	Destination  string `json:"destination"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorSMBEndpoint = "/monitor/smb"

type MonitorSMBResource struct {
	c f5.Client
}

func (r *MonitorSMBResource) ListAll() (*MonitorSMBConfigList, error) {
	var list MonitorSMBConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSMBEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSMBResource) Get(id string) (*MonitorSMBConfig, error) {
	var item MonitorSMBConfig
	if err := r.c.ReadQuery(BasePath+MonitorSMBEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSMBResource) Create(item MonitorSMBConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSMBEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSMBResource) Edit(id string, item MonitorSMBConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSMBEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSMBResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSMBEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
