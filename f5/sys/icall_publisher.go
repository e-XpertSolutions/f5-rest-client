// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ICallPublisherConfigList holds a list of ICallPublisher configuration.
type ICallPublisherConfigList struct {
	Items    []ICallPublisherConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// ICallPublisherConfig holds the configuration of a single ICallPublisher.
type ICallPublisherConfig struct {
}

// ICallPublisherEndpoint represents the REST resource for managing ICallPublisher.
const ICallPublisherEndpoint = "/icall/publisher"

// ICallPublisherResource provides an API to manage ICallPublisher configurations.
type ICallPublisherResource struct {
	c *f5.Client
}

// ListAll  lists all the ICallPublisher configurations.
func (r *ICallPublisherResource) ListAll() (*ICallPublisherConfigList, error) {
	var list ICallPublisherConfigList
	if err := r.c.ReadQuery(BasePath+ICallPublisherEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ICallPublisher configuration identified by id.
func (r *ICallPublisherResource) Get(id string) (*ICallPublisherConfig, error) {
	var item ICallPublisherConfig
	if err := r.c.ReadQuery(BasePath+ICallPublisherEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ICallPublisher configuration.
func (r *ICallPublisherResource) Create(item ICallPublisherConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ICallPublisherEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ICallPublisher configuration identified by id.
func (r *ICallPublisherResource) Edit(id string, item ICallPublisherConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ICallPublisherEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ICallPublisher configuration identified by id.
func (r *ICallPublisherResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ICallPublisherEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
