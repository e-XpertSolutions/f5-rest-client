// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DNSConfig holds the configuration of a single DNS.
type DNSConfig struct {
	Kind         string   `json:"kind,omitempty"`
	SelfLink     string   `json:"selfLink,omitempty"`
	Description  string   `json:"description,omitempty"`
	NameServers  []string `json:"nameServers,omitempty"`
	NumberOfDots int      `json:"numberOfDots,omitempty"`
	Search       []string `json:"search,omitempty"`
}

// DNSEndpoint represents the REST resource for managing DNS.
const DNSEndpoint = "/dns"

// DNSResource provides an API to manage DNS configurations.
type DNSResource struct {
	c *f5.Client
}

// Show display the current DNS configurations.
func (r *DNSResource) Show() (*DNSConfig, error) {
	var list DNSConfig
	if err := r.c.ReadQuery(BasePath+DNSEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *DNSResource) AddNameServers(ns ...string) error {
	if len(ns) == 0 {
		return nil
	}
	var item DNSConfig
	if err := r.c.ReadQuery(BasePath+DNSEndpoint, &item); err != nil {
		return err
	}

	item.NameServers = append(item.NameServers, ns...)

	if err := r.c.ModQuery("PUT", BasePath+DNSEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit the DNS configuration
func (r *DNSResource) Edit(item DNSConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DNSEndpoint, item); err != nil {
		return err
	}
	return nil
}
