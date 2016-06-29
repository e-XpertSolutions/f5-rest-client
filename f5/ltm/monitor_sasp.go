// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSASPConfigList struct {
	Items    []MonitorSASPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorSASPConfig struct {
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Interval    string `json:"interval"`
	Kind        string `json:"kind"`
	Mode        string `json:"mode"`
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	Protocol    string `json:"protocol"`
	SelfLink    string `json:"selfLink"`
	Service     string `json:"service"`
	TimeUntilUp int    `json:"timeUntilUp"`
}

const MonitorSASPEndpoint = "/monitor/sasp"

type MonitorSASPResource struct {
	c f5.Client
}

func (r *MonitorSASPResource) ListAll() (*MonitorSASPConfigList, error) {
	var list MonitorSASPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSASPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSASPResource) Get(id string) (*MonitorSASPConfig, error) {
	var item MonitorSASPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSASPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSASPResource) Create(item MonitorSASPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSASPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSASPResource) Edit(id string, item MonitorSASPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSASPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSASPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSASPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
