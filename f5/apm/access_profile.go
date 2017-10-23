// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AccessProfileConfigList holds a list of AccessProfile configuration.
type AccessProfileConfigList struct {
	Items    []AccessProfileConfig `json:"items,omitempty"`
	Kind     string                `json:"kind,omitempty"`
	SelfLink string                `json:"selflink,omitempty"`
}

// AccessProfileConfig holds the configuration of a single AccessProfile.
type AccessProfileConfig struct {
	AcceptLanguages       []string `json:"acceptLanguages,omitempty"`
	AccessPolicy          string   `json:"accessPolicy,omitempty"`
	AccessPolicyReference struct {
		Link string `json:"link,omitempty"`
	} `json:"accessPolicyReference,omitempty"`
	AccessPolicyTimeout         int    `json:"accessPolicyTimeout,omitempty"`
	CustomizationGroup          string `json:"customizationGroup,omitempty"`
	CustomizationGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"customizationGroupReference,omitempty"`
	DefaultLanguage       string `json:"defaultLanguage,omitempty"`
	DefaultsFrom          string `json:"defaultsFrom,omitempty"`
	DefaultsFromReference struct {
		Link string `json:"link,omitempty"`
	} `json:"defaultsFromReference,omitempty"`
	DomainGroupsReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"domainGroupsReference,omitempty"`
	DomainMode        string `json:"domainMode,omitempty"`
	EpsGroup          string `json:"epsGroup,omitempty"`
	EpsGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"epsGroupReference,omitempty"`
	ErrormapGroup          string `json:"errormapGroup,omitempty"`
	ErrormapGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"errormapGroupReference,omitempty"`
	FrameworkInstallationGroup          string `json:"frameworkInstallationGroup,omitempty"`
	FrameworkInstallationGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"frameworkInstallationGroupReference,omitempty"`
	FullPath                string `json:"fullPath,omitempty"`
	GeneralUIGroup          string `json:"generalUiGroup,omitempty"`
	GeneralUIGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"generalUiGroupReference,omitempty"`
	Generation           int      `json:"generation,omitempty"`
	GenerationAction     string   `json:"generationAction,omitempty"`
	HttponlyCookie       string   `json:"httponlyCookie,omitempty"`
	InactivityTimeout    int      `json:"inactivityTimeout,omitempty"`
	Kind                 string   `json:"kind,omitempty"`
	LogSettings          []string `json:"logSettings,omitempty"`
	LogSettingsReference []struct {
		Link string `json:"link,omitempty"`
	} `json:"logSettingsReference,omitempty"`
	LogoutURITimeout            int    `json:"logoutUriTimeout,omitempty"`
	MaxConcurrentSessions       int    `json:"maxConcurrentSessions,omitempty"`
	MaxConcurrentUsers          int    `json:"maxConcurrentUsers,omitempty"`
	MaxFailureDelay             int    `json:"maxFailureDelay,omitempty"`
	MaxInProgressSessions       int    `json:"maxInProgressSessions,omitempty"`
	MaxSessionTimeout           int    `json:"maxSessionTimeout,omitempty"`
	MinFailureDelay             int    `json:"minFailureDelay,omitempty"`
	ModifiedSinceLastPolicySync string `json:"modifiedSinceLastPolicySync,omitempty"`
	Name                        string `json:"name,omitempty"`
	Partition                   string `json:"partition,omitempty"`
	PersistentCookie            string `json:"persistentCookie,omitempty"`
	RestrictToSingleClientIP    string `json:"restrictToSingleClientIp,omitempty"`
	Scope                       string `json:"scope,omitempty"`
	SecureCookie                string `json:"secureCookie,omitempty"`
	SelfLink                    string `json:"selfLink,omitempty"`
	TmGeneration                int    `json:"tmGeneration,omitempty"`
	Type                        string `json:"type,omitempty"`
	UseHTTP503OnError           string `json:"useHttp_503OnError,omitempty"`
	UserIdentityMethod          string `json:"userIdentityMethod,omitempty"`
}

// AccessProfileEndpoint represents the REST resource for managing AccessProfile.
const AccessProfileEndpoint = "/profile/access"

// AccessProfileResource provides an API to manage AccessProfile configurations.
type AccessProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the AccessProfile configurations.
func (r *AccessProfileResource) ListAll() (*AccessProfileConfigList, error) {
	var list AccessProfileConfigList
	if err := r.c.ReadQuery(BasePath+AccessProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AccessProfile configuration identified by id.
func (r *AccessProfileResource) Get(id string) (*AccessProfileConfig, error) {
	var item AccessProfileConfig
	if err := r.c.ReadQuery(BasePath+AccessProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AccessProfile configuration.
func (r *AccessProfileResource) Create(item AccessProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AccessProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AccessProfile configuration identified by id.
func (r *AccessProfileResource) Edit(id string, item AccessProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AccessProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// ApplyPolicy applies APM policy identified by its id.
func (r *AccessProfileResource) ApplyPolicy(id string) error {
	data := AccessProfileConfig{GenerationAction: "increment"}
	if err := r.c.ModQuery("PATCH", BasePath+AccessProfileEndpoint+"/"+id, data); err != nil {
		return err
	}
	return nil
}

// Delete a single AccessProfile configuration identified by id.
func (r *AccessProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AccessProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
