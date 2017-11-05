// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideipList holds a list of WideipA configuration.
type WideipList struct {
	Items    []Wideip `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

// Wideip holds the configuration of a single WideipA.
type Wideip struct {
	Enabled              bool   `json:"enabled,omitempty"`
	FailureRcode         string `json:"failureRcode,omitempty"`
	FailureRcodeResponse string `json:"failureRcodeResponse,omitempty"`
	FailureRcodeTTL      int    `json:"failureRcodeTtl,omitempty"`
	FullPath             string `json:"fullPath,omitempty"`
	Generation           int    `json:"generation,omitempty"`
	Kind                 string `json:"kind,omitempty"`
	LastResortPool       string `json:"lastResortPool,omitempty"`
	MinimalResponse      string `json:"minimalResponse,omitempty"`
	Name                 string `json:"name,omitempty"`
	Partition            string `json:"partition,omitempty"`
	PersistCidrIpv4      int    `json:"persistCidrIpv4,omitempty"`
	PersistCidrIpv6      int    `json:"persistCidrIpv6,omitempty"`
	Persistence          string `json:"persistence,omitempty"`
	PoolLbMode           string `json:"poolLbMode,omitempty"`
	Pools                []struct {
		Name          string `json:"name,omitempty"`
		NameReference struct {
			Link string `json:"link,omitempty"`
		} `json:"nameReference,omitempty"`
		Order     int    `json:"order,omitempty"`
		Partition string `json:"partition,omitempty"`
		Ratio     int    `json:"ratio,omitempty"`
	} `json:"pools,omitempty"`
	SelfLink       string `json:"selfLink,omitempty"`
	TTLPersistence int    `json:"ttlPersistence,omitempty"`
}

// WideipAEndpoint represents the REST resource for managing WideipA.
const WideipAEndpoint = "/wideip/a"

// WideipAResource provides an API to manage WideipA configurations.
type WideipAResource struct {
	c *f5.Client
}

// ListAll  lists all the WideipA configurations.
func (r *WideipAResource) ListAll() (*WideipList, error) {
	var list WideipList
	if err := r.c.ReadQuery(BasePath+WideipAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideipA configuration identified by id.
func (r *WideipAResource) Get(id string) (*Wideip, error) {
	var item Wideip
	if err := r.c.ReadQuery(BasePath+WideipAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideipA configuration.
func (r *WideipAResource) Create(item Wideip) error {
	if err := r.c.ModQuery("POST", BasePath+WideipAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideipA configuration identified by id.
func (r *WideipAResource) Edit(id string, item Wideip) error {
	if err := r.c.ModQuery("PUT", BasePath+WideipAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideipA configuration identified by id.
func (r *WideipAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideipAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
