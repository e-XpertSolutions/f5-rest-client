// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import (
	"encoding/json"
	"net/http"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// A PoolMembersList holds a list of pool members uration.
type PoolMembersList struct {
	Items    []PoolMembers `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selfLink,omitempty"`
}

// A PoolMembers holds the uration for the members of a pool.
type PoolMembers struct {
	Name            string `json:"name,omitempty"`
	Address         string `json:"address,omitempty"`
	ConnectionLimit int64  `json:"connectionLimit,omitempty"`
	DynamicRatio    int64  `json:"dynamicRatio,omitempty"`
	Ephemeral       string `json:"ephemeral,omitempty"`
	Fqdn            struct {
		Autopopulate string `json:"autopopulate,omitempty"`
	} `json:"fqdn,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	Generation     int64  `json:"generation,omitempty"`
	InheritProfile string `json:"inheritProfile,omitempty"`
	Kind           string `json:"kind,omitempty"`
	Logging        string `json:"logging,omitempty"`
	Monitor        string `json:"monitor,omitempty"`
	Partition      string `json:"partition,omitempty"`
	PriorityGroup  int64  `json:"priorityGroup,omitempty"`
	RateLimit      string `json:"rateLimit,omitempty"`
	Ratio          int64  `json:"ratio,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	Session        string `json:"session,omitempty"`
	State          string `json:"state,omitempty"`
}

// PoolMembersEndpoint represents the REST resource for managing pool members.
const PoolMembersEndpoint = "/members"

// PoolMembersResource provides an API to manage pool members uration.
type PoolMembersResource struct {
	c *f5.Client
}

// ListAll lists all the pool members urations.
func (pmr *PoolMembersResource) ListAll(pool string) (*PoolMembersList, error) {
	resp, err := pmr.doRequest("GET", pool, "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return nil, err
	}
	var poolMembers PoolMembersList
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&poolMembers); err != nil {
		return nil, err
	}
	return &poolMembers, nil
}

// Get a single pool members uration identified by id.
func (pmr *PoolMembersResource) Get(pool string, id string) (*PoolMembers, error) {
	resp, err := pmr.doRequest("GET", pool, id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return nil, err
	}
	var item PoolMembers
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new pool members uration.
func (pmr *PoolMembersResource) Create(pool string, item PoolMembers) error {
	resp, err := pmr.doRequest("POST", pool, "", item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Edit a pool members uration indentified by id.
func (pmr *PoolMembersResource) Edit(pool string, id string, item PoolMembers) error {
	resp, err := pmr.doRequest("PUT", pool, id, item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Delete a single pool members uration identified by id.
func (pmr *PoolMembersResource) Delete(pool string, id string) error {
	resp, err := pmr.doRequest("DELETE", pool, id, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return err
	}
	return nil
}

// doRequest creates and send HTTP request using the F5 client.
//
// TODO(gilliek): decorate errors
func (pmr *PoolMembersResource) doRequest(method, pool string, id string, data interface{}) (*http.Response, error) {
	req, err := pmr.c.MakeRequest(method, BasePath+PoolEndpoint+"/"+pool+PoolMembersEndpoint+"/"+id, data)
	if err != nil {
		return nil, err
	}
	resp, err := pmr.c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// readError reads request error object from a HTTP response.
//
// TODO(gilliek): move this function into F5 package.
func (pmr *PoolMembersResource) readError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		errResp, err := f5.NewRequestError(resp.Body)
		if err != nil {
			return err
		}
		return errResp
	}
	return nil
}
