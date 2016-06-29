// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SyslogConfigList holds a list of Syslog configuration.
type SyslogConfigList struct {
	Items    []SyslogConfig `json:"items"`
	Kind     string         `json:"kind"`
	SelfLink string         `json:"selflink"`
}

// SyslogConfig holds the configuration of a single Syslog.
type SyslogConfig struct {
}

// SyslogEndpoint represents the REST resource for managing Syslog.
const SyslogEndpoint = "/tm/sys/syslog"

// SyslogResource provides an API to manage Syslog configurations.
type SyslogResource struct {
	c f5.Client
}

// ListAll  lists all the Syslog configurations.
func (r *SyslogResource) ListAll() (*SyslogConfigList, error) {
	var list SyslogConfigList
	if err := r.c.ReadQuery(BasePath+SyslogEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Syslog configuration identified by id.
func (r *SyslogResource) Get(id string) (*SyslogConfig, error) {
	var item SyslogConfig
	if err := r.c.ReadQuery(BasePath+SyslogEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Syslog configuration.
func (r *SyslogResource) Create(item SyslogConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SyslogEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Syslog configuration identified by id.
func (r *SyslogResource) Edit(id string, item SyslogConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SyslogEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Syslog configuration identified by id.
func (r *SyslogResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SyslogEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
