// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorRadiusAccountingConfigList struct {
	Items    []MonitorRadiusAccountingConfig `json:"items"`
	Kind     string                          `json:"kind"`
	SelfLink string                          `json:"selflink"`
}

type MonitorRadiusAccountingConfig struct {
	Debug        string `json:"debug"`
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

const MonitorRadiusAccountingEndpoint = "/monitor/rarius-accounting"

type MonitorRadiusAccountingResource struct {
	c f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorRadiusAccountingEndpoint, &item); err != nil {
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
