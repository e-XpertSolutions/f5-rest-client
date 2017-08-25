// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ClusterConfigList holds a list of Cluster configuration.
type ClusterConfigList struct {
	Items    []ClusterConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// ClusterConfig holds the configuration of a single Cluster.
type ClusterConfig struct {
}

// ClusterEndpoint represents the REST resource for managing Cluster.
const ClusterEndpoint = "/cluster"

// ClusterResource provides an API to manage Cluster configurations.
type ClusterResource struct {
	c *f5.Client
}

// ListAll  lists all the Cluster configurations.
func (r *ClusterResource) ListAll() (*ClusterConfigList, error) {
	var list ClusterConfigList
	if err := r.c.ReadQuery(BasePath+ClusterEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Cluster configuration identified by id.
func (r *ClusterResource) Get(id string) (*ClusterConfig, error) {
	var item ClusterConfig
	if err := r.c.ReadQuery(BasePath+ClusterEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Cluster configuration.
func (r *ClusterResource) Create(item ClusterConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ClusterEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Cluster configuration identified by id.
func (r *ClusterResource) Edit(id string, item ClusterConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ClusterEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Cluster configuration identified by id.
func (r *ClusterResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ClusterEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
