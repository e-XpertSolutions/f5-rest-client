// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GlobalSettingsMetricsConfigList holds a list of GlobalSettingsMetrics configuration.
type GlobalSettingsMetricsConfigList struct {
	Items    []GlobalSettingsMetricsConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// GlobalSettingsMetricsConfig holds the configuration of a single GlobalSettingsMetrics.
type GlobalSettingsMetricsConfig struct {
}

// GlobalSettingsMetricsEndpoint represents the REST resource for managing GlobalSettingsMetrics.
const GlobalSettingsMetricsEndpoint = "/global-settings/metrics"

// GlobalSettingsMetricsResource provides an API to manage GlobalSettingsMetrics configurations.
type GlobalSettingsMetricsResource struct {
	c f5.Client
}

// ListAll  lists all the GlobalSettingsMetrics configurations.
func (r *GlobalSettingsMetricsResource) ListAll() (*GlobalSettingsMetricsConfigList, error) {
	var list GlobalSettingsMetricsConfigList
	if err := r.c.ReadQuery(BasePath+GlobalSettingsMetricsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single GlobalSettingsMetrics configuration identified by id.
func (r *GlobalSettingsMetricsResource) Get(id string) (*GlobalSettingsMetricsConfig, error) {
	var item GlobalSettingsMetricsConfig
	if err := r.c.ReadQuery(BasePath+GlobalSettingsMetricsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new GlobalSettingsMetrics configuration.
func (r *GlobalSettingsMetricsResource) Create(item GlobalSettingsMetricsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+GlobalSettingsMetricsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a GlobalSettingsMetrics configuration identified by id.
func (r *GlobalSettingsMetricsResource) Edit(id string, item GlobalSettingsMetricsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+GlobalSettingsMetricsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single GlobalSettingsMetrics configuration identified by id.
func (r *GlobalSettingsMetricsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+GlobalSettingsMetricsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
