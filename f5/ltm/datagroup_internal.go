// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type DataGroupInternalConfigList struct {
	Items    []DataGroupInternalConfig `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selfLink"`
}

type DataGroupInternalConfig struct {
	Description string `json:"description"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Records     []struct {
		Data string `json:"data"`
		Name string `json:"name"`
	} `json:"records"`
	SelfLink  string `json:"selfLink"`
	Type      string `json:"type"`
	Partition string `json:"partition"`
}

const DataGroupInternalEndpoint = "/data-group/internal"

type DataGroupInternalResource struct {
	c *f5.Client
}

func (nr *DataGroupInternalResource) ListAll() (*DataGroupInternalConfigList, error) {
	var list DataGroupInternalConfigList
	if err := nr.c.ReadQuery(BasePath+DataGroupInternalEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (nr *DataGroupInternalResource) Get(id string) (*DataGroupInternalConfig, error) {
	var item DataGroupInternalConfig
	if err := nr.c.ReadQuery(BasePath+DataGroupInternalEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (nr *DataGroupInternalResource) Create(item DataGroupInternalConfig) error {
	if err := nr.c.ModQuery("POST", BasePath+DataGroupInternalEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (nr *DataGroupInternalResource) Edit(id string, item DataGroupInternalConfig) error {
	if err := nr.c.ModQuery("PUT", BasePath+DataGroupInternalEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (nr *DataGroupInternalResource) Delete(id string) error {
	if err := nr.c.ModQuery("DELETE", BasePath+DataGroupInternalEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
