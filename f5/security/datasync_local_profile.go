// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DatasyncLocalProfileConfigList holds a list of DatasyncLocalProfile configuration.
type DatasyncLocalProfileConfigList struct {
	Items    []DatasyncLocalProfileConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// DatasyncLocalProfileConfig holds the configuration of a single DatasyncLocalProfile.
type DatasyncLocalProfileConfig struct {
	BufSize         int    `json:"bufSize"`
	DsArea          string `json:"dsArea"`
	FullPath        string `json:"fullPath"`
	GenPauseSec     int    `json:"genPauseSec"`
	GenTimeoutSec   int    `json:"genTimeoutSec"`
	Generation      int    `json:"generation"`
	KeepConfFiles   int    `json:"keepConfFiles"`
	Kind            string `json:"kind"`
	MaxGenRows      string `json:"maxGenRows"`
	MinCPUPercent   int    `json:"minCpuPercent"`
	MinMemMb        int    `json:"minMemMb"`
	Name            string `json:"name"`
	OfflineUntilGen string `json:"offlineUntilGen"`
	RowsBulk        int    `json:"rowsBulk"`
	SelfLink        string `json:"selfLink"`
}

// DatasyncLocalProfileEndpoint represents the REST resource for managing DatasyncLocalProfile.
const DatasyncLocalProfileEndpoint = "/datasync/local-profile"

// DatasyncLocalProfileResource provides an API to manage DatasyncLocalProfile configurations.
type DatasyncLocalProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the DatasyncLocalProfile configurations.
func (r *DatasyncLocalProfileResource) ListAll() (*DatasyncLocalProfileConfigList, error) {
	var list DatasyncLocalProfileConfigList
	if err := r.c.ReadQuery(BasePath+DatasyncLocalProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DatasyncLocalProfile configuration identified by id.
func (r *DatasyncLocalProfileResource) Get(id string) (*DatasyncLocalProfileConfig, error) {
	var item DatasyncLocalProfileConfig
	if err := r.c.ReadQuery(BasePath+DatasyncLocalProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DatasyncLocalProfile configuration.
func (r *DatasyncLocalProfileResource) Create(item DatasyncLocalProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DatasyncLocalProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DatasyncLocalProfile configuration identified by id.
func (r *DatasyncLocalProfileResource) Edit(id string, item DatasyncLocalProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DatasyncLocalProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DatasyncLocalProfile configuration identified by id.
func (r *DatasyncLocalProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DatasyncLocalProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
