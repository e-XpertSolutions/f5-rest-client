// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ScriptdConfigList holds a list of Scriptd configuration.
type ScriptdConfigList struct {
	Items    []ScriptdConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// ScriptdConfig holds the configuration of a single Scriptd.
type ScriptdConfig struct {
}

// ScriptdEndpoint represents the REST resource for managing Scriptd.
const ScriptdEndpoint = "/scriptd"

// ScriptdResource provides an API to manage Scriptd configurations.
type ScriptdResource struct {
	c *f5.Client
}

// ListAll  lists all the Scriptd configurations.
func (r *ScriptdResource) ListAll() (*ScriptdConfigList, error) {
	var list ScriptdConfigList
	if err := r.c.ReadQuery(BasePath+ScriptdEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Scriptd configuration identified by id.
func (r *ScriptdResource) Get(id string) (*ScriptdConfig, error) {
	var item ScriptdConfig
	if err := r.c.ReadQuery(BasePath+ScriptdEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Scriptd configuration.
func (r *ScriptdResource) Create(item ScriptdConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ScriptdEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Scriptd configuration identified by id.
func (r *ScriptdResource) Edit(id string, item ScriptdConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ScriptdEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Scriptd configuration identified by id.
func (r *ScriptdResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ScriptdEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
