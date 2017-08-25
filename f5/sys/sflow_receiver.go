// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SFlowReceiverConfigList holds a list of SFlowReceiver configuration.
type SFlowReceiverConfigList struct {
	Items    []SFlowReceiverConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// SFlowReceiverConfig holds the configuration of a single SFlowReceiver.
type SFlowReceiverConfig struct {
}

// SFlowReceiverEndpoint represents the REST resource for managing SFlowReceiver.
const SFlowReceiverEndpoint = "/sflow/receiver"

// SFlowReceiverResource provides an API to manage SFlowReceiver configurations.
type SFlowReceiverResource struct {
	c *f5.Client
}

// ListAll  lists all the SFlowReceiver configurations.
func (r *SFlowReceiverResource) ListAll() (*SFlowReceiverConfigList, error) {
	var list SFlowReceiverConfigList
	if err := r.c.ReadQuery(BasePath+SFlowReceiverEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SFlowReceiver configuration identified by id.
func (r *SFlowReceiverResource) Get(id string) (*SFlowReceiverConfig, error) {
	var item SFlowReceiverConfig
	if err := r.c.ReadQuery(BasePath+SFlowReceiverEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SFlowReceiver configuration.
func (r *SFlowReceiverResource) Create(item SFlowReceiverConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SFlowReceiverEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SFlowReceiver configuration identified by id.
func (r *SFlowReceiverResource) Edit(id string, item SFlowReceiverConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SFlowReceiverEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SFlowReceiver configuration identified by id.
func (r *SFlowReceiverResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SFlowReceiverEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
