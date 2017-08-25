// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// DBConfigList holds a list of DB configuration.
type DBConfigList struct {
	Items    []DBConfig `json:"items"`
	Kind     string     `json:"kind"`
	SelfLink string     `json:"selflink"`
}

// DBConfig holds the configuration of a single DB.
type DBConfig struct {
	DefaultValue string `json:"defaultValue"`
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Kind         string `json:"kind"`
	Name         string `json:"name"`
	ScfConfig    string `json:"scfConfig"`
	SelfLink     string `json:"selfLink"`
	Value        string `json:"value"`
	ValueRange   string `json:"valueRange"`
}

// DBEndpoint represents the REST resource for managing DB.
const DBEndpoint = "/db"

// DBResource provides an API to manage DB configurations.
type DBResource struct {
	c *f5.Client
}

// ListAll  lists all the DB configurations.
func (r *DBResource) ListAll() (*DBConfigList, error) {
	var list DBConfigList
	if err := r.c.ReadQuery(BasePath+DBEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single DB configuration identified by id.
func (r *DBResource) Get(id string) (*DBConfig, error) {
	var item DBConfig
	if err := r.c.ReadQuery(BasePath+DBEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new DB configuration.
func (r *DBResource) Create(item DBConfig) error {
	if err := r.c.ModQuery("POST", BasePath+DBEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a DB configuration identified by id.
func (r *DBResource) Edit(id string, item DBConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+DBEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single DB configuration identified by id.
func (r *DBResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+DBEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
