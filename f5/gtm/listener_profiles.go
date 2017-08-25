// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ListenerProfilesConfigList holds a list of ListenerProfiles configuration.
type ListenerProfilesConfigList struct {
	Items    []ListenerProfilesConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// ListenerProfilesConfig holds the configuration of a single ListenerProfiles.
type ListenerProfilesConfig struct {
}

// ListenerProfilesEndpoint represents the REST resource for managing ListenerProfiles.
const ListenerProfilesEndpoint = "/listener_profiles"

// ListenerProfilesResource provides an API to manage ListenerProfiles configurations.
type ListenerProfilesResource struct {
	c *f5.Client
}

// ListAll  lists all the ListenerProfiles configurations.
func (r *ListenerProfilesResource) ListAll() (*ListenerProfilesConfigList, error) {
	var list ListenerProfilesConfigList
	if err := r.c.ReadQuery(BasePath+ListenerProfilesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ListenerProfiles configuration identified by id.
func (r *ListenerProfilesResource) Get(id string) (*ListenerProfilesConfig, error) {
	var item ListenerProfilesConfig
	if err := r.c.ReadQuery(BasePath+ListenerProfilesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ListenerProfiles configuration.
func (r *ListenerProfilesResource) Create(item ListenerProfilesConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ListenerProfilesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ListenerProfiles configuration identified by id.
func (r *ListenerProfilesResource) Edit(id string, item ListenerProfilesConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ListenerProfilesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ListenerProfiles configuration identified by id.
func (r *ListenerProfilesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ListenerProfilesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
