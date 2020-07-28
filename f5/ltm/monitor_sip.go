// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorSIPConfigList struct {
	Items    []MonitorSIPConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorSIPConfig struct {
	AppService    string `json:"appService,omitempty"`
	Cert          string `json:"cert,omitempty"`
	Cipherlist    string `json:"cipherlist,omitempty"`
	Compatibility string `json:"compatibility,omitempty"`
	Debug         string `json:"debug,omitempty"`
	DefaultsFrom  string `json:"defaultsFrom,omitempty"`
	Description   string `json:"description,omitempty"`
	Destination   string `json:"destination,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Filter        string `json:"filter,omitempty"`
	FilterNeg     string `json:"filterNeg,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Headers       string `json:"headers,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Key           string `json:"key,omitempty"`
	Kind          string `json:"kind,omitempty"`
	ManualResume  string `json:"manualResume,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	Request       string `json:"request,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	TimeUntilUp   int    `json:"timeUntilUp,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	UpInterval    int    `json:"upInterval,omitempty"`
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
	if err := r.c.ReadQuery(BasePath+MonitorSIPEndpoint+"/"+id, &item); err != nil {
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
