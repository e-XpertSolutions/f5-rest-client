// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DOSBotSignatureConfigList holds a list of DOSBotSignature configuration.
type DOSBotSignatureConfigList struct {
	Items    []DOSBotSignatureConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

// DOSBotSignatureConfig holds the configuration of a single DOSBotSignature.
type DOSBotSignatureConfig struct {
}

// DOSBotSignatureEndpoint represents the REST resource for managing DOSBotSignature.
const DOSBotSignatureEndpoint = "/dos/bot-signature"

// DOSBotSignatureResource provides an API to manage DOSBotSignature configurations.
type DOSBotSignatureResource struct {
	c *f5.Client
}

// ListAll  lists all the DOSBotSignature configurations.
func (r *DOSBotSignatureResource) ListAll() (*DOSBotSignatureConfigList, error) {
	var list DOSBotSignatureConfigList
	if err := r.c.ReadQuery(BasePath+DOSBotSignatureEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DOSBotSignature configuration identified by id.
func (r *DOSBotSignatureResource) Get(id string) (*DOSBotSignatureConfig, error) {
	var item DOSBotSignatureConfig
	if err := r.c.ReadQuery(BasePath+DOSBotSignatureEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DOSBotSignature configuration.
func (r *DOSBotSignatureResource) Create(item DOSBotSignatureConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DOSBotSignatureEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DOSBotSignature configuration identified by id.
func (r *DOSBotSignatureResource) Edit(id string, item DOSBotSignatureConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DOSBotSignatureEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DOSBotSignature configuration identified by id.
func (r *DOSBotSignatureResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DOSBotSignatureEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
