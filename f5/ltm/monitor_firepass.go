// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorFirepassConfigList struct {
	Items    []MonitorFirepassConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

type MonitorFirepassConfig struct {
	Cipherlist       string `json:"cipherlist"`
	ConcurrencyLimit int    `json:"concurrencyLimit"`
	Destination      string `json:"destination"`
	FullPath         string `json:"fullPath"`
	Generation       int    `json:"generation"`
	Interval         int    `json:"interval"`
	Kind             string `json:"kind"`
	MaxLoadAverage   int    `json:"maxLoadAverage"`
	Name             string `json:"name"`
	Partition        string `json:"partition"`
	SelfLink         string `json:"selfLink"`
	TimeUntilUp      int    `json:"timeUntilUp"`
	Timeout          int    `json:"timeout"`
	UpInterval       int    `json:"upInterval"`
	Username         string `json:"username"`
}

const MonitorFirepassEndpoint = "/monitor/firepass"

type MonitorFirepassResource struct {
	c f5.Client
}

func (r *MonitorFirepassResource) ListAll() (*MonitorFirepassConfigList, error) {
	var list MonitorFirepassConfigList
	if err := r.c.ReadQuery(BasePath+MonitorFirepassEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorFirepassResource) Get(id string) (*MonitorFirepassConfig, error) {
	var item MonitorFirepassConfig
	if err := r.c.ReadQuery(BasePath+MonitorFirepassEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorFirepassResource) Create(item MonitorFirepassConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorFirepassEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorFirepassResource) Edit(id string, item MonitorFirepassConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorFirepassEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorFirepassResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorFirepassEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
