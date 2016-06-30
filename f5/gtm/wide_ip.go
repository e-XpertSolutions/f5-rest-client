// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideIPConfigList holds a list of WideIP configuration.
type WideIPConfigList struct {
	Items    []WideIPConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// WideIPConfig holds the configuration of a single WideIP.
type WideIPConfig struct {
}

// WideIPEndpoint represents the REST resource for managing WideIP.
const WideIPEndpoint = "/wideip"

// WideIPResource provides an API to manage WideIP configurations.
type WideIPResource struct {
	c f5.Client
}

// ListAll  lists all the WideIP configurations.
func (r *WideIPResource) ListAll() (*WideIPConfigList, error) {
	var list WideIPConfigList
	if err := r.c.ReadQuery(BasePath+WideIPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideIP configuration identified by id.
func (r *WideIPResource) Get(id string) (*WideIPConfig, error) {
	var item WideIPConfig
	if err := r.c.ReadQuery(BasePath+WideIPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideIP configuration.
func (r *WideIPResource) Create(item WideIPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+WideIPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideIP configuration identified by id.
func (r *WideIPResource) Edit(id string, item WideIPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+WideIPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideIP configuration identified by id.
func (r *WideIPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideIPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
