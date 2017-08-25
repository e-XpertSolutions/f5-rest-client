// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A RouteDomainList holds a list of RouteDomain.
type RouteDomainList struct {
	Items    []RouteDomain `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selfLink,omitempty"`
}

// A RouteDomain hold the configuration for a route domain.
type RouteDomain struct {
	ConnectionLimit int      `json:"connectionLimit,omitempty"`
	FullPath        string   `json:"fullPath,omitempty"`
	Generation      int      `json:"generation,omitempty"`
	ID              int      `json:"id,omitempty"`
	Kind            string   `json:"kind,omitempty"`
	Name            string   `json:"name,omitempty"`
	SelfLink        string   `json:"selfLink,omitempty"`
	Strict          string   `json:"strict,omitempty"`
	Vlans           []string `json:"vlans,omitempty"`
	VlansReference  []struct {
		Link string `json:"link,omitempty"`
	} `json:"vlansReference,omitempty"`
}

// RouteDomainEndpoint represents the REST resource for managing a route domain.
const RouteDomainEndpoint = "/route-domain"

// A RouteDomainResource provides API to manage route domain configuration.
type RouteDomainResource struct {
	c *f5.Client
}

// ListAll lists all the route domain configurations.
func (pr *RouteDomainResource) ListAll() (*RouteDomainList, error) {
	var list RouteDomainList
	if err := pr.c.ReadQuery(BasePath+RouteDomainEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single route domain configuration identified by id.
func (pr *RouteDomainResource) Get(id string) (*RouteDomain, error) {
	var item RouteDomain
	if err := pr.c.ReadQuery(BasePath+RouteDomainEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new route domain configuration.
func (pr *RouteDomainResource) Create(item RouteDomain) error {
	if err := pr.c.ModQuery("POST", BasePath+RouteDomainEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a route domain configuration identified by id.
func (pr *RouteDomainResource) Edit(id string, item RouteDomain) error {
	if err := pr.c.ModQuery("PUT", BasePath+RouteDomainEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single route domain configuration identified by id.
func (pr *RouteDomainResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+RouteDomainEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
