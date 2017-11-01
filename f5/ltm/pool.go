// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A PoolConfigList holds a list of PoolConfig.
type PoolConfigList struct {
	Items    []PoolConfig `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selfLink,omitempty"`
}

// A PoolConfig hold the configuration for a pool.
type PoolConfig struct {
	AllowNat              string   `json:"allowNat,omitempty"`
	AllowSnat             string   `json:"allowSnat,omitempty"`
	FullPath              string   `json:"fullPath,omitempty"`
	Generation            int64    `json:"generation,omitempty"`
	IgnorePersistedWeight string   `json:"ignorePersistedWeight,omitempty"`
	IPTosToClient         string   `json:"ipTosToClient,omitempty"`
	IPTosToServer         string   `json:"ipTosToServer,omitempty"`
	Kind                  string   `json:"kind,omitempty"`
	LinkQosToClient       string   `json:"linkQosToClient,omitempty"`
	LinkQosToServer       string   `json:"linkQosToServer,omitempty"`
	LoadBalancingMode     string   `json:"loadBalancingMode,omitempty"`
	Members               []string `json:"members,omitempty"`
	MembersReference      struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"membersReference,omitempty"`
	MinActiveMembers       int64  `json:"minActiveMembers,omitempty"`
	MinUpMembers           int64  `json:"minUpMembers,omitempty"`
	MinUpMembersAction     string `json:"minUpMembersAction,omitempty"`
	MinUpMembersChecking   string `json:"minUpMembersChecking,omitempty"`
	Monitor                string `json:"monitor,omitempty"`
	Name                   string `json:"name,omitempty"`
	QueueDepthLimit        int64  `json:"queueDepthLimit,omitempty"`
	QueueOnConnectionLimit string `json:"queueOnConnectionLimit,omitempty"`
	QueueTimeLimit         int64  `json:"queueTimeLimit,omitempty"`
	ReselectTries          int64  `json:"reselectTries,omitempty"`
	SelfLink               string `json:"selfLink,omitempty"`
	ServiceDownAction      string `json:"serviceDownAction,omitempty"`
	SlowRampTime           int64  `json:"slowRampTime,omitempty"`
}

// PoolEndpoint represents the REST resource for managing a pool.
const PoolEndpoint = "/pool"

// A PoolResource provides API to manage pool configuration.
type PoolResource struct {
	c *f5.Client
}

// ListAll lists all the pool configurations.
func (pr *PoolResource) ListAll() (*PoolConfigList, error) {
	var list PoolConfigList
	if err := pr.c.ReadQuery(BasePath+PoolEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single pool configuration identified by id.
func (pr *PoolResource) Get(id string) (*PoolConfig, error) {
	var item PoolConfig
	if err := pr.c.ReadQuery(BasePath+PoolEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new pool configuration.
func (pr *PoolResource) Create(item PoolConfig) error {
	if err := pr.c.ModQuery("POST", BasePath+PoolEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a pool configuration identified by id.
func (pr *PoolResource) Edit(id string, item PoolConfig) error {
	if err := pr.c.ModQuery("PUT", BasePath+PoolEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single pool configuration identified by id.
func (pr *PoolResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+PoolEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}

// GetMembers gets all the members associated to the pool identified by id.
func (pr *PoolResource) GetMembers(id string) (*PoolMembersConfigList, error) {
	var poolMembersConfig PoolMembersConfigList
	if err := pr.c.ReadQuery(BasePath+PoolEndpoint+"/"+id+"/members", &poolMembersConfig); err != nil {
		return nil, err
	}
	return &poolMembersConfig, nil
}

func (nr *PoolResource) ShowStats(id string) (*PoolStats, error) {
	var item PoolStats
	if err := nr.c.ReadQuery(BasePath+PoolEndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
