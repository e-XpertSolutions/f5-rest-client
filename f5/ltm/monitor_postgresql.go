// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorPostgreSQLConfigList struct {
	Items    []MonitorPostgreSQLConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

type MonitorPostgreSQLConfig struct {
	Count        string `json:"count"`
	Debug        string `json:"debug"`
	Destination  string `json:"destination"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
}

const MonitorPostgreSQLEndpoint = "/monitor/postgresql"

type MonitorPostgreSQLResource struct {
	c *f5.Client
}

func (r *MonitorPostgreSQLResource) ListAll() (*MonitorPostgreSQLConfigList, error) {
	var list MonitorPostgreSQLConfigList
	if err := r.c.ReadQuery(BasePath+MonitorPostgreSQLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorPostgreSQLResource) Get(id string) (*MonitorPostgreSQLConfig, error) {
	var item MonitorPostgreSQLConfig
	if err := r.c.ReadQuery(BasePath+MonitorPostgreSQLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorPostgreSQLResource) Create(item MonitorPostgreSQLConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorPostgreSQLEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorPostgreSQLResource) Edit(id string, item MonitorPostgreSQLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorPostgreSQLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorPostgreSQLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorPostgreSQLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
