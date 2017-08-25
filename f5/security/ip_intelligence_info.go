// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPIntelligenceInfoConfigList holds a list of IPIntelligenceInfo configuration.
type IPIntelligenceInfoConfigList struct {
	Items    []IPIntelligenceInfoConfig `json:"items"`
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selflink"`
}

// IPIntelligenceInfoConfig holds the configuration of a single IPIntelligenceInfo.
type IPIntelligenceInfoConfig struct {
}

// IPIntelligenceInfoEndpoint represents the REST resource for managing IPIntelligenceInfo.
const IPIntelligenceInfoEndpoint = "/ip-intelligence/info"

// IPIntelligenceInfoResource provides an API to manage IPIntelligenceInfo configurations.
type IPIntelligenceInfoResource struct {
	c *f5.Client
}

// ListAll  lists all the IPIntelligenceInfo configurations.
func (r *IPIntelligenceInfoResource) ListAll() (*IPIntelligenceInfoConfigList, error) {
	var list IPIntelligenceInfoConfigList
	if err := r.c.ReadQuery(BasePath+IPIntelligenceInfoEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPIntelligenceInfo configuration identified by id.
func (r *IPIntelligenceInfoResource) Get(id string) (*IPIntelligenceInfoConfig, error) {
	var item IPIntelligenceInfoConfig
	if err := r.c.ReadQuery(BasePath+IPIntelligenceInfoEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPIntelligenceInfo configuration.
func (r *IPIntelligenceInfoResource) Create(item IPIntelligenceInfoConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPIntelligenceInfoEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPIntelligenceInfo configuration identified by id.
func (r *IPIntelligenceInfoResource) Edit(id string, item IPIntelligenceInfoConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPIntelligenceInfoEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPIntelligenceInfo configuration identified by id.
func (r *IPIntelligenceInfoResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPIntelligenceInfoEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
