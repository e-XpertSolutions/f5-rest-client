// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SecurityConfigList holds a list of Security configuration.
type SecurityConfigList struct {
	Items    []SecurityConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// SecurityConfig holds the configuration of a single Security.
type SecurityConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// SecurityEndpoint represents the REST resource for managing Security.
const SecurityEndpoint = ""

// SecurityResource provides an API to manage Security configurations.
type SecurityResource struct {
	c *f5.Client
}

// ListAll  lists all the Security configurations.
func (r *SecurityResource) ListAll() (*SecurityConfigList, error) {
	var list SecurityConfigList
	if err := r.c.ReadQuery(BasePath+SecurityEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Security configuration identified by id.
func (r *SecurityResource) Get(id string) (*SecurityConfig, error) {
	var item SecurityConfig
	if err := r.c.ReadQuery(BasePath+SecurityEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Security configuration.
func (r *SecurityResource) Create(item SecurityConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SecurityEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Security configuration identified by id.
func (r *SecurityResource) Edit(id string, item SecurityConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SecurityEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Security configuration identified by id.
func (r *SecurityResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SecurityEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
