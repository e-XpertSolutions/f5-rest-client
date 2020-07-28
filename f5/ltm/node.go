// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type NodeList struct {
	Items    []Node `json:"items"`
	Kind     string `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string `json:"selfLink,omitempty" pretty:",expanded"`
}

type Node struct {
	Address         string `json:"address,omitempty"`
	ConnectionLimit int    `json:"connectionLimit,omitempty"`
	Description     string `json:"description,omitempty"`
	DynamicRatio    int    `json:"dynamicRatio,omitempty"`
	Ephemeral       string `json:"ephemeral,omitempty"`
	Fqdn            struct {
		AddressFamily string `json:"addressFamily,omitempty"`
		Autopopulate  string `json:"autopopulate,omitempty"`
		DownInterval  int    `json:"downInterval,omitempty"`
		Interval      string `json:"interval,omitempty"`
		Name          string `json:"tmName,omitempty"`
	} `json:"fqdn,omitempty"`
	FullPath   string `json:"fullPath,omitempty" pretty:",expanded"`
	Generation int    `json:"generation,omitempty" pretty:",expanded"`
	Kind       string `json:"kind,omitempty" pretty:",expanded"`
	Logging    string `json:"logging,omitempty"`
	Monitor    string `json:"monitor,omitempty"`
	Name       string `json:"name,omitempty"`
	Partition  string `json:"partition,omitempty"`
	RateLimit  string `json:"rateLimit,omitempty"`
	Ratio      int    `json:"ratio,omitempty"`
	SelfLink   string `json:"selfLink,omitempty" pretty:",expanded"`
	Session    string `json:"session,omitempty"`
	State      string `json:"state,omitempty"`
}

const NodeEndpoint = "/node"

type NodeResource struct {
	c *f5.Client
}

func (nr *NodeResource) ListAll() (*NodeList, error) {
	var list NodeList
	if err := nr.c.ReadQuery(BasePath+NodeEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (nr *NodeResource) Get(id string) (*Node, error) {
	var item Node
	if err := nr.c.ReadQuery(BasePath+NodeEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (nr *NodeResource) Create(item Node) error {
	if err := nr.c.ModQuery("POST", BasePath+NodeEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) Edit(id string, item Node) error {
	if err := nr.c.ModQuery("PUT", BasePath+NodeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) Enable(id string) error {
	item := Node{Session: "user-enabled", State: "user-up"}
	if err := nr.c.ModQuery("PUT", BasePath+NodeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) Disable(id string) error {
	item := Node{Session: "user-disabled", State: "user-up"}
	if err := nr.c.ModQuery("PUT", BasePath+NodeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (nr *NodeResource) ForceOffline(id string) error {
	item := Node{Session: "user-disabled", State: "user-down"}
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
