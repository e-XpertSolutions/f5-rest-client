// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "e-xpert_solutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/gtm"

type GTM struct {
	c f5.Client
}
