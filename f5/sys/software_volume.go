// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SoftwareVolumeConfigList holds a list of SoftwareVolume configuration.
type SoftwareVolumeConfigList struct {
	Items    []SoftwareVolumeConfig `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// SoftwareVolumeConfig holds the configuration of a single SoftwareVolume.
type SoftwareVolumeConfig struct {
	Active     bool   `json:"active"`
	Basebuild  string `json:"basebuild"`
	Build      string `json:"build"`
	FullPath   string `json:"fullPath"`
	Generation int    `json:"generation"`
	Kind       string `json:"kind"`
	Media      []struct {
		DefaultBootLocation bool   `json:"defaultBootLocation"`
		Media               string `json:"media"`
		Name                string `json:"name"`
		NameReference       struct {
			Link string `json:"link"`
		} `json:"nameReference"`
		Size string `json:"size"`
	} `json:"media"`
	Name     string `json:"name"`
	Product  string `json:"product"`
	SelfLink string `json:"selfLink"`
	Status   string `json:"status"`
	Version  string `json:"version"`
}

// SoftwareVolumeEndpoint represents the REST resource for managing SoftwareVolume.
const SoftwareVolumeEndpoint = "/software/volume"

// SoftwareVolumeResource provides an API to manage SoftwareVolume configurations.
type SoftwareVolumeResource struct {
	c *f5.Client
}

// ListAll  lists all the SoftwareVolume configurations.
func (r *SoftwareVolumeResource) ListAll() (*SoftwareVolumeConfigList, error) {
	var list SoftwareVolumeConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareVolumeEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SoftwareVolume configuration identified by id.
func (r *SoftwareVolumeResource) Get(id string) (*SoftwareVolumeConfig, error) {
	var item SoftwareVolumeConfig
	if err := r.c.ReadQuery(BasePath+SoftwareVolumeEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Get the running software version
func (r *SoftwareVolumeResource) GetVersion() (string, error) {
	var list SoftwareVolumeConfigList
	if err := r.c.ReadQuery(BasePath+SoftwareVolumeEndpoint, &list); err != nil {
		return "", err
	}

	var version string

	for _, vol := range list.Items {
		if vol.Active == true {
			version = vol.Version + " " + vol.Build
		}
	}

	return version, nil
}

// Create a new SoftwareVolume configuration.
func (r *SoftwareVolumeResource) Create(item SoftwareVolumeConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SoftwareVolumeEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SoftwareVolume configuration identified by id.
func (r *SoftwareVolumeResource) Edit(id string, item SoftwareVolumeConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SoftwareVolumeEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SoftwareVolume configuration identified by id.
func (r *SoftwareVolumeResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SoftwareVolumeEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
