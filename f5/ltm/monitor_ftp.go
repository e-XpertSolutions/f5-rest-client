// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorFTPConfigList struct {
	Items    []MonitorFTPConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

type MonitorFTPConfig struct {
	Adaptive                 string `json:"adaptive,omitempty"`
	AdaptiveDivergenceType   string `json:"adaptiveDivergenceType,omitempty"`
	AdaptiveDivergenceValue  int    `json:"adaptiveDivergenceValue,omitempty"`
	AdaptiveLimit            int    `json:"adaptiveLimit,omitempty"`
	AdaptiveSamplingTimespan int    `json:"adaptiveSamplingTimespan,omitempty"`
	AppService               string `json:"appService,omitempty"`
	Debug                    string `json:"debug,omitempty,omitempty"`
	DefaultsFrom             string `json:"defaultsFrom,omitempty"`
	Description              string `json:"description,omitempty"`
	Destination              string `json:"destination,omitempty,omitempty"`
	Filename                 string `json:"filename,omitempty,omitempty"`
	FullPath                 string `json:"fullPath,omitempty,omitempty"`
	Generation               int    `json:"generation,omitempty,omitempty"`
	Interval                 int    `json:"interval,omitempty,omitempty"`
	Kind                     string `json:"kind,omitempty,omitempty"`
	ManualResume             string `json:"manualResume,omitempty,omitempty"`
	Mode                     string `json:"mode,omitempty,omitempty"`
	Name                     string `json:"name,omitempty,omitempty"`
	Partition                string `json:"partition,omitempty,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty,omitempty"`
	TimeUntilUp              int    `json:"timeUntilUp,omitempty,omitempty"`
	Timeout                  int    `json:"timeout,omitempty,omitempty"`
	UpInterval               int    `json:"upInterval,omitempty,omitempty"`
}

const MonitorFTPEndpoint = "/monitor/ftp"

type MonitorFTPResource struct {
	c *f5.Client
}

func (r *MonitorFTPResource) ListAll() (*MonitorFTPConfigList, error) {
	var list MonitorFTPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorFTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorFTPResource) Get(id string) (*MonitorFTPConfig, error) {
	var item MonitorFTPConfig
	if err := r.c.ReadQuery(BasePath+MonitorFTPEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorFTPResource) Create(item MonitorFTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorFTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorFTPResource) Edit(id string, item MonitorFTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorFTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorFTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorFTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
