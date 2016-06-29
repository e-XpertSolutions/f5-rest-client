// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorGatewayICMPConfigList struct {
	Items    []MonitorGatewayICMPConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

type MonitorGatewayICMPConfig struct {
	Adaptive                 string `json:"adaptive"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue"`
	AdaptiveLimit            int    `json:"adaptiveLimit"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan"`
	Destination              string `json:"destination"`
	FullPath                 string `json:"fullPath"`
	Generation               int    `json:"generation"`
	Interval                 int    `json:"interval"`
	Kind                     string `json:"kind"`
	ManualResume             string `json:"manualResume"`
	Name                     string `json:"name"`
	Partition                string `json:"partition"`
	SelfLink                 string `json:"selfLink"`
	TimeUntilUp              int    `json:"timeUntilUp"`
	Timeout                  int    `json:"timeout"`
	Transparent              string `json:"transparent"`
	UpInterval               int    `json:"upInterval"`
}

const MonitorGatewayICMPEndpoint = "/monitor/gateway-icmp"

type MonitorGatewayICMPResource struct {
	c f5.Client
}

func (r *MonitorGatewayICMPResource) ListAll() (*MonitorGatewayICMPConfigList, error) {
	var list MonitorGatewayICMPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorGatewayICMPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorGatewayICMPResource) Get(id string) (*MonitorGatewayICMPConfig, error) {
	var item MonitorGatewayICMPConfig
	if err := r.c.ReadQuery(BasePath+MonitorGatewayICMPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorGatewayICMPResource) Create(item MonitorGatewayICMPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorGatewayICMPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorGatewayICMPResource) Edit(id string, item MonitorGatewayICMPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorGatewayICMPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorGatewayICMPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorGatewayICMPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
