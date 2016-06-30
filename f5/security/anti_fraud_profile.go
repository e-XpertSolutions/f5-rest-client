// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AntiFraudProfileConfigList holds a list of AntiFraudProfile configuration.
type AntiFraudProfileConfigList struct {
	Items    []AntiFraudProfileConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// AntiFraudProfileConfig holds the configuration of a single AntiFraudProfile.
type AntiFraudProfileConfig struct {
}

// AntiFraudProfileEndpoint represents the REST resource for managing AntiFraudProfile.
const AntiFraudProfileEndpoint = "/anti-fraud/profile"

// AntiFraudProfileResource provides an API to manage AntiFraudProfile configurations.
type AntiFraudProfileResource struct {
	c f5.Client
}

// ListAll  lists all the AntiFraudProfile configurations.
func (r *AntiFraudProfileResource) ListAll() (*AntiFraudProfileConfigList, error) {
	var list AntiFraudProfileConfigList
	if err := r.c.ReadQuery(BasePath+AntiFraudProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AntiFraudProfile configuration identified by id.
func (r *AntiFraudProfileResource) Get(id string) (*AntiFraudProfileConfig, error) {
	var item AntiFraudProfileConfig
	if err := r.c.ReadQuery(BasePath+AntiFraudProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AntiFraudProfile configuration.
func (r *AntiFraudProfileResource) Create(item AntiFraudProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AntiFraudProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AntiFraudProfile configuration identified by id.
func (r *AntiFraudProfileResource) Edit(id string, item AntiFraudProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AntiFraudProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AntiFraudProfile configuration identified by id.
func (r *AntiFraudProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AntiFraudProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
