// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SSHDConfigList holds a list of SSHD configuration.
type SSHDConfigList struct {
	Items    []SSHDConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selflink"`
}

// SSHDConfig holds the configuration of a single SSHD.
type SSHDConfig struct {
}

// SSHDEndpoint represents the REST resource for managing SSHD.
const SSHDEndpoint = "/sshd"

// SSHDResource provides an API to manage SSHD configurations.
type SSHDResource struct {
	c *f5.Client
}

// ListAll  lists all the SSHD configurations.
func (r *SSHDResource) ListAll() (*SSHDConfigList, error) {
	var list SSHDConfigList
	if err := r.c.ReadQuery(BasePath+SSHDEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SSHD configuration identified by id.
func (r *SSHDResource) Get(id string) (*SSHDConfig, error) {
	var item SSHDConfig
	if err := r.c.ReadQuery(BasePath+SSHDEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SSHD configuration.
func (r *SSHDResource) Create(item SSHDConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SSHDEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SSHD configuration identified by id.
func (r *SSHDResource) Edit(id string, item SSHDConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SSHDEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SSHD configuration identified by id.
func (r *SSHDResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SSHDEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
