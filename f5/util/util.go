// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import "github.com/e-XpertSolutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/util"

type Util struct {
	c *f5.Client

	bash BashResource
}

// New creates a new NET client.
func New(c *f5.Client) Util {
	return Util{
		c:    c,
		bash: BashResource{c: c},
	}
}

func (util Util) Bash() *BashResource {
	return &util.bash
}
