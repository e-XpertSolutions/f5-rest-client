// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A CertList holds a list of Cert.
type CertList struct {
	Items    []Cert `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// A Cert hold the configuration for a Cert.
type Cert struct {
	CertificateKeySize int    `json:"certificateKeySize,omitempty"`
	Checksum           string `json:"checksum,omitempty"`
	CreateTime         string `json:"createTime,omitempty"`
	CreatedBy          string `json:"createdBy,omitempty"`
	ExpirationDate     int    `json:"expirationDate,omitempty"`
	ExpirationString   string `json:"expirationString,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IsBundle           string `json:"isBundle,omitempty"`
	Issuer             string `json:"issuer,omitempty"`
	KeyType            string `json:"keyType,omitempty"`
	Kind               string `json:"kind,omitempty"`
	LastUpdateTime     string `json:"lastUpdateTime,omitempty"`
	Mode               int    `json:"mode,omitempty"`
	Name               string `json:"name,omitempty"`
	Revision           int    `json:"revision,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	SerialNumber       string `json:"serialNumber,omitempty"`
	Size               int    `json:"size,omitempty"`
	Subject            string `json:"subject,omitempty"`
	UpdatedBy          string `json:"updatedBy,omitempty"`
	Version            int    `json:"version,omitempty"`
}

// CertEndpoint represents the REST resource for managing a Cert.
const CertEndpoint = "/cert"

// A CertResource provides API to manage Certs configuration.
type CertResource struct {
	c *f5.Client
}

// ListAll lists all the Cert configurations.
func (pr *CertResource) ListAll() (*CertList, error) {
	var list CertList
	if err := pr.c.ReadQuery(BasePath+CertEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Cert configuration identified by id.
func (pr *CertResource) Get(id string) (*Cert, error) {
	var item Cert
	if err := pr.c.ReadQuery(BasePath+CertEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Cert configuration.
func (pr *CertResource) Create(item Cert) error {
	if err := pr.c.ModQuery("POST", BasePath+CertEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Cert configuration identified by id.
func (pr *CertResource) Edit(id string, item Cert) error {
	if err := pr.c.ModQuery("PUT", BasePath+CertEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Cert configuration identified by id.
func (pr *CertResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+CertEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
