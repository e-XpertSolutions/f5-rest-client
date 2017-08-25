// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// NTPRestrictConfigList holds a list of NTPRestrict configuration.
type NTPRestrictConfigList struct {
	Items    []NTPRestrictConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// NTPRestrictConfig holds the configuration of a single NTPRestrict.
type NTPRestrictConfig struct {
}

// NTPRestrictEndpoint represents the REST resource for managing NTPRestrict.
const NTPRestrictEndpoint = "/ntp_restrict"

// NTPRestrictResource provides an API to manage NTPRestrict configurations.
type NTPRestrictResource struct {
	c *f5.Client
}

// ListAll  lists all the NTPRestrict configurations.
func (r *NTPRestrictResource) ListAll() (*NTPRestrictConfigList, error) {
	var list NTPRestrictConfigList
	if err := r.c.ReadQuery(BasePath+NTPRestrictEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single NTPRestrict configuration identified by id.
func (r *NTPRestrictResource) Get(id string) (*NTPRestrictConfig, error) {
	var item NTPRestrictConfig
	if err := r.c.ReadQuery(BasePath+NTPRestrictEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new NTPRestrict configuration.
func (r *NTPRestrictResource) Create(item NTPRestrictConfig) error {
	if err := r.c.ModQuery("POST", BasePath+NTPRestrictEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a NTPRestrict configuration identified by id.
func (r *NTPRestrictResource) Edit(id string, item NTPRestrictConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+NTPRestrictEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single NTPRestrict configuration identified by id.
func (r *NTPRestrictResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+NTPRestrictEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
