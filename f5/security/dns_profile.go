// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DNSProfileConfigList holds a list of DNSProfile configuration.
type DNSProfileConfigList struct {
	Items    []DNSProfileConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// DNSProfileConfig holds the configuration of a single DNSProfile.
type DNSProfileConfig struct {
}

// DNSProfileEndpoint represents the REST resource for managing DNSProfile.
const DNSProfileEndpoint = "/dns/profile"

// DNSProfileResource provides an API to manage DNSProfile configurations.
type DNSProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the DNSProfile configurations.
func (r *DNSProfileResource) ListAll() (*DNSProfileConfigList, error) {
	var list DNSProfileConfigList
	if err := r.c.ReadQuery(BasePath+DNSProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DNSProfile configuration identified by id.
func (r *DNSProfileResource) Get(id string) (*DNSProfileConfig, error) {
	var item DNSProfileConfig
	if err := r.c.ReadQuery(BasePath+DNSProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DNSProfile configuration.
func (r *DNSProfileResource) Create(item DNSProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DNSProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DNSProfile configuration identified by id.
func (r *DNSProfileResource) Edit(id string, item DNSProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DNSProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DNSProfile configuration identified by id.
func (r *DNSProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DNSProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
