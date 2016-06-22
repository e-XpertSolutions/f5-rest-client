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

	virtual VirtualResource
}

// New creates a new LTM client.
func New(c f5.Client) LTM {
	return LTM{
		c:       c,
		virtual: VirtualResource{c: c},
	}
}

// Virtual returns a VirtualResource configured to query tm/ltm/virtual API.virtual API.
func (ltm LTM) Virtual() *VirtualResource {
	return &ltm.virtual
}
