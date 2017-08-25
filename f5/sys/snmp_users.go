// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SNMPUsersConfigList holds a list of SNMPUsers configuration.
type SNMPUsersConfigList struct {
	Items    []SNMPUsersConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// SNMPUsersConfig holds the configuration of a single SNMPUsers.
type SNMPUsersConfig struct {
}

// SNMPUsersEndpoint represents the REST resource for managing SNMPUsers.
const SNMPUsersEndpoint = "/snmp_users"

// SNMPUsersResource provides an API to manage SNMPUsers configurations.
type SNMPUsersResource struct {
	c *f5.Client
}

// ListAll  lists all the SNMPUsers configurations.
func (r *SNMPUsersResource) ListAll() (*SNMPUsersConfigList, error) {
	var list SNMPUsersConfigList
	if err := r.c.ReadQuery(BasePath+SNMPUsersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SNMPUsers configuration identified by id.
func (r *SNMPUsersResource) Get(id string) (*SNMPUsersConfig, error) {
	var item SNMPUsersConfig
	if err := r.c.ReadQuery(BasePath+SNMPUsersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SNMPUsers configuration.
func (r *SNMPUsersResource) Create(item SNMPUsersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SNMPUsersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SNMPUsers configuration identified by id.
func (r *SNMPUsersResource) Edit(id string, item SNMPUsersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SNMPUsersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SNMPUsers configuration identified by id.
func (r *SNMPUsersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SNMPUsersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
