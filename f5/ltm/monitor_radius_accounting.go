// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorRadiusAccountingConfigList struct {
	Items    []MonitorRadiusAccountingConfig `json:"items,omitempty"`
	Kind     string                          `json:"kind,omitempty"`
	SelfLink string                          `json:"selflink,omitempty"`
}

type MonitorRadiusAccountingConfig struct {
	AppService   string `json:"appService,omitempty"`
	Debug        string `json:"debug,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Interval     int    `json:"interval,omitempty"`
	Kind         string `json:"kind,omitempty"`
	ManualResume string `json:"manualResume,omitempty"`
	Name         string `json:"name,omitempty"`
	NasIpAddress string `json:"nasIpAddress,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
}

const MonitorRadiusAccountingEndpoint = "/monitor/rarius-accounting"

type MonitorRadiusAccountingResource struct {
	c *f5.Client
}

func (r *MonitorRadiusAccountingResource) ListAll() (*MonitorRadiusAccountingConfigList, error) {
	var list MonitorRadiusAccountingConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRadiusAccountingEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorRadiusAccountingResource) Get(id string) (*MonitorRadiusAccountingConfig, error) {
	var item MonitorRadiusAccountingConfig
	if err := r.c.ReadQuery(BasePath+MonitorRadiusAccountingEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorRadiusAccountingResource) Create(item MonitorRadiusAccountingConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRadiusAccountingEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRadiusAccountingResource) Edit(id string, item MonitorRadiusAccountingConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRadiusAccountingEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRadiusAccountingResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRadiusAccountingEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
