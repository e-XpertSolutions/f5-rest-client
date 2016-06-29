// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorRPCConfigList struct {
	Items    []MonitorRPCConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorRPCConfig struct {
	Debug        string `json:"debug"`
	Destination  string `json:"destination"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Mode         string `json:"mode"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorRPCEndpoint = "/monitor/rpc"

type MonitorRPCResource struct {
	c f5.Client
}

func (r *MonitorRPCResource) ListAll() (*MonitorRPCConfigList, error) {
	var list MonitorRPCConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRPCEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorRPCResource) Get(id string) (*MonitorRPCConfig, error) {
	var item MonitorRPCConfig
	if err := r.c.ReadQuery(BasePath+MonitorRPCEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorRPCResource) Create(item MonitorRPCConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRPCEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRPCResource) Edit(id string, item MonitorRPCConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRPCEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRPCResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRPCEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
