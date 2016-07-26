// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// LicenseConfigList holds a list of License configuration.
type LicenseConfigList struct {
	Items    []LicenseConfig `json:"items"`
	Kind     string          `json:"kind"`
	SelfLink string          `json:"selflink"`
}

// LicenseConfig holds the configuration of a single License.
type LicenseConfig struct {
	Entries  map[string]interface{} `json:"entries"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selfLink"`
}

// LicenseEndpoint represents the REST resource for managing License.
const LicenseEndpoint = "/license"

// LicenseResource provides an API to manage License configurations.
type LicenseResource struct {
	c f5.Client
}

// Get list the License configurations.
func (r *LicenseResource) Get() (*LicenseConfigList, error) {
	var list LicenseConfigList
	if err := r.c.ReadQuery(BasePath+LicenseEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
