// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"errors"
	"io"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// CryptoCRLConfigList holds a list of CryptoCRL configuration.
type CryptoCRLConfigList struct {
	Items    []CryptoCRLConfig `json:"items"`
	Kind     string            `json:"kind"`
	SelfLink string            `json:"selflink"`
}

// CryptoCRLConfig holds the configuration of a single CryptoCRL.
type CryptoCRLConfig struct {
	Name       string `json:"name,omitempty"`
	FullPath   string `json:"fullPath,omitempty"`
	Generation int64  `json:"generation,omitempty"`
}

// CryptoCRLEndpoint represents the REST resource for managing CryptoCRL.
const CryptoCRLEndpoint = "/crypto/crl"

// CryptoCRLResource provides an API to manage CryptoCRL configurations.
type CryptoCRLResource struct {
	c *f5.Client
}

// ListAll  lists all the CryptoCRL configurations.
func (r *CryptoCRLResource) ListAll() (*CryptoCRLConfigList, error) {
	var list CryptoCRLConfigList
	if err := r.c.ReadQuery(BasePath+CryptoCRLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single CryptoCRL configuration identified by id.
func (r *CryptoCRLResource) Get(id string) (*CryptoCRLConfig, error) {
	var item CryptoCRLConfig
	if err := r.c.ReadQuery(BasePath+CryptoCRLEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *CryptoCRLResource) Create(name string, item CryptoCRLConfig) error {
	return errors.New("not yet implemented")
}

// CreateFromFile upload a CRL file in PEM format and create a new CRL entry.
func (r *CryptoCRLResource) CreateFromFile(name string, crlPEMFile io.Reader, filesize int64) error {
	uploadResp, err := r.c.UploadFile(crlPEMFile, name+".crl", filesize)
	if err != nil {
		return err
	}
	data := map[string]string{
		"command":         "install",
		"name":            name,
		"from-local-file": uploadResp.LocalFilePath,
	}
	if err := r.c.ModQuery("POST", BasePath+CryptoCRLEndpoint, data); err != nil {
		return err
	}
	return nil
}

// Edit a CryptoCRL configuration identified by id.
func (r *CryptoCRLResource) Edit(id string, item CryptoCRLConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoCRLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single CryptoCRL configuration identified by id.
func (r *CryptoCRLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoCRLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
