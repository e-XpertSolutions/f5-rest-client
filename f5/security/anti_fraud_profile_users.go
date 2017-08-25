// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AntiFraudProfileUsersConfigList holds a list of AntiFraudProfileUsers configuration.
type AntiFraudProfileUsersConfigList struct {
	Items    []AntiFraudProfileUsersConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// AntiFraudProfileUsersConfig holds the configuration of a single AntiFraudProfileUsers.
type AntiFraudProfileUsersConfig struct {
}

// AntiFraudProfileUsersEndpoint represents the REST resource for managing AntiFraudProfileUsers.
const AntiFraudProfileUsersEndpoint = "/anti-fraud/profile_users"

// AntiFraudProfileUsersResource provides an API to manage AntiFraudProfileUsers configurations.
type AntiFraudProfileUsersResource struct {
	c *f5.Client
}

// ListAll  lists all the AntiFraudProfileUsers configurations.
func (r *AntiFraudProfileUsersResource) ListAll() (*AntiFraudProfileUsersConfigList, error) {
	var list AntiFraudProfileUsersConfigList
	if err := r.c.ReadQuery(BasePath+AntiFraudProfileUsersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AntiFraudProfileUsers configuration identified by id.
func (r *AntiFraudProfileUsersResource) Get(id string) (*AntiFraudProfileUsersConfig, error) {
	var item AntiFraudProfileUsersConfig
	if err := r.c.ReadQuery(BasePath+AntiFraudProfileUsersEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AntiFraudProfileUsers configuration.
func (r *AntiFraudProfileUsersResource) Create(item AntiFraudProfileUsersConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AntiFraudProfileUsersEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AntiFraudProfileUsers configuration identified by id.
func (r *AntiFraudProfileUsersResource) Edit(id string, item AntiFraudProfileUsersConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AntiFraudProfileUsersEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AntiFraudProfileUsers configuration identified by id.
func (r *AntiFraudProfileUsersResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AntiFraudProfileUsersEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
