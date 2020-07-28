// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MonitorLDAPConfigList struct {
	Items    []MonitorLDAPConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

type MonitorLDAPConfig struct {
	AppService          string `json:"appService,omitempty"`
	Base                string `json:"base,omitempty"`
	ChaseReferrals      string `json:"chaseReferrals,omitempty"`
	Debug               string `json:"debug,omitempty"`
	DefaultsFrom        string `json:"defaultsFrom,omitempty"`
	Description         string `json:"description,omitempty"`
	Destination         string `json:"destination,omitempty"`
	Filter              string `json:"filter,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	Interval            int    `json:"interval,omitempty"`
	Kind                string `json:"kind,omitempty"`
	MandatoryAttributes string `json:"mandatoryAttributes,omitempty"`
	ManualResume        string `json:"manualResume,omitempty"`
	Name                string `json:"name,omitempty"`
	Partition           string `json:"partition,omitempty"`
	Security            string `json:"security,omitempty"`
	SelfLink            string `json:"selfLink,omitempty"`
	TimeUntilUp         int    `json:"timeUntilUp,omitempty"`
	Timeout             int    `json:"timeout,omitempty"`
	UpInterval          int    `json:"upInterval,omitempty"`
}

const MonitorLDAPEndpoint = "/monitor/ldap"

type MonitorLDAPResource struct {
	c *f5.Client
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
	if err := r.c.ReadQuery(BasePath+MonitorLDAPEndpoint+"/"+id, &item); err != nil {
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
