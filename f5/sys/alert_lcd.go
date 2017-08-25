// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AlertLCDConfigList holds a list of AlertLCD configuration.
type AlertLCDConfigList struct {
	Items    []AlertLCDConfig `json:"items"`
	Kind     string           `json:"kind"`
	SelfLink string           `json:"selflink"`
}

// AlertLCDConfig holds the configuration of a single AlertLCD.
type AlertLCDConfig struct {
}

// AlertLCDEndpoint represents the REST resource for managing AlertLCD.
const AlertLCDEndpoint = "/alert/lcd"

// AlertLCDResource provides an API to manage AlertLCD configurations.
type AlertLCDResource struct {
	c *f5.Client
}

// ListAll  lists all the AlertLCD configurations.
func (r *AlertLCDResource) ListAll() (*AlertLCDConfigList, error) {
	var list AlertLCDConfigList
	if err := r.c.ReadQuery(BasePath+AlertLCDEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AlertLCD configuration identified by id.
func (r *AlertLCDResource) Get(id string) (*AlertLCDConfig, error) {
	var item AlertLCDConfig
	if err := r.c.ReadQuery(BasePath+AlertLCDEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AlertLCD configuration.
func (r *AlertLCDResource) Create(item AlertLCDConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AlertLCDEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AlertLCD configuration identified by id.
func (r *AlertLCDResource) Edit(id string, item AlertLCDConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AlertLCDEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AlertLCD configuration identified by id.
func (r *AlertLCDResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AlertLCDEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
