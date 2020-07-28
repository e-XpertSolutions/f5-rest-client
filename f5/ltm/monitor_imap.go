// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorIMAPConfigList struct {
	Items    []MonitorIMAPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorIMAPConfig struct {
	Debug        string `json:"debug"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination"`
	Folder       string `json:"folder"`
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

const MonitorIMAPEndpoint = "/monitor/imap"

type MonitorIMAPResource struct {
	c *f5.Client
}

func (r *MonitorIMAPResource) ListAll() (*MonitorIMAPConfigList, error) {
	var list MonitorIMAPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorIMAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorIMAPResource) Get(id string) (*MonitorIMAPConfig, error) {
	var item MonitorIMAPConfig
	if err := r.c.ReadQuery(BasePath+MonitorIMAPEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorIMAPResource) Create(item MonitorIMAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorIMAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorIMAPResource) Edit(id string, item MonitorIMAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorIMAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorIMAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorIMAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
