// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LogProfileConfigList holds a list of LogProfile configuration.
type LogProfileConfigList struct {
	Items    []LogProfileConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// LogProfileConfig holds the configuration of a single LogProfile.
type LogProfileConfig struct {
	ApplicationReference struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"applicationReference"`
	Description string `json:"description"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	SelfLink    string `json:"selfLink"`
}

// LogProfileEndpoint represents the REST resource for managing LogProfile.
const LogProfileEndpoint = "/log/profile"

// LogProfileResource provides an API to manage LogProfile configurations.
type LogProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the LogProfile configurations.
func (r *LogProfileResource) ListAll() (*LogProfileConfigList, error) {
	var list LogProfileConfigList
	if err := r.c.ReadQuery(BasePath+LogProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single LogProfile configuration identified by id.
func (r *LogProfileResource) Get(id string) (*LogProfileConfig, error) {
	var item LogProfileConfig
	if err := r.c.ReadQuery(BasePath+LogProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new LogProfile configuration.
func (r *LogProfileResource) Create(item LogProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+LogProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a LogProfile configuration identified by id.
func (r *LogProfileResource) Edit(id string, item LogProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+LogProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single LogProfile configuration identified by id.
func (r *LogProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LogProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
