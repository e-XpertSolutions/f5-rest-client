// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "e-xpert_solutions/f5-rest-client/f5"

type MonitorDiameterConfigList struct {
	Items    []MonitorDiameterConfig `json:"items"`
	Kind     string                  `json:"kind"`
	SelfLink string                  `json:"selflink"`
}

type MonitorDiameterConfig struct {
	FullPath     string `json:"fullPath"`
	Generation   int    `json:"generation"`
	Interval     int    `json:"interval"`
	Kind         string `json:"kind"`
	ManualResume string `json:"manualResume"`
	Name         string `json:"name"`
	OriginRealm  string `json:"originRealm"`
	Partition    string `json:"partition"`
	ProductName  string `json:"productName"`
	SelfLink     string `json:"selfLink"`
	TimeUntilUp  int    `json:"timeUntilUp"`
	Timeout      int    `json:"timeout"`
	UpInterval   int    `json:"upInterval"`
	VendorID     string `json:"vendorId"`
}

const MonitorDiameterEndpoint = "/monitor/diameter"

type MonitorDiameterResource struct {
	c f5.Client
}

func (r *MonitorDiameterResource) ListAll() (*MonitorDiameterConfigList, error) {
	var list MonitorDiameterConfigList
	if err := r.c.ReadQuery(BasePath+MonitorDiameterEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *MonitorDiameterResource) Get(id string) (*MonitorDiameterConfig, error) {
	var item MonitorDiameterConfig
	if err := r.c.ReadQuery(BasePath+MonitorDiameterEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *MonitorDiameterResource) Create(item MonitorDiameterConfig) error {
	if err := r.c.ModQuery("POST", BasePath+MonitorDiameterEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorDiameterResource) Edit(id string, item MonitorDiameterConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+MonitorDiameterEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *MonitorDiameterResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+MonitorDiameterEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
