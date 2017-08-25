// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AnalyticsConfigList holds a list of Analytics configuration.
type AnalyticsConfigList struct {
	Items    []AnalyticsConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// AnalyticsConfig holds the configuration of a single Analytics.
type AnalyticsConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// AnalyticsEndpoint represents the REST resource for managing Analytics.
const AnalyticsEndpoint = "/analytics"

// AnalyticsResource provides an API to manage Analytics configurations.
type AnalyticsResource struct {
	c *f5.Client
}

// ListAll  lists all the Analytics configurations.
func (r *AnalyticsResource) ListAll() (*AnalyticsConfigList, error) {
	var list AnalyticsConfigList
	if err := r.c.ReadQuery(BasePath+AnalyticsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Analytics configuration identified by id.
func (r *AnalyticsResource) Get(id string) (*AnalyticsConfig, error) {
	var item AnalyticsConfig
	if err := r.c.ReadQuery(BasePath+AnalyticsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Analytics configuration.
func (r *AnalyticsResource) Create(item AnalyticsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AnalyticsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Analytics configuration identified by id.
func (r *AnalyticsResource) Edit(id string, item AnalyticsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AnalyticsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Analytics configuration identified by id.
func (r *AnalyticsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AnalyticsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
