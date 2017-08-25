// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogProfileProtocolSIPConfigList holds a list of LogProfileProtocolSIP configuration.
type LogProfileProtocolSIPConfigList struct {
	Items    []LogProfileProtocolSIPConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// LogProfileProtocolSIPConfig holds the configuration of a single LogProfileProtocolSIP.
type LogProfileProtocolSIPConfig struct {
}

// LogProfileProtocolSIPEndpoint represents the REST resource for managing LogProfileProtocolSIP.
const LogProfileProtocolSIPEndpoint = "/log/profile_protocol-sip"

// LogProfileProtocolSIPResource provides an API to manage LogProfileProtocolSIP configurations.
type LogProfileProtocolSIPResource struct {
	c *f5.Client
}

// ListAll  lists all the LogProfileProtocolSIP configurations.
func (r *LogProfileProtocolSIPResource) ListAll() (*LogProfileProtocolSIPConfigList, error) {
	var list LogProfileProtocolSIPConfigList
	if err := r.c.ReadQuery(BasePath+LogProfileProtocolSIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogProfileProtocolSIP configuration identified by id.
func (r *LogProfileProtocolSIPResource) Get(id string) (*LogProfileProtocolSIPConfig, error) {
	var item LogProfileProtocolSIPConfig
	if err := r.c.ReadQuery(BasePath+LogProfileProtocolSIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogProfileProtocolSIP configuration.
func (r *LogProfileProtocolSIPResource) Create(item LogProfileProtocolSIPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogProfileProtocolSIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogProfileProtocolSIP configuration identified by id.
func (r *LogProfileProtocolSIPResource) Edit(id string, item LogProfileProtocolSIPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogProfileProtocolSIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogProfileProtocolSIP configuration identified by id.
func (r *LogProfileProtocolSIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogProfileProtocolSIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
