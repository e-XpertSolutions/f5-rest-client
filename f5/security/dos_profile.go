// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSProfileConfigList holds a list of DOSProfile configuration.
type DOSProfileConfigList struct {
	Items    []DOSProfileConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// DOSProfileConfig holds the configuration of a single DOSProfile.
type DOSProfileConfig struct {
	ApplicationReference struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"applicationReference"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	SelfLink   string `json:"selfLink"`
}

// DOSProfileEndpoint represents the REST resource for managing DOSProfile.
const DOSProfileEndpoint = "/dos/profile"

// DOSProfileResource provides an API to manage DOSProfile configurations.
type DOSProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSProfile configurations.
func (r *DOSProfileResource) ListAll() (*DOSProfileConfigList, error) {
	var list DOSProfileConfigList
	if err := r.c.ReadQuery(BasePath+DOSProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSProfile configuration identified by id.
func (r *DOSProfileResource) Get(id string) (*DOSProfileConfig, error) {
	var item DOSProfileConfig
	if err := r.c.ReadQuery(BasePath+DOSProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSProfile configuration.
func (r *DOSProfileResource) Create(item DOSProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSProfile configuration identified by id.
func (r *DOSProfileResource) Edit(id string, item DOSProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSProfile configuration identified by id.
func (r *DOSProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
