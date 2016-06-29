// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSNMPDCABaseConfigList struct {
	Items    []MonitorSNMPDCABaseConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

type MonitorSNMPDCABaseConfig struct {
	Community   string `json:"community"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Interval    int    `json:"interval"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	SelfLink    string `json:"selfLink"`
	TimeUntilUp int    `json:"timeUntilUp"`
	Timeout     int    `json:"timeout"`
	Version     string `json:"version"`
}

const MonitorSNMPDCABaseEndpoint = "/monitor/snmp-dca-base"

type MonitorSNMPDCABaseResource struct {
	c f5.Client
}

func (r *MonitorSNMPDCABaseResource) ListAll() (*MonitorSNMPDCABaseConfigList, error) {
	var list MonitorSNMPDCABaseConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSNMPDCABaseEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSNMPDCABaseResource) Get(id string) (*MonitorSNMPDCABaseConfig, error) {
	var item MonitorSNMPDCABaseConfig
	if err := r.c.ReadQuery(BasePath+MonitorSNMPDCABaseEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSNMPDCABaseResource) Create(item MonitorSNMPDCABaseConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSNMPDCABaseEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSNMPDCABaseResource) Edit(id string, item MonitorSNMPDCABaseConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSNMPDCABaseEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSNMPDCABaseResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSNMPDCABaseEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
