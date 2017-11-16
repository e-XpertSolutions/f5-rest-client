// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorHTTPConfigList struct {
	Items    []MonitorHTTPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorHTTPConfig struct {
	Name                     string `json:"name,omitempty"`
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	Destination              string `json:"destination,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	Interval                 int    `json:"interval,omitempty"`
	IPDscp                   int    `json:"ipDscp,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	ManualResume             string `json:"manualResume,omitempty"`
	Partition                string `json:"partition,omitempty"`
	Reverse                  string `json:"reverse,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	Send                     string `json:"send,omitempty"`
	Recv                     string `json:"recv,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	Transparent              string `json:"transparent,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty"`
}

const MonitorHTTPEndpoint = "/monitor/http"

type MonitorHTTPResource struct {
	c *f5.Client
}

func (r *MonitorHTTPResource) ListAll() (*MonitorHTTPConfigList, error) {
	var list MonitorHTTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorHTTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorHTTPResource) Get(id string) (*MonitorHTTPConfig, error) {
	var item MonitorHTTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorHTTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorHTTPResource) Create(item MonitorHTTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorHTTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorHTTPResource) Edit(id string, item MonitorHTTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorHTTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorHTTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorHTTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
