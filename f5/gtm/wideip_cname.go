// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideipCnameConfigList holds a list of WideipCname configuration.
type WideipCnameConfigList struct {
	Items    []WideipCnameConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// WideipCnameConfig holds the configuration of a single WideipCname.
type WideipCnameConfig struct {
}

// WideipCnameEndpoint represents the REST resource for managing WideipCname.
const WideipCnameEndpoint = "/wideip/cname"

// WideipCnameResource provides an API to manage WideipCname configurations.
type WideipCnameResource struct {
	c f5.Client
}

// ListAll  lists all the WideipCname configurations.
func (r *WideipCnameResource) ListAll() (*WideipCnameConfigList, error) {
	var list WideipCnameConfigList
	if err := r.c.ReadQuery(BasePath+WideipCnameEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideipCname configuration identified by id.
func (r *WideipCnameResource) Get(id string) (*WideipCnameConfig, error) {
	var item WideipCnameConfig
	if err := r.c.ReadQuery(BasePath+WideipCnameEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideipCname configuration.
func (r *WideipCnameResource) Create(item WideipCnameConfig) error {
	if err := r.c.ModQuery("POST", BasePath+WideipCnameEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideipCname configuration identified by id.
func (r *WideipCnameResource) Edit(id string, item WideipCnameConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+WideipCnameEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideipCname configuration identified by id.
func (r *WideipCnameResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideipCnameEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
