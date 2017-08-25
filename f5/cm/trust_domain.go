// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// A TrustDomainList holds a list of TrustDomain.
type TrustDomainList struct {
	Items    []TrustDomain `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selfLink,omitempty"`
}

// A TrustDomain hold the configuration for a TrustDomain.
type TrustDomain struct {
	CaCert                string `json:"caCert,omitempty"`
	CaCertBundle          string `json:"caCertBundle,omitempty"`
	CaCertBundleReference struct {
		Link string `json:"link,omitempty"`
	} `json:"caCertBundleReference,omitempty"`
	CaCertReference struct {
		Link string `json:"link,omitempty"`
	} `json:"caCertReference,omitempty"`
	CaDevices          []string `json:"caDevices,omitempty"`
	CaDevicesReference []struct {
		Link string `json:"link,omitempty"`
	} `json:"caDevicesReference,omitempty"`
	CaKey          string `json:"caKey,omitempty"`
	CaKeyReference struct {
		Link string `json:"link,omitempty"`
	} `json:"caKeyReference,omitempty"`
	FullPath            string `json:"fullPath,omitempty"`
	Generation          int    `json:"generation,omitempty"`
	Kind                string `json:"kind,omitempty"`
	Name                string `json:"name,omitempty"`
	SelfLink            string `json:"selfLink,omitempty"`
	Status              string `json:"status,omitempty"`
	TrustGroup          string `json:"trustGroup,omitempty"`
	TrustGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"trustGroupReference,omitempty"`
}

// TrustDomainEndpoint represents the REST resource for managing a TrustDomain.
const TrustDomainEndpoint = "/trust-domain"

// A TrustDomainResource provides API to manage TrustDomains configuration.
type TrustDomainResource struct {
	c *f5.Client
}

// ListAll lists all the TrustDomain configurations.
func (pr *TrustDomainResource) ListAll() (*TrustDomainList, error) {
	var list TrustDomainList
	if err := pr.c.ReadQuery(BasePath+TrustDomainEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single TrustDomain configuration identified by id.
func (pr *TrustDomainResource) Get(id string) (*TrustDomain, error) {
	var item TrustDomain
	if err := pr.c.ReadQuery(BasePath+TrustDomainEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new TrustDomain configuration.
func (pr *TrustDomainResource) Create(item TrustDomain) error {
	if err := pr.c.ModQuery("POST", BasePath+TrustDomainEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a TrustDomain configuration identified by id.
func (pr *TrustDomainResource) Edit(id string, item TrustDomain) error {
	if err := pr.c.ModQuery("PUT", BasePath+TrustDomainEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single TrustDomain configuration identified by id.
func (pr *TrustDomainResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+TrustDomainEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
