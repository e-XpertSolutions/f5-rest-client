// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SNMPCommunitiesConfigList holds a list of SNMPCommunities configuration.
type SNMPCommunitiesConfigList struct {
	Items    []SNMPCommunitiesConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// SNMPCommunitiesConfig holds the configuration of a single SNMPCommunities.
type SNMPCommunitiesConfig struct {
}

// SNMPCommunitiesEndpoint represents the REST resource for managing SNMPCommunities.
const SNMPCommunitiesEndpoint = "/snmp_communities"

// SNMPCommunitiesResource provides an API to manage SNMPCommunities configurations.
type SNMPCommunitiesResource struct {
	c *f5.Client
}

// ListAll  lists all the SNMPCommunities configurations.
func (r *SNMPCommunitiesResource) ListAll() (*SNMPCommunitiesConfigList, error) {
	var list SNMPCommunitiesConfigList
	if err := r.c.ReadQuery(BasePath+SNMPCommunitiesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SNMPCommunities configuration identified by id.
func (r *SNMPCommunitiesResource) Get(id string) (*SNMPCommunitiesConfig, error) {
	var item SNMPCommunitiesConfig
	if err := r.c.ReadQuery(BasePath+SNMPCommunitiesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SNMPCommunities configuration.
func (r *SNMPCommunitiesResource) Create(item SNMPCommunitiesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SNMPCommunitiesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SNMPCommunities configuration identified by id.
func (r *SNMPCommunitiesResource) Edit(id string, item SNMPCommunitiesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SNMPCommunitiesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SNMPCommunities configuration identified by id.
func (r *SNMPCommunitiesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SNMPCommunitiesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
