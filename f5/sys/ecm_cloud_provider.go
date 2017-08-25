// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ECMCloudProviderConfigList holds a list of ECMCloudProvider configuration.
type ECMCloudProviderConfigList struct {
	Items    []ECMCloudProviderConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// ECMCloudProviderConfig holds the configuration of a single ECMCloudProvider.
type ECMCloudProviderConfig struct {
	Description      string `json:"description"`
	FullPath         string `json:"fullPath"`
	Generation       int    `json:"generation"`
	Kind             string `json:"kind"`
	Name             string `json:"name"`
	Partition        string `json:"partition"`
	PropertyTemplate []struct {
		Name string `json:"name"`
	} `json:"propertyTemplate"`
	SelfLink string `json:"selfLink"`
}

// ECMCloudProviderEndpoint represents the REST resource for managing ECMCloudProvider.
const ECMCloudProviderEndpoint = "/ecm/cloud-provider"

// ECMCloudProviderResource provides an API to manage ECMCloudProvider configurations.
type ECMCloudProviderResource struct {
	c *f5.Client
}

// ListAll  lists all the ECMCloudProvider configurations.
func (r *ECMCloudProviderResource) ListAll() (*ECMCloudProviderConfigList, error) {
	var list ECMCloudProviderConfigList
	if err := r.c.ReadQuery(BasePath+ECMCloudProviderEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ECMCloudProvider configuration identified by id.
func (r *ECMCloudProviderResource) Get(id string) (*ECMCloudProviderConfig, error) {
	var item ECMCloudProviderConfig
	if err := r.c.ReadQuery(BasePath+ECMCloudProviderEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ECMCloudProvider configuration.
func (r *ECMCloudProviderResource) Create(item ECMCloudProviderConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ECMCloudProviderEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ECMCloudProvider configuration identified by id.
func (r *ECMCloudProviderResource) Edit(id string, item ECMCloudProviderConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ECMCloudProviderEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ECMCloudProvider configuration identified by id.
func (r *ECMCloudProviderResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ECMCloudProviderEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
