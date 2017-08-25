// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ICallScriptConfigList holds a list of ICallScript configuration.
type ICallScriptConfigList struct {
	Items    []ICallScriptConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// ICallScriptConfig holds the configuration of a single ICallScript.
type ICallScriptConfig struct {
}

// ICallScriptEndpoint represents the REST resource for managing ICallScript.
const ICallScriptEndpoint = "/icall/script"

// ICallScriptResource provides an API to manage ICallScript configurations.
type ICallScriptResource struct {
	c *f5.Client
}

// ListAll  lists all the ICallScript configurations.
func (r *ICallScriptResource) ListAll() (*ICallScriptConfigList, error) {
	var list ICallScriptConfigList
	if err := r.c.ReadQuery(BasePath+ICallScriptEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ICallScript configuration identified by id.
func (r *ICallScriptResource) Get(id string) (*ICallScriptConfig, error) {
	var item ICallScriptConfig
	if err := r.c.ReadQuery(BasePath+ICallScriptEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ICallScript configuration.
func (r *ICallScriptResource) Create(item ICallScriptConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ICallScriptEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ICallScript configuration identified by id.
func (r *ICallScriptResource) Edit(id string, item ICallScriptConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ICallScriptEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ICallScript configuration identified by id.
func (r *ICallScriptResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ICallScriptEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
