// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorRadiusAccountingConfigList holds a list of MonitorRadiusAccounting configuration.
type MonitorRadiusAccountingConfigList struct {
	Items    []MonitorRadiusAccountingConfig `json:"items"`
	Kind     string                          `json:"kind"`
	SelfLink string                          `json:"selflink"`
}

// MonitorRadiusAccountingConfig holds the configuration of a single MonitorRadiusAccounting.
type MonitorRadiusAccountingConfig struct {
}

// MonitorRadiusAccountingEndpoint represents the REST resource for managing MonitorRadiusAccounting.
const MonitorRadiusAccountingEndpoint = "/monitor/radius-accounting"

// MonitorRadiusAccountingResource provides an API to manage MonitorRadiusAccounting configurations.
type MonitorRadiusAccountingResource struct {
	c f5.Client
}

// ListAll  lists all the MonitorRadiusAccounting configurations.
func (r *MonitorRadiusAccountingResource) ListAll() (*MonitorRadiusAccountingConfigList, error) {
	var list MonitorRadiusAccountingConfigList
	if err := r.c.ReadQuery(BasePath+MonitorRadiusAccountingEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorRadiusAccounting configuration identified by id.
func (r *MonitorRadiusAccountingResource) Get(id string) (*MonitorRadiusAccountingConfig, error) {
	var item MonitorRadiusAccountingConfig
	if err := r.c.ReadQuery(BasePath+MonitorRadiusAccountingEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorRadiusAccounting configuration.
func (r *MonitorRadiusAccountingResource) Create(item MonitorRadiusAccountingConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorRadiusAccountingEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorRadiusAccounting configuration identified by id.
func (r *MonitorRadiusAccountingResource) Edit(id string, item MonitorRadiusAccountingConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorRadiusAccountingEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorRadiusAccounting configuration identified by id.
func (r *MonitorRadiusAccountingResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorRadiusAccountingEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
