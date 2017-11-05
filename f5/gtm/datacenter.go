// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DatacenterList holds a list of Datacenter configuration.
type DatacenterList struct {
	Items    []Datacenter `json:"items,omitempty"`
	Kind     string       `json:"kind,omitempty"`
	SelfLink string       `json:"selflink,omitempty"`
}

// Datacenter holds the configuration of a single Datacenter.
type Datacenter struct {
	Contact          string `json:"contact,omitempty"`
	Description      string `json:"description,omitempty"`
	Enabled          bool   `json:"enabled,omitempty"`
	FullPath         string `json:"fullPath,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	Kind             string `json:"kind,omitempty"`
	Location         string `json:"location,omitempty"`
	Name             string `json:"name,omitempty"`
	Partition        string `json:"partition,omitempty"`
	ProberFallback   string `json:"proberFallback,omitempty"`
	ProberPreference string `json:"proberPreference,omitempty"`
	SelfLink         string `json:"selfLink,omitempty"`
}

// DatacenterEndpoint represents the REST resource for managing Datacenter.
const DatacenterEndpoint = "/datacenter"

// DatacenterResource provides an API to manage Datacenter configurations.
type DatacenterResource struct {
	c *f5.Client
}

// ListAll  lists all the Datacenter configurations.
func (r *DatacenterResource) ListAll() (*DatacenterList, error) {
	var list DatacenterList
	if err := r.c.ReadQuery(BasePath+DatacenterEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Datacenter configuration identified by id.
func (r *DatacenterResource) Get(id string) (*Datacenter, error) {
	var item Datacenter
	if err := r.c.ReadQuery(BasePath+DatacenterEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Datacenter configuration.
func (r *DatacenterResource) Create(item Datacenter) error {
	if err := r.c.ModQuery("POST", BasePath+DatacenterEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Datacenter configuration identified by id.
func (r *DatacenterResource) Edit(id string, item Datacenter) error {
	if err := r.c.ModQuery("PUT", BasePath+DatacenterEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Datacenter configuration identified by id.
func (r *DatacenterResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DatacenterEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
