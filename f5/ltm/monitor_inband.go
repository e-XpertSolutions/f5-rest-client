// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorInbandConfigList struct {
	Items    []MonitorInbandConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

type MonitorInbandConfig struct {
	FailureInterval int    `json:"failureInterval"`
	Failures        int    `json:"failures"`
	FullPath        string `json:"fullPath"`
	Generation      int    `json:"generation"`
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	Partition       string `json:"partition"`
	ResponseTime    int    `json:"responseTime"`
	RetryTime       int    `json:"retryTime"`
	SelfLink        string `json:"selfLink"`
}

const MonitorInbandEndpoint = "/monitor/inband"

type MonitorInbandResource struct {
	c *f5.Client
}

func (r *MonitorInbandResource) ListAll() (*MonitorInbandConfigList, error) {
	var list MonitorInbandConfigList
	if err := r.c.ReadQuery(BasePath+MonitorInbandEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorInbandResource) Get(id string) (*MonitorInbandConfig, error) {
	var item MonitorInbandConfig
	if err := r.c.ReadQuery(BasePath+MonitorInbandEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorInbandResource) Create(item MonitorInbandConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorInbandEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorInbandResource) Edit(id string, item MonitorInbandConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorInbandEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorInbandResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorInbandEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
