// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SSHProfileAuthInfoConfigList holds a list of SSHProfileAuthInfo configuration.
type SSHProfileAuthInfoConfigList struct {
	Items    []SSHProfileAuthInfoConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// SSHProfileAuthInfoConfig holds the configuration of a single SSHProfileAuthInfo.
type SSHProfileAuthInfoConfig struct {
}

// SSHProfileAuthInfoEndpoint represents the REST resource for managing SSHProfileAuthInfo.
const SSHProfileAuthInfoEndpoint = "/ssh/profile_auth-info"

// SSHProfileAuthInfoResource provides an API to manage SSHProfileAuthInfo configurations.
type SSHProfileAuthInfoResource struct {
	c *f5.Client
}

// ListAll  lists all the SSHProfileAuthInfo configurations.
func (r *SSHProfileAuthInfoResource) ListAll() (*SSHProfileAuthInfoConfigList, error) {
	var list SSHProfileAuthInfoConfigList
	if err := r.c.ReadQuery(BasePath+SSHProfileAuthInfoEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SSHProfileAuthInfo configuration identified by id.
func (r *SSHProfileAuthInfoResource) Get(id string) (*SSHProfileAuthInfoConfig, error) {
	var item SSHProfileAuthInfoConfig
	if err := r.c.ReadQuery(BasePath+SSHProfileAuthInfoEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SSHProfileAuthInfo configuration.
func (r *SSHProfileAuthInfoResource) Create(item SSHProfileAuthInfoConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SSHProfileAuthInfoEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SSHProfileAuthInfo configuration identified by id.
func (r *SSHProfileAuthInfoResource) Edit(id string, item SSHProfileAuthInfoConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SSHProfileAuthInfoEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SSHProfileAuthInfo configuration identified by id.
func (r *SSHProfileAuthInfoResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SSHProfileAuthInfoEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
