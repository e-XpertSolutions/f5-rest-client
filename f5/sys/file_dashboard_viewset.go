// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileDashboardViewsetConfigList holds a list of FileDashboardViewset configuration.
type FileDashboardViewsetConfigList struct {
	Items    []FileDashboardViewsetConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// FileDashboardViewsetConfig holds the configuration of a single FileDashboardViewset.
type FileDashboardViewsetConfig struct {
}

// FileDashboardViewsetEndpoint represents the REST resource for managing FileDashboardViewset.
const FileDashboardViewsetEndpoint = "/file/dashboard-viewset"

// FileDashboardViewsetResource provides an API to manage FileDashboardViewset configurations.
type FileDashboardViewsetResource struct {
	c *f5.Client
}

// ListAll  lists all the FileDashboardViewset configurations.
func (r *FileDashboardViewsetResource) ListAll() (*FileDashboardViewsetConfigList, error) {
	var list FileDashboardViewsetConfigList
	if err := r.c.ReadQuery(BasePath+FileDashboardViewsetEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileDashboardViewset configuration identified by id.
func (r *FileDashboardViewsetResource) Get(id string) (*FileDashboardViewsetConfig, error) {
	var item FileDashboardViewsetConfig
	if err := r.c.ReadQuery(BasePath+FileDashboardViewsetEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileDashboardViewset configuration.
func (r *FileDashboardViewsetResource) Create(item FileDashboardViewsetConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileDashboardViewsetEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileDashboardViewset configuration identified by id.
func (r *FileDashboardViewsetResource) Edit(id string, item FileDashboardViewsetConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileDashboardViewsetEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileDashboardViewset configuration identified by id.
func (r *FileDashboardViewsetResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileDashboardViewsetEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
