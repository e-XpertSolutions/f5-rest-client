// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorTCPHalfOpenConfigList struct {
	Items    []MonitorTCPHalfOpenConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

type MonitorTCPHalfOpenConfig struct {
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
	Transparent  string `json:"transparent"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorTCPHalfOpenEndpoint = "/monitor/tcp-half-open"

type MonitorTCPHalfOpenResource struct {
	c f5.Client
}

func (r *MonitorTCPHalfOpenResource) ListAll() (*MonitorTCPHalfOpenConfigList, error) {
	var list MonitorTCPHalfOpenConfigList
	if err := r.c.ReadQuery(BasePath+MonitorTCPHalfOpenEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorTCPHalfOpenResource) Get(id string) (*MonitorTCPHalfOpenConfig, error) {
	var item MonitorTCPHalfOpenConfig
	if err := r.c.ReadQuery(BasePath+MonitorTCPHalfOpenEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorTCPHalfOpenResource) Create(item MonitorTCPHalfOpenConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorTCPHalfOpenEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorTCPHalfOpenResource) Edit(id string, item MonitorTCPHalfOpenConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorTCPHalfOpenEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorTCPHalfOpenResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorTCPHalfOpenEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
