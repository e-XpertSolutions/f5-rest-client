// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcmp

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// TrafficProfileConfigList holds a list of TrafficProfile configuration.
type TrafficProfileConfigList struct {
	Items    []TrafficProfileConfig `json:"items,omitempty"`
	Kind     string                 `json:"kind,omitempty"`
	SelfLink string                 `json:"selflink,omitempty"`
}

// TrafficProfileConfig holds the configuration of a single TrafficProfile.
type TrafficProfileConfig struct {
	Name         string `json:"name,omitempty"`
	AppService   string `json:"appService,omitempty"`
	ColorPolicer string `json:"colorPolicer,omitempty"`
}

// TrafficProfileEndpoint represents the REST resource for managing TrafficProfile.
const TrafficProfileEndpoint = "/traffic-profile"

// TrafficProfileResource provides an API to manage TrafficProfile configurations.
type TrafficProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the TrafficProfile configurations.
func (r *TrafficProfileResource) ListAll() (*TrafficProfileConfigList, error) {
	var list TrafficProfileConfigList
	if err := r.c.ReadQuery(BasePath+TrafficProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single TrafficProfile configuration identified by id.
func (r *TrafficProfileResource) Get(id string) (*TrafficProfileConfig, error) {
	var item TrafficProfileConfig
	if err := r.c.ReadQuery(BasePath+TrafficProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new TrafficProfile configuration.
func (r *TrafficProfileResource) Create(item TrafficProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+TrafficProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a TrafficProfile configuration identified by id.
func (r *TrafficProfileResource) Edit(id string, item TrafficProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+TrafficProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single TrafficProfile configuration identified by id.
func (r *TrafficProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+TrafficProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
