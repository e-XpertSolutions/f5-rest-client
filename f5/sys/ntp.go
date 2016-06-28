// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "e-xpert_solutions/f5-rest-client/f5"

// NTPConfigList holds a list of NTP configuration.
type NTPConfigList struct {
	Items    []NTPConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// NTPConfig holds the configuration of a single NTP.
type NTPConfig struct {
}

// NTPEndpoint represents the REST resource for managing NTP.
const NTPEndpoint = "/tm/sys/ntp"

// NTPResource provides an API to manage NTP configurations.
type NTPResource struct {
	c f5.Client
}

// ListAll  lists all the NTP configurations.
func (r *NTPResource) ListAll() (*NTPConfigList, error) {
	var list NTPConfigList
	if err := r.c.ReadQuery(BasePath+NTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single NTP configuration identified by id.
func (r *NTPResource) Get(id string) (*NTPConfig, error) {
	var item NTPConfig
	if err := r.c.ReadQuery(BasePath+NTPEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new NTP configuration.
func (r *NTPResource) Create(item NTPConfig) error {
	if err := r.c.ModQuery("POST", BasePath+NTPEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a NTP configuration identified by id.
func (r *NTPResource) Edit(id string, item NTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+NTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single NTP configuration identified by id.
func (r *NTPResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+NTPEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
