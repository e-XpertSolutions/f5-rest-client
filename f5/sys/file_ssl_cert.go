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

// FileSSLCertConfigList holds a list of FileSSLCert configuration.
type FileSSLCertConfigList struct {
	Items    []FileSSLCertConfig `json:"items,omitempty"`
	Kind     string              `json:"kind,omitempty"`
	SelfLink string              `json:"selflink,omitempty"`
}

// FileSSLCertConfig holds the configuration of a single FileSSLCert.
type FileSSLCertConfig struct {
	BundleCertificatesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"bundleCertificatesReference,omitempty"`
	CertificateKeyCurveName string `json:"certificateKeyCurveName,omitempty"`
	CertificateKeySize      int    `json:"certificateKeySize,omitempty"`
	Checksum                string `json:"checksum,omitempty"`
	CreateTime              string `json:"createTime,omitempty"`
	CreatedBy               string `json:"createdBy,omitempty"`
	ExpirationDate          int64  `json:"expirationDate,omitempty"`
	ExpirationString        string `json:"expirationString,omitempty"`
	FullPath                string `json:"fullPath,omitempty"`
	Generation              int    `json:"generation,omitempty"`
	IsBundle                string `json:"isBundle,omitempty"`
	Issuer                  string `json:"issuer,omitempty"`
	KeyType                 string `json:"keyType,omitempty"`
	Kind                    string `json:"kind,omitempty"`
	LastUpdateTime          string `json:"lastUpdateTime,omitempty"`
	Mode                    int    `json:"mode,omitempty"`
	Name                    string `json:"name,omitempty"`
	Partition               string `json:"partition,omitempty"`
	Revision                int    `json:"revision,omitempty"`
	SelfLink                string `json:"selfLink,omitempty"`
	SerialNumber            string `json:"serialNumber,omitempty"`
	Size                    int    `json:"size,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	SystemPath              string `json:"systemPath,omitempty"`
	UpdatedBy               string `json:"updatedBy,omitempty"`
	Version                 int    `json:"version,omitempty"`
}

// FileSSLCertEndpoint represents the REST resource for managing FileSSLCert.
const FileSSLCertEndpoint = "/file/ssl-cert"

// FileSSLCertResource provides an API to manage FileSSLCert configurations.
type FileSSLCertResource struct {
	c *f5.Client
}

// ListAll  lists all the FileSSLCert configurations.
func (r *FileSSLCertResource) ListAll() (*FileSSLCertConfigList, error) {
	var list FileSSLCertConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLCertEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// ListExpired lists all expired certificates.
func (r *FileSSLCertResource) ListExpired() (*FileSSLCertConfigList, error) {
	var list FileSSLCertConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLCertEndpoint, &list); err != nil {
		return nil, err
	}

	var expiredList []FileSSLCertConfig

	for _, cert := range list.Items {
		if time.Now().After(time.Unix(cert.ExpirationDate, 0)) {
			expiredList = append(expiredList, cert)
		}
	}

	expConfigList := FileSSLCertConfigList{
		Items:    expiredList,
		Kind:     list.Kind,
		SelfLink: list.SelfLink,
	}

	return &expConfigList, nil
}

// ListExpiring lists all expiring certificates.
func (r *FileSSLCertResource) ListExpiring(sec int64) (*FileSSLCertConfigList, error) {
	var list FileSSLCertConfigList
	if err := r.c.ReadQuery(BasePath+FileSSLCertEndpoint, &list); err != nil {
		return nil, err
	}

	var expiringList []FileSSLCertConfig

	for _, cert := range list.Items {
		if time.Now().After(time.Unix(cert.ExpirationDate-sec, 0)) && time.Now().Before(time.Unix(cert.ExpirationDate, 0)) {
			expiringList = append(expiringList, cert)
		}
	}

	expConfigList := FileSSLCertConfigList{
		Items:    expiringList,
		Kind:     list.Kind,
		SelfLink: list.SelfLink,
	}

	return &expConfigList, nil
}

// Get a single FileSSLCert configuration identified by id.
func (r *FileSSLCertResource) Get(id string) (*FileSSLCertConfig, error) {
	var item FileSSLCertConfig
	if err := r.c.ReadQuery(BasePath+FileSSLCertEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new FileSSLCert configuration.
func (r *FileSSLCertResource) Create(name, path string) error {
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
		return fmt.Errorf("failed to upload ssl certificate file: %v", err)
	}

	data := map[string]string{
		"name":        name,
		"source-path": "file://localhost/var/config/rest/downloads/" + filepath.Base(path),
	}
	if err := r.c.ModQuery("POST", BasePath+FileSSLCertEndpoint, data); err != nil {
		return fmt.Errorf("failed to create FileSSLCert configuration: %v", err)
	}

	return nil
}

// Edit a FileSSLCert configuration identified by id.
func (r *FileSSLCertResource) Edit(id, path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("failed to gather information about '%s': %v", path, err)
	}
	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read file from path: %v", err)
	}
	defer f.Close()

	if _, err := r.c.UploadFile(f, filepath.Base(path), info.Size()); err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}

	data := map[string]string{
		"source-path": "file://localhost/var/config/rest/downloads/" + filepath.Base(path),
	}
	if err := r.c.ModQuery("PUT", BasePath+FileSSLCertEndpoint+"/"+id, data); err != nil {
		return fmt.Errorf("failed to create FileSSLCert configuration: %v", err)
	}

	return nil
}

// CreateFromFile uploads a CRL file in PEM and create or update its value.
func (r *FileSSLCertResource) CreateFromFile(name string, certPEMFile io.Reader, filesize int64) error {
	uploadResp, err := r.c.UploadFile(certPEMFile, name+".crt", filesize)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	data := map[string]string{
		"name":        name + ".crt",
		"source-path": "file:" + uploadResp.LocalFilePath,
	}
	if err := r.c.ModQuery("POST", BasePath+FileSSLCertEndpoint, data); err != nil {
		return fmt.Errorf("failed to import certificate file: %v", err)
	}
	return nil
}

func (r *FileSSLCertResource) EditFromFile(name string, certPEMFile io.Reader, filesize int64) error {
	uploadResp, err := r.c.UploadFile(certPEMFile, name+".crt", filesize)
	if err != nil {
		return fmt.Errorf("failed to create upload request: %v", err)
	}
	data := map[string]string{
		"name":        name + ".crt",
		"source-path": "file:" + uploadResp.LocalFilePath,
	}
	if err := r.c.ModQuery("PUT", BasePath+FileSSLCertEndpoint+"/"+name+".crt", data); err != nil {
		return fmt.Errorf("failed to update imported crl file: %v", err)
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
