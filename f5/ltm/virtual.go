// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import (
	"encoding/json"
	"net/http"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

type Persistence struct {
	Name      string `json:"name,omitempty"`
	Partition string `json:"partition,omitempty"`
	TMDefault string `json:"tmDefault,omitempty"`
}

// VirtualServerList holds a list of virtual server uration.
type VirtualServerList struct {
	Items    []VirtualServer `json:"items,omitempty"`
	Kind     string          `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string          `json:"selfLink,omitempty" pretty:",expanded"`
}

// VirtualServer holds the uration of a single virtual server.
type VirtualServer struct {
	AddressStatus       string        `json:"addressStatus,omitempty"`
	AutoLasthop         string        `json:"autoLasthop,omitempty"`
	CmpEnabled          string        `json:"cmpEnabled,omitempty"`
	ConnectionLimit     int64         `json:"connectionLimit,omitempty"`
	Description         string        `json:"description,omitempty"`
	Destination         string        `json:"destination,omitempty"`
	Disabled            bool          `json:"disabled,omitempty"`
	Enabled             bool          `json:"enabled,omitempty"`
	FallbackPersistence string        `json:"fallbackPersistence,omitempty"`
	FullPath            string        `json:"fullPath,omitempty" pretty:",expanded"`
	FwEnforcedPolicy    string        `json:"fwEnforcedPolicy,omitempty"`
	Generation          int64         `json:"generation,omitempty" pretty:",expanded"`
	GtmScore            int64         `json:"gtmScore,omitempty" pretty:",expanded"`
	IPProtocol          string        `json:"ipProtocol,omitempty"`
	Kind                string        `json:"kind,omitempty" pretty:",expanded"`
	Mask                string        `json:"mask,omitempty"`
	Mirror              string        `json:"mirror,omitempty"`
	MobileAppTunnel     string        `json:"mobileAppTunnel,omitempty" pretty:",expanded"`
	Name                string        `json:"name,omitempty"`
	Nat64               string        `json:"nat64,omitempty" pretty:",expanded"`
	Partition           string        `json:"partition,omitempty"`
	Persistences        []Persistence `json:"persist,omitempty"`
	PoliciesReference   struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"policiesReference,omitempty"`
	Pool              string   `json:"pool,omitempty"`
	Profiles          []string `json:"profiles,omitempty"` // only used to link existing profiles a creation or update
	ProfilesReference struct {
		IsSubcollection bool      `json:"isSubcollection,omitempty"`
		Link            string    `json:"link,omitempty"`
		Profiles        []Profile `json:"items,omitempty"`
	} `json:"profilesReference,omitempty"`
	RateLimit                string                   `json:"rateLimit,omitempty" pretty:",expanded"`
	RateLimitDstMask         int64                    `json:"rateLimitDstMask,omitempty" pretty:",expanded"`
	RateLimitMode            string                   `json:"rateLimitMode,omitempty" pretty:",expanded"`
	RateLimitSrcMask         int64                    `json:"rateLimitSrcMask,omitempty" pretty:",expanded"`
	Rules                    []string                 `json:"rules,omitempty"`
	SelfLink                 string                   `json:"selfLink,omitempty" pretty:",expanded"`
	Source                   string                   `json:"source,omitempty"`
	SourceAddressTranslation SourceAddressTranslation `json:"sourceAddressTranslation,omitempty"`
	SourcePort               string                   `json:"sourcePort,omitempty"`
	SynCookieStatus          string                   `json:"synCookieStatus,omitempty"`
	TranslateAddress         string                   `json:"translateAddress,omitempty"`
	TranslatePort            string                   `json:"translatePort,omitempty"`
	Vlans                    []string                 `json:"vlans,omitempty"`
	VlansDisabled            bool                     `json:"vlansDisabled,omitempty"`
	VlansEnabled             bool                     `json:"vlansEnabled,omitempty"`
	VsIndex                  int64                    `json:"vsIndex,omitempty" pretty:",expanded"`
}

type SourceAddressTranslation struct {
	Type string `json:"type,omitempty"`
}

type Profile struct {
	Name    string `json:"name,omitempty"`
	Context string `json:"context,omitempty"`
}

// VirtualEndpoint represents the REST resource for managing virtual server.
const VirtualEndpoint = "/virtual"

// VirtualResponse provide a simple mechanism to read paginated results.
//
// TODO(gilliek): use VirtualResponse object where pagination is needed.
type VirtualResponse struct {
}

// VirtualResource provides an API to manage virtual server urations.
type VirtualResource struct {
	c *f5.Client
}

// ListAll lists all the virtual server urations.
func (vr *VirtualResource) ListAll() (*VirtualServerList, error) {
	resp, err := vr.doRequest("GET", "", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return nil, err
	}
	var vsc VirtualServerList
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&vsc); err != nil {
		return nil, err
	}
	return &vsc, nil
}

// Get a single virtual server uration identified by id.
func (vr *VirtualResource) Get(id string) (*VirtualServer, error) {
	resp, err := vr.doRequest("GET", id, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return nil, err
	}
	var vsci VirtualServer
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&vsci); err != nil {
		return nil, err
	}
	return &vsci, nil
}

// Create a new virtual server uration.
func (vr *VirtualResource) Create(item VirtualServer) error {
	if err := vr.c.ModQuery("POST", BasePath+VirtualEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a virtual server uration identified by id.
func (vr *VirtualResource) Edit(id string, item VirtualServer) error {
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

// Enabling a virtual server item identified by id.
func (vr *VirtualResource) Enable(id string) error {
	item := VirtualServer{Enabled: true}
	resp, err := vr.doRequest("PATCH", id, item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Disabling a virtual server item identified by id.
func (vr *VirtualResource) Disable(id string) error {
	item := VirtualServer{Disabled: true}
	resp, err := vr.doRequest("PATCH", id, item)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// Delete a single server uration identified by id.
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

// Rules gets the iRules uration for a virtual server identified by id.
func (vr *VirtualResource) Rules(id string) ([]Rule, error) {
	resp, err := vr.doRequest("GET", id+"/rule", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return nil, err
	}
	var rules []Rule
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&rules); err != nil {
		return nil, err
	}
	return rules, nil
}

// AddRule adds an iRule to the virtual server identified by id.
func (vr *VirtualResource) AddRule(id string, rule Rule) error {
	resp, err := vr.doRequest("POST", id+"/rule", rule)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if err := vr.readError(resp); err != nil {
		return err
	}
	return nil
}

// RemoveRule removes a single  iRule from the virtual server identified by id.
func (vr *VirtualResource) RemoveRule(vsID, ruleID string) error {
	resp, err := vr.doRequest("DELETE", vsID+"/rule/"+ruleID, nil)
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

func (vr *VirtualResource) ShowStats(id string) (*VirtualStatsList, error) {
	var item VirtualStatsList
	if err := vr.c.ReadQuery(BasePath+VirtualEndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
