// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "e-xpert_solutions/f5-rest-client/f5"

// A TrunkList holds a list of Trunks.
type TrunkList struct {
	Items    []Trunk `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selfLink,omitempty"`
}

// A Trunk hold the configuration for a trunk.
type Trunk struct {
	Bandwidth           int      `json:"bandwidth"`
	CfgMbrCount         int      `json:"cfgMbrCount"`
	DistributionHash    string   `json:"distributionHash"`
	FullPath            string   `json:"fullPath"`
	Generation          int      `json:"generation"`
	ID                  int      `json:"id"`
	Interfaces          []string `json:"interfaces"`
	InterfacesReference []struct {
		Link string `json:"link"`
	} `json:"interfacesReference"`
	Kind             string `json:"kind"`
	Lacp             string `json:"lacp"`
	LacpMode         string `json:"lacpMode"`
	LacpTimeout      string `json:"lacpTimeout"`
	LinkSelectPolicy string `json:"linkSelectPolicy"`
	MacAddress       string `json:"macAddress"`
	Media            string `json:"media"`
	Name             string `json:"name"`
	QinqEthertype    string `json:"qinqEthertype"`
	SelfLink         string `json:"selfLink"`
	Stp              string `json:"stp"`
	Type             string `json:"type"`
	WorkingMbrCount  int    `json:"workingMbrCount"`
}

// TrunkEndpoint represents the REST resource for managing a trunk.
const TrunkEndpoint = "/trunk"

// A TrunkResource provides API to manage trunks configuration.
type TrunkResource struct {
	c f5.Client
}

// ListAll lists all the trunk configurations.
func (pr *TrunkResource) ListAll() (*TrunkList, error) {
	var list TrunkList
	if err := pr.c.ReadQuery(BasePath+TrunkEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single trunk configuration identified by id.
func (pr *TrunkResource) Get(id string) (*Trunk, error) {
	var item Trunk
	if err := pr.c.ReadQuery(BasePath+TrunkEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new trunk configuration.
func (pr *TrunkResource) Create(item Trunk) error {
	if err := pr.c.ModQuery("POST", BasePath+TrunkEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a trunk configuration identified by id.
func (pr *TrunkResource) Edit(id string, item Trunk) error {
	if err := pr.c.ModQuery("PUT", BasePath+TrunkEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single trunk configuration identified by id.
func (pr *TrunkResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+TrunkEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
