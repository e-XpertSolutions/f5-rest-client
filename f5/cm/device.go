// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A DeviceList holds a list of Device.
type DeviceList struct {
	Items    []Device `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selfLink,omitempty"`
}

// A Device hold the configuration for a device.
type Device struct {
	ActiveModules []string `json:"activeModules,omitempty"`
	BaseMac       string   `json:"baseMac,omitempty"`
	Build         string   `json:"build,omitempty"`
	Cert          string   `json:"cert,omitempty"`
	CertReference struct {
		Link string `json:"link,omitempty"`
	} `json:"certReference,omitempty"`
	ChassisID     string `json:"chassisId,omitempty"`
	ChassisType   string `json:"chassisType,omitempty"`
	ConfigsyncIP  string `json:"configsyncIp,omitempty"`
	Edition       string `json:"edition,omitempty"`
	FailoverState string `json:"failoverState,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	HaCapacity    int    `json:"haCapacity,omitempty"`
	Hostname      string `json:"hostname,omitempty"`
	Key           string `json:"key,omitempty"`
	KeyReference  struct {
		Link string `json:"link,omitempty"`
	} `json:"keyReference,omitempty"`
	Kind              string   `json:"kind,omitempty"`
	ManagementIP      string   `json:"managementIp,omitempty"`
	MarketingName     string   `json:"marketingName,omitempty"`
	MirrorIP          string   `json:"mirrorIp,omitempty"`
	MirrorSecondaryIP string   `json:"mirrorSecondaryIp,omitempty"`
	MulticastIP       string   `json:"multicastIp,omitempty"`
	MulticastPort     int      `json:"multicastPort,omitempty"`
	Name              string   `json:"name,omitempty"`
	OptionalModules   []string `json:"optionalModules,omitempty"`
	PlatformID        string   `json:"platformId,omitempty"`
	Product           string   `json:"product,omitempty"`
	SelfDevice        string   `json:"selfDevice,omitempty"`
	SelfLink          string   `json:"selfLink,omitempty"`
	TimeZone          string   `json:"timeZone,omitempty"`
	Version           string   `json:"version,omitempty"`
}

// DeviceEndpoint represents the REST resource for managing a device.
const DeviceEndpoint = "/device"

// A DeviceResource provides API to manage Devices configuration.
type DeviceResource struct {
	c *f5.Client
}

// ListAll lists all the Device configurations.
func (pr *DeviceResource) ListAll() (*DeviceList, error) {
	var list DeviceList
	if err := pr.c.ReadQuery(BasePath+DeviceEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Device configuration identified by id.
func (pr *DeviceResource) Get(id string) (*Device, error) {
	var item Device
	if err := pr.c.ReadQuery(BasePath+DeviceEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Device configuration.
func (pr *DeviceResource) Create(item Device) error {
	if err := pr.c.ModQuery("POST", BasePath+DeviceEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Device configuration identified by id.
func (pr *DeviceResource) Edit(id string, item Device) error {
	if err := pr.c.ModQuery("PUT", BasePath+DeviceEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Device configuration identified by id.
func (pr *DeviceResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+DeviceEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
