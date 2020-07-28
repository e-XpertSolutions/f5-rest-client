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
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	Cipherlist               string `json:"cipherlist,omitempty"`
	Compatibility            string `json:"compatibility,omitempty"`
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

const MonitorHTTPSEndpoint = "/monitor/https"

type MonitorHTTPSResource struct {
	c *f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorHTTPSEndpoint+"/"+id, &item); err != nil {
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
