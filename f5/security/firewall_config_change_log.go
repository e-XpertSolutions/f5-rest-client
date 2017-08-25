// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallConfigChangeLogConfigList holds a list of FirewallConfigChangeLog configuration.
type FirewallConfigChangeLogConfigList struct {
	Items    []FirewallConfigChangeLogConfig `json:"items"`
	Kind     string                          `json:"kind"`
	SelfLink string                          `json:"selflink"`
}

// FirewallConfigChangeLogConfig holds the configuration of a single FirewallConfigChangeLog.
type FirewallConfigChangeLogConfig struct {
}

// FirewallConfigChangeLogEndpoint represents the REST resource for managing FirewallConfigChangeLog.
const FirewallConfigChangeLogEndpoint = "/firewall/config-change-log"

// FirewallConfigChangeLogResource provides an API to manage FirewallConfigChangeLog configurations.
type FirewallConfigChangeLogResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallConfigChangeLog configurations.
func (r *FirewallConfigChangeLogResource) ListAll() (*FirewallConfigChangeLogConfigList, error) {
	var list FirewallConfigChangeLogConfigList
	if err := r.c.ReadQuery(BasePath+FirewallConfigChangeLogEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallConfigChangeLog configuration identified by id.
func (r *FirewallConfigChangeLogResource) Get(id string) (*FirewallConfigChangeLogConfig, error) {
	var item FirewallConfigChangeLogConfig
	if err := r.c.ReadQuery(BasePath+FirewallConfigChangeLogEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallConfigChangeLog configuration.
func (r *FirewallConfigChangeLogResource) Create(item FirewallConfigChangeLogConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallConfigChangeLogEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallConfigChangeLog configuration identified by id.
func (r *FirewallConfigChangeLogResource) Edit(id string, item FirewallConfigChangeLogConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallConfigChangeLogEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallConfigChangeLog configuration identified by id.
func (r *FirewallConfigChangeLogResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallConfigChangeLogEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
