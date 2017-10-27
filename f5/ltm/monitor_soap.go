// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSOAPConfigList struct {
	Items    []MonitorSOAPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorSOAPConfig struct {
	Debug        string `json:"debug,omitempty"`
	Destination  string `json:"destination,omitempty"`
	ExpectFault  string `json:"expectFault,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	Protocol     string `json:"protocol,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const MonitorSOAPEndpoint = "/monitor/soap"

type MonitorSOAPResource struct {
	c *f5.Client
}

func (r *MonitorSOAPResource) ListAll() (*MonitorSOAPConfigList, error) {
	var list MonitorSOAPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSOAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSOAPResource) Get(id string) (*MonitorSOAPConfig, error) {
	var item MonitorSOAPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSOAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSOAPResource) Create(item MonitorSOAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSOAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSOAPResource) Edit(id string, item MonitorSOAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSOAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSOAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSOAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
