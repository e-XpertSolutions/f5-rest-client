// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ProberPoolConfigList holds a list of ProberPool configuration.
type ProberPoolConfigList struct {
	Items    []ProberPoolConfig `json:"items"`
	Kind     string             `json:"kind"`
	SelfLink string             `json:"selflink"`
}

// ProberPoolConfig holds the configuration of a single ProberPool.
type ProberPoolConfig struct {
}

// ProberPoolEndpoint represents the REST resource for managing ProberPool.
const ProberPoolEndpoint = "/prober-pool"

// ProberPoolResource provides an API to manage ProberPool configurations.
type ProberPoolResource struct {
	c *f5.Client
}

// ListAll  lists all the ProberPool configurations.
func (r *ProberPoolResource) ListAll() (*ProberPoolConfigList, error) {
	var list ProberPoolConfigList
	if err := r.c.ReadQuery(BasePath+ProberPoolEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single ProberPool configuration identified by id.
func (r *ProberPoolResource) Get(id string) (*ProberPoolConfig, error) {
	var item ProberPoolConfig
	if err := r.c.ReadQuery(BasePath+ProberPoolEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new ProberPool configuration.
func (r *ProberPoolResource) Create(item ProberPoolConfig) error {
	if err := r.c.ModQuery("POST", BasePath+ProberPoolEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a ProberPool configuration identified by id.
func (r *ProberPoolResource) Edit(id string, item ProberPoolConfig) error {
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
