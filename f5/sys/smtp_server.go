// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SMTPServerConfigList holds a list of SMTPServer configuration.
type SMTPServerConfigList struct {
	Items    []SMTPServerConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// SMTPServerConfig holds the configuration of a single SMTPServer.
type SMTPServerConfig struct {
}

// SMTPServerEndpoint represents the REST resource for managing SMTPServer.
const SMTPServerEndpoint = "/smtp-server"

// SMTPServerResource provides an API to manage SMTPServer configurations.
type SMTPServerResource struct {
	c *f5.Client
}

// ListAll  lists all the SMTPServer configurations.
func (r *SMTPServerResource) ListAll() (*SMTPServerConfigList, error) {
	var list SMTPServerConfigList
	if err := r.c.ReadQuery(BasePath+SMTPServerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SMTPServer configuration identified by id.
func (r *SMTPServerResource) Get(id string) (*SMTPServerConfig, error) {
	var item SMTPServerConfig
	if err := r.c.ReadQuery(BasePath+SMTPServerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SMTPServer configuration.
func (r *SMTPServerResource) Create(item SMTPServerConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SMTPServerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SMTPServer configuration identified by id.
func (r *SMTPServerResource) Edit(id string, item SMTPServerConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SMTPServerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SMTPServer configuration identified by id.
func (r *SMTPServerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SMTPServerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
