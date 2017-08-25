// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SyncSysFilesConfigList holds a list of SyncSysFiles configuration.
type SyncSysFilesConfigList struct {
	Items    []SyncSysFilesConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// SyncSysFilesConfig holds the configuration of a single SyncSysFiles.
type SyncSysFilesConfig struct {
}

// SyncSysFilesEndpoint represents the REST resource for managing SyncSysFiles.
const SyncSysFilesEndpoint = "/sync-sys-files"

// SyncSysFilesResource provides an API to manage SyncSysFiles configurations.
type SyncSysFilesResource struct {
	c *f5.Client
}

// ListAll  lists all the SyncSysFiles configurations.
func (r *SyncSysFilesResource) ListAll() (*SyncSysFilesConfigList, error) {
	var list SyncSysFilesConfigList
	if err := r.c.ReadQuery(BasePath+SyncSysFilesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SyncSysFiles configuration identified by id.
func (r *SyncSysFilesResource) Get(id string) (*SyncSysFilesConfig, error) {
	var item SyncSysFilesConfig
	if err := r.c.ReadQuery(BasePath+SyncSysFilesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SyncSysFiles configuration.
func (r *SyncSysFilesResource) Create(item SyncSysFilesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SyncSysFilesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SyncSysFiles configuration identified by id.
func (r *SyncSysFilesResource) Edit(id string, item SyncSysFilesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SyncSysFilesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SyncSysFiles configuration identified by id.
func (r *SyncSysFilesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SyncSysFilesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
