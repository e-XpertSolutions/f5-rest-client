// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SSHProfileConfigList holds a list of SSHProfile configuration.
type SSHProfileConfigList struct {
	Items    []SSHProfileConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// SSHProfileConfig holds the configuration of a single SSHProfile.
type SSHProfileConfig struct {
}

// SSHProfileEndpoint represents the REST resource for managing SSHProfile.
const SSHProfileEndpoint = "/ssh/profile"

// SSHProfileResource provides an API to manage SSHProfile configurations.
type SSHProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the SSHProfile configurations.
func (r *SSHProfileResource) ListAll() (*SSHProfileConfigList, error) {
	var list SSHProfileConfigList
	if err := r.c.ReadQuery(BasePath+SSHProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SSHProfile configuration identified by id.
func (r *SSHProfileResource) Get(id string) (*SSHProfileConfig, error) {
	var item SSHProfileConfig
	if err := r.c.ReadQuery(BasePath+SSHProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SSHProfile configuration.
func (r *SSHProfileResource) Create(item SSHProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SSHProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SSHProfile configuration identified by id.
func (r *SSHProfileResource) Edit(id string, item SSHProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SSHProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SSHProfile configuration identified by id.
func (r *SSHProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SSHProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
