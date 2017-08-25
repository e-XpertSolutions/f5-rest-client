// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A SelfList holds a list of Self.
type SelfList struct {
	Items    []Self `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// A Self hold the configuration for a Self IP.
type Self struct {
	Address               string `json:"address,omitempty"`
	AddressSource         string `json:"addressSource,omitempty"`
	Floating              string `json:"floating,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int    `json:"generation,omitempty"`
	InheritedTrafficGroup string `json:"inheritedTrafficGroup,omitempty"`
	Kind                  string `json:"kind,omitempty"`
	Name                  string `json:"name,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	TrafficGroup          string `json:"trafficGroup,omitempty"`
	TrafficGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"trafficGroupReference,omitempty"`
	Unit          int    `json:"unit,omitempty"`
	Vlan          string `json:"vlan,omitempty"`
	VlanReference struct {
		Link string `json:"link,omitempty"`
	} `json:"vlanReference,omitempty"`
}

// SelfEndpoint represents the REST resource for managing a self IP.
const SelfEndpoint = "/self"

// A SelfResource provides API to manage self ip configuration.
type SelfResource struct {
	c *f5.Client
}

// ListAll lists all the self ip configurations.
func (pr *SelfResource) ListAll() (*SelfList, error) {
	var list SelfList
	if err := pr.c.ReadQuery(BasePath+SelfEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single self ip configuration identified by id.
func (pr *SelfResource) Get(id string) (*Self, error) {
	var item Self
	if err := pr.c.ReadQuery(BasePath+SelfEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new self ip configuration.
func (pr *SelfResource) Create(item Self) error {
	if err := pr.c.ModQuery("POST", BasePath+SelfEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a self ip configuration identified by id.
func (pr *SelfResource) Edit(id string, item Self) error {
	if err := pr.c.ModQuery("PUT", BasePath+SelfEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single self ip configuration identified by id.
func (pr *SelfResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+SelfEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
