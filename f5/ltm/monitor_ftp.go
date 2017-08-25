// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorFTPConfigList struct {
	Items    []MonitorFTPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorFTPConfig struct {
	Debug        string `json:"debug"`
	Destination  string `json:"destination"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Mode         string `json:"mode"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorFTPEndpoint = "/monitor/ftp"

type MonitorFTPResource struct {
	c *f5.Client
}

func (r *MonitorFTPResource) ListAll() (*MonitorFTPConfigList, error) {
	var list MonitorFTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorFTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorFTPResource) Get(id string) (*MonitorFTPConfig, error) {
	var item MonitorFTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorFTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorFTPResource) Create(item MonitorFTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorFTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorFTPResource) Edit(id string, item MonitorFTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorFTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorFTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorFTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
