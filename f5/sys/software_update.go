// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SoftwareUpdateConfigList holds a list of SoftwareUpdate configuration.
type SoftwareUpdateConfigList struct {
	Items    []SoftwareUpdateConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// SoftwareUpdateConfig holds the configuration of a single SoftwareUpdate.
type SoftwareUpdateConfig struct {
}

// SoftwareUpdateEndpoint represents the REST resource for managing SoftwareUpdate.
const SoftwareUpdateEndpoint = "/software/update"

// SoftwareUpdateResource provides an API to manage SoftwareUpdate configurations.
type SoftwareUpdateResource struct {
	c *f5.Client
}

// ListAll  lists all the SoftwareUpdate configurations.
func (r *SoftwareUpdateResource) ListAll() (*SoftwareUpdateConfigList, error) {
	var list SoftwareUpdateConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareUpdateEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SoftwareUpdate configuration identified by id.
func (r *SoftwareUpdateResource) Get(id string) (*SoftwareUpdateConfig, error) {
	var item SoftwareUpdateConfig
	if err := r.c.ReadQuery(BasePath+SoftwareUpdateEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SoftwareUpdate configuration.
func (r *SoftwareUpdateResource) Create(item SoftwareUpdateConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SoftwareUpdateEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SoftwareUpdate configuration identified by id.
func (r *SoftwareUpdateResource) Edit(id string, item SoftwareUpdateConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SoftwareUpdateEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SoftwareUpdate configuration identified by id.
func (r *SoftwareUpdateResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SoftwareUpdateEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
