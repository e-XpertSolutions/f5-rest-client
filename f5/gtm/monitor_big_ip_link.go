// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// MonitorBigIPLinkList holds a list of MonitorBigIPLink uration.
type MonitorBigIPLinkList struct {
	Items    []MonitorBigIPLink `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty"`
	SelfLink string             `json:"selflink,omitempty"`
}

// MonitorBigIPLink holds the uration of a single MonitorBigIPLink.
type MonitorBigIPLink struct {
	Destination        string `json:"destination,omitempty"`
	FullPath           string `json:"fullPath,omitempty"`
	Generation         int    `json:"generation,omitempty"`
	IgnoreDownResponse string `json:"ignoreDownResponse,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Kind               string `json:"kind,omitempty"`
	Name               string `json:"name,omitempty"`
	Partition          string `json:"partition,omitempty"`
	SelfLink           string `json:"selfLink,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`
}

// MonitorBigIPLinkEndpoint represents the REST resource for managing MonitorBigIPLink.
const MonitorBigIPLinkEndpoint = "/monitor/bigip-link"

// MonitorBigIPLinkResource provides an API to manage MonitorBigIPLink urations.
type MonitorBigIPLinkResource struct {
	c *f5.Client
}

// ListAll  lists all the MonitorBigIPLink urations.
func (r *MonitorBigIPLinkResource) ListAll() (*MonitorBigIPLinkList, error) {
	var list MonitorBigIPLinkList
	if err := r.c.ReadQuery(BasePath+MonitorBigIPLinkEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single MonitorBigIPLink uration identified by id.
func (r *MonitorBigIPLinkResource) Get(id string) (*MonitorBigIPLink, error) {
	var item MonitorBigIPLink
	if err := r.c.ReadQuery(BasePath+MonitorBigIPLinkEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new MonitorBigIPLink uration.
func (r *MonitorBigIPLinkResource) Create(item MonitorBigIPLink) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorBigIPLinkEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a MonitorBigIPLink uration identified by id.
func (r *MonitorBigIPLinkResource) Edit(id string, item MonitorBigIPLink) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorBigIPLinkEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single MonitorBigIPLink uration identified by id.
func (r *MonitorBigIPLinkResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorBigIPLinkEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
