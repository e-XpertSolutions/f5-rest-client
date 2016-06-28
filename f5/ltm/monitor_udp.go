// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorUDPConfigList struct {
	Items    []MonitorUDPConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorUDPConfig struct {
	Adaptive                 string `json:"adaptive"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue"`
	AdaptiveLimit            int    `json:"adaptiveLimit"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan"`
	Debug                    string `json:"debug"`
	Destination              string `json:"destination"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	Interval                 int    `json:"interval"`
	Kind                     string `json:"kind"`
	ManualResume             string `json:"manualResume"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	Reverse                  string `json:"reverse"`
	SelfLink                 string `json:"selfLink"`
	Send                     string `json:"send"`
	TimeUntilUp              int    `json:"timeUntilUp"`
	Timeout                  int    `json:"timeout"`
	Transparent              string `json:"transparent"`
	UpInterval               int    `json:"upInterval"`
}

const MonitorUDPEndpoint = "/monitor/udp"

type MonitorUDPResource struct {
	c f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorUDPEndpoint, &item); err != nil {
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
