// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogProfileProtocolDNSConfigList holds a list of LogProfileProtocolDNS configuration.
type LogProfileProtocolDNSConfigList struct {
	Items    []LogProfileProtocolDNSConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// LogProfileProtocolDNSConfig holds the configuration of a single LogProfileProtocolDNS.
type LogProfileProtocolDNSConfig struct {
}

// LogProfileProtocolDNSEndpoint represents the REST resource for managing LogProfileProtocolDNS.
const LogProfileProtocolDNSEndpoint = "/log/profile_protocol-dns"

// LogProfileProtocolDNSResource provides an API to manage LogProfileProtocolDNS configurations.
type LogProfileProtocolDNSResource struct {
	c *f5.Client
}

// ListAll  lists all the LogProfileProtocolDNS configurations.
func (r *LogProfileProtocolDNSResource) ListAll() (*LogProfileProtocolDNSConfigList, error) {
	var list LogProfileProtocolDNSConfigList
	if err := r.c.ReadQuery(BasePath+LogProfileProtocolDNSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogProfileProtocolDNS configuration identified by id.
func (r *LogProfileProtocolDNSResource) Get(id string) (*LogProfileProtocolDNSConfig, error) {
	var item LogProfileProtocolDNSConfig
	if err := r.c.ReadQuery(BasePath+LogProfileProtocolDNSEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogProfileProtocolDNS configuration.
func (r *LogProfileProtocolDNSResource) Create(item LogProfileProtocolDNSConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogProfileProtocolDNSEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogProfileProtocolDNS configuration identified by id.
func (r *LogProfileProtocolDNSResource) Edit(id string, item LogProfileProtocolDNSConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogProfileProtocolDNSEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogProfileProtocolDNS configuration identified by id.
func (r *LogProfileProtocolDNSResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogProfileProtocolDNSEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
