// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallAddressListConfigList holds a list of FirewallAddressList configuration.
type FirewallAddressListConfigList struct {
	Items    []FirewallAddressListConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// FirewallAddressListConfig holds the configuration of a single FirewallAddressList.
type FirewallAddressListConfig struct {
}

// FirewallAddressListEndpoint represents the REST resource for managing FirewallAddressList.
const FirewallAddressListEndpoint = "/firewall/address-list"

// FirewallAddressListResource provides an API to manage FirewallAddressList configurations.
type FirewallAddressListResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallAddressList configurations.
func (r *FirewallAddressListResource) ListAll() (*FirewallAddressListConfigList, error) {
	var list FirewallAddressListConfigList
	if err := r.c.ReadQuery(BasePath+FirewallAddressListEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallAddressList configuration identified by id.
func (r *FirewallAddressListResource) Get(id string) (*FirewallAddressListConfig, error) {
	var item FirewallAddressListConfig
	if err := r.c.ReadQuery(BasePath+FirewallAddressListEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallAddressList configuration.
func (r *FirewallAddressListResource) Create(item FirewallAddressListConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallAddressListEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallAddressList configuration identified by id.
func (r *FirewallAddressListResource) Edit(id string, item FirewallAddressListConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallAddressListEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallAddressList configuration identified by id.
func (r *FirewallAddressListResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallAddressListEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
