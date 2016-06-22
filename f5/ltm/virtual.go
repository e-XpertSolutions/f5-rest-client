// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import (
	"encoding/json"
	"net/http"

	"e-xpert_solutions/f5-rest-client/f5"
)

// VirtualServerConfig holds a list of virtual server configuration.
type VirtualServerConfig struct {
	Items    []VirtualServerConfigItem `json:"items"`
	Kind     string                    `json:"kind"`
	SelfLink string                    `json:"selfLink"`
}

// VirtualServerConfigItem  holds the configuration of a single virtual server.
type VirtualServerConfigItem struct {
	AddressStatus     string `json:"addressStatus"`
	AutoLasthop       string `json:"autoLasthop"`
	CmpEnabled        string `json:"cmpEnabled"`
	ConnectionLimit   int64  `json:"connectionLimit"`
	Description       string `json:"description"`
	Destination       string `json:"destination"`
	Disabled          bool   `json:"disabled"`
	Enabled           bool   `json:"enabled"`
	FullPath          string `json:"fullPath"`
	Generation        int64  `json:"generation"`
	GtmScore          int64  `json:"gtmScore"`
	IPProtocol        string `json:"ipProtocol"`
	Kind              string `json:"kind"`
	Mask              string `json:"mask"`
	Mirror            string `json:"mirror"`
	MobileAppTunnel   string `json:"mobileAppTunnel"`
	Name              string `json:"name"`
	Nat64             string `json:"nat64"`
	Partition         string `json:"partition"`
	PoliciesReference struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"policiesReference"`
	Pool              string `json:"pool"`
	ProfilesReference struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"profilesReference"`
	RateLimit                string   `json:"rateLimit"`
	RateLimitDstMask         int64    `json:"rateLimitDstMask"`
	RateLimitMode            string   `json:"rateLimitMode"`
	RateLimitSrcMask         int64    `json:"rateLimitSrcMask"`
	Rules                    []string `json:"rules"`
	SelfLink                 string   `json:"selfLink"`
	Source                   string   `json:"source"`
	SourceAddressTranslation struct {
		Type string `json:"type"`
	} `json:"sourceAddressTranslation"`
	SourcePort       string   `json:"sourcePort"`
	SynCookieStatus  string   `json:"synCookieStatus"`
	TranslateAddress string   `json:"translateAddress"`
	TranslatePort    string   `json:"translatePort"`
	Vlans            []string `json:"vlans"`
	VlansDisabled    bool     `json:"vlansDisabled"`
	VlansEnabled     bool     `json:"vlansEnabled"`
	VsIndex          int64    `json:"vsIndex"`
}

// VirtualEndpoint represents the REST resource for managing virtual server.
const VirtualEndpoint = "/virtual"

// VirtualResponse provide a simple mechanism to read paginated results.
//
// TODO(gilliek): use VirtualResponse object where pagination is needed.
type VirtualResponse struct {
}

// VirtualResource provides API to manage virtual server configurations.
type VirtualResource struct {
	c f5.Client
}

// ListAll lists all the virtual server configurations.
func (vr *VirtualResource) ListAll() (*VirtualServerConfig, error) {
	resp, err := vr.doRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return nil, err
	}
	var vsc VirtualServerConfig
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&vsc); err != nil {
		return nil, err
	}
	return &vsc, nil
}

// Get a single virtual server configurations identified by id.
func (vr *VirtualResource) Get(id string) (*VirtualServerConfigItem, error) {
	resp, err := vr.doRequest("GET", id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return nil, err
	}
	var vsci VirtualServerConfigItem
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&vsci); err != nil {
		return nil, err
	}
	return &vsci, nil
}

// Create a new virtual server configuration.
func (vr *VirtualResource) Create(item VirtualServerConfigItem) error {
	resp, err := vr.doRequest("POST", "", item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Edit a virtual server configuration identified by id.
func (vr *VirtualResource) Edit(id string, item VirtualServerConfigItem) error {
	resp, err := vr.doRequest("PUT", id, item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Delete a single server configuration identified by id.
func (vr *VirtualResource) Delete(id string) error {
	resp, err := vr.doRequest("DELETE", id, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// doRequest creates and send HTTP request using the F5 client.
//
// TODO(gilliek): decorate errors
func (vr *VirtualResource) doRequest(method, id string, data interface{}) (*http.Response, error) {
	req, err := vr.c.MakeRequest(method, BasePath+VirtualEndpoint+"/"+id, data)
	if err != nil {
		return nil, err
	}
	resp, err := vr.c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// readError reads request error object from a HTTP response.
//
// TODO(gilliek): move this function into F5 package.
func (vr *VirtualResource) readError(resp *http.Response) error {
	if resp.StatusCode >= 400 {
		errResp, err := f5.NewRequestError(resp.Body)
		if err != nil {
			return err
		}
		return errResp
	}
	return nil
}
