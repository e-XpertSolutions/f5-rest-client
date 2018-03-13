// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"time"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// License holds the configuration of a single License.
type License struct {
	Kind     string `json:"kind"`
	SelfLink string `json:"selfLink"`
	Entries  map[string]struct {
		NestedStats struct {
			Entries map[string]struct {
				Description string `json:"description"`
			}
		} `json:"nestedStats"`
	} `json:"entries"`
}

func (l License) ServiceCheckDate() time.Time {
	for _, entry := range l.Entries {
		innerEntry, ok := entry.NestedStats.Entries["serviceCheckDate"]
		if !ok {
			continue
		}
		t, err := time.Parse("2006/01/02", innerEntry.Description)
		if err != nil {
			break
		}
		return t
	}
	return time.Time{}
}

func (l License) RegistrationKey() string {
	for _, entry := range l.Entries {
		innerEntry, ok := entry.NestedStats.Entries["registrationKey"]
		if !ok {
			continue
		}
		return innerEntry.Description
	}
	return ""
}

// LicenseEndpoint represents the REST resource for managing License.
const LicenseEndpoint = "/license"

// LicenseResource provides an API to manage License configurations.
type LicenseResource struct {
	c *f5.Client
}

// Get list the License configurations.
func (r *LicenseResource) Get() (*License, error) {
	var l License
	if err := r.c.ReadQuery(BasePath+LicenseEndpoint, &l); err != nil {
		return nil, err
	}
	return &l, nil
}

// Activate license.
func (r *LicenseResource) Activate(registrationKey string) error {
	data := struct {
		Command         string `json:"command"`
		RegistrationKey string `json:"registrationKey"`
	}{
		Command:         "install",
		RegistrationKey: registrationKey,
	}
	return r.c.ModQuery("POST", BasePath+LicenseEndpoint, &data)
}
