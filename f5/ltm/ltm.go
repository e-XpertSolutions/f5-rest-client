// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/ltm"

type LTM struct {
	c f5.Client

	virtual VirtualResource
}

func New(c f5.Client) LTM {
	return LTM{
		c:       c,
		virtual: VirtualResource{c: c},
	}
}

func (ltm LTM) Virtual() VirtualResource {
	return ltm.virtual
}
