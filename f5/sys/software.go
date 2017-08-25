// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SoftwareConfigList holds a list of Software configuration.
type SoftwareConfigList struct {
	Items    []SoftwareConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// SoftwareConfig holds the configuration of a single Software.
type SoftwareConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// SoftwareEndpoint represents the REST resource for managing Software.
const SoftwareEndpoint = "/software"

// SoftwareResource provides an API to manage Software configurations.
type SoftwareResource struct {
	c *f5.Client
}

// ListAll  lists all the Software configurations.
func (r *SoftwareResource) ListAll() (*SoftwareConfigList, error) {
	var list SoftwareConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Software configuration identified by id.
func (r *SoftwareResource) Get(id string) (*SoftwareConfig, error) {
	var item SoftwareConfig
	if err := r.c.ReadQuery(BasePath+SoftwareEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Software configuration.
func (r *SoftwareResource) Create(item SoftwareConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SoftwareEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Software configuration identified by id.
func (r *SoftwareResource) Edit(id string, item SoftwareConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SoftwareEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Software configuration identified by id.
func (r *SoftwareResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SoftwareEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
