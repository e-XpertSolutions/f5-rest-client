// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DistributedAppList holds a list of DistributedApp uration.
type DistributedAppList struct {
	Items    []DistributedApp `json:"items,omitempty"`
	Kind     string           `json:"kind,omitempty"`
	SelfLink string           `json:"selflink,omitempty"`
}

// DistributedApp holds the uration of a single DistributedApp.
type DistributedApp struct {
	DependencyLevel string `json:"dependencyLevel,omitempty"`
	FullPath        string `json:"fullPath,omitempty"`
	Generation      int    `json:"generation,omitempty"`
	Kind            string `json:"kind,omitempty"`
	Name            string `json:"name,omitempty"`
	Partition       string `json:"partition,omitempty"`
	PersistCidrIpv4 int    `json:"persistCidrIpv4,omitempty"`
	PersistCidrIpv6 int    `json:"persistCidrIpv6,omitempty"`
	Persistence     string `json:"persistence,omitempty"`
	SelfLink        string `json:"selfLink,omitempty"`
	TTLPersistence  int    `json:"ttlPersistence,omitempty"`
	Wideips         []struct {
		Name string `json:"name,omitempty"`
	} `json:"wideips,omitempty"`
}

// DistributedAppEndpoint represents the REST resource for managing DistributedApp.
const DistributedAppEndpoint = "/distributed-app"

// DistributedAppResource provides an API to manage DistributedApp urations.
type DistributedAppResource struct {
	c *f5.Client
}

// ListAll  lists all the DistributedApp urations.
func (r *DistributedAppResource) ListAll() (*DistributedAppList, error) {
	var list DistributedAppList
	if err := r.c.ReadQuery(BasePath+DistributedAppEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DistributedApp uration identified by id.
func (r *DistributedAppResource) Get(id string) (*DistributedApp, error) {
	var item DistributedApp
	if err := r.c.ReadQuery(BasePath+DistributedAppEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DistributedApp uration.
func (r *DistributedAppResource) Create(item DistributedApp) error {
	if err := r.c.ModQuery("POST", BasePath+DistributedAppEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DistributedApp uration identified by id.
func (r *DistributedAppResource) Edit(id string, item DistributedApp) error {
	if err := r.c.ModQuery("PUT", BasePath+DistributedAppEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DistributedApp uration identified by id.
func (r *DistributedAppResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DistributedAppEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
