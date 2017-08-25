// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SSHSSHPluginStatsConfigList holds a list of SSHSSHPluginStats configuration.
type SSHSSHPluginStatsConfigList struct {
	Items    []SSHSSHPluginStatsConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selflink"`
}

// SSHSSHPluginStatsConfig holds the configuration of a single SSHSSHPluginStats.
type SSHSSHPluginStatsConfig struct {
}

// SSHSSHPluginStatsEndpoint represents the REST resource for managing SSHSSHPluginStats.
const SSHSSHPluginStatsEndpoint = "/ssh/sshplugin-stats"

// SSHSSHPluginStatsResource provides an API to manage SSHSSHPluginStats configurations.
type SSHSSHPluginStatsResource struct {
	c *f5.Client
}

// ListAll  lists all the SSHSSHPluginStats configurations.
func (r *SSHSSHPluginStatsResource) ListAll() (*SSHSSHPluginStatsConfigList, error) {
	var list SSHSSHPluginStatsConfigList
	if err := r.c.ReadQuery(BasePath+SSHSSHPluginStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single SSHSSHPluginStats configuration identified by id.
func (r *SSHSSHPluginStatsResource) Get(id string) (*SSHSSHPluginStatsConfig, error) {
	var item SSHSSHPluginStatsConfig
	if err := r.c.ReadQuery(BasePath+SSHSSHPluginStatsEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new SSHSSHPluginStats configuration.
func (r *SSHSSHPluginStatsResource) Create(item SSHSSHPluginStatsConfig) error {
	if err := r.c.ModQuery("POST", BasePath+SSHSSHPluginStatsEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a SSHSSHPluginStats configuration identified by id.
func (r *SSHSSHPluginStatsResource) Edit(id string, item SSHSSHPluginStatsConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SSHSSHPluginStatsEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single SSHSSHPluginStats configuration identified by id.
func (r *SSHSSHPluginStatsResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+SSHSSHPluginStatsEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
