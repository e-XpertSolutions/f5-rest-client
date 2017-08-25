// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DeviceDeviceContextNATRuleConfigList holds a list of DeviceDeviceContextNATRule configuration.
type DeviceDeviceContextNATRuleConfigList struct {
	Items    []DeviceDeviceContextNATRuleConfig `json:"items"`
	Kind     string                             `json:"kind"`
	SelfLink string                             `json:"selflink"`
}

// DeviceDeviceContextNATRuleConfig holds the configuration of a single DeviceDeviceContextNATRule.
type DeviceDeviceContextNATRuleConfig struct {
}

// DeviceDeviceContextNATRuleEndpoint represents the REST resource for managing DeviceDeviceContextNATRule.
const DeviceDeviceContextNATRuleEndpoint = "/device/device-context_nat-rules"

// DeviceDeviceContextNATRuleResource provides an API to manage DeviceDeviceContextNATRule configurations.
type DeviceDeviceContextNATRuleResource struct {
	c *f5.Client
}

// ListAll  lists all the DeviceDeviceContextNATRule configurations.
func (r *DeviceDeviceContextNATRuleResource) ListAll() (*DeviceDeviceContextNATRuleConfigList, error) {
	var list DeviceDeviceContextNATRuleConfigList
	if err := r.c.ReadQuery(BasePath+DeviceDeviceContextNATRuleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DeviceDeviceContextNATRule configuration identified by id.
func (r *DeviceDeviceContextNATRuleResource) Get(id string) (*DeviceDeviceContextNATRuleConfig, error) {
	var item DeviceDeviceContextNATRuleConfig
	if err := r.c.ReadQuery(BasePath+DeviceDeviceContextNATRuleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DeviceDeviceContextNATRule configuration.
func (r *DeviceDeviceContextNATRuleResource) Create(item DeviceDeviceContextNATRuleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DeviceDeviceContextNATRuleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DeviceDeviceContextNATRule configuration identified by id.
func (r *DeviceDeviceContextNATRuleResource) Edit(id string, item DeviceDeviceContextNATRuleConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DeviceDeviceContextNATRuleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DeviceDeviceContextNATRule configuration identified by id.
func (r *DeviceDeviceContextNATRuleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DeviceDeviceContextNATRuleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
