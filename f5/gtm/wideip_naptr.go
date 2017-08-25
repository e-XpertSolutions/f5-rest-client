// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideipNaptrConfigList holds a list of WideipNaptr configuration.
type WideipNaptrConfigList struct {
	Items    []WideipNaptrConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// WideipNaptrConfig holds the configuration of a single WideipNaptr.
type WideipNaptrConfig struct {
}

// WideipNaptrEndpoint represents the REST resource for managing WideipNaptr.
const WideipNaptrEndpoint = "/wideip/naptr"

// WideipNaptrResource provides an API to manage WideipNaptr configurations.
type WideipNaptrResource struct {
	c *f5.Client
}

// ListAll  lists all the WideipNaptr configurations.
func (r *WideipNaptrResource) ListAll() (*WideipNaptrConfigList, error) {
	var list WideipNaptrConfigList
	if err := r.c.ReadQuery(BasePath+WideipNaptrEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideipNaptr configuration identified by id.
func (r *WideipNaptrResource) Get(id string) (*WideipNaptrConfig, error) {
	var item WideipNaptrConfig
	if err := r.c.ReadQuery(BasePath+WideipNaptrEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideipNaptr configuration.
func (r *WideipNaptrResource) Create(item WideipNaptrConfig) error {
	if err := r.c.ModQuery("POST", BasePath+WideipNaptrEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideipNaptr configuration identified by id.
func (r *WideipNaptrResource) Edit(id string, item WideipNaptrConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+WideipNaptrEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideipNaptr configuration identified by id.
func (r *WideipNaptrResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideipNaptrEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
