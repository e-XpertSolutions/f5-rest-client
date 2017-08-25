// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SFlowConfigList holds a list of SFlow configuration.
type SFlowConfigList struct {
	Items    []SFlowConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// SFlowConfig holds the configuration of a single SFlow.
type SFlowConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// SFlowEndpoint represents the REST resource for managing SFlow.
const SFlowEndpoint = "/sflow"

// SFlowResource provides an API to manage SFlow configurations.
type SFlowResource struct {
	c *f5.Client
}

// ListAll  lists all the SFlow configurations.
func (r *SFlowResource) ListAll() (*SFlowConfigList, error) {
	var list SFlowConfigList
	if err := r.c.ReadQuery(BasePath+SFlowEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SFlow configuration identified by id.
func (r *SFlowResource) Get(id string) (*SFlowConfig, error) {
	var item SFlowConfig
	if err := r.c.ReadQuery(BasePath+SFlowEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SFlow configuration.
func (r *SFlowResource) Create(item SFlowConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SFlowEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SFlow configuration identified by id.
func (r *SFlowResource) Edit(id string, item SFlowConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SFlowEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SFlow configuration identified by id.
func (r *SFlowResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SFlowEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
