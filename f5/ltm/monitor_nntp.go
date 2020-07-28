// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorNNTPConfigList struct {
	Items    []MonitorNNTPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorNNTPConfig struct {
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
	Newsgroup    string `json:"newsgroup,omitempty"`
	Partition    string `json:"partition,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	TimeUntilUp  int    `json:"timeUntilUp,omitempty"`
	Timeout      int    `json:"timeout,omitempty"`
	UpInterval   int    `json:"upInterval,omitempty"`
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
	if err := r.c.ReadQuery(BasePath+MonitorNNTPEndpoint+"/"+id, &item); err != nil {
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
