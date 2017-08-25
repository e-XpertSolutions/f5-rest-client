// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A VlanList holds a list of Vlan.
type VlanList struct {
	Items    []Vlan `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// A Vlan hold the configuration for a vlan.
type Vlan struct {
	AutoLasthop         string `json:"autoLasthop,omitempty"`
	CmpHash             string `json:"cmpHash,omitempty"`
	DagRoundRobin       string `json:"dagRoundRobin,omitempty"`
	DagTunnel           string `json:"dagTunnel,omitempty"`
	Failsafe            string `json:"failsafe,omitempty"`
	FailsafeAction      string `json:"failsafeAction,omitempty"`
	FailsafeTimeout     int    `json:"failsafeTimeout,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	IfIndex             int    `json:"ifIndex,omitempty"`
	InterfacesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"interfacesReference,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Learning string `json:"learning,omitempty"`
	Mtu      int    `json:"mtu,omitempty"`
	Name     string `json:"name,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
	Sflow    struct {
		PollInterval       int    `json:"pollInterval,omitempty"`
		PollIntervalGlobal string `json:"pollIntervalGlobal,omitempty"`
		SamplingRate       int    `json:"samplingRate,omitempty"`
		SamplingRateGlobal string `json:"samplingRateGlobal,omitempty"`
	} `json:"sflow,omitempty"`
	SourceChecking string `json:"sourceChecking,omitempty"`
	Tag            int    `json:"tag,omitempty"`
}

type AssignedInterfaceList struct {
	Items    []AssignedInterface `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selfLink,omitempty"`
}

type AssignedInterface struct {
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Name       string `json:"name,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	TagMode    string `json:"tagMode,omitempty"`
	Untagged   bool   `json:"untagged,omitempty"`
}

// VlanEndpoint represents the REST resource for managing a vlan.
const VlanEndpoint = "/vlan"

// A VlanResource provides API to manage vlan configuration.
type VlanResource struct {
	c *f5.Client
}

// ListAll lists all the vlan configurations.
func (pr *VlanResource) ListAll() (*VlanList, error) {
	var list VlanList
	if err := pr.c.ReadQuery(BasePath+VlanEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single vlan configuration identified by id.
func (pr *VlanResource) Get(id string) (*Vlan, error) {
	var item Vlan
	if err := pr.c.ReadQuery(BasePath+VlanEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new vlan configuration.
func (pr *VlanResource) Create(item Vlan) error {
	if err := pr.c.ModQuery("POST", BasePath+VlanEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a vlan configuration identified by id.
func (pr *VlanResource) Edit(id string, item Vlan) error {
	if err := pr.c.ModQuery("PUT", BasePath+VlanEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single vlan configuration identified by id.
func (pr *VlanResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+VlanEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}

// GetInterfaces gets all interfaces associated to the vlan identified by id.
func (pr *VlanResource) GetInterfaces(id string) (*AssignedInterfaceList, error) {
	var list AssignedInterfaceList
	if err := pr.c.ReadQuery(BasePath+VlanEndpoint+"/"+id+"/interfaces", &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Edit a vlan configuration identified by id.
func (pr *VlanResource) AddInterface(id string, item AssignedInterface) error {
	if err := pr.c.ModQuery("PUT", BasePath+VlanEndpoint+"/"+id+"/interfaces", item); err != nil {
		return err
	}
	return nil
}
