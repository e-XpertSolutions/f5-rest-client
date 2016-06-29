// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileSSLKeyConfigList holds a list of FileSSLKey configuration.
type FileSSLKeyConfigList struct {
	Items    []FileSSLKeyConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// FileSSLKeyConfig holds the configuration of a single FileSSLKey.
type FileSSLKeyConfig struct {
	Checksum       string `json:"checksum"`
	CreateTime     string `json:"createTime"`
	CreatedBy      string `json:"createdBy"`
	CurveName      string `json:"curveName"`
	FullPath       string `json:"fullPath"`
	Generation     int    `json:"generation"`
	KeySize        int    `json:"keySize"`
	KeyType        string `json:"keyType"`
	Kind           string `json:"kind"`
	LastUpdateTime string `json:"lastUpdateTime"`
	Mode           int    `json:"mode"`
	Name           string `json:"name"`
	Partition      string `json:"partition"`
	Revision       int    `json:"revision"`
	SecurityType   string `json:"securityType"`
	SelfLink       string `json:"selfLink"`
	Size           int    `json:"size"`
	SystemPath     string `json:"systemPath"`
	UpdatedBy      string `json:"updatedBy"`
}

// FileSSLKeyEndpoint represents the REST resource for managing FileSSLKey.
const FileSSLKeyEndpoint = "/tm/sys/file/ssl-key"

// FileSSLKeyResource provides an API to manage FileSSLKey configurations.
type FileSSLKeyResource struct {
	c f5.Client
}

// ListAll  lists all the FileSSLKey configurations.
func (r *FileSSLKeyResource) ListAll() (*FileSSLKeyConfigList, error) {
	var list FileSSLKeyConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLKeyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSSLKey configuration identified by id.
func (r *FileSSLKeyResource) Get(id string) (*FileSSLKeyConfig, error) {
	var item FileSSLKeyConfig
	if err := r.c.ReadQuery(BasePath+FileSSLKeyEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSSLKey configuration.
func (r *FileSSLKeyResource) Create(item FileSSLKeyConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileSSLKeyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileSSLKey configuration identified by id.
func (r *FileSSLKeyResource) Edit(id string, item FileSSLKeyConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileSSLKeyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileSSLKey configuration identified by id.
func (r *FileSSLKeyResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSSLKeyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
