// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorLDAPConfigList struct {
	Items    []MonitorLDAPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

type MonitorLDAPConfig struct {
	ChaseReferrals string `json:"chaseReferrals"`
	Debug          string `json:"debug"`
	Destination    string `json:"destination"`
	FullPath       string `json:"fullPath"`
	Generation     int    `json:"generation"`
	Interval       int    `json:"interval"`
	Kind           string `json:"kind"`
	ManualResume   string `json:"manualResume"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	SelfLink       string `json:"selfLink"`
	TimeUntilUp    int    `json:"timeUntilUp"`
	Timeout        int    `json:"timeout"`
	UpInterval     int    `json:"upInterval"`
}

const MonitorLDAPEndpoint = "/monitor/ldap"

type MonitorLDAPResource struct {
	c f5.Client
}

func (r *MonitorLDAPResource) ListAll() (*MonitorLDAPConfigList, error) {
	var list MonitorLDAPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorLDAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorLDAPResource) Get(id string) (*MonitorLDAPConfig, error) {
	var item MonitorLDAPConfig
	if err := r.c.ReadQuery(BasePath+MonitorLDAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorLDAPResource) Create(item MonitorLDAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorLDAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorLDAPResource) Edit(id string, item MonitorLDAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorLDAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorLDAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorLDAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
