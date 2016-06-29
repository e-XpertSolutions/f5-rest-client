// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorHTTPSConfigList struct {
	Items    []MonitorHTTPSConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

type MonitorHTTPSConfig struct {
	Adaptive                 string `json:"adaptive"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue"`
	AdaptiveLimit            int    `json:"adaptiveLimit"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan"`
	Cipherlist               string `json:"cipherlist"`
	Compatibility            string `json:"compatibility"`
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

const MonitorHTTPSEndpoint = "/monitor/https"

type MonitorHTTPSResource struct {
	c f5.Client
}

func (r *MonitorHTTPSResource) ListAll() (*MonitorHTTPSConfigList, error) {
	var list MonitorHTTPSConfigList
	if err := r.c.ReadQuery(BasePath+MonitorHTTPSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorHTTPSResource) Get(id string) (*MonitorHTTPSConfig, error) {
	var item MonitorHTTPSConfig
	if err := r.c.ReadQuery(BasePath+MonitorHTTPSEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorHTTPSResource) Create(item MonitorHTTPSConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorHTTPSEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorHTTPSResource) Edit(id string, item MonitorHTTPSConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorHTTPSEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorHTTPSResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorHTTPSEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
