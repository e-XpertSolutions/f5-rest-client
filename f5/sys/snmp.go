// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SNMPConfigList holds a list of SNMP configuration.
type SNMPConfigList struct {
	Items    []SNMPConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// SNMPConfig holds the configuration of a single SNMP.
type SNMPConfig struct {
}

// SNMPEndpoint represents the REST resource for managing SNMP.
const SNMPEndpoint = "/snmp"

// SNMPResource provides an API to manage SNMP configurations.
type SNMPResource struct {
	c *f5.Client
}

// ListAll  lists all the SNMP configurations.
func (r *SNMPResource) ListAll() (*SNMPConfigList, error) {
	var list SNMPConfigList
	if err := r.c.ReadQuery(BasePath+SNMPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SNMP configuration identified by id.
func (r *SNMPResource) Get(id string) (*SNMPConfig, error) {
	var item SNMPConfig
	if err := r.c.ReadQuery(BasePath+SNMPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SNMP configuration.
func (r *SNMPResource) Create(item SNMPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SNMPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SNMP configuration identified by id.
func (r *SNMPResource) Edit(id string, item SNMPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SNMPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SNMP configuration identified by id.
func (r *SNMPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SNMPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
