// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ltm provides a REST client for the /tm/ltm F5 BigIP API.
package ltm

import "e-xpert_solutions/f5-rest-client/f5"

// BasePath is the base path of the ltm API.
const BasePath = "mgmt/tm/ltm"

// LTM implement a REST client for the F5 BigIP LTM API.
type LTM struct {
	c f5.Client

	virtual     VirtualResource
	rule        RuleResource
	pool        PoolResource
	poolMembers PoolMembersResource
	node        NodeResource
}

// New creates a new LTM client.
func New(c f5.Client) LTM {
	return LTM{
		c:           c,
		virtual:     VirtualResource{c: c},
		rule:        RuleResource{c: c},
		pool:        PoolResource{c: c},
		poolMembers: PoolMembersResource{c: c},
		node:        NodeResource{c: c},
	}
}

// Virtual returns a VirtualResource configured to query tm/ltm/virtual API.
func (ltm LTM) Virtual() *VirtualResource {
	return &ltm.virtual
}

// Rule returns a RuleResource configured to query tm/ltm/rule API.
func (ltm LTM) Rule() *RuleResource {
	return &ltm.rule
}

// Pool returns a PoolResource configured to query /tm/ltm/pool API.
func (ltm LTM) Pool() *PoolResource {
	return &ltm.pool
}

// PoolMembers returns a PoolResource configured to query /tm/ltm/pool API.
func (ltm LTM) PoolMembers() *PoolMembersResource {
	return &ltm.poolMembers
}

// Node returns a NodeResource configured to query /tm/ltm/node API.
func (ltm LTM) Node() *NodeResource {
	return &ltm.node
}
