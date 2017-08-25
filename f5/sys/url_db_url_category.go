// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// URLDBURLCategoryConfigList holds a list of URLDBURLCategory configuration.
type URLDBURLCategoryConfigList struct {
	Items    []URLDBURLCategoryConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// URLDBURLCategoryConfig holds the configuration of a single URLDBURLCategory.
type URLDBURLCategoryConfig struct {
	CatNumber       int    `json:"catNumber"`
	DefaultAction   string `json:"defaultAction"`
	Description     string `json:"description"`
	DisplayName     string `json:"displayName"`
	FullPath        string `json:"fullPath"`
	Generation      int    `json:"generation"`
	IsCustom        string `json:"isCustom"`
	IsRecategory    string `json:"isRecategory"`
	Kind            string `json:"kind"`
	Name            string `json:"name"`
	ParentCatNumber int    `json:"parentCatNumber"`
	Partition       string `json:"partition"`
	SelfLink        string `json:"selfLink"`
	SeverityLevel   int    `json:"severityLevel"`
}

// URLDBURLCategoryEndpoint represents the REST resource for managing URLDBURLCategory.
const URLDBURLCategoryEndpoint = "/url-db/url-category"

// URLDBURLCategoryResource provides an API to manage URLDBURLCategory configurations.
type URLDBURLCategoryResource struct {
	c *f5.Client
}

// ListAll  lists all the URLDBURLCategory configurations.
func (r *URLDBURLCategoryResource) ListAll() (*URLDBURLCategoryConfigList, error) {
	var list URLDBURLCategoryConfigList
	if err := r.c.ReadQuery(BasePath+URLDBURLCategoryEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single URLDBURLCategory configuration identified by id.
func (r *URLDBURLCategoryResource) Get(id string) (*URLDBURLCategoryConfig, error) {
	var item URLDBURLCategoryConfig
	if err := r.c.ReadQuery(BasePath+URLDBURLCategoryEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new URLDBURLCategory configuration.
func (r *URLDBURLCategoryResource) Create(item URLDBURLCategoryConfig) error {
	if err := r.c.ModQuery("POST", BasePath+URLDBURLCategoryEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a URLDBURLCategory configuration identified by id.
func (r *URLDBURLCategoryResource) Edit(id string, item URLDBURLCategoryConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+URLDBURLCategoryEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single URLDBURLCategory configuration identified by id.
func (r *URLDBURLCategoryResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+URLDBURLCategoryEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
