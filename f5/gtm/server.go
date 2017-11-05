// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// ServerList holds a list of Server configuration.
type ServerList struct {
	Items    []Server `json:"items,omitempty"`
	Kind     string   `json:"kind,omitempty"`
	SelfLink string   `json:"selflink,omitempty"`
}

// Server holds the configuration of a single Server.
type Server struct {
	Addresses []struct {
		DeviceName  string `json:"deviceName,omitempty"`
		Name        string `json:"name,omitempty"`
		Translation string `json:"translation,omitempty"`
	} `json:"addresses,omitempty"`
	Datacenter          string `json:"datacenter,omitempty"`
	DatacenterReference struct {
		Link string `json:"link,omitempty"`
	} `json:"datacenterReference,omitempty"`
	DevicesReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"devicesReference,omitempty"`
	Enabled                   bool   `json:"enabled,omitempty"`
	ExposeRouteDomains        string `json:"exposeRouteDomains,omitempty"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	IqAllowPath               string `json:"iqAllowPath,omitempty"`
	IqAllowServiceCheck       string `json:"iqAllowServiceCheck,omitempty"`
	IqAllowSnmp               string `json:"iqAllowSnmp,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	LimitCPUUsage             int    `json:"limitCpuUsage,omitempty"`
	LimitCPUUsageStatus       string `json:"limitCpuUsageStatus,omitempty"`
	LimitMaxBps               int    `json:"limitMaxBps,omitempty"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus,omitempty"`
	LimitMaxConnections       int    `json:"limitMaxConnections,omitempty"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus,omitempty"`
	LimitMaxPps               int    `json:"limitMaxPps,omitempty"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus,omitempty"`
	LimitMemAvail             int    `json:"limitMemAvail,omitempty"`
	LimitMemAvailStatus       string `json:"limitMemAvailStatus,omitempty"`
	LinkDiscovery             string `json:"linkDiscovery,omitempty"`
	Monitor                   string `json:"monitor,omitempty"`
	Name                      string `json:"name,omitempty"`
	Partition                 string `json:"partition,omitempty"`
	ProberFallback            string `json:"proberFallback,omitempty"`
	ProberPreference          string `json:"proberPreference,omitempty"`
	Product                   string `json:"product,omitempty"`
	SelfLink                  string `json:"selfLink,omitempty"`
	VirtualServerDiscovery    string `json:"virtualServerDiscovery,omitempty"`
	VirtualServersReference   struct {
		VirtualServers  []ServerVirtualServers `json:"items,omitempty"`
		IsSubcollection bool                   `json:"isSubcollection,omitempty"`
		Link            string                 `json:"link,omitempty"`
	} `json:"virtualServersReference,omitempty"`
}

// ServerVirtualServersList holds a list of ServerVirtualServers configuration.
type ServerVirtualServersList struct {
	Items    []ServerVirtualServers `json:"items"`
	Kind     string                 `json:"kind"`
	SelfLink string                 `json:"selflink"`
}

// ServerVirtualServers holds the configuration of a single ServerVirtualServers.
type ServerVirtualServers struct {
	LimitMaxPpsStatus        string `json:"limitMaxPpsStatus,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	LimitMaxBps              int    `json:"limitMaxBps,omitempty"`
	Destination              string `json:"destination,omitempty"`
	LimitMaxConnections      string `json:"limitMaxConnections,omitempty"`
	Enabled                  bool   `json:"enabled,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	LimitMaxConnectionStatus string `json:"limitMaxConnectionStatus,omitempty"`
	TranslationPort          int    `json:"translationPort,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	LimitMaxPps              int    `json:"limitMaxPps,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	LimitMaxBpsStatus        string `json:"limitMaxBpsStatus,omitempty"`
	TranslationAddress       string `json:"translationAddress,omitempty"`
	Name                     string `json:"name,omitempty"`
}

// ServerVirtualServersEndpoint represents the REST resource for managing ServerVirtualServers.
const ServerVirtualServersEndpoint = "/virtual-servers"

// ServerEndpoint represents the REST resource for managing Server.
const ServerEndpoint = "/server"

// ServerResource provides an API to manage Server configurations.
type ServerResource struct {
	c *f5.Client
}

// ListAll  lists all the Server configurations.
func (r *ServerResource) ListAll() (*ServerList, error) {
	var list ServerList
	if err := r.c.ReadQuery(BasePath+ServerEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Server configuration identified by id.
func (r *ServerResource) Get(id string) (*Server, error) {
	var item Server
	if err := r.c.ReadQuery(BasePath+ServerEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// GetVirtualServers lists all the ServerVirtualServers configurations.
func (r *ServerResource) GetVirtualServers(id string) (*ServerVirtualServersList, error) {
	var list ServerVirtualServersList
	if err := r.c.ReadQuery(BasePath+ServerEndpoint+"/"+id+ServerVirtualServersEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Create a new Server configuration.
func (r *ServerResource) Create(item Server) error {
	if err := r.c.ModQuery("POST", BasePath+ServerEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Server configuration identified by id.
func (r *ServerResource) Edit(id string, item Server) error {
	if err := r.c.ModQuery("PUT", BasePath+ServerEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Server configuration identified by id.
func (r *ServerResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ServerEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
