// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DatasyncGlobalProfileConfigList holds a list of DatasyncGlobalProfile configuration.
type DatasyncGlobalProfileConfigList struct {
	Items    []DatasyncGlobalProfileConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// DatasyncGlobalProfileConfig holds the configuration of a single DatasyncGlobalProfile.
type DatasyncGlobalProfileConfig struct {
	ActivationEpoch   string `json:"activationEpoch"`
	CreateTimestamp   int    `json:"createTimestamp"`
	DeactivationEpoch string `json:"deactivationEpoch"`
	FullPath          string `json:"fullPath"`
	Generation        int    `json:"generation"`
	GraceTime         string `json:"graceTime"`
	Kind              string `json:"kind"`
	MasterKey         string `json:"masterKey"`
	MaxRows           string `json:"maxRows"`
	MinRows           string `json:"minRows"`
	Name              string `json:"name"`
	Partition         string `json:"partition"`
	RegenInterval     string `json:"regenInterval"`
	RegenTimeOffset   string `json:"regenTimeOffset"`
	RsaBits           string `json:"rsaBits"`
	RsaExp            string `json:"rsaExp"`
	SelfLink          string `json:"selfLink"`
	SubPath           string `json:"subPath"`
	Table             string `json:"table"`
}

// DatasyncGlobalProfileEndpoint represents the REST resource for managing DatasyncGlobalProfile.
const DatasyncGlobalProfileEndpoint = "/datasync/global-profile"

// DatasyncGlobalProfileResource provides an API to manage DatasyncGlobalProfile configurations.
type DatasyncGlobalProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the DatasyncGlobalProfile configurations.
func (r *DatasyncGlobalProfileResource) ListAll() (*DatasyncGlobalProfileConfigList, error) {
	var list DatasyncGlobalProfileConfigList
	if err := r.c.ReadQuery(BasePath+DatasyncGlobalProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DatasyncGlobalProfile configuration identified by id.
func (r *DatasyncGlobalProfileResource) Get(id string) (*DatasyncGlobalProfileConfig, error) {
	var item DatasyncGlobalProfileConfig
	if err := r.c.ReadQuery(BasePath+DatasyncGlobalProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DatasyncGlobalProfile configuration.
func (r *DatasyncGlobalProfileResource) Create(item DatasyncGlobalProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DatasyncGlobalProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DatasyncGlobalProfile configuration identified by id.
func (r *DatasyncGlobalProfileResource) Edit(id string, item DatasyncGlobalProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DatasyncGlobalProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DatasyncGlobalProfile configuration identified by id.
func (r *DatasyncGlobalProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DatasyncGlobalProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
