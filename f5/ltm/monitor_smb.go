// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSMBConfigList struct {
	Items    []MonitorSMBConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorSMBConfig struct {
	AppService   string `json:"appService,omitempty"`
	Debug        string `json:"debug,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Get          string `json:"get,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	Server       string `json:"server,omitempty"`
	Service      string `json:"service,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const MonitorSMBEndpoint = "/monitor/smb"

type MonitorSMBResource struct {
	c *f5.Client
}

func (r *MonitorSMBResource) ListAll() (*MonitorSMBConfigList, error) {
	var list MonitorSMBConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSMBEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSMBResource) Get(id string) (*MonitorSMBConfig, error) {
	var item MonitorSMBConfig
	if err := r.c.ReadQuery(BasePath+MonitorSMBEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSMBResource) Create(item MonitorSMBConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSMBEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSMBResource) Edit(id string, item MonitorSMBConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSMBEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSMBResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSMBEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
