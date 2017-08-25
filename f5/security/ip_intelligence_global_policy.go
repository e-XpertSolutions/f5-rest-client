// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPIntelligenceGlobalPolicyConfigList holds a list of IPIntelligenceGlobalPolicy configuration.
type IPIntelligenceGlobalPolicyConfigList struct {
	Items    []IPIntelligenceGlobalPolicyConfig `json:"items"`
	Kind     string                             `json:"kind"`
	SelfLink string                             `json:"selflink"`
}

// IPIntelligenceGlobalPolicyConfig holds the configuration of a single IPIntelligenceGlobalPolicy.
type IPIntelligenceGlobalPolicyConfig struct {
}

// IPIntelligenceGlobalPolicyEndpoint represents the REST resource for managing IPIntelligenceGlobalPolicy.
const IPIntelligenceGlobalPolicyEndpoint = "/ip-intelligence/global-policy"

// IPIntelligenceGlobalPolicyResource provides an API to manage IPIntelligenceGlobalPolicy configurations.
type IPIntelligenceGlobalPolicyResource struct {
	c *f5.Client
}

// ListAll  lists all the IPIntelligenceGlobalPolicy configurations.
func (r *IPIntelligenceGlobalPolicyResource) ListAll() (*IPIntelligenceGlobalPolicyConfigList, error) {
	var list IPIntelligenceGlobalPolicyConfigList
	if err := r.c.ReadQuery(BasePath+IPIntelligenceGlobalPolicyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPIntelligenceGlobalPolicy configuration identified by id.
func (r *IPIntelligenceGlobalPolicyResource) Get(id string) (*IPIntelligenceGlobalPolicyConfig, error) {
	var item IPIntelligenceGlobalPolicyConfig
	if err := r.c.ReadQuery(BasePath+IPIntelligenceGlobalPolicyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPIntelligenceGlobalPolicy configuration.
func (r *IPIntelligenceGlobalPolicyResource) Create(item IPIntelligenceGlobalPolicyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPIntelligenceGlobalPolicyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPIntelligenceGlobalPolicy configuration identified by id.
func (r *IPIntelligenceGlobalPolicyResource) Edit(id string, item IPIntelligenceGlobalPolicyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPIntelligenceGlobalPolicyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPIntelligenceGlobalPolicy configuration identified by id.
func (r *IPIntelligenceGlobalPolicyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPIntelligenceGlobalPolicyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
