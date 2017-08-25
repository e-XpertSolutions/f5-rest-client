// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PFManConsumerConfigList holds a list of PFManConsumer configuration.
type PFManConsumerConfigList struct {
	Items    []PFManConsumerConfig `json:"items"`
	Kind     string                `json:"kind"`
	SelfLink string                `json:"selflink"`
}

// PFManConsumerConfig holds the configuration of a single PFManConsumer.
type PFManConsumerConfig struct {
}

// PFManConsumerEndpoint represents the REST resource for managing PFManConsumer.
const PFManConsumerEndpoint = "/pfman/consumer"

// PFManConsumerResource provides an API to manage PFManConsumer configurations.
type PFManConsumerResource struct {
	c *f5.Client
}

// ListAll  lists all the PFManConsumer configurations.
func (r *PFManConsumerResource) ListAll() (*PFManConsumerConfigList, error) {
	var list PFManConsumerConfigList
	if err := r.c.ReadQuery(BasePath+PFManConsumerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PFManConsumer configuration identified by id.
func (r *PFManConsumerResource) Get(id string) (*PFManConsumerConfig, error) {
	var item PFManConsumerConfig
	if err := r.c.ReadQuery(BasePath+PFManConsumerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PFManConsumer configuration.
func (r *PFManConsumerResource) Create(item PFManConsumerConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PFManConsumerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PFManConsumer configuration identified by id.
func (r *PFManConsumerResource) Edit(id string, item PFManConsumerConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PFManConsumerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PFManConsumer configuration identified by id.
func (r *PFManConsumerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PFManConsumerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
