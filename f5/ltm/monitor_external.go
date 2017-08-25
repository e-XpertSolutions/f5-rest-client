// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorExternalConfigList struct {
	Items    []MonitorExternalConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

type MonitorExternalConfig struct {
	Destination  string `json:"destination"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorExternalEndpoint = "/monitor/external"

type MonitorExternalResource struct {
	c *f5.Client
}

func (r *MonitorExternalResource) ListAll() (*MonitorExternalConfigList, error) {
	var list MonitorExternalConfigList
	if err := r.c.ReadQuery(BasePath+MonitorExternalEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorExternalResource) Get(id string) (*MonitorExternalConfig, error) {
	var item MonitorExternalConfig
	if err := r.c.ReadQuery(BasePath+MonitorExternalEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorExternalResource) Create(item MonitorExternalConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorExternalEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorExternalResource) Edit(id string, item MonitorExternalConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorExternalEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorExternalResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorExternalEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
