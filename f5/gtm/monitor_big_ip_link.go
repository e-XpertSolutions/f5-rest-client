// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorBigIPLinkConfigList holds a list of MonitorBigIPLink configuration.
type MonitorBigIPLinkConfigList struct {
	Items    []MonitorBigIPLinkConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// MonitorBigIPLinkConfig holds the configuration of a single MonitorBigIPLink.
type MonitorBigIPLinkConfig struct {
}

// MonitorBigIPLinkEndpoint represents the REST resource for managing MonitorBigIPLink.
const MonitorBigIPLinkEndpoint = "/monitor/bigip-link"

// MonitorBigIPLinkResource provides an API to manage MonitorBigIPLink configurations.
type MonitorBigIPLinkResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorBigIPLink configurations.
func (r *MonitorBigIPLinkResource) ListAll() (*MonitorBigIPLinkConfigList, error) {
	var list MonitorBigIPLinkConfigList
	if err := r.c.ReadQuery(BasePath+MonitorBigIPLinkEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorBigIPLink configuration identified by id.
func (r *MonitorBigIPLinkResource) Get(id string) (*MonitorBigIPLinkConfig, error) {
	var item MonitorBigIPLinkConfig
	if err := r.c.ReadQuery(BasePath+MonitorBigIPLinkEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorBigIPLink configuration.
func (r *MonitorBigIPLinkResource) Create(item MonitorBigIPLinkConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorBigIPLinkEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorBigIPLink configuration identified by id.
func (r *MonitorBigIPLinkResource) Edit(id string, item MonitorBigIPLinkConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorBigIPLinkEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorBigIPLink configuration identified by id.
func (r *MonitorBigIPLinkResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorBigIPLinkEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
