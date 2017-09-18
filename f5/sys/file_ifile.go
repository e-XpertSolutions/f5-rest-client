// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"io"

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
	c *f5.Client
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
	if err := r.c.ReadQuery(BasePath+FileIFileEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// CreateFromFile uploads an iFile and create a new entry in the
// iFiles list.
func (r *FileIFileResource) CreateFromFile(filename string, file io.Reader, filesize int64) error {
	uploadResp, err := r.c.UploadFile(file, filename, filesize)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	data := map[string]string{
		"name":        filename,
		"source-path": "file:" + uploadResp.LocalFilePath,
	}
	if err := r.c.ModQuery("POST", BasePath+FileIFileEndpoint, data); err != nil {
		return fmt.Errorf("failed to import ifile: %v", err)
	}
	return nil
}

// EditFromFile uploads an iFile and update an existing entry in the
// iFiles list.
func (r *FileIFileResource) EditFromFile(filename string, file io.Reader, filesize int64) error {
	uploadResp, err := r.c.UploadFile(file, filename, filesize)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	data := map[string]string{
		"name":        filename,
		"source-path": "file:" + uploadResp.LocalFilePath,
	}
	if err := r.c.ModQuery("PUT", BasePath+FileIFileEndpoint+"/"+filename, data); err != nil {
		return fmt.Errorf("failed to update ifile: %v", err)
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
