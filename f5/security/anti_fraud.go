// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AntiFraudConfigList holds a list of AntiFraud configuration.
type AntiFraudConfigList struct {
	Items    []AntiFraudConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// AntiFraudConfig holds the configuration of a single AntiFraud.
type AntiFraudConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// AntiFraudEndpoint represents the REST resource for managing AntiFraud.
const AntiFraudEndpoint = "/anti-fraud"

// AntiFraudResource provides an API to manage AntiFraud configurations.
type AntiFraudResource struct {
	c *f5.Client
}

// ListAll  lists all the AntiFraud configurations.
func (r *AntiFraudResource) ListAll() (*AntiFraudConfigList, error) {
	var list AntiFraudConfigList
	if err := r.c.ReadQuery(BasePath+AntiFraudEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AntiFraud configuration identified by id.
func (r *AntiFraudResource) Get(id string) (*AntiFraudConfig, error) {
	var item AntiFraudConfig
	if err := r.c.ReadQuery(BasePath+AntiFraudEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AntiFraud configuration.
func (r *AntiFraudResource) Create(item AntiFraudConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AntiFraudEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AntiFraud configuration identified by id.
func (r *AntiFraudResource) Edit(id string, item AntiFraudConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AntiFraudEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AntiFraud configuration identified by id.
func (r *AntiFraudResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AntiFraudEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
