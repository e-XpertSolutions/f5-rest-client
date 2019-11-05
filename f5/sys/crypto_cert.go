// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// CryptoCertConfigList holds a list of CryptoCert configuration.
type CryptoCertConfigList struct {
	Items    []CryptoCertConfig `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

// CryptoCertConfig holds the configuration of a single CryptoCert.
type CryptoCertConfig struct {
	APIRawValues struct {
		CertificateKeySize string `json:"certificateKeySize,omitempty"`
		Expiration         string `json:"expiration,omitempty"`
		PublicKeyType      string `json:"publicKeyType,omitempty"`
	} `json:"apiRawValues,omitempty"`
	Country      string `json:"country,omitempty"`
	CommonName   string `json:"commonName,omitempty"`
	Fingerprint  string `json:"fingerprint,omitempty"`
	FullPath     string `json:"fullPath,omitempty"`
	Generation   int    `json:"generation,omitempty"`
	Kind         string `json:"kind,omitempty"`
	Name         string `json:"name,omitempty"`
	Organization string `json:"organization,omitempty"`
	Ou           string `json:"ou,omitempty"`
	SelfLink     string `json:"selfLink,omitempty"`
}

// CryptoCertEndpoint represents the REST resource for managing CryptoCert.
const CryptoCertEndpoint = "/crypto/cert"

// CryptoCertResource provides an API to manage CryptoCert configurations.
type CryptoCertResource struct {
	c *f5.Client
}

// ListAll  lists all the CryptoCert configurations.
func (r *CryptoCertResource) ListAll() (*CryptoCertConfigList, error) {
	var list CryptoCertConfigList
	if err := r.c.ReadQuery(BasePath+CryptoCertEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single CryptoCert configuration identified by id.
func (r *CryptoCertResource) Get(id string) (*CryptoCertConfig, error) {
	var item CryptoCertConfig
	if err := r.c.ReadQuery(BasePath+CryptoCertEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new CryptoCert configuration.
func (r *CryptoCertResource) Create(item CryptoCertConfig) error {
	if err := r.c.ModQuery("POST", BasePath+CryptoCertEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a CryptoCert configuration identified by id.
func (r *CryptoCertResource) Edit(id string, item CryptoCertConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+CryptoCertEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single CryptoCert configuration identified by id.
func (r *CryptoCertResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+CryptoCertEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
