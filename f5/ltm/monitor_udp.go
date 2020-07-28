// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorUDPConfigList struct {
	Items    []MonitorUDPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorUDPConfig struct {
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	AppService               string `json:"appService,omitempty"`
	Debug                    string `json:"debug,omitempty"`
	DefaultsFrom             string `json:"defaultsFrom,omitempty"`
	Description              string `json:"description,omitempty"`
	Destination              string `json:"destination,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	Interval                 int    `json:"interval,omitempty"`
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

const MonitorUDPEndpoint = "/monitor/udp"

type MonitorUDPResource struct {
	c *f5.Client
}

func (r *MonitorUDPResource) ListAll() (*MonitorUDPConfigList, error) {
	var list MonitorUDPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorUDPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorUDPResource) Get(id string) (*MonitorUDPConfig, error) {
	var item MonitorUDPConfig
	if err := r.c.ReadQuery(BasePath+MonitorUDPEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorUDPResource) Create(item MonitorUDPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorUDPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorUDPResource) Edit(id string, item MonitorUDPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorUDPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorUDPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorUDPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
