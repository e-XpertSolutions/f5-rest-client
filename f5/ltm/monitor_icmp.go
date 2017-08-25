// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorICMPConfigList struct {
	Items    []MonitorICMPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorICMPConfig struct {
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
