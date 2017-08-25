// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorRadiusConfigList struct {
	Items    []MonitorRadiusConfig `json:"items,omitempty"`
	Kind     string                `json:"kind,omitempty"`
	SelfLink string                `json:"selflink,omitempty"`
}

type MonitorRadiusConfig struct {
	Debug        string `json:"debug,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const MonitorRadiusEndpoint = "/monitor/radius"

type MonitorRadiusResource struct {
	c *f5.Client
}

func (r *MonitorRadiusResource) ListAll() (*MonitorRadiusConfigList, error) {
	var list MonitorRadiusConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRadiusEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorRadiusResource) Get(id string) (*MonitorRadiusConfig, error) {
	var item MonitorRadiusConfig
	if err := r.c.ReadQuery(BasePath+MonitorRadiusEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorRadiusResource) Create(item MonitorRadiusConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRadiusEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRadiusResource) Edit(id string, item MonitorRadiusConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRadiusEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRadiusResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRadiusEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
