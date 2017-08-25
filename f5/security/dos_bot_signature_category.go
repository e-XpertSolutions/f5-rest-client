// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSBotSignatureCategoryConfigList holds a list of DOSBotSignatureCategory configuration.
type DOSBotSignatureCategoryConfigList struct {
	Items    []DOSBotSignatureCategoryConfig `json:"items"`
	Kind     string                          `json:"kind"`
	SelfLink string                          `json:"selflink"`
}

// DOSBotSignatureCategoryConfig holds the configuration of a single DOSBotSignatureCategory.
type DOSBotSignatureCategoryConfig struct {
}

// DOSBotSignatureCategoryEndpoint represents the REST resource for managing DOSBotSignatureCategory.
const DOSBotSignatureCategoryEndpoint = "/dos/bot-signature-category"

// DOSBotSignatureCategoryResource provides an API to manage DOSBotSignatureCategory configurations.
type DOSBotSignatureCategoryResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSBotSignatureCategory configurations.
func (r *DOSBotSignatureCategoryResource) ListAll() (*DOSBotSignatureCategoryConfigList, error) {
	var list DOSBotSignatureCategoryConfigList
	if err := r.c.ReadQuery(BasePath+DOSBotSignatureCategoryEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSBotSignatureCategory configuration identified by id.
func (r *DOSBotSignatureCategoryResource) Get(id string) (*DOSBotSignatureCategoryConfig, error) {
	var item DOSBotSignatureCategoryConfig
	if err := r.c.ReadQuery(BasePath+DOSBotSignatureCategoryEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSBotSignatureCategory configuration.
func (r *DOSBotSignatureCategoryResource) Create(item DOSBotSignatureCategoryConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSBotSignatureCategoryEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSBotSignatureCategory configuration identified by id.
func (r *DOSBotSignatureCategoryResource) Edit(id string, item DOSBotSignatureCategoryConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSBotSignatureCategoryEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSBotSignatureCategory configuration identified by id.
func (r *DOSBotSignatureCategoryResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSBotSignatureCategoryEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
