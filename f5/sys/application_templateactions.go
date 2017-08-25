// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ApplicationTemplateActionsConfigList holds a list of ApplicationTemplateActions configuration.
type ApplicationTemplateActionsConfigList struct {
	Items    []ApplicationTemplateActionsConfig `json:"items"`
	Kind     string                             `json:"kind"`
	SelfLink string                             `json:"selflink"`
}

// ApplicationTemplateActionsConfig holds the configuration of a single ApplicationTemplateActions.
type ApplicationTemplateActionsConfig struct {
}

// ApplicationTemplateActionsEndpoint represents the REST resource for managing ApplicationTemplateActions.
const ApplicationTemplateActionsEndpoint = "/application/template_actions"

// ApplicationTemplateActionsResource provides an API to manage ApplicationTemplateActions configurations.
type ApplicationTemplateActionsResource struct {
	c *f5.Client
}

// ListAll  lists all the ApplicationTemplateActions configurations.
func (r *ApplicationTemplateActionsResource) ListAll() (*ApplicationTemplateActionsConfigList, error) {
	var list ApplicationTemplateActionsConfigList
	if err := r.c.ReadQuery(BasePath+ApplicationTemplateActionsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ApplicationTemplateActions configuration identified by id.
func (r *ApplicationTemplateActionsResource) Get(id string) (*ApplicationTemplateActionsConfig, error) {
	var item ApplicationTemplateActionsConfig
	if err := r.c.ReadQuery(BasePath+ApplicationTemplateActionsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ApplicationTemplateActions configuration.
func (r *ApplicationTemplateActionsResource) Create(item ApplicationTemplateActionsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ApplicationTemplateActionsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ApplicationTemplateActions configuration identified by id.
func (r *ApplicationTemplateActionsResource) Edit(id string, item ApplicationTemplateActionsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ApplicationTemplateActionsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ApplicationTemplateActions configuration identified by id.
func (r *ApplicationTemplateActionsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ApplicationTemplateActionsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
