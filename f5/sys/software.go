// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"fmt"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// SoftwareConfigList holds a list of Software configuration.
type SoftwareConfigList struct {
	Items    []SoftwareConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// SoftwareConfig holds the configuration of a single Software.
type SoftwareConfig struct {
	Reference struct {
		Link string `json:"link"`
	} `json:"reference"`
}

// SoftwareEndpoint represents the REST resource for managing Software.
const SoftwareEndpoint = "/software"

// SoftwareResource provides an API to manage Software configurations.
type SoftwareResource struct {
	c *f5.Client
}

// Install installs a software image on a specified volume.
func (r *SoftwareResource) Install(imageName, volumeName string) error {
	data := struct {
		Command string `json:"command"`
		Name    string `json:"name"`
		Volume  string `json:"volume"`
	}{
		Command: "install",
		Name:    imageName,
		Volume:  volumeName,
	}
	if err := r.c.ModQuery("POST", BasePath+SoftwareEndpoint, &data); err != nil {
		return fmt.Errorf("install software image command failed: %v", err)
	}
	return nil
}
