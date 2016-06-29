// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorScriptedConfigList struct {
	Items    []MonitorScriptedConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

type MonitorScriptedConfig struct {
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

const MonitorScriptedEndpoint = "/monitor/scripted"

type MonitorScriptedResource struct {
	c f5.Client
}

func (r *MonitorScriptedResource) ListAll() (*MonitorScriptedConfigList, error) {
	var list MonitorScriptedConfigList
	if err := r.c.ReadQuery(BasePath+MonitorScriptedEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorScriptedResource) Get(id string) (*MonitorScriptedConfig, error) {
	var item MonitorScriptedConfig
	if err := r.c.ReadQuery(BasePath+MonitorScriptedEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorScriptedResource) Create(item MonitorScriptedConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorScriptedEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorScriptedResource) Edit(id string, item MonitorScriptedConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorScriptedEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorScriptedResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorScriptedEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
