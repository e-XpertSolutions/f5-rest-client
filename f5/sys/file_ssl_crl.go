// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// FileSSLCRLConfigList holds a list of FileSSLCRL configuration.
type FileSSLCRLConfigList struct {
	Items    []FileSSLCRLConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// FileSSLCRLConfig holds the configuration of a single FileSSLCRL.
type FileSSLCRLConfig struct {
	Name           string    `json:"name,omitempty"`
	Partition      string    `json:"partition,omitempty"`
	FullPath       string    `json:"fullPath,omitempty"`
	Generator      int64     `json:"generator,omitempty"`
	Checksum       string    `json:"checksum,omitempty"`
	CreateTime     time.Time `json:"createTime,omitempty"`
	CreatedBy      string    `json:"createdBy,omitempty"`
	LastUpdateTime time.Time `json:"lastUpdateTime,omitempty"`
	Mode           int64     `json:"mode,omitempty"`
	Revision       int64     `json:"revision,omitempty"`
	Size           int64     `json:"size,omitempty"`
	UpdatedBy      string    `json:"updatedBy,omitempty"`
}

// FileSSLCRLEndpoint represents the REST resource for managing FileSSLCRL.
const FileSSLCRLEndpoint = "/file/ssl-crl"

// FileSSLCRLResource provides an API to manage FileSSLCRL configurations.
type FileSSLCRLResource struct {
	c *f5.Client
}

// ListAll  lists all the FileSSLCRL configurations.
func (r *FileSSLCRLResource) ListAll() (*FileSSLCRLConfigList, error) {
	var list FileSSLCRLConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLCRLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSSLCRL configuration identified by id.
func (r *FileSSLCRLResource) Get(id string) (*FileSSLCRLConfig, error) {
	var item FileSSLCRLConfig
	if err := r.c.ReadQuery(BasePath+FileSSLCRLEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// CreateFromFile uploads a CRL file in PEM and create or update its value.
func (r *FileSSLCRLResource) CreateFromFile(name string, crlPEMFile io.Reader, filesize int64) error {
	uploadResp, err := r.c.UploadFile(crlPEMFile, name+".crl", filesize)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	data := map[string]string{
		"name":        name + ".crl",
		"source-path": "file:" + uploadResp.LocalFilePath,
	}
	if err := r.c.ModQuery("POST", BasePath+FileSSLCRLEndpoint, data); err != nil {
		return fmt.Errorf("failed to import crl file: %v", err)
	}
	return nil
}

func (r *FileSSLCRLResource) EditFromFile(name string, crlPEMFile io.Reader, filesize int64) error {
	uploadResp, err := r.c.UploadFile(crlPEMFile, name+".crl", filesize)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	data := map[string]string{
		"name":        name + ".crl",
		"source-path": "file:" + uploadResp.LocalFilePath,
	}
	if err := r.c.ModQuery("PUT", BasePath+FileSSLCRLEndpoint+"/"+name+".crl", data); err != nil {
		return fmt.Errorf("failed to update imported crl file: %v", err)
	}
	return nil
}

// Edit a FileSSLCRL configuration identified by id.
func (r *FileSSLCRLResource) Edit(id, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to gather information about '%s': %v", path, err)
	}
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read file from path: %v", err)
	}
	defer f.Close()

	uploadResp, err := r.c.UploadFile(f, filepath.Base(path), info.Size())
	if err != nil {
		return fmt.Errorf("failed to upload file %q: %v", path, err)
	}

	data := map[string]string{
		"source-path": "file:" + uploadResp.LocalFilePath,
	}
	if err := r.c.ModQuery("PUT", BasePath+FileSSLCRLEndpoint+"/"+id, data); err != nil {
		return fmt.Errorf("failed to create FileSSLCRL configuration: %v", err)
	}

	return nil
}

// Delete a single FileSSLCRL configuration identified by id.
func (r *FileSSLCRLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSSLCRLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
