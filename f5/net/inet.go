// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A InterfaceList holds a list of Interface.
type InterfaceList struct {
	Items    []Interface `json:"items,omitempty"`
	Kind     string      `json:"kind,omitempty"`
	SelfLink string      `json:"selfLink,omitempty"`
}

type Interface struct {
	Bundle                 string `json:"bundle,omitempty"`
	BundleSpeed            string `json:"bundleSpeed,omitempty"`
	Enabled                bool   `json:"enabled,omitempty"`
	FlowControl            string `json:"flowControl,omitempty"`
	ForceGigabitFiber      string `json:"forceGigabitFiber,omitempty"`
	ForwardErrorCorrection string `json:"forwardErrorCorrection,omitempty"`
	FullPath               string `json:"fullPath,omitempty"`
	Generation             int    `json:"generation,omitempty"`
	IfIndex                int    `json:"ifIndex,omitempty"`
	Kind                   string `json:"kind,omitempty"`
	LldpAdmin              string `json:"lldpAdmin,omitempty"`
	LldpTlvmap             int    `json:"lldpTlvmap,omitempty"`
	MacAddress             string `json:"macAddress,omitempty"`
	MediaActive            string `json:"mediaActive,omitempty"`
	MediaFixed             string `json:"mediaFixed,omitempty"`
	MediaSfp               string `json:"mediaSfp,omitempty"`
	Mtu                    int    `json:"mtu,omitempty"`
	Name                   string `json:"name,omitempty"`
	PreferPort             string `json:"preferPort,omitempty"`
	QinqEthertype          string `json:"qinqEthertype,omitempty"`
	SelfLink               string `json:"selfLink,omitempty"`
	Sflow                  struct {
		PollInterval       int    `json:"pollInterval,omitempty"`
		PollIntervalGlobal string `json:"pollIntervalGlobal,omitempty"`
	} `json:"sflow,omitempty"`
	Stp             string `json:"stp,omitempty"`
	StpAutoEdgePort string `json:"stpAutoEdgePort,omitempty"`
	StpEdgePort     string `json:"stpEdgePort,omitempty"`
	StpLinkType     string `json:"stpLinkType,omitempty"`
}

// InterfaceEndpoint represents the REST resource for managing interfaces.
const InterfaceEndpoint = "/interface"

// A InetResource provides an API to manage Interface configurations.
type InetResource struct {
	c *f5.Client
}

// ListAll lists all interfaces configuration.
func (pr *InetResource) ListAll() (*InterfaceList, error) {
	var list InterfaceList
	if err := pr.c.ReadQuery(BasePath+InterfaceEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single interface configuration identified by id.
func (pr *InetResource) Get(id string) (*Interface, error) {
	var item Interface
	if err := pr.c.ReadQuery(BasePath+InterfaceEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (pr *InetResource) ShowStats(id string) (*InterfaceStatsList, error) {
	var item InterfaceStatsList
	if err := pr.c.ReadQuery(BasePath+InterfaceEndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
