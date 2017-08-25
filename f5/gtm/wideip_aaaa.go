// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideipAAAAConfigList holds a list of WideipAAAA configuration.
type WideipAAAAConfigList struct {
	Items    []WideipAAAAConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// WideipAAAAConfig holds the configuration of a single WideipAAAA.
type WideipAAAAConfig struct {
}

// WideipAAAAEndpoint represents the REST resource for managing WideipAAAA.
const WideipAAAAEndpoint = "/wideip/aaaa"

// WideipAAAAResource provides an API to manage WideipAAAA configurations.
type WideipAAAAResource struct {
	c *f5.Client
}

// ListAll  lists all the WideipAAAA configurations.
func (r *WideipAAAAResource) ListAll() (*WideipAAAAConfigList, error) {
	var list WideipAAAAConfigList
	if err := r.c.ReadQuery(BasePath+WideipAAAAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideipAAAA configuration identified by id.
func (r *WideipAAAAResource) Get(id string) (*WideipAAAAConfig, error) {
	var item WideipAAAAConfig
	if err := r.c.ReadQuery(BasePath+WideipAAAAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideipAAAA configuration.
func (r *WideipAAAAResource) Create(item WideipAAAAConfig) error {
	if err := r.c.ModQuery("POST", BasePath+WideipAAAAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideipAAAA configuration identified by id.
func (r *WideipAAAAResource) Edit(id string, item WideipAAAAConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+WideipAAAAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideipAAAA configuration identified by id.
func (r *WideipAAAAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideipAAAAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
