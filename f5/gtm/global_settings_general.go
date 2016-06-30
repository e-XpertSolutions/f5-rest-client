// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GlobalSettingsGeneralConfigList holds a list of GlobalSettingsGeneral configuration.
type GlobalSettingsGeneralConfigList struct {
	Items    []GlobalSettingsGeneralConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// GlobalSettingsGeneralConfig holds the configuration of a single GlobalSettingsGeneral.
type GlobalSettingsGeneralConfig struct {
}

// GlobalSettingsGeneralEndpoint represents the REST resource for managing GlobalSettingsGeneral.
const GlobalSettingsGeneralEndpoint = "/global-settings/general"

// GlobalSettingsGeneralResource provides an API to manage GlobalSettingsGeneral configurations.
type GlobalSettingsGeneralResource struct {
	c f5.Client
}

// ListAll  lists all the GlobalSettingsGeneral configurations.
func (r *GlobalSettingsGeneralResource) ListAll() (*GlobalSettingsGeneralConfigList, error) {
	var list GlobalSettingsGeneralConfigList
	if err := r.c.ReadQuery(BasePath+GlobalSettingsGeneralEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single GlobalSettingsGeneral configuration identified by id.
func (r *GlobalSettingsGeneralResource) Get(id string) (*GlobalSettingsGeneralConfig, error) {
	var item GlobalSettingsGeneralConfig
	if err := r.c.ReadQuery(BasePath+GlobalSettingsGeneralEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new GlobalSettingsGeneral configuration.
func (r *GlobalSettingsGeneralResource) Create(item GlobalSettingsGeneralConfig) error {
	if err := r.c.ModQuery("POST", BasePath+GlobalSettingsGeneralEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a GlobalSettingsGeneral configuration identified by id.
func (r *GlobalSettingsGeneralResource) Edit(id string, item GlobalSettingsGeneralConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+GlobalSettingsGeneralEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single GlobalSettingsGeneral configuration identified by id.
func (r *GlobalSettingsGeneralResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+GlobalSettingsGeneralEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
