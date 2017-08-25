// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ICallConfigList holds a list of ICall configuration.
type ICallConfigList struct {
	Items    []ICallConfig `json:"items"`
	Kind     string        `json:"kind"`
	SelfLink string        `json:"selflink"`
}

// ICallConfig holds the configuration of a single ICall.
type ICallConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// ICallEndpoint represents the REST resource for managing ICall.
const ICallEndpoint = "/icall"

// ICallResource provides an API to manage ICall configurations.
type ICallResource struct {
	c *f5.Client
}

// ListAll  lists all the ICall configurations.
func (r *ICallResource) ListAll() (*ICallConfigList, error) {
	var list ICallConfigList
	if err := r.c.ReadQuery(BasePath+ICallEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ICall configuration identified by id.
func (r *ICallResource) Get(id string) (*ICallConfig, error) {
	var item ICallConfig
	if err := r.c.ReadQuery(BasePath+ICallEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ICall configuration.
func (r *ICallResource) Create(item ICallConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ICallEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ICall configuration identified by id.
func (r *ICallResource) Edit(id string, item ICallConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ICallEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ICall configuration identified by id.
func (r *ICallResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ICallEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
