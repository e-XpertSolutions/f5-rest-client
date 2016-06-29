// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// FileSSLCertConfigList holds a list of FileSSLCert configuration.
type FileSSLCertConfigList struct {
	Items    []FileSSLCertConfig `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// FileSSLCertConfig holds the configuration of a single FileSSLCert.
type FileSSLCertConfig struct {
	BundleCertificatesReference struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"bundleCertificatesReference"`
	CertificateKeyCurveName string `json:"certificateKeyCurveName"`
	CertificateKeySize      int    `json:"certificateKeySize"`
	Checksum                string `json:"checksum"`
	CreateTime              string `json:"createTime"`
	CreatedBy               string `json:"createdBy"`
	ExpirationDate          int    `json:"expirationDate"`
	ExpirationString        string `json:"expirationString"`
	FullPath                string `json:"fullPath"`
	Generation              int    `json:"generation"`
	IsBundle                string `json:"isBundle"`
	Issuer                  string `json:"issuer"`
	KeyType                 string `json:"keyType"`
	Kind                    string `json:"kind"`
	LastUpdateTime          string `json:"lastUpdateTime"`
	Mode                    int    `json:"mode"`
	Name                    string `json:"name"`
	Partition               string `json:"partition"`
	Revision                int    `json:"revision"`
	SelfLink                string `json:"selfLink"`
	SerialNumber            string `json:"serialNumber"`
	Size                    int    `json:"size"`
	Subject                 string `json:"subject"`
	SystemPath              string `json:"systemPath"`
	UpdatedBy               string `json:"updatedBy"`
	Version                 int    `json:"version"`
}

// FileSSLCertEndpoint represents the REST resource for managing FileSSLCert.
const FileSSLCertEndpoint = "/tm/sys/file/ssl-cert"

// FileSSLCertResource provides an API to manage FileSSLCert configurations.
type FileSSLCertResource struct {
	c f5.Client
}

// ListAll  lists all the FileSSLCert configurations.
func (r *FileSSLCertResource) ListAll() (*FileSSLCertConfigList, error) {
	var list FileSSLCertConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLCertEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single FileSSLCert configuration identified by id.
func (r *FileSSLCertResource) Get(id string) (*FileSSLCertConfig, error) {
	var item FileSSLCertConfig
	if err := r.c.ReadQuery(BasePath+FileSSLCertEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSSLCert configuration.
func (r *FileSSLCertResource) Create(item FileSSLCertConfig) error {
	if err := r.c.ModQuery("POST", BasePath+FileSSLCertEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a FileSSLCert configuration identified by id.
func (r *FileSSLCertResource) Edit(id string, item FileSSLCertConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+FileSSLCertEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single FileSSLCert configuration identified by id.
func (r *FileSSLCertResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+FileSSLCertEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
