// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ApplicationCustomStatConfigList holds a list of ApplicationCustomStat configuration.
type ApplicationCustomStatConfigList struct {
	Items    []ApplicationCustomStatConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// ApplicationCustomStatConfig holds the configuration of a single ApplicationCustomStat.
type ApplicationCustomStatConfig struct {
}

// ApplicationCustomStatEndpoint represents the REST resource for managing ApplicationCustomStat.
const ApplicationCustomStatEndpoint = "/application/custom-stat"

// ApplicationCustomStatResource provides an API to manage ApplicationCustomStat configurations.
type ApplicationCustomStatResource struct {
	c *f5.Client
}

// ListAll  lists all the ApplicationCustomStat configurations.
func (r *ApplicationCustomStatResource) ListAll() (*ApplicationCustomStatConfigList, error) {
	var list ApplicationCustomStatConfigList
	if err := r.c.ReadQuery(BasePath+ApplicationCustomStatEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ApplicationCustomStat configuration identified by id.
func (r *ApplicationCustomStatResource) Get(id string) (*ApplicationCustomStatConfig, error) {
	var item ApplicationCustomStatConfig
	if err := r.c.ReadQuery(BasePath+ApplicationCustomStatEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ApplicationCustomStat configuration.
func (r *ApplicationCustomStatResource) Create(item ApplicationCustomStatConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ApplicationCustomStatEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ApplicationCustomStat configuration identified by id.
func (r *ApplicationCustomStatResource) Edit(id string, item ApplicationCustomStatConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ApplicationCustomStatEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ApplicationCustomStat configuration identified by id.
func (r *ApplicationCustomStatResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ApplicationCustomStatEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
