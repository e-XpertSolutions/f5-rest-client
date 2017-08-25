// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorLDAPConfigList holds a list of MonitorLDAP configuration.
type MonitorLDAPConfigList struct {
	Items    []MonitorLDAPConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// MonitorLDAPConfig holds the configuration of a single MonitorLDAP.
type MonitorLDAPConfig struct {
	ChaseReferrals     string `json:"chaseReferrals"`
	Debug              string `json:"debug"`
	Destination        string `json:"destination"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreDownResponse string `json:"ignoreDownResponse"`
	Interval           int    `json:"interval"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	ProbeTimeout       int    `json:"probeTimeout"`
	SelfLink           string `json:"selfLink"`
	Timeout            int    `json:"timeout"`
}

// MonitorLDAPEndpoint represents the REST resource for managing MonitorLDAP.
const MonitorLDAPEndpoint = "/monitor/ldap"

// MonitorLDAPResource provides an API to manage MonitorLDAP configurations.
type MonitorLDAPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorLDAP configurations.
func (r *MonitorLDAPResource) ListAll() (*MonitorLDAPConfigList, error) {
	var list MonitorLDAPConfigList
	if err := r.c.ReadQuery(BasePath+MonitorLDAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorLDAP configuration identified by id.
func (r *MonitorLDAPResource) Get(id string) (*MonitorLDAPConfig, error) {
	var item MonitorLDAPConfig
	if err := r.c.ReadQuery(BasePath+MonitorLDAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorLDAP configuration.
func (r *MonitorLDAPResource) Create(item MonitorLDAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorLDAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorLDAP configuration identified by id.
func (r *MonitorLDAPResource) Edit(id string, item MonitorLDAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorLDAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorLDAP configuration identified by id.
func (r *MonitorLDAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorLDAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
