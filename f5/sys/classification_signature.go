// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ClassificationSignatureConfigList holds a list of ClassificationSignature configuration.
type ClassificationSignatureConfigList struct {
	Items    []ClassificationSignatureConfig `json:"items"`
	Kind     string                          `json:"kind"`
	SelfLink string                          `json:"selflink"`
}

// ClassificationSignatureConfig holds the configuration of a single ClassificationSignature.
type ClassificationSignatureConfig struct {
}

// ClassificationSignatureEndpoint represents the REST resource for managing ClassificationSignature.
const ClassificationSignatureEndpoint = "/classification-signature"

// ClassificationSignatureResource provides an API to manage ClassificationSignature configurations.
type ClassificationSignatureResource struct {
	c *f5.Client
}

// ListAll  lists all the ClassificationSignature configurations.
func (r *ClassificationSignatureResource) ListAll() (*ClassificationSignatureConfigList, error) {
	var list ClassificationSignatureConfigList
	if err := r.c.ReadQuery(BasePath+ClassificationSignatureEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ClassificationSignature configuration identified by id.
func (r *ClassificationSignatureResource) Get(id string) (*ClassificationSignatureConfig, error) {
	var item ClassificationSignatureConfig
	if err := r.c.ReadQuery(BasePath+ClassificationSignatureEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ClassificationSignature configuration.
func (r *ClassificationSignatureResource) Create(item ClassificationSignatureConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ClassificationSignatureEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ClassificationSignature configuration identified by id.
func (r *ClassificationSignatureResource) Edit(id string, item ClassificationSignatureConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+ClassificationSignatureEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ClassificationSignature configuration identified by id.
func (r *ClassificationSignatureResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ClassificationSignatureEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
