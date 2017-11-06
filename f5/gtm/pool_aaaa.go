// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolAAAAEndpoint represents the REST resource for managing PoolAAAA.
const PoolAAAAEndpoint = "/pool/aaaa"

// PoolAAAAResource provides an API to manage PoolAAAA configurations.
type PoolAAAAResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolAAAA configurations.
func (r *PoolAAAAResource) ListAll() (*PoolList, error) {
	var list PoolList
	if err := r.c.ReadQuery(BasePath+PoolAAAAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolAAAA configuration identified by id.
func (r *PoolAAAAResource) Get(id string) (*Pool, error) {
	var item Pool
	if err := r.c.ReadQuery(BasePath+PoolAAAAEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetAAAAMembers  lists all the PoolMembers configurations.
func (r *PoolAResource) GetAAAAMembers(id string) (*PoolMembersList, error) {
	var list PoolMembersList
	if err := r.c.ReadQuery(BasePath+PoolAAAAEndpoint+"/"+id+"/members", &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Create a new PoolAAAA configuration.
func (r *PoolAAAAResource) Create(item Pool) error {
	if err := r.c.ModQuery("POST", BasePath+PoolAAAAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolAAAA configuration identified by id.
func (r *PoolAAAAResource) Edit(id string, item Pool) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolAAAAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolAAAA configuration identified by id.
func (r *PoolAAAAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolAAAAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}

func (pr *PoolAAAAResource) ShowAAAAStats(id string) (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolAAAAEndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (pr *PoolAAAAResource) ShowAllAAAAStats() (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolAAAAEndpoint+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
