// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SNMPTrapsConfigList holds a list of SNMPTraps configuration.
type SNMPTrapsConfigList struct {
	Items    []SNMPTrapsConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// SNMPTrapsConfig holds the configuration of a single SNMPTraps.
type SNMPTrapsConfig struct {
}

// SNMPTrapsEndpoint represents the REST resource for managing SNMPTraps.
const SNMPTrapsEndpoint = "/snmp_traps"

// SNMPTrapsResource provides an API to manage SNMPTraps configurations.
type SNMPTrapsResource struct {
	c *f5.Client
}

// ListAll  lists all the SNMPTraps configurations.
func (r *SNMPTrapsResource) ListAll() (*SNMPTrapsConfigList, error) {
	var list SNMPTrapsConfigList
	if err := r.c.ReadQuery(BasePath+SNMPTrapsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SNMPTraps configuration identified by id.
func (r *SNMPTrapsResource) Get(id string) (*SNMPTrapsConfig, error) {
	var item SNMPTrapsConfig
	if err := r.c.ReadQuery(BasePath+SNMPTrapsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SNMPTraps configuration.
func (r *SNMPTrapsResource) Create(item SNMPTrapsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SNMPTrapsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SNMPTraps configuration identified by id.
func (r *SNMPTrapsResource) Edit(id string, item SNMPTrapsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SNMPTrapsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SNMPTraps configuration identified by id.
func (r *SNMPTrapsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SNMPTrapsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
