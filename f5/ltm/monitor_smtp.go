// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSMTPConfigList struct {
	Items    []MonitorSMTPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorSMTPConfig struct {
	Debug        string `json:"debug,omitempty"`
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

const MonitorSMTPEndpoint = "/monitor/smtp"

type MonitorSMTPResource struct {
	c *f5.Client
}

func (r *MonitorSMTPResource) ListAll() (*MonitorSMTPConfigList, error) {
	var list MonitorSMTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSMTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSMTPResource) Get(id string) (*MonitorSMTPConfig, error) {
	var item MonitorSMTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSMTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSMTPResource) Create(item MonitorSMTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSMTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSMTPResource) Edit(id string, item MonitorSMTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSMTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSMTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSMTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
