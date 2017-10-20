// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

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
const FileSSLKeyEndpoint = "/file/ssl-key"

// FileSSLKeyResource provides an API to manage FileSSLKey configurations.
type FileSSLKeyResource struct {
	c *f5.Client
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
func (r *FileSSLKeyResource) Create(name, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to gather information about '%s': %v", path, err)
	}
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read file from path: %v", err)
	}
	defer f.Close()

	if _, err = r.c.UploadFile(f, filepath.Base(path), info.Size()); err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}

	data := map[string]string{
		"name":        name,
		"source-path": "file://localhost/var/config/rest/downloads/" + filepath.Base(path),
	}
	if err := r.c.ModQuery("POST", BasePath+FileSSLKeyEndpoint, data); err != nil {
		return fmt.Errorf("failed to create FileSSLCRL configuration: %v", err)
	}

	return nil
}

// Edit a FileSSLKey configuration identified by id.
func (r *FileSSLKeyResource) Edit(id, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to gather information about '%s': %v", path, err)
	}
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read file from path: %v", err)
	}
	defer f.Close()

	if _, err = r.c.UploadFile(f, filepath.Base(path), info.Size()); err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}

	data := map[string]string{
		"source-path": "file://localhost/var/config/rest/downloads/" + filepath.Base(path),
	}
	if err := r.c.ModQuery("PUT", BasePath+FileSSLKeyEndpoint+"/"+id, data); err != nil {
		return fmt.Errorf("failed to create FileSSLCRL configuration: %v", err)
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
