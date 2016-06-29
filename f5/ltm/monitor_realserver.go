// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorRealServerConfigList struct {
	Items    []MonitorRealServerConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

type MonitorRealServerConfig struct {
	Agent       string `json:"agent"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Interval    int    `json:"interval"`
	Kind        string `json:"kind"`
	Method      string `json:"method"`
	Metrics     string `json:"metrics"`
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	SelfLink    string `json:"selfLink"`
	TimeUntilUp int    `json:"timeUntilUp"`
	Timeout     int    `json:"timeout"`
	TmCommand   string `json:"tmCommand"`
}

const MonitorRealServerEndpoint = "/monitor/real-server"

type MonitorRealServerResource struct {
	c f5.Client
}

func (r *MonitorRealServerResource) ListAll() (*MonitorRealServerConfigList, error) {
	var list MonitorRealServerConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRealServerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorRealServerResource) Get(id string) (*MonitorRealServerConfig, error) {
	var item MonitorRealServerConfig
	if err := r.c.ReadQuery(BasePath+MonitorRealServerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorRealServerResource) Create(item MonitorRealServerConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRealServerEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRealServerResource) Edit(id string, item MonitorRealServerConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRealServerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRealServerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRealServerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
