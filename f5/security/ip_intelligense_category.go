// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPIntelligenseCategoryConfigList holds a list of IPIntelligenseCategory configuration.
type IPIntelligenseCategoryConfigList struct {
	Items    []IPIntelligenseCategoryConfig `json:"items"`
	Kind     string                         `json:"kind"`
	SelfLink string                         `json:"selflink"`
}

// IPIntelligenseCategoryConfig holds the configuration of a single IPIntelligenseCategory.
type IPIntelligenseCategoryConfig struct {
}

// IPIntelligenseCategoryEndpoint represents the REST resource for managing IPIntelligenseCategory.
const IPIntelligenseCategoryEndpoint = "/ip-intelligence/category"

// IPIntelligenseCategoryResource provides an API to manage IPIntelligenseCategory configurations.
type IPIntelligenseCategoryResource struct {
	c *f5.Client
}

// ListAll  lists all the IPIntelligenseCategory configurations.
func (r *IPIntelligenseCategoryResource) ListAll() (*IPIntelligenseCategoryConfigList, error) {
	var list IPIntelligenseCategoryConfigList
	if err := r.c.ReadQuery(BasePath+IPIntelligenseCategoryEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPIntelligenseCategory configuration identified by id.
func (r *IPIntelligenseCategoryResource) Get(id string) (*IPIntelligenseCategoryConfig, error) {
	var item IPIntelligenseCategoryConfig
	if err := r.c.ReadQuery(BasePath+IPIntelligenseCategoryEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPIntelligenseCategory configuration.
func (r *IPIntelligenseCategoryResource) Create(item IPIntelligenseCategoryConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPIntelligenseCategoryEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPIntelligenseCategory configuration identified by id.
func (r *IPIntelligenseCategoryResource) Edit(id string, item IPIntelligenseCategoryConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPIntelligenseCategoryEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPIntelligenseCategory configuration identified by id.
func (r *IPIntelligenseCategoryResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPIntelligenseCategoryEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
