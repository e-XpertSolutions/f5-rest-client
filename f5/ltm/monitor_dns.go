// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorDNSConfigList struct {
	Items    []MonitorDNSConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

type MonitorDNSConfig struct {
	AcceptRcode              string `json:"acceptRcode"`
	Adaptive                 string `json:"adaptive"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue"`
	AdaptiveLimit            int    `json:"adaptiveLimit"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan"`
	AnswerContains           string `json:"answerContains"`
	Destination              string `json:"destination"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	Interval                 int    `json:"interval"`
	Kind                     string `json:"kind"`
	ManualResume             string `json:"manualResume"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	Qtype                    string `json:"qtype"`
	Reverse                  string `json:"reverse"`
	SelfLink                 string `json:"selfLink"`
	TimeUntilUp              int    `json:"timeUntilUp"`
	Timeout                  int    `json:"timeout"`
	Transparent              string `json:"transparent"`
	UpInterval               int    `json:"upInterval"`
}

const MonitorDNSEndpoint = "/monitor/dns"

type MonitorDNSResource struct {
	c *f5.Client
}

func (r *MonitorDNSResource) ListAll() (*MonitorDNSConfigList, error) {
	var list MonitorDNSConfigList
	if err := r.c.ReadQuery(BasePath+MonitorDNSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorDNSResource) Get(id string) (*MonitorDNSConfig, error) {
	var item MonitorDNSConfig
	if err := r.c.ReadQuery(BasePath+MonitorDNSEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorDNSResource) Create(item MonitorDNSConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorDNSEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorDNSResource) Edit(id string, item MonitorDNSConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorDNSEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorDNSResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorDNSEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
