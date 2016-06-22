// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import (
	"e-xpert_solutions/f5-rest-client/f5"
	"encoding/json"
	"net/http"
)

type ObjectConfig struct {
	AllowNat              string `json:"allowNat,omitempty"`
	AllowSnat             string `json:"allowSnat,omitempty"`
	FullPath              string `json:"fullPath,omitempty"`
	Generation            int64  `json:"generation,omitempty"`
	IgnorePersistedWeight string `json:"ignorePersistedWeight,omitempty"`
	IpTosToClient         string `json:"ipTosToClient,omitempty"`
	IpTosToServer         string `json:"ipTosToServer,omitempty"`
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

const PoolEndpoint = "/pool"

type PoolResource struct {
	c f5.Client
}

func (pr *PoolResource) ListAll() ([]ObjectConfig, error) {
	resp, err := pr.doRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return nil, err
	}
	var objConfList []ObjectConfig
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&objConfList); err != nil {
		return nil, err
	}
	return objConfList, nil
}

func (pr *PoolResource) Get(id string) (*ObjectConfig, error) {
	resp, err := pr.doRequest("GET", id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return nil, err
	}
	var objConf ObjectConfig
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&objConf); err != nil {
		return nil, err
	}
	return &objConf, nil
}

func (pr *PoolResource) Create(objConf ObjectConfig) error {
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

func (pr *PoolResource) Edit(id string, objConf ObjectConfig) error {
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

func (pr *PoolResource) GetMembers(id string) (*PoolMembersConfig, error) {
	resp, err := pr.doRequest("GET", id+"/members", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := pr.readError(resp); err != nil {
		return nil, err
	}
	var poolMembersConfig PoolMembersConfig
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
