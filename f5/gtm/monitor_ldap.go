// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorLDAPList holds a list of MonitorLDAP uration.
type MonitorLDAPList struct {
	Items    []MonitorLDAP `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// MonitorLDAP holds the uration of a single MonitorLDAP.
type MonitorLDAP struct {
	ChaseReferrals     string `json:"chaseReferrals,omitempty"`
	Debug              string `json:"debug,omitempty"`
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	ProbeTimeout       int    `json:"probeTimeout,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// MonitorLDAPEndpoint represents the REST resource for managing MonitorLDAP.
const MonitorLDAPEndpoint = "/monitor/ldap"

// MonitorLDAPResource provides an API to manage MonitorLDAP urations.
type MonitorLDAPResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorLDAP urations.
func (r *MonitorLDAPResource) ListAll() (*MonitorLDAPList, error) {
	var list MonitorLDAPList
	if err := r.c.ReadQuery(BasePath+MonitorLDAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorLDAP uration identified by id.
func (r *MonitorLDAPResource) Get(id string) (*MonitorLDAP, error) {
	var item MonitorLDAP
	if err := r.c.ReadQuery(BasePath+MonitorLDAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorLDAP uration.
func (r *MonitorLDAPResource) Create(item MonitorLDAP) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorLDAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorLDAP uration identified by id.
func (r *MonitorLDAPResource) Edit(id string, item MonitorLDAP) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorLDAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorLDAP uration identified by id.
func (r *MonitorLDAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorLDAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
