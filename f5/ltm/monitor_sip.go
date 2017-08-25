// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSIPConfigList struct {
	Items    []MonitorSIPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorSIPConfig struct {
	Cipherlist    string `json:"cipherlist"`
	Compatibility string `json:"compatibility"`
	Debug         string `json:"debug"`
	Destination   string `json:"destination"`
	FullPath      string `json:"fullPath"`
	Generation    int    `json:"generation"`
	Interval      int    `json:"interval"`
	Kind          string `json:"kind"`
	ManualResume  string `json:"manualResume"`
	Mode          string `json:"mode"`
	Name          string `json:"name"`
	Partition     string `json:"partition"`
	SelfLink      string `json:"selfLink"`
	TimeUntilUp   int    `json:"timeUntilUp"`
	Timeout       int    `json:"timeout"`
	UpInterval    int    `json:"upInterval"`
}

const MonitorSIPEndpoint = "/monitor/sip"

type MonitorSIPResource struct {
	c *f5.Client
}

func (r *MonitorSIPResource) ListAll() (*MonitorSIPConfigList, error) {
	var list MonitorSIPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSIPResource) Get(id string) (*MonitorSIPConfig, error) {
	var item MonitorSIPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSIPResource) Create(item MonitorSIPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSIPResource) Edit(id string, item MonitorSIPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
