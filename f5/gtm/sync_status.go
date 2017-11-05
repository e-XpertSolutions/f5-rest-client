// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SyncStatus holds the configuration of a single SyncStatus.
type SyncStatus struct {
	Big3dDateTime    string `json:"big3dDateTime,omitempty"`
	Generation       int    `json:"generation,omitempty"`
	LastUpdateMicros int    `json:"lastUpdateMicros,omitempty"`
	Syncing          bool   `json:"syncing,omitempty"`
}

// SyncStatusEndpoint represents the REST resource for managing SyncStatus.
const SyncStatusEndpoint = "/sync-status"

// SyncStatusResource provides an API to manage SyncStatus configurations.
type SyncStatusResource struct {
	c *f5.Client
}

// Show  lists all the SyncStatus configurations.
func (r *SyncStatusResource) Show() (*SyncStatus, error) {
	var list SyncStatus
	if err := r.c.ReadQuery(BasePath+SyncStatusEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
