// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorInbandConfigList struct {
	Items    []MonitorInbandConfig `json:"items,omitempty"`
	Kind     string                `json:"kind,omitempty"`
	SelfLink string                `json:"selflink,omitempty"`
}

type MonitorInbandConfig struct {
	AppService      string `json:"appService,omitempty"`
	DefaultsFrom    string `json:"defaultsFrom,omitempty"`
	Description     string `json:"description,omitempty"`
	FailureInterval int    `json:"failureInterval,omitempty"`
	Failures        int    `json:"failures,omitempty"`
	FullPath        string `json:"fullPath,omitempty"`
	Generation      int    `json:"generation,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Partition       string `json:"partition,omitempty"`
	ResponseTime    int    `json:"responseTime,omitempty"`
	RetryTime       int    `json:"retryTime,omitempty"`
	SelfLink        string `json:"selfLink,omitempty"`
}

const MonitorInbandEndpoint = "/monitor/inband"

type MonitorInbandResource struct {
	c *f5.Client
}

func (r *MonitorInbandResource) ListAll() (*MonitorInbandConfigList, error) {
	var list MonitorInbandConfigList
	if err := r.c.ReadQuery(BasePath+MonitorInbandEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorInbandResource) Get(id string) (*MonitorInbandConfig, error) {
	var item MonitorInbandConfig
	if err := r.c.ReadQuery(BasePath+MonitorInbandEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorInbandResource) Create(item MonitorInbandConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorInbandEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorInbandResource) Edit(id string, item MonitorInbandConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorInbandEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorInbandResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorInbandEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
