// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IControlSOAPConfigList holds a list of IControlSOAP configuration.
type IControlSOAPConfigList struct {
	Items    []IControlSOAPConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// IControlSOAPConfig holds the configuration of a single IControlSOAP.
type IControlSOAPConfig struct {
}

// IControlSOAPEndpoint represents the REST resource for managing IControlSOAP.
const IControlSOAPEndpoint = "/icontrol-soap"

// IControlSOAPResource provides an API to manage IControlSOAP configurations.
type IControlSOAPResource struct {
	c *f5.Client
}

// ListAll  lists all the IControlSOAP configurations.
func (r *IControlSOAPResource) ListAll() (*IControlSOAPConfigList, error) {
	var list IControlSOAPConfigList
	if err := r.c.ReadQuery(BasePath+IControlSOAPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IControlSOAP configuration identified by id.
func (r *IControlSOAPResource) Get(id string) (*IControlSOAPConfig, error) {
	var item IControlSOAPConfig
	if err := r.c.ReadQuery(BasePath+IControlSOAPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IControlSOAP configuration.
func (r *IControlSOAPResource) Create(item IControlSOAPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IControlSOAPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IControlSOAP configuration identified by id.
func (r *IControlSOAPResource) Edit(id string, item IControlSOAPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IControlSOAPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IControlSOAP configuration identified by id.
func (r *IControlSOAPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IControlSOAPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
