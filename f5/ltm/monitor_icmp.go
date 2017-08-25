// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorICMPConfigList struct {
	Items    []MonitorICMPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorICMPConfig struct {
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

const MonitorICMPEndpoint = "/monitor/icmp"

type MonitorICMPResource struct {
	c *f5.Client
}

func (r *MonitorICMPResource) ListAll() (*MonitorICMPConfigList, error) {
	var list MonitorICMPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorICMPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorICMPResource) Get(id string) (*MonitorICMPConfig, error) {
	var item MonitorICMPConfig
	if err := r.c.ReadQuery(BasePath+MonitorICMPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorICMPResource) Create(item MonitorICMPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorICMPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorICMPResource) Edit(id string, item MonitorICMPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorICMPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorICMPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorICMPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
