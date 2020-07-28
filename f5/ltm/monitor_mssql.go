// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorMSSQLConfigList struct {
	Items    []MonitorMSSQLConfig `json:"items,omitempty"`
	Kind     string               `json:"kind,omitempty"`
	SelfLink string               `json:"selflink,omitempty"`
}

type MonitorMSSQLConfig struct {
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

const MonitorMSSQLEndpoint = "/monitor/mssql"

type MonitorMSSQLResource struct {
	c *f5.Client
}

func (r *MonitorMSSQLResource) ListAll() (*MonitorMSSQLConfigList, error) {
	var list MonitorMSSQLConfigList
	if err := r.c.ReadQuery(BasePath+MonitorMSSQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorMSSQLResource) Get(id string) (*MonitorMSSQLConfig, error) {
	var item MonitorMSSQLConfig
	if err := r.c.ReadQuery(BasePath+MonitorMSSQLEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorMSSQLResource) Create(item MonitorMSSQLConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorMSSQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorMSSQLResource) Edit(id string, item MonitorMSSQLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorMSSQLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorMSSQLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorMSSQLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
