// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ProberPoolList holds a list of ProberPool configuration.
type ProberPoolList struct {
	Items    []ProberPool `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// ProberPool holds the configuration of a single ProberPool.
type ProberPool struct {
	Enabled           bool   `json:"enabled,omitempty"`
	FullPath          string `json:"fullPath,omitempty"`
	Generation        int    `json:"generation,omitempty"`
	Kind              string `json:"kind,omitempty"`
	LoadBalancingMode string `json:"loadBalancingMode,omitempty"`
	MembersReference  struct {
		Members         []ProberPoolMembers `json:"items,omitempty"`
		IsSubcollection bool                `json:"isSubcollection,omitempty"`
		Link            string              `json:"link,omitempty"`
	} `json:"membersReference,omitempty"`
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
	SelfLink  string `json:"selfLink,omitempty"`
}

// ProberPoolMembersList holds a list of ProberPoolMembers configuration.
type ProberPoolMembersList struct {
	Items    []ProberPoolMembers `json:"items"`
	Kind     string              `json:"kind"`
	SelfLink string              `json:"selflink"`
}

// ProberPoolMembers holds the configuration of a single ProberPoolMembers.
type ProberPoolMembers struct {
	Enabled    bool   `json:"enabled"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Order      int    `json:"order"`
	SelfLink   string `json:"selfLink"`
}

// ProberPoolMembersEndpoint represents the REST resource for managing ProberPoolMembers.
const ProberPoolMembersEndpoint = "/members"

// ProberPoolEndpoint represents the REST resource for managing ProberPool.
const ProberPoolEndpoint = "/prober-pool"

// ProberPoolResource provides an API to manage ProberPool configurations.
type ProberPoolResource struct {
	c *f5.Client
}

// ListAll  lists all the ProberPool configurations.
func (r *ProberPoolResource) ListAll() (*ProberPoolList, error) {
	var list ProberPoolList
	if err := r.c.ReadQuery(BasePath+ProberPoolEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ProberPool configuration identified by id.
func (r *ProberPoolResource) Get(id string) (*ProberPool, error) {
	var item ProberPool
	if err := r.c.ReadQuery(BasePath+ProberPoolEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetMembers  lists all the ProberPoolMembers configurations.
func (r *ProberPoolResource) GetMembers(id string) (*ProberPoolMembersList, error) {
	var list ProberPoolMembersList
	if err := r.c.ReadQuery(BasePath+ProberPoolEndpoint+"/"+id+ProberPoolMembersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Create a new ProberPool configuration.
func (r *ProberPoolResource) Create(item ProberPool) error {
	if err := r.c.ModQuery("POST", BasePath+ProberPoolEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ProberPool configuration identified by id.
func (r *ProberPoolResource) Edit(id string, item ProberPool) error {
	if err := r.c.ModQuery("PUT", BasePath+ProberPoolEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single ProberPool configuration identified by id.
func (r *ProberPoolResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ProberPoolEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
