// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// FileIFileConfigList holds a list of FileIFile configuration.
type FileIFileConfigList struct {
	Items    []FileIFileConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// FileIFileConfig holds the configuration of a single FileIFile.
type FileIFileConfig struct {
	Kind           string `json:"kind,omitempty"`
	Name           string `json:"name,omitempty"`
	Partition      string `json:"partition,omitempty"`
	FullPath       string `json:"fullPath,omitempty"`
	Generation     int    `json:"generation,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	Checksum       string `json:"checksum,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	Mode           int    `json:"mode,omitempty"`
	Revision       int    `json:"revision,omitempty"`
	Size           int    `json:"size,omitempty"`
	UpdatedBy      string `json:"updatedBy,omitempty"`
}

// FileIFileEndpoint represents the REST resource for managing FileIFile.
const FileIFileEndpoint = "/file/ifile"

// FileIFileResource provides an API to manage FileIFile configurations.
type FileIFileResource struct {
	c f5.Client
}

// ListAll  lists all the FileIFile configurations.
func (r *FileIFileResource) ListAll() (*FileIFileConfigList, error) {
	var list FileIFileConfigList
	if err := r.c.ReadQuery(BasePath+FileIFileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileIFile configuration identified by id.
func (r *FileIFileResource) Get(id string) (*FileIFileConfig, error) {
	var item FileIFileConfig
	if err := r.c.ReadQuery(BasePath+FileIFileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileIFile configuration.
func (r *FileIFileResource) Create(name, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to gather information about '%s': %v", path, err)
	}
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read file from path: %v", err)
	}
	defer f.Close()

	req, err := r.c.MakeUploadRequest(f5.UploadRESTPath+"/"+filepath.Base(path), f, info.Size())
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	resp, err := r.c.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload file '%s': %v", path, err)
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	log.Print("DEBUG resp=", string(bs))

	data := map[string]string{
		"name":        name,
		"source-path": "file://localhost/var/config/rest/downloads/" + filepath.Base(path),
	}
	if err := r.c.ModQuery("POST", BasePath+FileIFileEndpoint, data); err != nil {
		return fmt.Errorf("failed to create ifile configuration: %v", err)
	}

	return nil
}

// Edit a FileIFile configuration identified by id.
func (r *FileIFileResource) Edit(id, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to gather information about '%s': %v", path, err)
	}
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read file from path: %v", err)
	}
	defer f.Close()

	req, err := r.c.MakeUploadRequest(f5.UploadRESTPath+"/"+filepath.Base(path), f, info.Size())
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	resp, err := r.c.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload file '%s': %v", path, err)
	}
	defer resp.Body.Close()

	bs, _ := ioutil.ReadAll(resp.Body)
	log.Print("DEBUG resp=", string(bs))

	data := map[string]string{
		"source-path": "file://localhost/var/config/rest/downloads/" + filepath.Base(path),
	}
	if err := r.c.ModQuery("PUT", BasePath+FileIFileEndpoint+"/"+id, data); err != nil {
		return fmt.Errorf("failed to create ifile configuration: %v", err)
	}

	return nil
}

// Delete a single FileIFile configuration identified by id.
func (r *FileIFileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileIFileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
