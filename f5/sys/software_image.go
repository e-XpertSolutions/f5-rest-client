// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SoftwareImageConfigList holds a list of SoftwareImage configuration.
type SoftwareImageConfigList struct {
	Items    []SoftwareImageConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// SoftwareImageConfig holds the configuration of a single SoftwareImage.
type SoftwareImageConfig struct {
	Build        string `json:"build"`
	BuildDate    string `json:"buildDate"`
	Checksum     string `json:"checksum"`
	FileSize     string `json:"fileSize"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Kind         string `json:"kind"`
	LastModified string `json:"lastModified"`
	Name         string `json:"name"`
	Product      string `json:"product"`
	SelfLink     string `json:"selfLink"`
	Verified     string `json:"verified"`
	Version      string `json:"version"`
}

// SoftwareImageEndpoint represents the REST resource for managing SoftwareImage.
const SoftwareImageEndpoint = "/software/image"

// SoftwareImageResource provides an API to manage SoftwareImage configurations.
type SoftwareImageResource struct {
	c *f5.Client
}

// ListAll  lists all the SoftwareImage configurations.
func (r *SoftwareImageResource) ListAll() (*SoftwareImageConfigList, error) {
	var list SoftwareImageConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareImageEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SoftwareImage configuration identified by id.
func (r *SoftwareImageResource) Get(id string) (*SoftwareImageConfig, error) {
	var item SoftwareImageConfig
	if err := r.c.ReadQuery(BasePath+SoftwareImageEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SoftwareImage configuration.
func (r *SoftwareImageResource) Create(item SoftwareImageConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SoftwareImageEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SoftwareImage configuration identified by id.
func (r *SoftwareImageResource) Edit(id string, item SoftwareImageConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SoftwareImageEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SoftwareImage configuration identified by id.
func (r *SoftwareImageResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SoftwareImageEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
