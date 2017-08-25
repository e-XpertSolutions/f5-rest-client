// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorTCPEchoConfigList struct {
	Items    []MonitorTCPEchoConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

type MonitorTCPEchoConfig struct {
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
