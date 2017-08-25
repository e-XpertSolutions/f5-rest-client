// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPFixElementConfigList holds a list of IPFixElement configuration.
type IPFixElementConfigList struct {
	Items    []IPFixElementConfig `json:"items"`
	Kind     string               `json:"kind"`
	SelfLink string               `json:"selflink"`
}

// IPFixElementConfig holds the configuration of a single IPFixElement.
type IPFixElementConfig struct {
	DataType     string `json:"dataType"`
	EnterpriseID int    `json:"enterpriseId"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	ID           int    `json:"id"`
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	Partition    string `json:"partition"`
	SelfLink     string `json:"selfLink"`
	Size         int    `json:"size"`
}

// IPFixElementEndpoint represents the REST resource for managing IPFixElement.
const IPFixElementEndpoint = "/ipfix/element"

// IPFixElementResource provides an API to manage IPFixElement configurations.
type IPFixElementResource struct {
	c *f5.Client
}

// ListAll  lists all the IPFixElement configurations.
func (r *IPFixElementResource) ListAll() (*IPFixElementConfigList, error) {
	var list IPFixElementConfigList
	if err := r.c.ReadQuery(BasePath+IPFixElementEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPFixElement configuration identified by id.
func (r *IPFixElementResource) Get(id string) (*IPFixElementConfig, error) {
	var item IPFixElementConfig
	if err := r.c.ReadQuery(BasePath+IPFixElementEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPFixElement configuration.
func (r *IPFixElementResource) Create(item IPFixElementConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPFixElementEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPFixElement configuration identified by id.
func (r *IPFixElementResource) Edit(id string, item IPFixElementConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPFixElementEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPFixElement configuration identified by id.
func (r *IPFixElementResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPFixElementEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
