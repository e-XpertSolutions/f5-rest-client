// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorTCPConfigList struct {
	Items    []MonitorTCPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorTCPConfig struct {
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	AppService               string `json:"appService,omitempty"`
	DefaultsFrom             string `json:"defaultsFrom,omitempty"`
	Description              string `json:"description,omitempty"`
	Destination              string `json:"destination,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	Interval                 int    `json:"interval,omitempty"`
	IPDscp                   int    `json:"ipDscp,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	ManualResume             string `json:"manualResume,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	Recv                     string `json:"recv,omitempty"`
	RecvDisable              string `json:"recvDisable,omitempty"`
	Reverse                  string `json:"reverse,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	Send                     string `json:"send,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	Transparent              string `json:"transparent,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty"`
}

const MonitorTCPEndpoint = "/monitor/tcp"

type MonitorTCPResource struct {
	c *f5.Client
}

func (r *MonitorTCPResource) ListAll() (*MonitorTCPConfigList, error) {
	var list MonitorTCPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorTCPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorTCPResource) Get(id string) (*MonitorTCPConfig, error) {
	var item MonitorTCPConfig
	if err := r.c.ReadQuery(BasePath+MonitorTCPEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorTCPResource) Create(item MonitorTCPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorTCPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorTCPResource) Edit(id string, item MonitorTCPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorTCPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorTCPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorTCPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
