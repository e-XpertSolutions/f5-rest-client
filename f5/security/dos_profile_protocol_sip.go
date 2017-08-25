// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSProfileProtocolSIPConfigList holds a list of DOSProfileProtocolSIP configuration.
type DOSProfileProtocolSIPConfigList struct {
	Items    []DOSProfileProtocolSIPConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// DOSProfileProtocolSIPConfig holds the configuration of a single DOSProfileProtocolSIP.
type DOSProfileProtocolSIPConfig struct {
}

// DOSProfileProtocolSIPEndpoint represents the REST resource for managing DOSProfileProtocolSIP.
const DOSProfileProtocolSIPEndpoint = "/dos/profile_protocol-sip"

// DOSProfileProtocolSIPResource provides an API to manage DOSProfileProtocolSIP configurations.
type DOSProfileProtocolSIPResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSProfileProtocolSIP configurations.
func (r *DOSProfileProtocolSIPResource) ListAll() (*DOSProfileProtocolSIPConfigList, error) {
	var list DOSProfileProtocolSIPConfigList
	if err := r.c.ReadQuery(BasePath+DOSProfileProtocolSIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSProfileProtocolSIP configuration identified by id.
func (r *DOSProfileProtocolSIPResource) Get(id string) (*DOSProfileProtocolSIPConfig, error) {
	var item DOSProfileProtocolSIPConfig
	if err := r.c.ReadQuery(BasePath+DOSProfileProtocolSIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSProfileProtocolSIP configuration.
func (r *DOSProfileProtocolSIPResource) Create(item DOSProfileProtocolSIPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSProfileProtocolSIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSProfileProtocolSIP configuration identified by id.
func (r *DOSProfileProtocolSIPResource) Edit(id string, item DOSProfileProtocolSIPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSProfileProtocolSIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSProfileProtocolSIP configuration identified by id.
func (r *DOSProfileProtocolSIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSProfileProtocolSIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
