// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorNNTPConfigList struct {
	Items    []MonitorNNTPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorNNTPConfig struct {
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

const MonitorNNTPEndpoint = "/monitor/nntp"

type MonitorNNTPResource struct {
	c *f5.Client
}

func (r *MonitorNNTPResource) ListAll() (*MonitorNNTPConfigList, error) {
	var list MonitorNNTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorNNTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorNNTPResource) Get(id string) (*MonitorNNTPConfig, error) {
	var item MonitorNNTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorNNTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorNNTPResource) Create(item MonitorNNTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorNNTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorNNTPResource) Edit(id string, item MonitorNNTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorNNTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorNNTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorNNTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
