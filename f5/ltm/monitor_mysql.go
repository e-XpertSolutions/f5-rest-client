// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorMySQLConfigList struct {
	Items    []MonitorMySQLConfig `json:"items,omitempty"`
	Kind     string               `json:"kind,omitempty"`
	SelfLink string               `json:"selflink,omitempty"`
}

type MonitorMySQLConfig struct {
	AppService   string `json:"appService,omitempty"`
	Count        string `json:"count,omitempty"`
	Database     string `json:"database,omitempty"`
	Debug        string `json:"debug,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	Recv         string `json:"recv,omitempty"`
	RecvColumn   string `json:"recvColumn,omitempty"`
	RecvRow      string `json:"recvRow,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	Send         string `json:"send,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const MonitorMySQLEndpoint = "/monitor/mysql"

type MonitorMySQLResource struct {
	c *f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorMySQLEndpoint+"/"+id, &item); err != nil {
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
