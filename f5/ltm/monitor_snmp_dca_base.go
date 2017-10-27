// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSNMPDCABaseConfigList struct {
	Items    []MonitorSNMPDCABaseConfig `json:"items,omitempty"`
	Kind     string                     `json:"kind,omitempty"`
	SelfLink string                     `json:"selflink,omitempty"`
}

type MonitorSNMPDCABaseConfig struct {
	Community   string `json:"community,omitempty"`
	FullPath    string `json:"fullPath,omitempty"`
	Generation  int    `json:"generation,omitempty"`
	Interval    int    `json:"interval,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
	Partition   string `json:"partition,omitempty"`
	SelfLink    string `json:"selfLink,omitempty"`
	TimeUntilUp int    `json:"timeUntilUp,omitempty"`
	Timeout     int    `json:"timeout,omitempty"`
	Version     string `json:"version,omitempty"`
}

const MonitorSNMPDCABaseEndpoint = "/monitor/snmp-dca-base"

type MonitorSNMPDCABaseResource struct {
	c *f5.Client
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
