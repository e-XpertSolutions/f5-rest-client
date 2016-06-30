// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideipAConfigList holds a list of WideipA configuration.
type WideipAConfigList struct {
	Items    []WideipAConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// WideipAConfig holds the configuration of a single WideipA.
type WideipAConfig struct {
}

// WideipAEndpoint represents the REST resource for managing WideipA.
const WideipAEndpoint = "/wideip/a"

// WideipAResource provides an API to manage WideipA configurations.
type WideipAResource struct {
	c f5.Client
}

// ListAll  lists all the WideipA configurations.
func (r *WideipAResource) ListAll() (*WideipAConfigList, error) {
	var list WideipAConfigList
	if err := r.c.ReadQuery(BasePath+WideipAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideipA configuration identified by id.
func (r *WideipAResource) Get(id string) (*WideipAConfig, error) {
	var item WideipAConfig
	if err := r.c.ReadQuery(BasePath+WideipAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideipA configuration.
func (r *WideipAResource) Create(item WideipAConfig) error {
	if err := r.c.ModQuery("POST", BasePath+WideipAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideipA configuration identified by id.
func (r *WideipAResource) Edit(id string, item WideipAConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+WideipAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideipA configuration identified by id.
func (r *WideipAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideipAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
