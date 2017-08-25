// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// NTPConfig holds the configuration of a single NTP.
type NTPConfig struct {
	Kind              string `json:"kind,omitempty"`
	RestrictReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"restrictReference,omitempty"`
	SelfLink string   `json:"selfLink,omitempty"`
	Servers  []string `json:"servers,omitempty"`
	Timezone string   `json:"timezone,omitempty"`
}

// NTPEndpoint represents the REST resource for managing NTP.
const NTPEndpoint = "/ntp"

// NTPResource provides an API to manage NTP configurations.
type NTPResource struct {
	c *f5.Client
}

// ListAll  lists all the NTP configurations.
func (r *NTPResource) Show() (*NTPConfig, error) {
	var list NTPConfig
	if err := r.c.ReadQuery(BasePath+NTPEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Edit a NTP configuration identified by id.
func (r *NTPResource) Edit(id string, item NTPConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+NTPEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *NTPResource) AddServers(rs ...string) error {
	if len(rs) == 0 {
		return nil
	}
	var item NTPConfig
	if err := r.c.ReadQuery(BasePath+NTPEndpoint, &item); err != nil {
		return err
	}

	item.Servers = append(item.Servers, rs...)

	if err := r.c.ModQuery("PUT", BasePath+NTPEndpoint, item); err != nil {
		return err
	}
	return nil
}
