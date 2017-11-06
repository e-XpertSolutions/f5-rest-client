// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolNAPTREndpoint represents the REST resource for managing PoolNAPTR.
const PoolNAPTREndpoint = "/pool/naptr"

// PoolNAPTRResource provides an API to manage PoolNAPTR configurations.
type PoolNAPTRResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolNAPTR configurations.
func (r *PoolNAPTRResource) ListAll() (*PoolList, error) {
	var list PoolList
	if err := r.c.ReadQuery(BasePath+PoolNAPTREndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolNAPTR configuration identified by id.
func (r *PoolNAPTRResource) Get(id string) (*Pool, error) {
	var item Pool
	if err := r.c.ReadQuery(BasePath+PoolNAPTREndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetNAPTRMembers  lists all the PoolMembers configurations.
func (r *PoolAResource) GetNAPTRMembers(id string) (*PoolMembersList, error) {
	var list PoolMembersList
	if err := r.c.ReadQuery(BasePath+PoolNAPTREndpoint+"/"+id+"/members", &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Create a new PoolNAPTR configuration.
func (r *PoolNAPTRResource) Create(item Pool) error {
	if err := r.c.ModQuery("POST", BasePath+PoolNAPTREndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolNAPTR configuration identified by id.
func (r *PoolNAPTRResource) Edit(id string, item Pool) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolNAPTREndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolNAPTR configuration identified by id.
func (r *PoolNAPTRResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolNAPTREndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}

func (pr *PoolNAPTRResource) ShowNAPTRStats(id string) (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolNAPTREndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (pr *PoolNAPTRResource) ShowAllNAPTRStats() (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolNAPTREndpoint+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
