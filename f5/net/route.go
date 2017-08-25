// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A RouteList holds a list of Route.
type RouteList struct {
	Items    []Route `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selfLink,omitempty"`
}

// A Route hold the configuration for a route.
type Route struct {
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	Gw         string `json:"gw,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Mtu        int    `json:"mtu,omitempty"`
	Name       string `json:"name,omitempty"`
	Network    string `json:"network,omitempty"`
	Partition  string `json:"partition,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
}

// RouteEndpoint represents the REST resource for managing a route.
const RouteEndpoint = "/route"

// A RouteResource provides API to manage routes configuration.
type RouteResource struct {
	c *f5.Client
}

// ListAll lists all the route configurations.
func (pr *RouteResource) ListAll() (*RouteList, error) {
	var list RouteList
	if err := pr.c.ReadQuery(BasePath+RouteEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single route configuration identified by id.
func (pr *RouteResource) Get(id string) (*Route, error) {
	var item Route
	if err := pr.c.ReadQuery(BasePath+RouteEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new route configuration.
func (pr *RouteResource) Create(item Route) error {
	if err := pr.c.ModQuery("POST", BasePath+RouteEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a route configuration identified by id.
func (pr *RouteResource) Edit(id string, item Route) error {
	if err := pr.c.ModQuery("PUT", BasePath+RouteEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single route configuration identified by id.
func (pr *RouteResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+RouteEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
