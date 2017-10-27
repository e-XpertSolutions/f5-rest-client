// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorOracleConfigList struct {
	Items    []MonitorOracleConfig `json:"items,omitempty"`
	Kind     string                `json:"kind,omitempty"`
	SelfLink string                `json:"selflink,omitempty"`
}

type MonitorOracleConfig struct {
	Count        string `json:"count,omitempty"`
	Database     string `json:"database,omitempty"`
	Debug        string `json:"debug,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
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
