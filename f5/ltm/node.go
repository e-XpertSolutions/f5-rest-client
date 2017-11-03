// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type NodeConfigList struct {
	Items    []NodeConfig `json:"items"`
	Kind     string       `json:"kind"`
	SelfLink string       `json:"selfLink"`
}

type NodeConfig struct {
	Address         string `json:"address,omitempty"`
	ConnectionLimit int    `json:"connectionLimit,omitempty"`
	DynamicRatio    int    `json:"dynamicRatio,omitempty"`
	Ephemeral       string `json:"ephemeral,omitempty"`
	Fqdn            struct {
		AddressFamily string `json:"addressFamily,omitempty"`
		Autopopulate  string `json:"autopopulate,omitempty"`
		DownInterval  int    `json:"downInterval,omitempty"`
		Interval      string `json:"interval,omitempty"`
	} `json:"fqdn,omitempty"`
	FullPath   string `json:"fullPath,omitempty"`
	Generation int    `json:"generation,omitempty"`
	Kind       string `json:"kind,omitempty"`
	Logging    string `json:"logging,omitempty"`
	Monitor    string `json:"monitor,omitempty"`
	Name       string `json:"name,omitempty"`
	Partition  string `json:"partition,omitempty"`
	RateLimit  string `json:"rateLimit,omitempty"`
	Ratio      int    `json:"ratio,omitempty"`
	SelfLink   string `json:"selfLink,omitempty"`
	Session    string `json:"session,omitempty"`
	State      string `json:"state,omitempty"`
}

const NodeEndpoint = "/node"

type NodeResource struct {
	c *f5.Client
}

func (nr *NodeResource) ListAll() (*NodeConfigList, error) {
	var list NodeConfigList
	if err := nr.c.ReadQuery(BasePath+NodeEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (nr *NodeResource) Get(id string) (*NodeConfig, error) {
	var item NodeConfig
	if err := nr.c.ReadQuery(BasePath+NodeEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (nr *NodeResource) Create(item NodeConfig) error {
	if err := nr.c.ModQuery("POST", BasePath+NodeEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) Edit(id string, item NodeConfig) error {
	if err := nr.c.ModQuery("PUT", BasePath+NodeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) Enable(id string) error {
	item := NodeConfig{Session: "unchecked", State: "user-enabled"}
	if err := nr.c.ModQuery("PUT", BasePath+NodeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) Disable(id string) error {
	item := NodeConfig{Session: "unchecked", State: "user-disabled"}
	if err := nr.c.ModQuery("PUT", BasePath+NodeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) ForceOffline(id string) error {
	item := NodeConfig{Session: "user-down", State: "user-disabled"}
	if err := nr.c.ModQuery("PUT", BasePath+NodeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) Delete(id string) error {
	if err := nr.c.ModQuery("DELETE", BasePath+NodeEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) ShowStats(id string) (*NodeStatsList, error) {
	var item NodeStatsList
	if err := nr.c.ReadQuery(BasePath+NodeEndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
