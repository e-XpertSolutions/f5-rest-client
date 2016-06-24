// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import (
	"encoding/json"
	"net/http"

	"e-xpert_solutions/f5-rest-client/f5"
)

// A PoolConfigList holds a list of PoolConfig.
type PoolConfigList struct {
	Items    []VirtualServerConfig `json:"items,omitempty"`
	Kind     string                `json:"kind,omitempty"`
	SelfLink string                `json:"selfLink,omitempty"`
}

// A PoolConfig hold the configuration for a pool.
type PoolConfig struct {
	AllowNat              string `json:"allowNat,omitempty"`
	AllowSnat             string `json:"allowSnat,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int64  `json:"generation,omitempty"`
	IgnorePersistedWeight string `json:"ignorePersistedWeight,omitempty"`
	IPTosToClient         string `json:"ipTosToClient,omitempty"`
	IPTosToServer         string `json:"ipTosToServer,omitempty"`
	Kind                  string `json:"kind,omitempty"`
	LinkQosToClient       string `json:"linkQosToClient,omitempty"`
	LinkQosToServer       string `json:"linkQosToServer,omitempty"`
	LoadBalancingMode     string `json:"loadBalancingMode,omitempty"`
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
	c f5.Client
}

// ListAll lists all the pool configurations.
func (pr *PoolResource) ListAll() (*PoolConfigList, error) {
	resp, err := pr.doRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return nil, err
	}
	var list PoolConfigList
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single pool configuration identified by id.
func (pr *PoolResource) Get(id string) (*PoolConfig, error) {
	resp, err := pr.doRequest("GET", id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return nil, err
	}
	var objConf PoolConfig
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&objConf); err != nil {
		return nil, err
	}
	return &objConf, nil
}

// Create a new pool configuration.
func (pr *PoolResource) Create(objConf PoolConfig) error {
	resp, err := pr.doRequest("POST", "", objConf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Edit a pool configuration identified by id.
func (pr *PoolResource) Edit(id string, objConf PoolConfig) error {
	resp, err := pr.doRequest("PUT", id, objConf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Delete a single pool configuration identified by id.
func (pr *PoolResource) Delete(id string) error {
	resp, err := pr.doRequest("DELETE", id, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return err
	}
	return nil
}

// GetMembers gets all the members associated to the pool identified by id.
func (pr *PoolResource) GetMembers(id string) (*PoolMembersConfigList, error) {
	resp, err := pr.doRequest("GET", id+"/members", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return nil, err
	}
	var poolMembersConfig PoolMembersConfigList
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&poolMembersConfig); err != nil {
		return nil, err
	}
	return &poolMembersConfig, nil
}

// doRequest creates and send HTTP request using the F5 client.
//
// TODO(gilliek): decorate errors
func (pr *PoolResource) doRequest(method, id string, data interface{}) (*http.Response, error) {
	req, err := pr.c.MakeRequest(method, BasePath+PoolEndpoint+"/"+id, data)
	if err != nil {
		return nil, err
	}
	resp, err := pr.c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// readError reads request error object from a HTTP response.
//
// TODO(gilliek): move this function into F5 package.
func (pr *PoolResource) readError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		errResp, err := f5.NewRequestError(resp.Body)
		if err != nil {
			return err
		}
		return errResp
	}
	return nil
}
