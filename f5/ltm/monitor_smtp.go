// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSMTPConfigList struct {
	Items    []MonitorSMTPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorSMTPConfig struct {
	AppService   string `json:"appService,omitempty"`
	Debug        string `json:"debug,omitempty"`
	DefaultsFrom string `json:"defaultsFrom,omitempty"`
	Description  string `json:"description,omitempty"`
	Destination  string `json:"destination,omitempty,omitempty"`
	Domain       string `json:"domain,omitempty,omitempty"`
	FullPath     string `json:"fullPath,omitempty,omitempty"`
	Generation   int    `json:"generation,omitempty,omitempty"`
	Interval     int    `json:"interval,omitempty,omitempty"`
	Kind         string `json:"kind,omitempty,omitempty"`
	ManualResume string `json:"manualResume,omitempty,omitempty"`
	Name         string `json:"name,omitempty,omitempty"`
	Partition    string `json:"partition,omitempty,omitempty"`
	SelfLink     string `json:"selfLink,omitempty,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty,omitempty"`
	Timeout      int    `json:"timeout,omitempty,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty,omitempty"`
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
	if err := r.c.ReadQuery(BasePath+MonitorSMTPEndpoint+"/"+id, &item); err != nil {
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
