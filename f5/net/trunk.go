// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A TrunkList holds a list of Trunks.
type TrunkList struct {
	Items    []Trunk `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selfLink,omitempty"`
}

// A Trunk hold the configuration for a trunk.
type Trunk struct {
	Bandwidth           int      `json:"bandwidth,omitempty"`
	CfgMbrCount         int      `json:"cfgMbrCount,omitempty"`
	DistributionHash    string   `json:"distributionHash,omitempty"`
	FullPath            string   `json:"fullPath,omitempty"`
	Generation          int      `json:"generation,omitempty"`
	ID                  int      `json:"id,omitempty"`
	Interfaces          []string `json:"interfaces,omitempty"`
	InterfacesReference []struct {
		Link string `json:"link,omitempty"`
	} `json:"interfacesReference,omitempty"`
	Kind             string `json:"kind,omitempty"`
	Lacp             string `json:"lacp,omitempty"`
	LacpMode         string `json:"lacpMode,omitempty"`
	LacpTimeout      string `json:"lacpTimeout,omitempty"`
	LinkSelectPolicy string `json:"linkSelectPolicy,omitempty"`
	MacAddress       string `json:"macAddress,omitempty"`
	Media            string `json:"media,omitempty"`
	Name             string `json:"name,omitempty"`
	QinqEthertype    string `json:"qinqEthertype,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
	Stp              string `json:"stp,omitempty"`
	Type             string `json:"type,omitempty"`
	WorkingMbrCount  int    `json:"workingMbrCount,omitempty"`
}

// TrunkEndpoint represents the REST resource for managing a trunk.
const TrunkEndpoint = "/trunk"

// A TrunkResource provides API to manage trunks configuration.
type TrunkResource struct {
	c *f5.Client
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
