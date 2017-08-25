// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SoftwareUpdateStatusConfigList holds a list of SoftwareUpdateStatus configuration.
type SoftwareUpdateStatusConfigList struct {
	Items    []SoftwareUpdateStatusConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// SoftwareUpdateStatusConfig holds the configuration of a single SoftwareUpdateStatus.
type SoftwareUpdateStatusConfig struct {
}

// SoftwareUpdateStatusEndpoint represents the REST resource for managing SoftwareUpdateStatus.
const SoftwareUpdateStatusEndpoint = "/software/update-status"

// SoftwareUpdateStatusResource provides an API to manage SoftwareUpdateStatus configurations.
type SoftwareUpdateStatusResource struct {
	c *f5.Client
}

// ListAll  lists all the SoftwareUpdateStatus configurations.
func (r *SoftwareUpdateStatusResource) ListAll() (*SoftwareUpdateStatusConfigList, error) {
	var list SoftwareUpdateStatusConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareUpdateStatusEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SoftwareUpdateStatus configuration identified by id.
func (r *SoftwareUpdateStatusResource) Get(id string) (*SoftwareUpdateStatusConfig, error) {
	var item SoftwareUpdateStatusConfig
	if err := r.c.ReadQuery(BasePath+SoftwareUpdateStatusEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SoftwareUpdateStatus configuration.
func (r *SoftwareUpdateStatusResource) Create(item SoftwareUpdateStatusConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SoftwareUpdateStatusEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SoftwareUpdateStatus configuration identified by id.
func (r *SoftwareUpdateStatusResource) Edit(id string, item SoftwareUpdateStatusConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SoftwareUpdateStatusEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SoftwareUpdateStatus configuration identified by id.
func (r *SoftwareUpdateStatusResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SoftwareUpdateStatusEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
