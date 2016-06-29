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
	Adaptive                 string `json:"adaptive"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue"`
	AdaptiveLimit            int    `json:"adaptiveLimit"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan"`
	Destination              string `json:"destination"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	Interval                 int    `json:"interval"`
	IPDscp                   int    `json:"ipDscp"`
	Kind                     string `json:"kind"`
	ManualResume             string `json:"manualResume"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	Reverse                  string `json:"reverse"`
	SelfLink                 string `json:"selfLink"`
	TimeUntilUp              int    `json:"timeUntilUp"`
	Timeout                  int    `json:"timeout"`
	Transparent              string `json:"transparent"`
	UpInterval               int    `json:"upInterval"`
}

const MonitorTCPEndpoint = "/monitor/tcp"

type MonitorTCPResource struct {
	c f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorTCPEndpoint, &item); err != nil {
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
