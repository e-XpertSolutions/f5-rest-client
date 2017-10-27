// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorExternalConfigList struct {
	Items    []MonitorExternalConfig `json:"items,omitempty"`
	Kind     string                  `json:"kind,omitempty"`
	SelfLink string                  `json:"selflink,omitempty"`
}

type MonitorExternalConfig struct {
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
