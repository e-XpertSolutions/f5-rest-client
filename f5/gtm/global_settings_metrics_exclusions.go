// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GlobalSettingsMetricsExclusionsConfigList holds a list of GlobalSettingsMetricsExclusions configuration.
type GlobalSettingsMetricsExclusionsConfigList struct {
	Items    []GlobalSettingsMetricsExclusionsConfig `json:"items"`
	Kind     string                                  `json:"kind"`
	SelfLink string                                  `json:"selflink"`
}

// GlobalSettingsMetricsExclusionsConfig holds the configuration of a single GlobalSettingsMetricsExclusions.
type GlobalSettingsMetricsExclusionsConfig struct {
}

// GlobalSettingsMetricsExclusionsEndpoint represents the REST resource for managing GlobalSettingsMetricsExclusions.
const GlobalSettingsMetricsExclusionsEndpoint = "/global-settings/metrics-exclusions"

// GlobalSettingsMetricsExclusionsResource provides an API to manage GlobalSettingsMetricsExclusions configurations.
type GlobalSettingsMetricsExclusionsResource struct {
	c *f5.Client
}

// ListAll  lists all the GlobalSettingsMetricsExclusions configurations.
func (r *GlobalSettingsMetricsExclusionsResource) ListAll() (*GlobalSettingsMetricsExclusionsConfigList, error) {
	var list GlobalSettingsMetricsExclusionsConfigList
	if err := r.c.ReadQuery(BasePath+GlobalSettingsMetricsExclusionsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single GlobalSettingsMetricsExclusions configuration identified by id.
func (r *GlobalSettingsMetricsExclusionsResource) Get(id string) (*GlobalSettingsMetricsExclusionsConfig, error) {
	var item GlobalSettingsMetricsExclusionsConfig
	if err := r.c.ReadQuery(BasePath+GlobalSettingsMetricsExclusionsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new GlobalSettingsMetricsExclusions configuration.
func (r *GlobalSettingsMetricsExclusionsResource) Create(item GlobalSettingsMetricsExclusionsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+GlobalSettingsMetricsExclusionsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a GlobalSettingsMetricsExclusions configuration identified by id.
func (r *GlobalSettingsMetricsExclusionsResource) Edit(id string, item GlobalSettingsMetricsExclusionsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+GlobalSettingsMetricsExclusionsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single GlobalSettingsMetricsExclusions configuration identified by id.
func (r *GlobalSettingsMetricsExclusionsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+GlobalSettingsMetricsExclusionsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
