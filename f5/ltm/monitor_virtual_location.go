// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorVirtualLocationConfigList struct {
	Items    []MonitorVirtualLocationConfig `json:"items,omitempty"`
	Kind     string                         `json:"kind,omitempty"`
	SelfLink string                         `json:"selflink,omitempty"`
}

type MonitorVirtualLocationConfig struct {
	Debug       string `json:"debug,omitempty"`
	FullPath    string `json:"fullPath,omitempty"`
	Generation  int    `json:"generation,omitempty"`
	Interval    int    `json:"interval,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
	Partition   string `json:"partition,omitempty"`
	SelfLink    string `json:"selfLink,omitempty"`
	TimeUntilUp int    `json:"timeUntilUp,omitempty"`
	Timeout     int    `json:"timeout,omitempty"`
	UpInterval  int    `json:"upInterval,omitempty"`
}

const MonitorVirtualLocationEndpoint = "/monitor/virtual-location"

type MonitorVirtualLocationResource struct {
	c *f5.Client
}

func (r *MonitorVirtualLocationResource) ListAll() (*MonitorVirtualLocationConfigList, error) {
	var list MonitorVirtualLocationConfigList
	if err := r.c.ReadQuery(BasePath+MonitorVirtualLocationEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorVirtualLocationResource) Get(id string) (*MonitorVirtualLocationConfig, error) {
	var item MonitorVirtualLocationConfig
	if err := r.c.ReadQuery(BasePath+MonitorVirtualLocationEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorVirtualLocationResource) Create(item MonitorVirtualLocationConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorVirtualLocationEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorVirtualLocationResource) Edit(id string, item MonitorVirtualLocationConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorVirtualLocationEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorVirtualLocationResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorVirtualLocationEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
