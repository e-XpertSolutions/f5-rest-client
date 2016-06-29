// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSOAPConfigList struct {
	Items    []MonitorSOAPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorSOAPConfig struct {
	Debug        string `json:"debug"`
	Destination  string `json:"destination"`
	ExpectFault  string `json:"expectFault"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	Protocol     string `json:"protocol"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorSOAPEndpoint = "/monitor/soap"

type MonitorSOAPResource struct {
	c f5.Client
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
