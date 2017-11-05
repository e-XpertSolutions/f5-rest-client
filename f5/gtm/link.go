// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LinkList holds a list of Link configuration.
type LinkList struct {
	Items    []Link `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// Link holds the configuration of a single Link.
type Link struct {
	Datacenter          string `json:"datacenter,omitempty"`
	DatacenterReference struct {
		Link string `json:"link,omitempty"`
	} `json:"datacenterReference,omitempty"`
	DuplexBilling             string `json:"duplexBilling,omitempty"`
	Enabled                   bool   `json:"enabled,omitempty"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	LimitMaxInboundBps        int    `json:"limitMaxInboundBps,omitempty"`
	LimitMaxInboundBpsStatus  string `json:"limitMaxInboundBpsStatus,omitempty"`
	LimitMaxOutboundBps       int    `json:"limitMaxOutboundBps,omitempty"`
	LimitMaxOutboundBpsStatus string `json:"limitMaxOutboundBpsStatus,omitempty"`
	LimitMaxTotalBps          int    `json:"limitMaxTotalBps,omitempty"`
	LimitMaxTotalBpsStatus    string `json:"limitMaxTotalBpsStatus,omitempty"`
	LinkRatio                 int    `json:"linkRatio,omitempty"`
	Monitor                   string `json:"monitor,omitempty"`
	Name                      string `json:"name,omitempty"`
	Partition                 string `json:"partition,omitempty"`
	PrepaidSegment            int    `json:"prepaidSegment,omitempty"`
	RouterAddresses           []struct {
		Name        string `json:"name,omitempty"`
		Translation string `json:"translation,omitempty"`
	} `json:"routerAddresses,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	UplinkAddress string `json:"uplinkAddress,omitempty"`
	Weighting     string `json:"weighting,omitempty"`
}

// LinkEndpoint represents the REST resource for managing Link.
const LinkEndpoint = "/link"

// LinkResource provides an API to manage Link configurations.
type LinkResource struct {
	c *f5.Client
}

// ListAll  lists all the Link configurations.
func (r *LinkResource) ListAll() (*LinkList, error) {
	var list LinkList
	if err := r.c.ReadQuery(BasePath+LinkEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Link configuration identified by id.
func (r *LinkResource) Get(id string) (*Link, error) {
	var item Link
	if err := r.c.ReadQuery(BasePath+LinkEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Link configuration.
func (r *LinkResource) Create(item Link) error {
	if err := r.c.ModQuery("POST", BasePath+LinkEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Link configuration identified by id.
func (r *LinkResource) Edit(id string, item Link) error {
	if err := r.c.ModQuery("PUT", BasePath+LinkEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Link configuration identified by id.
func (r *LinkResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+LinkEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
