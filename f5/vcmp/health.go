// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcmp

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// HealthConfigList holds a list of Health configuration.
type HealthConfigList struct {
	Items    []HealthConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// HealthConfig holds the configuration of a single Health.
type HealthConfig struct {
}

// HealthEndpoint represents the REST resource for managing Health.
const HealthEndpoint = "/health"

// HealthResource provides an API to manage Health configurations.
type HealthResource struct {
	c *f5.Client
}

// ListAll  lists all the Health configurations.
func (r *HealthResource) ListAll() (*HealthConfigList, error) {
	var list HealthConfigList
	if err := r.c.ReadQuery(BasePath+HealthEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Health configuration identified by id.
func (r *HealthResource) Get(id string) (*HealthConfig, error) {
	var item HealthConfig
	if err := r.c.ReadQuery(BasePath+HealthEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Health configuration.
func (r *HealthResource) Create(item HealthConfig) error {
	if err := r.c.ModQuery("POST", BasePath+HealthEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Health configuration identified by id.
func (r *HealthResource) Edit(id string, item HealthConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+HealthEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Health configuration identified by id.
func (r *HealthResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+HealthEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
