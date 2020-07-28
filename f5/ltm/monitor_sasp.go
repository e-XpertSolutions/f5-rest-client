// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSASPConfigList struct {
	Items    []MonitorSASPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorSASPConfig struct {
	AppService       string `json:"appService,omitempty"`
	DefaultsFrom     string `json:"defaultsFrom,omitempty"`
	Description      string `json:"description,omitempty"`
	Destination      string `json:"destination,omitempty"`
	FullPath         string `json:"fullPath,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	Interval         string `json:"interval,omitempty"`
	Kind             string `json:"kind,omitempty"`
	Mode             string `json:"mode,omitempty"`
	MonInterval      int    `json:"monInterval,omitempty"`
	Name             string `json:"name,omitempty"`
	Partition        string `json:"partition,omitempty"`
	PrimaryAddress   string `json:"primaryAddress,omitempty"`
	Protocol         string `json:"protocol,omitempty"`
	SecondaryAddress string `json:"secondaryAddress,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	Service          string `json:"service,omitempty"`
	Timeout          int    `json:"timeout,omitempty"`
	TimeUntilUp      int    `json:"timeUntilUp,omitempty"`
}

const MonitorSASPEndpoint = "/monitor/sasp"

type MonitorSASPResource struct {
	c *f5.Client
}

func (r *MonitorSASPResource) ListAll() (*MonitorSASPConfigList, error) {
	var list MonitorSASPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorSASPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorSASPResource) Get(id string) (*MonitorSASPConfig, error) {
	var item MonitorSASPConfig
	if err := r.c.ReadQuery(BasePath+MonitorSASPEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorSASPResource) Create(item MonitorSASPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorSASPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSASPResource) Edit(id string, item MonitorSASPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorSASPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorSASPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorSASPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
