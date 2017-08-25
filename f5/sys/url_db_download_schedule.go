// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// URLDBDownloadScheduleConfigList holds a list of URLDBDownloadSchedule configuration.
type URLDBDownloadScheduleConfigList struct {
	Items    []URLDBDownloadScheduleConfig `json:"items"`
	Kind     string                        `json:"kind"`
	SelfLink string                        `json:"selflink"`
}

// URLDBDownloadScheduleConfig holds the configuration of a single URLDBDownloadSchedule.
type URLDBDownloadScheduleConfig struct {
	DownloadNow string `json:"downloadNow"`
	EndTime     string `json:"endTime"`
	FullPath    string `json:"fullPath"`
	Generation  int    `json:"generation"`
	Kind        string `json:"kind"`
	Name        string `json:"name"`
	Partition   string `json:"partition"`
	SelfLink    string `json:"selfLink"`
	StartTime   string `json:"startTime"`
	Status      string `json:"status"`
}

// URLDBDownloadScheduleEndpoint represents the REST resource for managing URLDBDownloadSchedule.
const URLDBDownloadScheduleEndpoint = "/url-db/download-schedule"

// URLDBDownloadScheduleResource provides an API to manage URLDBDownloadSchedule configurations.
type URLDBDownloadScheduleResource struct {
	c *f5.Client
}

// ListAll  lists all the URLDBDownloadSchedule configurations.
func (r *URLDBDownloadScheduleResource) ListAll() (*URLDBDownloadScheduleConfigList, error) {
	var list URLDBDownloadScheduleConfigList
	if err := r.c.ReadQuery(BasePath+URLDBDownloadScheduleEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single URLDBDownloadSchedule configuration identified by id.
func (r *URLDBDownloadScheduleResource) Get(id string) (*URLDBDownloadScheduleConfig, error) {
	var item URLDBDownloadScheduleConfig
	if err := r.c.ReadQuery(BasePath+URLDBDownloadScheduleEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new URLDBDownloadSchedule configuration.
func (r *URLDBDownloadScheduleResource) Create(item URLDBDownloadScheduleConfig) error {
	if err := r.c.ModQuery("POST", BasePath+URLDBDownloadScheduleEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a URLDBDownloadSchedule configuration identified by id.
func (r *URLDBDownloadScheduleResource) Edit(id string, item URLDBDownloadScheduleConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+URLDBDownloadScheduleEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single URLDBDownloadSchedule configuration identified by id.
func (r *URLDBDownloadScheduleResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+URLDBDownloadScheduleEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
