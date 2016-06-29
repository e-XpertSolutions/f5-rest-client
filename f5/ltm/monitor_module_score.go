// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorModuleScoreConfigList struct {
	Items    []MonitorModuleScoreConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

type MonitorModuleScoreConfig struct {
	Debug         string `json:"debug"`
	FullPath      string `json:"fullPath"`
	Generation    int    `json:"generation"`
	Interval      int    `json:"interval"`
	Kind          string `json:"kind"`
	Name          string `json:"name"`
	Partition     string `json:"partition"`
	SelfLink      string `json:"selfLink"`
	SnmpCommunity string `json:"snmpCommunity"`
	SnmpPort      int    `json:"snmpPort"`
	SnmpVersion   string `json:"snmpVersion"`
	TimeUntilUp   int    `json:"timeUntilUp"`
	Timeout       int    `json:"timeout"`
	UpInterval    int    `json:"upInterval"`
}

const MonitorModuleScoreEndpoint = "/monitor/module-score"

type MonitorModuleScoreResource struct {
	c f5.Client
}

func (r *MonitorModuleScoreResource) ListAll() (*MonitorModuleScoreConfigList, error) {
	var list MonitorModuleScoreConfigList
	if err := r.c.ReadQuery(BasePath+MonitorModuleScoreEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorModuleScoreResource) Get(id string) (*MonitorModuleScoreConfig, error) {
	var item MonitorModuleScoreConfig
	if err := r.c.ReadQuery(BasePath+MonitorModuleScoreEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorModuleScoreResource) Create(item MonitorModuleScoreConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorModuleScoreEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorModuleScoreResource) Edit(id string, item MonitorModuleScoreConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorModuleScoreEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorModuleScoreResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorModuleScoreEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
