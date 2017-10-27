// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DiskApplicationVolumeConfigList holds a list of DiskApplicationVolume configuration.
type DiskApplicationVolumeConfigList struct {
	Items    []DiskApplicationVolumeConfig `json:"items,omitempty"`
	Kind     string                        `json:"kind,omitempty"`
	SelfLink string                        `json:"selflink,omitempty"`
}

// DiskApplicationVolumeConfig holds the configuration of a single DiskApplicationVolume.
type DiskApplicationVolumeConfig struct {
	FullPath             string `json:"fullPath,omitempty"`
	Generation           int    `json:"generation,omitempty"`
	Kind                 string `json:"kind,omitempty"`
	LogicalDisk          string `json:"logicalDisk,omitempty"`
	LogicalDiskReference struct {
		Link string `json:"link,omitempty"`
	} `json:"logicalDiskReference,omitempty"`
	Name                         string `json:"name,omitempty"`
	Owner                        string `json:"owner,omitempty"`
	Preservability               string `json:"preservability,omitempty"`
	Resizeable                   string `json:"resizeable,omitempty"`
	SelfLink                     string `json:"selfLink,omitempty"`
	Size                         int    `json:"size,omitempty"`
	VolumeSetVisibilityRestraint string `json:"volumeSetVisibilityRestraint,omitempty"`
}

// DiskApplicationVolumeEndpoint represents the REST resource for managing DiskApplicationVolume.
const DiskApplicationVolumeEndpoint = "/disk/application-volume"

// DiskApplicationVolumeResource provides an API to manage DiskApplicationVolume configurations.
type DiskApplicationVolumeResource struct {
	c *f5.Client
}

// ListAll  lists all the DiskApplicationVolume configurations.
func (r *DiskApplicationVolumeResource) ListAll() (*DiskApplicationVolumeConfigList, error) {
	var list DiskApplicationVolumeConfigList
	if err := r.c.ReadQuery(BasePath+DiskApplicationVolumeEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DiskApplicationVolume configuration identified by id.
func (r *DiskApplicationVolumeResource) Get(id string) (*DiskApplicationVolumeConfig, error) {
	var item DiskApplicationVolumeConfig
	if err := r.c.ReadQuery(BasePath+DiskApplicationVolumeEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DiskApplicationVolume configuration.
func (r *DiskApplicationVolumeResource) Create(item DiskApplicationVolumeConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DiskApplicationVolumeEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DiskApplicationVolume configuration identified by id.
func (r *DiskApplicationVolumeResource) Edit(id string, item DiskApplicationVolumeConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DiskApplicationVolumeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DiskApplicationVolume configuration identified by id.
func (r *DiskApplicationVolumeResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DiskApplicationVolumeEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
