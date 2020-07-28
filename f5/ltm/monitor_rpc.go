// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorRPCConfigList struct {
	Items    []MonitorRPCConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorRPCConfig struct {
	AppService    string `json:"appService,omitempty"`
	Debug         string `json:"debug,omitempty"`
	DefaultsFrom  string `json:"defaultsFrom,omitempty"`
	Description   string `json:"description,omitempty"`
	Destination   string `json:"destination,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Interval      int    `json:"interval,omitempty"`
	Kind          string `json:"kind,omitempty"`
	ManualResume  string `json:"manualResume,omitempty"`
	Mode          string `json:"mode,omitempty"`
	Name          string `json:"name,omitempty"`
	Partition     string `json:"partition,omitempty"`
	ProgramNumber string `json:"programNumber,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	TimeUntilUp   int    `json:"timeUntilUp,omitempty"`
	Timeout       int    `json:"timeout,omitempty"`
	UpInterval    int    `json:"upInterval,omitempty"`
	VersionNumber string `json:"versionNumber,omitempty"`
}

const MonitorRPCEndpoint = "/monitor/rpc"

type MonitorRPCResource struct {
	c *f5.Client
}

func (r *MonitorRPCResource) ListAll() (*MonitorRPCConfigList, error) {
	var list MonitorRPCConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRPCEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorRPCResource) Get(id string) (*MonitorRPCConfig, error) {
	var item MonitorRPCConfig
	if err := r.c.ReadQuery(BasePath+MonitorRPCEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorRPCResource) Create(item MonitorRPCConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRPCEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRPCResource) Edit(id string, item MonitorRPCConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRPCEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorRPCResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRPCEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
