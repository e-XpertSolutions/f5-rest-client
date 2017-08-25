// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A TrafficGroupList holds a list of TrafficGroup.
type TrafficGroupList struct {
	Items    []TrafficGroup `json:"items,omitempty"`
	Kind     string         `json:"kind,omitempty"`
	SelfLink string         `json:"selfLink,omitempty"`
}

// A TrafficGroup hold the configuration for a TrafficGroup.
type TrafficGroup struct {
	AutoFailbackEnabled string `json:"autoFailbackEnabled,omitempty"`
	AutoFailbackTime    int    `json:"autoFailbackTime,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	HaLoadFactor        int    `json:"haLoadFactor,omitempty"`
	IsFloating          string `json:"isFloating,omitempty"`
	Kind                string `json:"kind,omitempty"`
	Mac                 string `json:"mac,omitempty"`
	Name                string `json:"name,omitempty"`
	SelfLink            string `json:"selfLink,omitempty"`
	UnitID              int    `json:"unitId,omitempty"`
}

// TrafficGroupEndpoint represents the REST resource for managing a TrafficGroup.
const TrafficGroupEndpoint = "/traffic-group"

// A TrafficGroupResource provides API to manage TrafficGroups configuration.
type TrafficGroupResource struct {
	c *f5.Client
}

// ListAll lists all the TrafficGroup configurations.
func (pr *TrafficGroupResource) ListAll() (*TrafficGroupList, error) {
	var list TrafficGroupList
	if err := pr.c.ReadQuery(BasePath+TrafficGroupEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single TrafficGroup configuration identified by id.
func (pr *TrafficGroupResource) Get(id string) (*TrafficGroup, error) {
	var item TrafficGroup
	if err := pr.c.ReadQuery(BasePath+TrafficGroupEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new TrafficGroup configuration.
func (pr *TrafficGroupResource) Create(item TrafficGroup) error {
	if err := pr.c.ModQuery("POST", BasePath+TrafficGroupEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a TrafficGroup configuration identified by id.
func (pr *TrafficGroupResource) Edit(id string, item TrafficGroup) error {
	if err := pr.c.ModQuery("PUT", BasePath+TrafficGroupEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single TrafficGroup configuration identified by id.
func (pr *TrafficGroupResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+TrafficGroupEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
