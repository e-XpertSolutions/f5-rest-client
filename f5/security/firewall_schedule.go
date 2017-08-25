// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FirewallScheduleConfigList holds a list of FirewallSchedule configuration.
type FirewallScheduleConfigList struct {
	Items    []FirewallScheduleConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// FirewallScheduleConfig holds the configuration of a single FirewallSchedule.
type FirewallScheduleConfig struct {
}

// FirewallScheduleEndpoint represents the REST resource for managing FirewallSchedule.
const FirewallScheduleEndpoint = "/firewall/schedule"

// FirewallScheduleResource provides an API to manage FirewallSchedule configurations.
type FirewallScheduleResource struct {
	c *f5.Client
}

// ListAll  lists all the FirewallSchedule configurations.
func (r *FirewallScheduleResource) ListAll() (*FirewallScheduleConfigList, error) {
	var list FirewallScheduleConfigList
	if err := r.c.ReadQuery(BasePath+FirewallScheduleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FirewallSchedule configuration identified by id.
func (r *FirewallScheduleResource) Get(id string) (*FirewallScheduleConfig, error) {
	var item FirewallScheduleConfig
	if err := r.c.ReadQuery(BasePath+FirewallScheduleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FirewallSchedule configuration.
func (r *FirewallScheduleResource) Create(item FirewallScheduleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FirewallScheduleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FirewallSchedule configuration identified by id.
func (r *FirewallScheduleResource) Edit(id string, item FirewallScheduleConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FirewallScheduleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FirewallSchedule configuration identified by id.
func (r *FirewallScheduleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FirewallScheduleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
