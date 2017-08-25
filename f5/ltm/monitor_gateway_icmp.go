// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorGatewayICMPConfigList struct {
	Items    []MonitorGatewayICMPConfig `json:"items,omitempty"`
	Kind     string                     `json:"kind,omitempty"`
	SelfLink string                     `json:"selflink,omitempty"`
}

type MonitorGatewayICMPConfig struct {
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	Destination              string `json:"destination,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	Interval                 int    `json:"interval,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	ManualResume             string `json:"manualResume,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty"`
	Timeout                  int    `json:"timeout,omitempty"`
	Transparent              string `json:"transparent,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty"`
}

const MonitorGatewayICMPEndpoint = "/monitor/gateway-icmp"

type MonitorGatewayICMPResource struct {
	c *f5.Client
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
