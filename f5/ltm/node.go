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
	Address         string `json:"address"`
	ConnectionLimit int    `json:"connectionLimit"`
	DynamicRatio    int    `json:"dynamicRatio"`
	Ephemeral       string `json:"ephemeral"`
	Fqdn            struct {
		AddressFamily string `json:"addressFamily"`
		Autopopulate  string `json:"autopopulate"`
		DownInterval  int    `json:"downInterval"`
		Interval      string `json:"interval"`
	} `json:"fqdn"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Logging    string `json:"logging"`
	Monitor    string `json:"monitor"`
	Name       string `json:"name"`
	Partition  string `json:"partition"`
	RateLimit  string `json:"rateLimit"`
	Ratio      int    `json:"ratio"`
	SelfLink   string `json:"selfLink"`
	Session    string `json:"session"`
	State      string `json:"state"`
}

const NodeEndpoint = "/node"

type NodeResource struct {
	c f5.Client
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

func (nr *NodeResource) Delete(id string) error {
	if err := nr.c.ModQuery("DELETE", BasePath+NodeEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
