// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorRealServerConfigList struct {
	Items    []MonitorRealServerConfig `json:"items,omitempty"`
	Kind     string                    `json:"kind,omitempty"`
	SelfLink string                    `json:"selflink,omitempty"`
}

type MonitorRealServerConfig struct {
	Agent       string `json:"agent,omitempty"`
	FullPath    string `json:"fullPath,omitempty"`
	Generation  int    `json:"generation,omitempty"`
	Interval    int    `json:"interval,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Method      string `json:"method,omitempty"`
	Metrics     string `json:"metrics,omitempty"`
	Name        string `json:"name,omitempty"`
	Partition   string `json:"partition,omitempty"`
	SelfLink    string `json:"selfLink,omitempty"`
	TimeUntilUp int    `json:"timeUntilUp,omitempty"`
	Timeout     int    `json:"timeout,omitempty"`
	TmCommand   string `json:"tmCommand,omitempty"`
}

const MonitorRealServerEndpoint = "/monitor/real-server"

type MonitorRealServerResource struct {
	c *f5.Client
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
