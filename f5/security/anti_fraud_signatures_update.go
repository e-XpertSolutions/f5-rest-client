// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AntiFraudSignaturesUpdateConfigList holds a list of AntiFraudSignaturesUpdate configuration.
type AntiFraudSignaturesUpdateConfigList struct {
	Items    []AntiFraudSignaturesUpdateConfig `json:"items"`
	Kind     string                            `json:"kind"`
	SelfLink string                            `json:"selflink"`
}

// AntiFraudSignaturesUpdateConfig holds the configuration of a single AntiFraudSignaturesUpdate.
type AntiFraudSignaturesUpdateConfig struct {
}

// AntiFraudSignaturesUpdateEndpoint represents the REST resource for managing AntiFraudSignaturesUpdate.
const AntiFraudSignaturesUpdateEndpoint = "/anti-fraud/signatures-update"

// AntiFraudSignaturesUpdateResource provides an API to manage AntiFraudSignaturesUpdate configurations.
type AntiFraudSignaturesUpdateResource struct {
	c *f5.Client
}

// ListAll  lists all the AntiFraudSignaturesUpdate configurations.
func (r *AntiFraudSignaturesUpdateResource) ListAll() (*AntiFraudSignaturesUpdateConfigList, error) {
	var list AntiFraudSignaturesUpdateConfigList
	if err := r.c.ReadQuery(BasePath+AntiFraudSignaturesUpdateEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AntiFraudSignaturesUpdate configuration identified by id.
func (r *AntiFraudSignaturesUpdateResource) Get(id string) (*AntiFraudSignaturesUpdateConfig, error) {
	var item AntiFraudSignaturesUpdateConfig
	if err := r.c.ReadQuery(BasePath+AntiFraudSignaturesUpdateEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AntiFraudSignaturesUpdate configuration.
func (r *AntiFraudSignaturesUpdateResource) Create(item AntiFraudSignaturesUpdateConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AntiFraudSignaturesUpdateEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AntiFraudSignaturesUpdate configuration identified by id.
func (r *AntiFraudSignaturesUpdateResource) Edit(id string, item AntiFraudSignaturesUpdateConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AntiFraudSignaturesUpdateEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AntiFraudSignaturesUpdate configuration identified by id.
func (r *AntiFraudSignaturesUpdateResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AntiFraudSignaturesUpdateEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
