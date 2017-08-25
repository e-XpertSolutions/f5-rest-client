// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// URLDBDownloadResultConfigList holds a list of URLDBDownloadResult configuration.
type URLDBDownloadResultConfigList struct {
	Items    []URLDBDownloadResultConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// URLDBDownloadResultConfig holds the configuration of a single URLDBDownloadResult.
type URLDBDownloadResultConfig struct {
}

// URLDBDownloadResultEndpoint represents the REST resource for managing URLDBDownloadResult.
const URLDBDownloadResultEndpoint = "/url-db/download-result"

// URLDBDownloadResultResource provides an API to manage URLDBDownloadResult configurations.
type URLDBDownloadResultResource struct {
	c *f5.Client
}

// ListAll  lists all the URLDBDownloadResult configurations.
func (r *URLDBDownloadResultResource) ListAll() (*URLDBDownloadResultConfigList, error) {
	var list URLDBDownloadResultConfigList
	if err := r.c.ReadQuery(BasePath+URLDBDownloadResultEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single URLDBDownloadResult configuration identified by id.
func (r *URLDBDownloadResultResource) Get(id string) (*URLDBDownloadResultConfig, error) {
	var item URLDBDownloadResultConfig
	if err := r.c.ReadQuery(BasePath+URLDBDownloadResultEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new URLDBDownloadResult configuration.
func (r *URLDBDownloadResultResource) Create(item URLDBDownloadResultConfig) error {
	if err := r.c.ModQuery("POST", BasePath+URLDBDownloadResultEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a URLDBDownloadResult configuration identified by id.
func (r *URLDBDownloadResultResource) Edit(id string, item URLDBDownloadResultConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+URLDBDownloadResultEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single URLDBDownloadResult configuration identified by id.
func (r *URLDBDownloadResultResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+URLDBDownloadResultEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
