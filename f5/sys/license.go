// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LicenseConfigList holds a list of License configuration.
type LicenseConfigList struct {
	Items    []LicenseConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// LicenseConfig holds the configuration of a single License.
type LicenseConfig struct {
}

// LicenseEndpoint represents the REST resource for managing License.
const LicenseEndpoint = "/tm/sys/license"

// LicenseResource provides an API to manage License configurations.
type LicenseResource struct {
	c f5.Client
}

// ListAll  lists all the License configurations.
func (r *LicenseResource) ListAll() (*LicenseConfigList, error) {
	var list LicenseConfigList
	if err := r.c.ReadQuery(BasePath+LicenseEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single License configuration identified by id.
func (r *LicenseResource) Get(id string) (*LicenseConfig, error) {
	var item LicenseConfig
	if err := r.c.ReadQuery(BasePath+LicenseEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new License configuration.
func (r *LicenseResource) Create(item LicenseConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LicenseEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a License configuration identified by id.
func (r *LicenseResource) Edit(id string, item LicenseConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LicenseEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single License configuration identified by id.
func (r *LicenseResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LicenseEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
