// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DNSConfigList holds a list of DNS configuration.
type DNSConfigList struct {
	Items    []DNSConfig `json:"items"`
	Kind     string      `json:"kind"`
	SelfLink string      `json:"selflink"`
}

// DNSConfig holds the configuration of a single DNS.
type DNSConfig struct {
}

// DNSEndpoint represents the REST resource for managing DNS.
const DNSEndpoint = "/dns"

// DNSResource provides an API to manage DNS configurations.
type DNSResource struct {
	c *f5.Client
}

// ListAll  lists all the DNS configurations.
func (r *DNSResource) ListAll() (*DNSConfigList, error) {
	var list DNSConfigList
	if err := r.c.ReadQuery(BasePath+DNSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DNS configuration identified by id.
func (r *DNSResource) Get(id string) (*DNSConfig, error) {
	var item DNSConfig
	if err := r.c.ReadQuery(BasePath+DNSEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DNS configuration.
func (r *DNSResource) Create(item DNSConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DNSEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DNS configuration identified by id.
func (r *DNSResource) Edit(id string, item DNSConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DNSEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DNS configuration identified by id.
func (r *DNSResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DNSEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
