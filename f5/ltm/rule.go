// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import (
	"encoding/json"
	"net/http"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// A RuleList holds a list of iRule configurations.
type RuleList struct {
	Items    []Rule `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string `json:"selfLink,omitempty" pretty:",expanded"`
}

// A Rule holds an iRule configuration.
type Rule struct {
	Action              string `json:"action,omitempty"`
	AppService          string `json:"appService,omitempty"`
	DefinitionChecksum  string `json:"definitionChecksum,omitempty"`
	DefinitionSignature string `json:"definitionSignature,omitempty"`
	Hidden              bool   `json:"hidden,omitempty"`
	IgnoreVerification  string `json:"ignoreVerification,omitempty"`
	NoDelete            bool   `json:"noDelete,omitempty"`
	NoWrite             bool   `json:"noWrite,omitempty"`
	TMPartition         string `json:"tmPartition,omitempty"`
	Plugin              string `json:"plugin,omitempty"`
	PublicCert          string `json:"publicCert,omitempty"`
	SigningKey          string `json:"signingKey,omitempty"`

	Name         string `json:"name,omitempty"`
	Partition    string `json:"partition,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
	ApiAnonymous string `json:"apiAnonymous,omitempty"`
}

// RuleEndpoint represents the REST resource for managing iRule configurations.
const RuleEndpoint = "/rule"

// A RuleResource provides an API to manage iRule configurations.
type RuleResource struct {
	c *f5.Client
}

// ListAll lists all the iRule configurations.
func (rr *RuleResource) ListAll() (*RuleList, error) {
	resp, err := rr.doRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := rr.readError(resp); err != nil {
		return nil, err
	}
	var rules RuleList
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&rules); err != nil {
		return nil, err
	}
	return &rules, nil
}

// Get a single iRule configuration identified by id.
func (rr *RuleResource) Get(id string) (*Rule, error) {
	resp, err := rr.doRequest("GET", id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := rr.readError(resp); err != nil {
		return nil, err
	}
	var rule Rule
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&rule); err != nil {
		return nil, err
	}
	return &rule, nil
}

// Create a new iRule configuration.
func (rr *RuleResource) Create(rule Rule) error {
	resp, err := rr.doRequest("POST", "", rule)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := rr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Edit an iRule configuration identified by id.
func (rr *RuleResource) Edit(id string, rule Rule) error {
	resp, err := rr.doRequest("PUT", id, rule)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := rr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Delete a single iRule configuration identified by id.
func (rr *RuleResource) Delete(id string) error {
	resp, err := rr.doRequest("DELETE", id, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := rr.readError(resp); err != nil {
		return err
	}
	return nil
}

// doRequest creates and send HTTP request using the F5 client.
//
// TODO(gilliek): decorate errors
func (rr *RuleResource) doRequest(method, id string, data interface{}) (*http.Response, error) {
	req, err := rr.c.MakeRequest(method, BasePath+RuleEndpoint+"/"+id, data)
	if err != nil {
		return nil, err
	}
	resp, err := rr.c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// readError reads request error object from a HTTP response.
//
// TODO(gilliek): move this function into F5 package.
func (rr *RuleResource) readError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		errResp, err := f5.NewRequestError(resp.Body)
		if err != nil {
			return err
		}
		return errResp
	}
	return nil
}
