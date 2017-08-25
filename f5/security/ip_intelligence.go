// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPIntelligenceConfigList holds a list of IPIntelligence configuration.
type IPIntelligenceConfigList struct {
	Items    []IPIntelligenceConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// IPIntelligenceConfig holds the configuration of a single IPIntelligence.
type IPIntelligenceConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// IPIntelligenceEndpoint represents the REST resource for managing IPIntelligence.
const IPIntelligenceEndpoint = "/ip-intelligence"

// IPIntelligenceResource provides an API to manage IPIntelligence configurations.
type IPIntelligenceResource struct {
	c *f5.Client
}

// ListAll  lists all the IPIntelligence configurations.
func (r *IPIntelligenceResource) ListAll() (*IPIntelligenceConfigList, error) {
	var list IPIntelligenceConfigList
	if err := r.c.ReadQuery(BasePath+IPIntelligenceEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPIntelligence configuration identified by id.
func (r *IPIntelligenceResource) Get(id string) (*IPIntelligenceConfig, error) {
	var item IPIntelligenceConfig
	if err := r.c.ReadQuery(BasePath+IPIntelligenceEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPIntelligence configuration.
func (r *IPIntelligenceResource) Create(item IPIntelligenceConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPIntelligenceEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPIntelligence configuration identified by id.
func (r *IPIntelligenceResource) Edit(id string, item IPIntelligenceConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPIntelligenceEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPIntelligence configuration identified by id.
func (r *IPIntelligenceResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPIntelligenceEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
