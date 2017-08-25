// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SSHConfigList holds a list of SSH configuration.
type SSHConfigList struct {
	Items    []SSHConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// SSHConfig holds the configuration of a single SSH.
type SSHConfig struct {
}

// SSHEndpoint represents the REST resource for managing SSH.
const SSHEndpoint = "/ssh"

// SSHResource provides an API to manage SSH configurations.
type SSHResource struct {
	c *f5.Client
}

// ListAll  lists all the SSH configurations.
func (r *SSHResource) ListAll() (*SSHConfigList, error) {
	var list SSHConfigList
	if err := r.c.ReadQuery(BasePath+SSHEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SSH configuration identified by id.
func (r *SSHResource) Get(id string) (*SSHConfig, error) {
	var item SSHConfig
	if err := r.c.ReadQuery(BasePath+SSHEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SSH configuration.
func (r *SSHResource) Create(item SSHConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SSHEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SSH configuration identified by id.
func (r *SSHResource) Edit(id string, item SSHConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SSHEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SSH configuration identified by id.
func (r *SSHResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SSHEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
