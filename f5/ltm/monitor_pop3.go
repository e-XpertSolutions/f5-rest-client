// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorPOP3ConfigList struct {
	Items    []MonitorPOP3Config `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorPOP3Config struct {
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

const MonitorPOP3Endpoint = "/monitor/pop3"

type MonitorPOP3Resource struct {
	c *f5.Client
}

func (r *MonitorPOP3Resource) ListAll() (*MonitorPOP3ConfigList, error) {
	var list MonitorPOP3ConfigList
	if err := r.c.ReadQuery(BasePath+MonitorPOP3Endpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorPOP3Resource) Get(id string) (*MonitorPOP3Config, error) {
	var item MonitorPOP3Config
	if err := r.c.ReadQuery(BasePath+MonitorPOP3Endpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorPOP3Resource) Create(item MonitorPOP3Config) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorPOP3Endpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorPOP3Resource) Edit(id string, item MonitorPOP3Config) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorPOP3Endpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorPOP3Resource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorPOP3Endpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
