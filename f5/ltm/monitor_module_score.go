// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorModuleScoreConfigList struct {
	Items    []MonitorModuleScoreConfig `json:"items,omitempty"`
	Kind     string                     `json:"kind,omitempty"`
	SelfLink string                     `json:"selflink,omitempty"`
}

type MonitorModuleScoreConfig struct {
	AppService    string `json:"appService,omitempty"`
	Debug         string `json:"debug,omitempty"`
	DefaultsFrom  string `json:"defaultsFrom,omitempty"`
	Description   string `json:"description,omitempty"`
	Destination   string `json:"destination,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Kind          string `json:"kind,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	Pool          string `json:"pool,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	SnmpCommunity string `json:"snmpCommunity,omitempty"`
	SnmpIpAddress string `json:"snmpIpAddress,omitempty"`
	SnmpPort      int    `json:"snmpPort,omitempty"`
	SnmpVersion   string `json:"snmpVersion,omitempty"`
	TimeUntilUp   int    `json:"timeUntilUp,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	UpInterval    int    `json:"upInterval,omitempty"`
}

const MonitorModuleScoreEndpoint = "/monitor/module-score"

type MonitorModuleScoreResource struct {
	c *f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorModuleScoreEndpoint+"/"+id, &item); err != nil {
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
