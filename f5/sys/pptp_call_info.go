// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PPTPCallInfoConfigList holds a list of PPTPCallInfo configuration.
type PPTPCallInfoConfigList struct {
	Items    []PPTPCallInfoConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// PPTPCallInfoConfig holds the configuration of a single PPTPCallInfo.
type PPTPCallInfoConfig struct {
}

// PPTPCallInfoEndpoint represents the REST resource for managing PPTPCallInfo.
const PPTPCallInfoEndpoint = "/pptp-call-info"

// PPTPCallInfoResource provides an API to manage PPTPCallInfo configurations.
type PPTPCallInfoResource struct {
	c *f5.Client
}

// ListAll  lists all the PPTPCallInfo configurations.
func (r *PPTPCallInfoResource) ListAll() (*PPTPCallInfoConfigList, error) {
	var list PPTPCallInfoConfigList
	if err := r.c.ReadQuery(BasePath+PPTPCallInfoEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PPTPCallInfo configuration identified by id.
func (r *PPTPCallInfoResource) Get(id string) (*PPTPCallInfoConfig, error) {
	var item PPTPCallInfoConfig
	if err := r.c.ReadQuery(BasePath+PPTPCallInfoEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new PPTPCallInfo configuration.
func (r *PPTPCallInfoResource) Create(item PPTPCallInfoConfig) error {
	if err := r.c.ModQuery("POST", BasePath+PPTPCallInfoEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PPTPCallInfo configuration identified by id.
func (r *PPTPCallInfoResource) Edit(id string, item PPTPCallInfoConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+PPTPCallInfoEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PPTPCallInfo configuration identified by id.
func (r *PPTPCallInfoResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PPTPCallInfoEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
