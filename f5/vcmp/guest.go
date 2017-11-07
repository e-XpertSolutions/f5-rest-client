// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcmp

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GuestList holds a list of Guests.
type GuestList struct {
	Items    []Guest `json:"items,omitempty"`
	Kind     string  `json:"kind,omitempty"`
	SelfLink string  `json:"selflink,omitempty"`
}

// Guest holds the configuration of a single Guest.
type Guest struct {
	Name              string `json:"name,omitempty"`
	AllowedSlots      int    `json:"allowedSlots,omitempty"`
	AppService        string `json:"appService,omitempty"`
	AssignedSlots     int    `json:"assignedSlots,omitempty"`
	CoresPerSlot      int    `json:"coresPerSlot,omitempty"`
	Hostname          string `json:"hostname,omitempty"`
	InitialHotfix     string `json:"initialHotfix,omitempty"`
	InitialImage      string `json:"initialImage,omitempty"`
	ManagementGW      string `json:"managementGw,omitempty"`
	ManagementIP      string `json:"managementIp,omitempty"`
	ManagementNetwork string `json:"managementNetwork,omitempty"`
	MinSlots          int    `json:"minSlots,omitempty"`
	PreferredSLots    int    `json:"preferredSlots,omitempty"`
	Slots             int    `json:"slots,omitempty"`
	SSLMode           string `json:"sslMode,omitempty"`
	State             string `json:"state,omitempty"`
	SymUnitKey        string `json:"symUnitKey,omitempty"`
	TrafficProfile    string `json:"trafficProfile,omitempty"`
	VirtualDisk       string `json:"virtualDisk,omitempty"`
}

// GuestEndpoint represents the REST resource for managing Guest.
const GuestEndpoint = "/guest"

// GuestResource provides an API to manage Guest configurations.
type GuestResource struct {
	c *f5.Client
}

// ListAll  lists all the Guest configurations.
func (r *GuestResource) ListAll() (*GuestList, error) {
	var list GuestList
	if err := r.c.ReadQuery(BasePath+GuestEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Guest configuration identified by id.
func (r *GuestResource) Get(id string) (*Guest, error) {
	var item Guest
	if err := r.c.ReadQuery(BasePath+GuestEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Guest configuration.
func (r *GuestResource) Create(item Guest) error {
	if err := r.c.ModQuery("POST", BasePath+GuestEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Guest configuration identified by id.
func (r *GuestResource) Edit(id string, item Guest) error {
	if err := r.c.ModQuery("PUT", BasePath+GuestEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Guest configuration identified by id.
func (r *GuestResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+GuestEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
