// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// WideipNaptrEndpoint represents the REST resource for managing WideipNaptr.
const WideipNaptrEndpoint = "/wideip/naptr"

// WideipNaptrResource provides an API to manage WideipNaptr configurations.
type WideipNaptrResource struct {
	c *f5.Client
}

// ListAll  lists all the WideipNaptr configurations.
func (r *WideipNaptrResource) ListAll() (*WideipList, error) {
	var list WideipList
	if err := r.c.ReadQuery(BasePath+WideipNaptrEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single WideipNaptr configuration identified by id.
func (r *WideipNaptrResource) Get(id string) (*Wideip, error) {
	var item Wideip
	if err := r.c.ReadQuery(BasePath+WideipNaptrEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new WideipNaptr configuration.
func (r *WideipNaptrResource) Create(item Wideip) error {
	if err := r.c.ModQuery("POST", BasePath+WideipNaptrEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a WideipNaptr configuration identified by id.
func (r *WideipNaptrResource) Edit(id string, item Wideip) error {
	if err := r.c.ModQuery("PUT", BasePath+WideipNaptrEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single WideipNaptr configuration identified by id.
func (r *WideipNaptrResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+WideipNaptrEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
