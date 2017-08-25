// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorOracleConfigList struct {
	Items    []MonitorOracleConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

type MonitorOracleConfig struct {
	Count        string `json:"count"`
	Database     string `json:"database"`
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

const MonitorOracleEndpoint = "/monitor/oracle"

type MonitorOracleResource struct {
	c *f5.Client
}

func (r *MonitorOracleResource) ListAll() (*MonitorOracleConfigList, error) {
	var list MonitorOracleConfigList
	if err := r.c.ReadQuery(BasePath+MonitorOracleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorOracleResource) Get(id string) (*MonitorOracleConfig, error) {
	var item MonitorOracleConfig
	if err := r.c.ReadQuery(BasePath+MonitorOracleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorOracleResource) Create(item MonitorOracleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorOracleEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorOracleResource) Edit(id string, item MonitorOracleConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorOracleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorOracleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorOracleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
