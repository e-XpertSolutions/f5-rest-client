// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSProfileProtocolDNSConfigList holds a list of DOSProfileProtocolDNS configuration.
type DOSProfileProtocolDNSConfigList struct {
	Items    []DOSProfileProtocolDNSConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// DOSProfileProtocolDNSConfig holds the configuration of a single DOSProfileProtocolDNS.
type DOSProfileProtocolDNSConfig struct {
}

// DOSProfileProtocolDNSEndpoint represents the REST resource for managing DOSProfileProtocolDNS.
const DOSProfileProtocolDNSEndpoint = "/dos/profile_protocol-dns"

// DOSProfileProtocolDNSResource provides an API to manage DOSProfileProtocolDNS configurations.
type DOSProfileProtocolDNSResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSProfileProtocolDNS configurations.
func (r *DOSProfileProtocolDNSResource) ListAll() (*DOSProfileProtocolDNSConfigList, error) {
	var list DOSProfileProtocolDNSConfigList
	if err := r.c.ReadQuery(BasePath+DOSProfileProtocolDNSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSProfileProtocolDNS configuration identified by id.
func (r *DOSProfileProtocolDNSResource) Get(id string) (*DOSProfileProtocolDNSConfig, error) {
	var item DOSProfileProtocolDNSConfig
	if err := r.c.ReadQuery(BasePath+DOSProfileProtocolDNSEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSProfileProtocolDNS configuration.
func (r *DOSProfileProtocolDNSResource) Create(item DOSProfileProtocolDNSConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSProfileProtocolDNSEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSProfileProtocolDNS configuration identified by id.
func (r *DOSProfileProtocolDNSResource) Edit(id string, item DOSProfileProtocolDNSConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSProfileProtocolDNSEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSProfileProtocolDNS configuration identified by id.
func (r *DOSProfileProtocolDNSResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSProfileProtocolDNSEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
