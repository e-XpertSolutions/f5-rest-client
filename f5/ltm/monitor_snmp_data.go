// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSNMPDataConfigList struct {
	Items    []MonitorSNMPDataConfig `json:"items,omitempty"`
	Kind     string                  `json:"kind,omitempty"`
	SelfLink string                  `json:"selflink,omitempty"`
}

type MonitorSNMPDataConfig struct {
}

const MonitorSNMPDataEndpoint = "/monitor/snmp-data"

type MonitorSNMPDataResource struct {
	c *f5.Client
}

func (r *MonitorSNMPDataResource) ListAll() (*MonitorSNMPDataConfigList, error) {
	var list MonitorSNMPDataConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSNMPDataEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSNMPDataResource) Get(id string) (*MonitorSNMPDataConfig, error) {
	var item MonitorSNMPDataConfig
	if err := r.c.ReadQuery(BasePath+MonitorSNMPDataEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSNMPDataResource) Create(item MonitorSNMPDataConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSNMPDataEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSNMPDataResource) Edit(id string, item MonitorSNMPDataConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSNMPDataEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSNMPDataResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSNMPDataEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
