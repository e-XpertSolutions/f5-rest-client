// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorHTTPConfigList struct {
	Items    []MonitorHTTPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorHTTPConfig struct {
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
	Send                     string `json:"send"`
	TimeUntilUp              int    `json:"timeUntilUp"`
	Timeout                  int    `json:"timeout"`
	Transparent              string `json:"transparent"`
	UpInterval               int    `json:"upInterval"`
}

const MonitorHTTPEndpoint = "/monitor/http"

type MonitorHTTPResource struct {
	c f5.Client
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
