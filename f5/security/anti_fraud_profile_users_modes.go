// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AntiFraudProfileUsersModesConfigList holds a list of AntiFraudProfileUsersModes configuration.
type AntiFraudProfileUsersModesConfigList struct {
	Items    []AntiFraudProfileUsersModesConfig `json:"items"`
	Kind     string                             `json:"kind"`
	SelfLink string                             `json:"selflink"`
}

// AntiFraudProfileUsersModesConfig holds the configuration of a single AntiFraudProfileUsersModes.
type AntiFraudProfileUsersModesConfig struct {
}

// AntiFraudProfileUsersModesEndpoint represents the REST resource for managing AntiFraudProfileUsersModes.
const AntiFraudProfileUsersModesEndpoint = "/anti-fraud/profile_users_modes"

// AntiFraudProfileUsersModesResource provides an API to manage AntiFraudProfileUsersModes configurations.
type AntiFraudProfileUsersModesResource struct {
	c *f5.Client
}

// ListAll  lists all the AntiFraudProfileUsersModes configurations.
func (r *AntiFraudProfileUsersModesResource) ListAll() (*AntiFraudProfileUsersModesConfigList, error) {
	var list AntiFraudProfileUsersModesConfigList
	if err := r.c.ReadQuery(BasePath+AntiFraudProfileUsersModesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AntiFraudProfileUsersModes configuration identified by id.
func (r *AntiFraudProfileUsersModesResource) Get(id string) (*AntiFraudProfileUsersModesConfig, error) {
	var item AntiFraudProfileUsersModesConfig
	if err := r.c.ReadQuery(BasePath+AntiFraudProfileUsersModesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AntiFraudProfileUsersModes configuration.
func (r *AntiFraudProfileUsersModesResource) Create(item AntiFraudProfileUsersModesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AntiFraudProfileUsersModesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AntiFraudProfileUsersModes configuration identified by id.
func (r *AntiFraudProfileUsersModesResource) Edit(id string, item AntiFraudProfileUsersModesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AntiFraudProfileUsersModesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AntiFraudProfileUsersModes configuration identified by id.
func (r *AntiFraudProfileUsersModesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AntiFraudProfileUsersModesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
