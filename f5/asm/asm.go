// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/asm"

type ASM struct {
	c *f5.Client

	policy PolicyResource
}

// New creates a new ASM client.
func New(c *f5.Client) ASM {
	return ASM{
		c:      c,
		policy: PolicyResource{c: c},
	}
}

// Policy returns a PolicyResource configured to query tm/asm/policy API.
func (asm ASM) Policy() *PolicyResource {
	return &asm.policy
}
