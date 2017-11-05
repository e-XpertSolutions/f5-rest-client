// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ListenerList holds a list of Listener uration.
type ListenerList struct {
	Items    []Listener `json:"items,omitempty"`
	Kind     string     `json:"kind,omitempty"`
	SelfLink string     `json:"selflink,omitempty"`
}

// Listener holds the uration of a single Listener.
type Listener struct {
	ProfilesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"profilesReference,omitempty"`
	Advertise                string `json:"advertise,omitempty"`
	TranslateAddress         string `json:"translateAddress,omitempty"`
	SourceAddressTranslation struct {
		Type string `json:"type,omitempty"`
	} `json:"sourceAddressTranslation,omitempty"`
	SelfLink      string `json:"selfLink,omitempty"`
	VlansDisabled bool   `json:"vlansDisabled,omitempty"`
	Name          string `json:"name,omitempty"`
	IpProtocol    string `json:"ipProtocol,omitempty"`
	FullPath      string `json:"fullPath,omitempty"`
	SourcePort    string `json:"sourcePort,omitempty"`
	Kind          string `json:"kind,omitempty"`
	TranslatePort string `json:"translatePort,omitempty"`
	Address       string `json:"address,omitempty"`
	Generation    int    `json:"generation,omitempty"`
	Port          int    `json:"port,omitempty"`
	Mask          string `json:"mask,omitempty"`
	Enabled       bool   `json:"enabled,omitempty"`
	AutoLasthop   string `json:"autoLasthop,omitempty"`
}

// ListenerEndpoint represents the REST resource for managing Listener.
const ListenerEndpoint = "/listener"

// ListenerResource provides an API to manage Listener urations.
type ListenerResource struct {
	c *f5.Client
}

// ListAll  lists all the Listener urations.
func (r *ListenerResource) ListAll() (*ListenerList, error) {
	var list ListenerList
	if err := r.c.ReadQuery(BasePath+ListenerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Listener uration identified by id.
func (r *ListenerResource) Get(id string) (*Listener, error) {
	var item Listener
	if err := r.c.ReadQuery(BasePath+ListenerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Listener uration.
func (r *ListenerResource) Create(item Listener) error {
	if err := r.c.ModQuery("POST", BasePath+ListenerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Listener uration identified by id.
func (r *ListenerResource) Edit(id string, item Listener) error {
	if err := r.c.ModQuery("PUT", BasePath+ListenerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Listener uration identified by id.
func (r *ListenerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ListenerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
