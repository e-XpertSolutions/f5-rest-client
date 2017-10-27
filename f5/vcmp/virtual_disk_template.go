// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcmp

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// VirtualDiskTemplateConfigList holds a list of VirtualDiskTemplate configuration.
type VirtualDiskTemplateConfigList struct {
	Items    []VirtualDiskTemplateConfig `json:"items,omitempty"`
	Kind     string                      `json:"kind,omitempty"`
	SelfLink string                      `json:"selflink,omitempty"`
}

// VirtualDiskTemplateConfig holds the configuration of a single VirtualDiskTemplate.
type VirtualDiskTemplateConfig struct {
	ISOVersion      string `json:"isoVersion,omitempty"`
	OperatingSystem string `json:"operatingSystem,omitempty"`
}

// VirtualDiskTemplateEndpoint represents the REST resource for managing VirtualDiskTemplate.
const VirtualDiskTemplateEndpoint = "/virtual-disk-template"

// VirtualDiskTemplateResource provides an API to manage VirtualDiskTemplate configurations.
type VirtualDiskTemplateResource struct {
	c *f5.Client
}

// ListAll  lists all the VirtualDiskTemplate configurations.
func (r *VirtualDiskTemplateResource) ListAll() (*VirtualDiskTemplateConfigList, error) {
	var list VirtualDiskTemplateConfigList
	if err := r.c.ReadQuery(BasePath+VirtualDiskTemplateEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single VirtualDiskTemplate configuration identified by id.
func (r *VirtualDiskTemplateResource) Get(id string) (*VirtualDiskTemplateConfig, error) {
	var item VirtualDiskTemplateConfig
	if err := r.c.ReadQuery(BasePath+VirtualDiskTemplateEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new VirtualDiskTemplate configuration.
func (r *VirtualDiskTemplateResource) Create(item VirtualDiskTemplateConfig) error {
	if err := r.c.ModQuery("POST", BasePath+VirtualDiskTemplateEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a VirtualDiskTemplate configuration identified by id.
func (r *VirtualDiskTemplateResource) Edit(id string, item VirtualDiskTemplateConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+VirtualDiskTemplateEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single VirtualDiskTemplate configuration identified by id.
func (r *VirtualDiskTemplateResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+VirtualDiskTemplateEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
