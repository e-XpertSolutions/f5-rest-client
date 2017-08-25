// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideipMxConfigList holds a list of WideipMx configuration.
type WideipMxConfigList struct {
	Items    []WideipMxConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// WideipMxConfig holds the configuration of a single WideipMx.
type WideipMxConfig struct {
}

// WideipMxEndpoint represents the REST resource for managing WideipMx.
const WideipMxEndpoint = "/wideip/mx"

// WideipMxResource provides an API to manage WideipMx configurations.
type WideipMxResource struct {
	c *f5.Client
}

// ListAll  lists all the WideipMx configurations.
func (r *WideipMxResource) ListAll() (*WideipMxConfigList, error) {
	var list WideipMxConfigList
	if err := r.c.ReadQuery(BasePath+WideipMxEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideipMx configuration identified by id.
func (r *WideipMxResource) Get(id string) (*WideipMxConfig, error) {
	var item WideipMxConfig
	if err := r.c.ReadQuery(BasePath+WideipMxEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideipMx configuration.
func (r *WideipMxResource) Create(item WideipMxConfig) error {
	if err := r.c.ModQuery("POST", BasePath+WideipMxEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideipMx configuration identified by id.
func (r *WideipMxResource) Edit(id string, item WideipMxConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+WideipMxEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideipMx configuration identified by id.
func (r *WideipMxResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideipMxEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
