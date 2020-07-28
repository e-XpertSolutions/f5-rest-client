// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A PoolList holds a list of Pool.
type PoolList struct {
	Items    []Pool `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string `json:"selfLink,omitempty" pretty:",expanded"`
}

// A Pool hold the uration for a pool.
type Pool struct {
	AllowNat              string   `json:"allowNat,omitempty" pretty:",expanded"`
	AllowSnat             string   `json:"allowSnat,omitempty" pretty:",expanded"`
	Description           string   `json:"description,omitempty"`
	FullPath              string   `json:"fullPath,omitempty" pretty:",expanded"`
	Generation            int64    `json:"generation,omitempty" pretty:",expanded"`
	IgnorePersistedWeight string   `json:"ignorePersistedWeight,omitempty" pretty:",expanded"`
	IPTosToClient         string   `json:"ipTosToClient,omitempty" pretty:",expanded"`
	IPTosToServer         string   `json:"ipTosToServer,omitempty" pretty:",expanded"`
	Kind                  string   `json:"kind,omitempty" pretty:",expanded"`
	LinkQosToClient       string   `json:"linkQosToClient,omitempty" pretty:",expanded"`
	LinkQosToServer       string   `json:"linkQosToServer,omitempty" pretty:",expanded"`
	LoadBalancingMode     string   `json:"loadBalancingMode,omitempty"`
	Members               []string `json:"members,omitempty"`
	MembersReference      struct {
		IsSubcollection bool          `json:"isSubcollection,omitempty"`
		Link            string        `json:"link,omitempty"`
		Members         []PoolMembers `json:"items,omitempty"`
	} `json:"membersReference,omitempty"`
	MinActiveMembers       int64  `json:"minActiveMembers,omitempty"`
	MinUpMembers           int64  `json:"minUpMembers,omitempty"`
	MinUpMembersAction     string `json:"minUpMembersAction,omitempty"`
	MinUpMembersChecking   string `json:"minUpMembersChecking,omitempty"`
	Monitor                string `json:"monitor,omitempty"`
	Name                   string `json:"name,omitempty"`
	QueueDepthLimit        int64  `json:"queueDepthLimit,omitempty" pretty:",expanded"`
	QueueOnConnectionLimit string `json:"queueOnConnectionLimit,omitempty" pretty:",expanded"`
	QueueTimeLimit         int64  `json:"queueTimeLimit,omitempty" pretty:",expanded"`
	ReselectTries          int64  `json:"reselectTries,omitempty"`
	SelfLink               string `json:"selfLink,omitempty" pretty:",expanded"`
	ServiceDownAction      string `json:"serviceDownAction,omitempty"`
	SlowRampTime           int64  `json:"slowRampTime,omitempty" pretty:",expanded"`
	Partition              string `json:"partition,omitempty"`
}

// PoolEndpoint represents the REST resource for managing a pool.
const PoolEndpoint = "/pool"

// A PoolResource provides API to manage pool uration.
type PoolResource struct {
	c *f5.Client
}

// ListAll lists all the pool urations.
func (pr *PoolResource) ListAll() (*PoolList, error) {
	var list PoolList
	if err := pr.c.ReadQuery(BasePath+PoolEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single pool uration identified by id.
func (pr *PoolResource) Get(id string) (*Pool, error) {
	var item Pool
	if err := pr.c.ReadQuery(BasePath+PoolEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new pool uration.
func (pr *PoolResource) Create(item Pool) error {
	if err := pr.c.ModQuery("POST", BasePath+PoolEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a pool uration identified by id.
func (pr *PoolResource) Edit(id string, item Pool) error {
	if err := pr.c.ModQuery("PUT", BasePath+PoolEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single pool uration identified by id.
func (pr *PoolResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+PoolEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}

// GetMembers gets all the members associated to the pool identified by id.
func (pr *PoolResource) GetMembers(id string) (*PoolMembersList, error) {
	var poolMembers PoolMembersList
	if err := pr.c.ReadQuery(BasePath+PoolEndpoint+"/"+id+"/members", &poolMembers); err != nil {
		return nil, err
	}
	return &poolMembers, nil
}

func (pr *PoolResource) AddMember(id string, poolMember PoolMembers) error {
	if err := pr.c.ModQuery("POST", BasePath+PoolEndpoint+"/"+id+"/members", poolMember); err != nil {
		return err
	}
	return nil
}

func (pr *PoolResource) ShowStats(id string) (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolEndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
