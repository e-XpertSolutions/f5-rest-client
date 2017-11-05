// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideipSrvEndpoint represents the REST resource for managing WideipSrv.
const WideipSrvEndpoint = "/wideip/srv"

// WideipSrvResource provides an API to manage WideipSrv configurations.
type WideipSrvResource struct {
	c *f5.Client
}

// ListAll  lists all the WideipSrv configurations.
func (r *WideipSrvResource) ListAll() (*WideipList, error) {
	var list WideipList
	if err := r.c.ReadQuery(BasePath+WideipSrvEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideipSrv configuration identified by id.
func (r *WideipSrvResource) Get(id string) (*Wideip, error) {
	var item Wideip
	if err := r.c.ReadQuery(BasePath+WideipSrvEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideipSrv configuration.
func (r *WideipSrvResource) Create(item Wideip) error {
	if err := r.c.ModQuery("POST", BasePath+WideipSrvEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideipSrv configuration identified by id.
func (r *WideipSrvResource) Edit(id string, item Wideip) error {
	if err := r.c.ModQuery("PUT", BasePath+WideipSrvEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideipSrv configuration identified by id.
func (r *WideipSrvResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideipSrvEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
