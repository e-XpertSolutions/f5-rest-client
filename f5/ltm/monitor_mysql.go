// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorMySQLConfigList struct {
	Items    []MonitorMySQLConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

type MonitorMySQLConfig struct {
	Count        string `json:"count"`
	Debug        string `json:"debug"`
	Destination  string `json:"destination"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorMySQLEndpoint = "/monitor/mysql"

type MonitorMySQLResource struct {
	c f5.Client
}

func (r *MonitorMySQLResource) ListAll() (*MonitorMySQLConfigList, error) {
	var list MonitorMySQLConfigList
	if err := r.c.ReadQuery(BasePath+MonitorMySQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorMySQLResource) Get(id string) (*MonitorMySQLConfig, error) {
	var item MonitorMySQLConfig
	if err := r.c.ReadQuery(BasePath+MonitorMySQLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorMySQLResource) Create(item MonitorMySQLConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorMySQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorMySQLResource) Edit(id string, item MonitorMySQLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorMySQLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorMySQLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorMySQLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
