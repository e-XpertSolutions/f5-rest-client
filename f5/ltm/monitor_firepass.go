// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorFirepassConfigList struct {
	Items    []MonitorFirepassConfig `json:"items,omitempty"`
	Kind     string                  `json:"kind,omitempty"`
	SelfLink string                  `json:"selflink,omitempty"`
}

type MonitorFirepassConfig struct {
	AppService       string `json:"appService,omitempty"`
	Cipherlist       string `json:"cipherlist,omitempty"`
	ConcurrencyLimit int    `json:"concurrencyLimit,omitempty"`
	DefaultsFrom     string `json:"defaultsFrom,omitempty"`
	Description      string `json:"description,omitempty"`
	Destination      string `json:"destination,omitempty"`
	FullPath         string `json:"fullPath,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	Interval         int    `json:"interval,omitempty"`
	Kind             string `json:"kind,omitempty"`
	MaxLoadAverage   int    `json:"maxLoadAverage,omitempty"`
	Name             string `json:"name,omitempty"`
	Partition        string `json:"partition,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	TimeUntilUp      int    `json:"timeUntilUp,omitempty"`
	Timeout          int    `json:"timeout,omitempty"`
	UpInterval       int    `json:"upInterval,omitempty"`
	Username         string `json:"username,omitempty"`
}

const MonitorFirepassEndpoint = "/monitor/firepass"

type MonitorFirepassResource struct {
	c *f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorFirepassEndpoint+"/"+id, &item); err != nil {
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
