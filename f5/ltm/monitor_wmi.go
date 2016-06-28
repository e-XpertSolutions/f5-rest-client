// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorWMIConfigList struct {
	Items    []MonitorWMIConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorWMIConfig struct {
	Agent       string `json:"agent"`
	Destination string `json:"destination"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Interval    int    `json:"interval"`
	Kind        string `json:"kind"`
	Method      string `json:"method"`
	Metrics     string `json:"metrics"`
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	Post        string `json:"post"`
	SelfLink    string `json:"selfLink"`
	TimeUntilUp int    `json:"timeUntilUp"`
	Timeout     int    `json:"timeout"`
	TmCommand   string `json:"tmCommand"`
	URL         string `json:"url"`
}

const MonitorWMIEndpoint = "/monitor/wmi"

type MonitorWMIResource struct {
	c f5.Client
}

func (r *MonitorWMIResource) ListAll() (*MonitorWMIConfigList, error) {
	var list MonitorWMIConfigList
	if err := r.c.ReadQuery(BasePath+MonitorWMIEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorWMIResource) Get(id string) (*MonitorWMIConfig, error) {
	var item MonitorWMIConfig
	if err := r.c.ReadQuery(BasePath+MonitorWMIEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorWMIResource) Create(item MonitorWMIConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorWMIEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorWMIResource) Edit(id string, item MonitorWMIConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorWMIEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorWMIResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorWMIEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
