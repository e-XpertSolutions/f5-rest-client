// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSNMPDCAConfigList struct {
	Items    []MonitorSNMPDCAConfig `json:"items,omitempty"`
	Kind     string                 `json:"kind,omitempty"`
	SelfLink string                 `json:"selflink,omitempty"`
}

type MonitorSNMPDCAConfig struct {
	AgentType         string `json:"agentType,omitempty"`
	Community         string `json:"community,omitempty"`
	CPUCoefficient    string `json:"cpuCoefficient,omitempty"`
	CPUThreshold      string `json:"cpuThreshold,omitempty"`
	DiskCoefficient   string `json:"diskCoefficient,omitempty"`
	DiskThreshold     string `json:"diskThreshold,omitempty"`
	FullPath          string `json:"fullPath,omitempty"`
	Generation        int    `json:"generation,omitempty"`
	Interval          int    `json:"interval,omitempty"`
	Kind              string `json:"kind,omitempty"`
	MemoryCoefficient string `json:"memoryCoefficient,omitempty"`
	MemoryThreshold   string `json:"memoryThreshold,omitempty"`
	Name              string `json:"name,omitempty"`
	Partition         string `json:"partition,omitempty"`
	SelfLink          string `json:"selfLink,omitempty"`
	TimeUntilUp       int    `json:"timeUntilUp,omitempty"`
	Timeout           int    `json:"timeout,omitempty"`
	Version           string `json:"version,omitempty"`
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
