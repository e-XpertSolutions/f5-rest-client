// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// Persist holds the uration of a single Persist.
type Persist struct {
	APIRawValues struct {
		APIAnonymous string `json:"apiAnonymous,omitempty"`
	} `json:"apiRawValues,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// PersistEndpoint represents the REST resource for managing Persist.
const PersistEndpoint = "/persist"

// PersistResource provides an API to manage Persist urations.
type PersistResource struct {
	c *f5.Client
}

// ListAll  lists all the Persist urations.
func (r *PersistResource) ListAll() (*Persist, error) {
	var list Persist
	if err := r.c.ReadQuery(BasePath+PersistEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
