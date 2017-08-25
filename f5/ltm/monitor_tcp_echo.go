// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorTCPEchoConfigList struct {
	Items    []MonitorTCPEchoConfig `json:"items,omitempty"`
	Kind     string                 `json:"kind,omitempty"`
	SelfLink string                 `json:"selflink,omitempty"`
}

type MonitorTCPEchoConfig struct {
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

const MonitorTCPEchoEndpoint = "/monitor/tcp-echo"

type MonitorTCPEchoResource struct {
	c *f5.Client
}

func (r *MonitorTCPEchoResource) ListAll() (*MonitorTCPEchoConfigList, error) {
	var list MonitorTCPEchoConfigList
	if err := r.c.ReadQuery(BasePath+MonitorTCPEchoEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorTCPEchoResource) Get(id string) (*MonitorTCPEchoConfig, error) {
	var item MonitorTCPEchoConfig
	if err := r.c.ReadQuery(BasePath+MonitorTCPEchoEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorTCPEchoResource) Create(item MonitorTCPEchoConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorTCPEchoEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorTCPEchoResource) Edit(id string, item MonitorTCPEchoConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorTCPEchoEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorTCPEchoResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorTCPEchoEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
