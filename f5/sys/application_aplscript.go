// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ApplicationAPLScriptConfigList holds a list of ApplicationAPLScript configuration.
type ApplicationAPLScriptConfigList struct {
	Items    []ApplicationAPLScriptConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// ApplicationAPLScriptConfig holds the configuration of a single ApplicationAPLScript.
type ApplicationAPLScriptConfig struct {
	APIAnonymous       string `json:"apiAnonymous"`
	AplSignature       string `json:"aplSignature"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	IgnoreVerification string `json:"ignoreVerification"`
	Kind               string `json:"kind"`
	Name               string `json:"name"`
	Partition          string `json:"partition"`
	SelfLink           string `json:"selfLink"`
	TotalSigningStatus string `json:"totalSigningStatus"`
	VerificationStatus string `json:"verificationStatus"`
}

// ApplicationAPLScriptEndpoint represents the REST resource for managing ApplicationAPLScript.
const ApplicationAPLScriptEndpoint = "/application/apl-script"

// ApplicationAPLScriptResource provides an API to manage ApplicationAPLScript configurations.
type ApplicationAPLScriptResource struct {
	c *f5.Client
}

// ListAll  lists all the ApplicationAPLScript configurations.
func (r *ApplicationAPLScriptResource) ListAll() (*ApplicationAPLScriptConfigList, error) {
	var list ApplicationAPLScriptConfigList
	if err := r.c.ReadQuery(BasePath+ApplicationAPLScriptEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ApplicationAPLScript configuration identified by id.
func (r *ApplicationAPLScriptResource) Get(id string) (*ApplicationAPLScriptConfig, error) {
	var item ApplicationAPLScriptConfig
	if err := r.c.ReadQuery(BasePath+ApplicationAPLScriptEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ApplicationAPLScript configuration.
func (r *ApplicationAPLScriptResource) Create(item ApplicationAPLScriptConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ApplicationAPLScriptEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ApplicationAPLScript configuration identified by id.
func (r *ApplicationAPLScriptResource) Edit(id string, item ApplicationAPLScriptConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ApplicationAPLScriptEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ApplicationAPLScript configuration identified by id.
func (r *ApplicationAPLScriptResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ApplicationAPLScriptEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
