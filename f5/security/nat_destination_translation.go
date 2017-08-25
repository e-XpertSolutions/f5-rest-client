// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// NATDestinationTranslationConfigList holds a list of NATDestinationTranslation configuration.
type NATDestinationTranslationConfigList struct {
	Items    []NATDestinationTranslationConfig `json:"items"`
	Kind     string                            `json:"kind"`
	SelfLink string                            `json:"selflink"`
}

// NATDestinationTranslationConfig holds the configuration of a single NATDestinationTranslation.
type NATDestinationTranslationConfig struct {
}

// NATDestinationTranslationEndpoint represents the REST resource for managing NATDestinationTranslation.
const NATDestinationTranslationEndpoint = "/nat/destination-translation"

// NATDestinationTranslationResource provides an API to manage NATDestinationTranslation configurations.
type NATDestinationTranslationResource struct {
	c *f5.Client
}

// ListAll  lists all the NATDestinationTranslation configurations.
func (r *NATDestinationTranslationResource) ListAll() (*NATDestinationTranslationConfigList, error) {
	var list NATDestinationTranslationConfigList
	if err := r.c.ReadQuery(BasePath+NATDestinationTranslationEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single NATDestinationTranslation configuration identified by id.
func (r *NATDestinationTranslationResource) Get(id string) (*NATDestinationTranslationConfig, error) {
	var item NATDestinationTranslationConfig
	if err := r.c.ReadQuery(BasePath+NATDestinationTranslationEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new NATDestinationTranslation configuration.
func (r *NATDestinationTranslationResource) Create(item NATDestinationTranslationConfig) error {
	if err := r.c.ModQuery("POST", BasePath+NATDestinationTranslationEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a NATDestinationTranslation configuration identified by id.
func (r *NATDestinationTranslationResource) Edit(id string, item NATDestinationTranslationConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+NATDestinationTranslationEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single NATDestinationTranslation configuration identified by id.
func (r *NATDestinationTranslationResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+NATDestinationTranslationEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
