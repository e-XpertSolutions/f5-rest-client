// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A DeviceGroupList holds a list of DeviceGroup.
type DeviceGroupList struct {
	Items    []DeviceGroup `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selfLink,omitempty"`
}

// A DeviceGroup hold the configuration for a DeviceGroup.
type DeviceGroup struct {
	AsmSync          string `json:"asmSync,omitempty"`
	AutoSync         string `json:"autoSync,omitempty"`
	DevicesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"devicesReference,omitempty"`
	FullLoadOnSync               string `json:"fullLoadOnSync,omitempty"`
	FullPath                     string `json:"fullPath,omitempty"`
	Generation                   int    `json:"generation,omitempty"`
	IncrementalConfigSyncSizeMax int    `json:"incrementalConfigSyncSizeMax,omitempty"`
	Kind                         string `json:"kind,omitempty"`
	Name                         string `json:"name,omitempty"`
	NetworkFailover              string `json:"networkFailover,omitempty"`
	Partition                    string `json:"partition,omitempty"`
	SaveOnAutoSync               string `json:"saveOnAutoSync,omitempty"`
	SelfLink                     string `json:"selfLink,omitempty"`
	Type                         string `json:"type,omitempty"`
}

// DeviceGroupEndpoint represents the REST resource for managing a DeviceGroup.
const DeviceGroupEndpoint = "/device-group"

// A DeviceGroupResource provides API to manage DeviceGroups configuration.
type DeviceGroupResource struct {
	c *f5.Client
}

// ListAll lists all the DeviceGroup configurations.
func (pr *DeviceGroupResource) ListAll() (*DeviceGroupList, error) {
	var list DeviceGroupList
	if err := pr.c.ReadQuery(BasePath+DeviceGroupEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DeviceGroup configuration identified by id.
func (pr *DeviceGroupResource) Get(id string) (*DeviceGroup, error) {
	var item DeviceGroup
	if err := pr.c.ReadQuery(BasePath+DeviceGroupEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DeviceGroup configuration.
func (pr *DeviceGroupResource) Create(item DeviceGroup) error {
	if err := pr.c.ModQuery("POST", BasePath+DeviceGroupEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DeviceGroup configuration identified by id.
func (pr *DeviceGroupResource) Edit(id string, item DeviceGroup) error {
	if err := pr.c.ModQuery("PUT", BasePath+DeviceGroupEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DeviceGroup configuration identified by id.
func (pr *DeviceGroupResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+DeviceGroupEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
