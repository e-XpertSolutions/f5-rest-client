// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolCNAMEEndpoint represents the REST resource for managing PoolCNAME.
const PoolCNAMEEndpoint = "/pool/cname"

// PoolCNAMEResource provides an API to manage PoolCNAME configurations.
type PoolCNAMEResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolCNAME configurations.
func (r *PoolCNAMEResource) ListAll() (*PoolList, error) {
	var list PoolList
	if err := r.c.ReadQuery(BasePath+PoolCNAMEEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single PoolCNAME configuration identified by id.
func (r *PoolCNAMEResource) Get(id string) (*Pool, error) {
	var item Pool
	if err := r.c.ReadQuery(BasePath+PoolCNAMEEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetCNAMEMembers  lists all the PoolMembers configurations.
func (r *PoolAResource) GetCNAMEMembers(id string) (*PoolMembersList, error) {
	var list PoolMembersList
	if err := r.c.ReadQuery(BasePath+PoolCNAMEEndpoint+"/"+id+"/members", &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Create a new PoolCNAME configuration.
func (r *PoolCNAMEResource) Create(item Pool) error {
	if err := r.c.ModQuery("POST", BasePath+PoolCNAMEEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a PoolCNAME configuration identified by id.
func (r *PoolCNAMEResource) Edit(id string, item Pool) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolCNAMEEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single PoolCNAME configuration identified by id.
func (r *PoolCNAMEResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolCNAMEEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}

func (pr *PoolCNAMEResource) ShowCNAMEStats(id string) (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolCNAMEEndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (pr *PoolCNAMEResource) ShowAllCNAMEStats() (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolCNAMEEndpoint+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
