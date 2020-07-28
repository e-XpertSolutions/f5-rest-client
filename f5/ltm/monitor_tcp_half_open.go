// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorTCPHalfOpenConfigList struct {
	Items    []MonitorTCPHalfOpenConfig `json:"items,omitempty"`
	Kind     string                     `json:"kind,omitempty"`
	SelfLink string                     `json:"selflink,omitempty"`
}

type MonitorTCPHalfOpenConfig struct {
	AppService   string `json:"appService,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
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
	Transparent  string `json:"transparent,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const MonitorTCPHalfOpenEndpoint = "/monitor/tcp-half-open"

type MonitorTCPHalfOpenResource struct {
	c *f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorTCPHalfOpenEndpoint+"/"+id, &item); err != nil {
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
