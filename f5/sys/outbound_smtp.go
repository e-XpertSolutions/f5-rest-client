// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// OutboundSMTPConfigList holds a list of OutboundSMTP configuration.
type OutboundSMTPConfigList struct {
	Items    []OutboundSMTPConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// OutboundSMTPConfig holds the configuration of a single OutboundSMTP.
type OutboundSMTPConfig struct {
}

// OutboundSMTPEndpoint represents the REST resource for managing OutboundSMTP.
const OutboundSMTPEndpoint = "/outbound-smtp"

// OutboundSMTPResource provides an API to manage OutboundSMTP configurations.
type OutboundSMTPResource struct {
	c *f5.Client
}

// ListAll  lists all the OutboundSMTP configurations.
func (r *OutboundSMTPResource) ListAll() (*OutboundSMTPConfigList, error) {
	var list OutboundSMTPConfigList
	if err := r.c.ReadQuery(BasePath+OutboundSMTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single OutboundSMTP configuration identified by id.
func (r *OutboundSMTPResource) Get(id string) (*OutboundSMTPConfig, error) {
	var item OutboundSMTPConfig
	if err := r.c.ReadQuery(BasePath+OutboundSMTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new OutboundSMTP configuration.
func (r *OutboundSMTPResource) Create(item OutboundSMTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+OutboundSMTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a OutboundSMTP configuration identified by id.
func (r *OutboundSMTPResource) Edit(id string, item OutboundSMTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+OutboundSMTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single OutboundSMTP configuration identified by id.
func (r *OutboundSMTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+OutboundSMTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
