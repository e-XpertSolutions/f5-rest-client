// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ApplicationTemplateConfigList holds a list of ApplicationTemplate configuration.
type ApplicationTemplateConfigList struct {
	Items    []ApplicationTemplateConfig `json:"items"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

// ApplicationTemplateConfig holds the configuration of a single ApplicationTemplate.
type ApplicationTemplateConfig struct {
	ActionsReference struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"actionsReference"`
	FullPath                string `json:"fullPath"`
	Generation              int    `json:"generation"`
	IgnoreVerification      string `json:"ignoreVerification"`
	Kind                    string `json:"kind"`
	Name                    string `json:"name"`
	Partition               string `json:"partition"`
	RequiresBigipVersionMin string `json:"requiresBigipVersionMin"`
	SelfLink                string `json:"selfLink"`
	TmplSignature           string `json:"tmplSignature"`
	TotalSigningStatus      string `json:"totalSigningStatus"`
	VerificationStatus      string `json:"verificationStatus"`
}

// ApplicationTemplateEndpoint represents the REST resource for managing ApplicationTemplate.
const ApplicationTemplateEndpoint = "/application/template"

// ApplicationTemplateResource provides an API to manage ApplicationTemplate configurations.
type ApplicationTemplateResource struct {
	c *f5.Client
}

// ListAll  lists all the ApplicationTemplate configurations.
func (r *ApplicationTemplateResource) ListAll() (*ApplicationTemplateConfigList, error) {
	var list ApplicationTemplateConfigList
	if err := r.c.ReadQuery(BasePath+ApplicationTemplateEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ApplicationTemplate configuration identified by id.
func (r *ApplicationTemplateResource) Get(id string) (*ApplicationTemplateConfig, error) {
	var item ApplicationTemplateConfig
	if err := r.c.ReadQuery(BasePath+ApplicationTemplateEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ApplicationTemplate configuration.
func (r *ApplicationTemplateResource) Create(item ApplicationTemplateConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ApplicationTemplateEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ApplicationTemplate configuration identified by id.
func (r *ApplicationTemplateResource) Edit(id string, item ApplicationTemplateConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ApplicationTemplateEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ApplicationTemplate configuration identified by id.
func (r *ApplicationTemplateResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ApplicationTemplateEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
