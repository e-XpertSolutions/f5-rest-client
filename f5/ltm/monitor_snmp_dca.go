// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSNMPDCAConfigList struct {
	Items    []MonitorSNMPDCAConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

type MonitorSNMPDCAConfig struct {
	AgentType         string `json:"agentType"`
	Community         string `json:"community"`
	CPUCoefficient    string `json:"cpuCoefficient"`
	CPUThreshold      string `json:"cpuThreshold"`
	DiskCoefficient   string `json:"diskCoefficient"`
	DiskThreshold     string `json:"diskThreshold"`
	FullPath          string `json:"fullPath"`
	Generation        int    `json:"generation"`
	Interval          int    `json:"interval"`
	Kind              string `json:"kind"`
	MemoryCoefficient string `json:"memoryCoefficient"`
	MemoryThreshold   string `json:"memoryThreshold"`
	Name              string `json:"name"`
	Partition         string `json:"partition"`
	SelfLink          string `json:"selfLink"`
	TimeUntilUp       int    `json:"timeUntilUp"`
	Timeout           int    `json:"timeout"`
	Version           string `json:"version"`
}

const MonitorSNMPDCAEndpoint = "/monitor/snmp-dca"

type MonitorSNMPDCAResource struct {
	c *f5.Client
}

func (r *MonitorSNMPDCAResource) ListAll() (*MonitorSNMPDCAConfigList, error) {
	var list MonitorSNMPDCAConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSNMPDCAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSNMPDCAResource) Get(id string) (*MonitorSNMPDCAConfig, error) {
	var item MonitorSNMPDCAConfig
	if err := r.c.ReadQuery(BasePath+MonitorSNMPDCAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSNMPDCAResource) Create(item MonitorSNMPDCAConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSNMPDCAEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSNMPDCAResource) Edit(id string, item MonitorSNMPDCAConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSNMPDCAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSNMPDCAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSNMPDCAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
