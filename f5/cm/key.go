// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A KeyList holds a list of Key.
type KeyList struct {
	Items    []Key  `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// A Key hold the configuration for a Key.
type Key struct {
	Checksum       string `json:"checksum,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	Generation     int    `json:"generation,omitempty"`
	KeySize        int    `json:"keySize,omitempty"`
	KeyType        string `json:"keyType,omitempty"`
	Kind           string `json:"kind,omitempty"`
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	Mode           int    `json:"mode,omitempty"`
	Name           string `json:"name,omitempty"`
	Partition      string `json:"partition,omitempty"`
	Revision       int    `json:"revision,omitempty"`
	SecurityType   string `json:"securityType,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	Size           int    `json:"size,omitempty"`
	UpdatedBy      string `json:"updatedBy,omitempty"`
}

// KeyEndpoint represents the REST resource for managing a Key.
const KeyEndpoint = "/key"

// A KeyResource provides API to manage Keys configuration.
type KeyResource struct {
	c *f5.Client
}

// ListAll lists all the Key configurations.
func (pr *KeyResource) ListAll() (*KeyList, error) {
	var list KeyList
	if err := pr.c.ReadQuery(BasePath+KeyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Key configuration identified by id.
func (pr *KeyResource) Get(id string) (*Key, error) {
	var item Key
	if err := pr.c.ReadQuery(BasePath+KeyEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Key configuration.
func (pr *KeyResource) Create(item Key) error {
	if err := pr.c.ModQuery("POST", BasePath+KeyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Key configuration identified by id.
func (pr *KeyResource) Edit(id string, item Key) error {
	if err := pr.c.ModQuery("PUT", BasePath+KeyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Key configuration identified by id.
func (pr *KeyResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+KeyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
