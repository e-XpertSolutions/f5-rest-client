// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorWAPConfigList struct {
	Items    []MonitorWAPConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorWAPConfig struct {
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

const MonitorWAPEndpoint = "/monitor/wap"

type MonitorWAPResource struct {
	c *f5.Client
}

func (r *MonitorWAPResource) ListAll() (*MonitorWAPConfigList, error) {
	var list MonitorWAPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorWAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorWAPResource) Get(id string) (*MonitorWAPConfig, error) {
	var item MonitorWAPConfig
	if err := r.c.ReadQuery(BasePath+MonitorWAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorWAPResource) Create(item MonitorWAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorWAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorWAPResource) Edit(id string, item MonitorWAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorWAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorWAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorWAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
