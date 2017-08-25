// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// NATSourceTranslationConfigList holds a list of NATSourceTranslation configuration.
type NATSourceTranslationConfigList struct {
	Items    []NATSourceTranslationConfig `json:"items"`
	Kind     string                       `json:"kind"`
	SelfLink string                       `json:"selflink"`
}

// NATSourceTranslationConfig holds the configuration of a single NATSourceTranslation.
type NATSourceTranslationConfig struct {
}

// NATSourceTranslationEndpoint represents the REST resource for managing NATSourceTranslation.
const NATSourceTranslationEndpoint = "/nat/source-translation"

// NATSourceTranslationResource provides an API to manage NATSourceTranslation configurations.
type NATSourceTranslationResource struct {
	c *f5.Client
}

// ListAll  lists all the NATSourceTranslation configurations.
func (r *NATSourceTranslationResource) ListAll() (*NATSourceTranslationConfigList, error) {
	var list NATSourceTranslationConfigList
	if err := r.c.ReadQuery(BasePath+NATSourceTranslationEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single NATSourceTranslation configuration identified by id.
func (r *NATSourceTranslationResource) Get(id string) (*NATSourceTranslationConfig, error) {
	var item NATSourceTranslationConfig
	if err := r.c.ReadQuery(BasePath+NATSourceTranslationEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new NATSourceTranslation configuration.
func (r *NATSourceTranslationResource) Create(item NATSourceTranslationConfig) error {
	if err := r.c.ModQuery("POST", BasePath+NATSourceTranslationEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a NATSourceTranslation configuration identified by id.
func (r *NATSourceTranslationResource) Edit(id string, item NATSourceTranslationConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+NATSourceTranslationEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single NATSourceTranslation configuration identified by id.
func (r *NATSourceTranslationResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+NATSourceTranslationEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
