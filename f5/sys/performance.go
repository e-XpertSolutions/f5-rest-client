// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PerformanceConfigList holds a list of Performance configuration.
type PerformanceConfigList struct {
	Items    []PerformanceConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// PerformanceConfig holds the configuration of a single Performance.
type PerformanceConfig struct {
}

// PerformanceEndpoint represents the REST resource for managing Performance.
const PerformanceEndpoint = "/performance"

// PerformanceResource provides an API to manage Performance configurations.
type PerformanceResource struct {
	c *f5.Client
}

// ListAll  lists all the Performance configurations.
func (r *PerformanceResource) ListAll() (*PerformanceConfigList, error) {
	var list PerformanceConfigList
	if err := r.c.ReadQuery(BasePath+PerformanceEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Performance configuration identified by id.
func (r *PerformanceResource) Get(id string) (*PerformanceConfig, error) {
	var item PerformanceConfig
	if err := r.c.ReadQuery(BasePath+PerformanceEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Performance configuration.
func (r *PerformanceResource) Create(item PerformanceConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PerformanceEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Performance configuration identified by id.
func (r *PerformanceResource) Edit(id string, item PerformanceConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PerformanceEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Performance configuration identified by id.
func (r *PerformanceResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PerformanceEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
