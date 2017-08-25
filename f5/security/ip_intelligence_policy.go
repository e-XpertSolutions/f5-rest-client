// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// IPIntelligencePolicyConfigList holds a list of IPIntelligencePolicy configuration.
type IPIntelligencePolicyConfigList struct {
	Items    []IPIntelligencePolicyConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// IPIntelligencePolicyConfig holds the configuration of a single IPIntelligencePolicy.
type IPIntelligencePolicyConfig struct {
	DefaultAction                   string `json:"defaultAction"`
	DefaultLogBlacklistHitOnly      string `json:"defaultLogBlacklistHitOnly"`
	DefaultLogBlacklistWhitelistHit string `json:"defaultLogBlacklistWhitelistHit"`
	FullPath                        string `json:"fullPath"`
	Generation                      int    `json:"generation"`
	Kind                            string `json:"kind"`
	Name                            string `json:"name"`
	Partition                       string `json:"partition"`
	SelfLink                        string `json:"selfLink"`
}

// IPIntelligencePolicyEndpoint represents the REST resource for managing IPIntelligencePolicy.
const IPIntelligencePolicyEndpoint = "/ip-intelligence/policy"

// IPIntelligencePolicyResource provides an API to manage IPIntelligencePolicy configurations.
type IPIntelligencePolicyResource struct {
	c *f5.Client
}

// ListAll  lists all the IPIntelligencePolicy configurations.
func (r *IPIntelligencePolicyResource) ListAll() (*IPIntelligencePolicyConfigList, error) {
	var list IPIntelligencePolicyConfigList
	if err := r.c.ReadQuery(BasePath+IPIntelligencePolicyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single IPIntelligencePolicy configuration identified by id.
func (r *IPIntelligencePolicyResource) Get(id string) (*IPIntelligencePolicyConfig, error) {
	var item IPIntelligencePolicyConfig
	if err := r.c.ReadQuery(BasePath+IPIntelligencePolicyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new IPIntelligencePolicy configuration.
func (r *IPIntelligencePolicyResource) Create(item IPIntelligencePolicyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+IPIntelligencePolicyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a IPIntelligencePolicy configuration identified by id.
func (r *IPIntelligencePolicyResource) Edit(id string, item IPIntelligencePolicyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+IPIntelligencePolicyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single IPIntelligencePolicy configuration identified by id.
func (r *IPIntelligencePolicyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+IPIntelligencePolicyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
