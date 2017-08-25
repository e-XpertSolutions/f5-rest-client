// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorDNSConfigList struct {
	Items    []MonitorDNSConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorDNSConfig struct {
	AcceptRcode              string `json:"acceptRcode,omitempty"`
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	AnswerContains           string `json:"answerContains,omitempty"`
	Destination              string `json:"destination,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	Interval                 int    `json:"interval,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	ManualResume             string `json:"manualResume,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	Qtype                    string `json:"qtype,omitempty"`
	Reverse                  string `json:"reverse,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	Transparent              string `json:"transparent,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty"`
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
