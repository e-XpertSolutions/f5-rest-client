// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallPortListConfigList holds a list of FirewallPortList configuration.
type FirewallPortListConfigList struct {
	Items    []FirewallPortListConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// FirewallPortListConfig holds the configuration of a single FirewallPortList.
type FirewallPortListConfig struct {
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	Ports      []struct {
		Name string `json:"name"`
	} `json:"ports"`
	SelfLink string `json:"selfLink"`
}

// FirewallPortListEndpoint represents the REST resource for managing FirewallPortList.
const FirewallPortListEndpoint = "/firewall/port-list"

// FirewallPortListResource provides an API to manage FirewallPortList configurations.
type FirewallPortListResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallPortList configurations.
func (r *FirewallPortListResource) ListAll() (*FirewallPortListConfigList, error) {
	var list FirewallPortListConfigList
	if err := r.c.ReadQuery(BasePath+FirewallPortListEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallPortList configuration identified by id.
func (r *FirewallPortListResource) Get(id string) (*FirewallPortListConfig, error) {
	var item FirewallPortListConfig
	if err := r.c.ReadQuery(BasePath+FirewallPortListEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallPortList configuration.
func (r *FirewallPortListResource) Create(item FirewallPortListConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallPortListEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallPortList configuration identified by id.
func (r *FirewallPortListResource) Edit(id string, item FirewallPortListConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallPortListEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallPortList configuration identified by id.
func (r *FirewallPortListResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallPortListEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
