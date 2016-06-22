// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import (
	"e-xpert_solutions/f5-rest-client/f5"
	"encoding/json"
	"net/http"
)

type PoolMembersConfig struct {
	Items    []PoolMembersConfigItem `json:"items,omitempty"`
	Kind     string                  `json:"kind,omitempty"`
	SelfLink string                  `json:"selfLink,omitempty"`
}

type PoolMembersConfigItem struct {
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
	Name           string `json:"name,omitempty"`
	Partition      string `json:"partition,omitempty"`
	PriorityGroup  int64  `json:"priorityGroup,omitempty"`
	RateLimit      string `json:"rateLimit,omitempty"`
	Ratio          int64  `json:"ratio,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	Session        string `json:"session,omitempty"`
	State          string `json:"state,omitempty"`
}

const PoolMembersEndpoint = "/pool_members"

type PoolMembersResource struct {
	c f5.Client
}

func (pmr *PoolMembersResource) ListAll() (*PoolMembersConfig, error) {
	resp, err := pmr.doRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return nil, err
	}
	var poolMembersConfig PoolMembersConfig
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&poolMembersConfig); err != nil {
		return nil, err
	}
	return &poolMembersConfig, nil
}

func (pmr *PoolMembersResource) Get(id string) (*PoolMembersConfigItem, error) {
	resp, err := pmr.doRequest("GET", id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return nil, err
	}
	var item PoolMembersConfigItem
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (pmr *PoolMembersResource) Create(item PoolMembersConfigItem) error {
	resp, err := pmr.doRequest("POST", "", item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return err
	}
	return nil
}

func (pmr *PoolMembersResource) Edit(id string, item PoolMembersConfigItem) error {
	resp, err := pmr.doRequest("PUT", id, item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := pmr.readError(resp); err != nil {
		return err
	}
	return nil
}

func (pmr *PoolMembersResource) Delete(id string) error {
	resp, err := pmr.doRequest("DELETE", id, nil)
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
func (pmr *PoolMembersResource) doRequest(method, id string, data interface{}) (*http.Response, error) {
	req, err := pmr.c.MakeRequest(method, BasePath+PoolEndpoint+"/"+id, data)
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
