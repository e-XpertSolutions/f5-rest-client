// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/e-XpertSolutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/net"

type NET struct {
	c *f5.Client

	inet        InetResource
	inetStats   InetStatsResource
	route       RouteResource
	vlan        VlanResource
	self        SelfResource
	routeDomain RouteDomainResource
	trunk       TrunkResource
}

// New creates a new NET client.
func New(c *f5.Client) NET {
	return NET{
		c:           c,
		inet:        InetResource{c: c},
		inetStats:   InetStatsResource{c: c},
		route:       RouteResource{c: c},
		vlan:        VlanResource{c: c},
		self:        SelfResource{c: c},
		routeDomain: RouteDomainResource{c: c},
		trunk:       TrunkResource{c: c},
	}
}

// Inet returns a InetResource configured to query tm/net/interface API.
func (net NET) Inet() *InetResource {
	return &net.inet
}

func (net NET) InetStats() *InetStatsResource {
	return &net.inetStats
}

// Route returns a RouteResource configured to query tm/net/route API.
func (net NET) Route() *RouteResource {
	return &net.route
}

// Vlan returns a VlanResource configured to query /tm/net/vlan API.
func (net NET) Vlan() *VlanResource {
	return &net.vlan
}

// Self returns a SelfResource configured to query /tm/net/self API.
func (net NET) Self() *SelfResource {
	return &net.self
}

// RouteDomain returns a RouteDomainResource configured to query /tm/net/route-domain API.
func (net NET) RouteDomain() *RouteDomainResource {
	return &net.routeDomain
}

// RouteDomain returns a RouteDomainResource configured to query /tm/net/route-domain API.
func (net NET) Trunk() *TrunkResource {
	return &net.trunk
}
