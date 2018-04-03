// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FolderConfigList holds a list of Folder configuration.
type FolderConfigList struct {
	Items    []FolderConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// FolderConfig holds the configuration of a single Folder.
type FolderConfig struct {
	DeviceGroup           string `json:"deviceGroup"`
	FullPath              string `json:"fullPath"`
	Generation            int    `json:"generation"`
	Hidden                string `json:"hidden"`
	InheritedDevicegroup  string `json:"inheritedDevicegroup"`
	InheritedTrafficGroup string `json:"inheritedTrafficGroup"`
	Kind                  string `json:"kind"`
	Name                  string `json:"name"`
	NoRefCheck            string `json:"noRefCheck"`
	SelfLink              string `json:"selfLink"`
	TrafficGroup          string `json:"trafficGroup"`
	TrafficGroupReference struct {
		Link string `json:"link"`
	} `json:"trafficGroupReference"`
}

type NewFolderConfig struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
}

// FolderEndpoint represents the REST resource for managing Folder.
const FolderEndpoint = "/folder"

// FolderResource provides an API to manage Folder configurations.
type FolderResource struct {
	c *f5.Client
}

// ListAll  lists all the Folder configurations.
func (r *FolderResource) ListAll() (*FolderConfigList, error) {
	var list FolderConfigList
	if err := r.c.ReadQuery(BasePath+FolderEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Folder configuration identified by id.
func (r *FolderResource) Get(id string) (*FolderConfig, error) {
	var item FolderConfig
	if err := r.c.ReadQuery(BasePath+FolderEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Folder configuration.
func (r *FolderResource) Create(item NewFolderConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FolderEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Folder configuration identified by id.
func (r *FolderResource) Edit(id string, item FolderConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FolderEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Folder configuration identified by id.
func (r *FolderResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FolderEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
