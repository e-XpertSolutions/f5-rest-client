// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ListenerProfilesList holds a list of ListenerProfiles uration.
type ListenerProfilesList struct {
	Items    []ListenerProfiles `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// ListenerProfiles holds the uration of a single ListenerProfiles.
type ListenerProfiles struct {
}

// ListenerProfilesEndpoint represents the REST resource for managing ListenerProfiles.
const ListenerProfilesEndpoint = "/listener_profiles"

// ListenerProfilesResource provides an API to manage ListenerProfiles urations.
type ListenerProfilesResource struct {
	c *f5.Client
}

// ListAll  lists all the ListenerProfiles urations.
func (r *ListenerProfilesResource) ListAll() (*ListenerProfilesList, error) {
	var list ListenerProfilesList
	if err := r.c.ReadQuery(BasePath+ListenerProfilesEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ListenerProfiles uration identified by id.
func (r *ListenerProfilesResource) Get(id string) (*ListenerProfiles, error) {
	var item ListenerProfiles
	if err := r.c.ReadQuery(BasePath+ListenerProfilesEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ListenerProfiles uration.
func (r *ListenerProfilesResource) Create(item ListenerProfiles) error {
	if err := r.c.ModQuery("POST", BasePath+ListenerProfilesEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ListenerProfiles uration identified by id.
func (r *ListenerProfilesResource) Edit(id string, item ListenerProfiles) error {
	if err := r.c.ModQuery("PUT", BasePath+ListenerProfilesEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ListenerProfiles uration identified by id.
func (r *ListenerProfilesResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ListenerProfilesEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
