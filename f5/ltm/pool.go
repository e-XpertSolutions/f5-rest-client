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
	AllowNat              string `json:"allowNat"`
	AllowSnat             string `json:"allowSnat"`
	FullPath              string `json:"fullPath"`
	Generation            int64  `json:"generation"`
	IgnorePersistedWeight string `json:"ignorePersistedWeight"`
	IpTosToClient         string `json:"ipTosToClient"`
	IpTosToServer         string `json:"ipTosToServer"`
	Kind                  string `json:"kind"`
	LinkQosToClient       string `json:"linkQosToClient"`
	LinkQosToServer       string `json:"linkQosToServer"`
	LoadBalancingMode     string `json:"loadBalancingMode"`
	MembersReference      struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"membersReference"`
	MinActiveMembers       int64  `json:"minActiveMembers"`
	MinUpMembers           int64  `json:"minUpMembers"`
	MinUpMembersAction     string `json:"minUpMembersAction"`
	MinUpMembersChecking   string `json:"minUpMembersChecking"`
	Monitor                string `json:"monitor"`
	Name                   string `json:"name"`
	QueueDepthLimit        int64  `json:"queueDepthLimit"`
	QueueOnConnectionLimit string `json:"queueOnConnectionLimit"`
	QueueTimeLimit         int64  `json:"queueTimeLimit"`
	ReselectTries          int64  `json:"reselectTries"`
	SelfLink               string `json:"selfLink"`
	ServiceDownAction      string `json:"serviceDownAction"`
	SlowRampTime           int64  `json:"slowRampTime"`
}

type PoolMembersConfig struct {
	Items []struct {
		Address         string `json:"address"`
		ConnectionLimit int64  `json:"connectionLimit"`
		DynamicRatio    int64  `json:"dynamicRatio"`
		Ephemeral       string `json:"ephemeral"`
		Fqdn            struct {
			Autopopulate string `json:"autopopulate"`
		} `json:"fqdn"`
		FullPath       string `json:"fullPath"`
		Generation     int64  `json:"generation"`
		InheritProfile string `json:"inheritProfile"`
		Kind           string `json:"kind"`
		Logging        string `json:"logging"`
		Monitor        string `json:"monitor"`
		Name           string `json:"name"`
		Partition      string `json:"partition"`
		PriorityGroup  int64  `json:"priorityGroup"`
		RateLimit      string `json:"rateLimit"`
		Ratio          int64  `json:"ratio"`
		SelfLink       string `json:"selfLink"`
		Session        string `json:"session"`
		State          string `json:"state"`
	} `json:"items"`
	Kind     string `json:"kind"`
	SelfLink string `json:"selfLink"`
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
