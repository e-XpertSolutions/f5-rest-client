// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorWMIConfigList struct {
	Items    []MonitorWMIConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorWMIConfig struct {
	Agent        string `json:"agent,omitempty"`
	AppService   string `json:"appService,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Method       string `json:"method,omitempty"`
	Metrics      string `json:"metrics,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	Post         string `json:"post,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	TmCommand    string `json:"tmCommand,omitempty"`
	URL          string `json:"url,omitempty"`
}

const MonitorWMIEndpoint = "/monitor/wmi"

type MonitorWMIResource struct {
	c *f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorWMIEndpoint+"/"+id, &item); err != nil {
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
