// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apm

import (
	"encoding/json"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// OAuthClientAppConfigList holds a list of OAuthClientApp configuration.
type OAuthClientAppConfigList struct {
	Items    []OAuthClientAppConfig `json:"items,omitempty"`
	Kind     string                 `json:"kind,omitempty"`
	SelfLink string                 `json:"selflink,omitempty"`
}

// OAuthClientAppConfig holds the configuration of a single OAuthClientApp.
type OAuthClientAppConfig struct {
	AccessTokenLifetime         int    `json:"accessTokenLifetime,omitempty"`
	AppName                     string `json:"appName,omitempty"`
	AuthCodeLifetime            int    `json:"authCodeLifetime,omitempty"`
	AuthType                    string `json:"authType,omitempty"`
	ClientID                    string `json:"clientId,omitempty"`
	ClientSecret                string `json:"clientSecret,omitempty"`
	CustomizationGroup          string `json:"customizationGroup,omitempty"`
	CustomizationGroupReference struct {
		Link string `json:"link,omitempty"`
	} `json:"customizationGroupReference,omitempty"`
	FullPath               string `json:"fullPath,omitempty"`
	GenerateRefreshToken   string `json:"generateRefreshToken,omitempty"`
	Generation             int    `json:"generation,omitempty"`
	GrantCode              string `json:"grantCode,omitempty"`
	GrantPassword          string `json:"grantPassword,omitempty"`
	GrantToken             string `json:"grantToken,omitempty"`
	Kind                   string `json:"kind,omitempty"`
	Name                   string `json:"name,omitempty"`
	Partition              string `json:"partition,omitempty"`
	RefreshTokenLifetime   int    `json:"refreshTokenLifetime,omitempty"`
	RefreshTokenUsageLimit int    `json:"refreshTokenUsageLimit,omitempty"`
	ReuseAccessToken       string `json:"reuseAccessToken,omitempty"`
	ReuseRefreshToken      string `json:"reuseRefreshToken,omitempty"`
	Scopes                 string `json:"scopes,omitempty"`
	ScopesReference        struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"scopesReference,omitempty"`
	SelfLink                    string `json:"selfLink,omitempty"`
	UseProfileTokenMgmtSettings string `json:"useProfileTokenMgmtSettings,omitempty"`
}

// OAuthClientAppEndpoint represents the REST resource for managing OAuthClientApp.
const OAuthClientAppEndpoint = "/oauth/oauth-client-app"

// OAuthClientAppResource provides an API to manage OAuthClientApp configurations.
type OAuthClientAppResource struct {
	c *f5.Client
}

// ListAll  lists all the OAuthClientApp configurations.
func (r *OAuthClientAppResource) ListAll() (*OAuthClientAppConfigList, error) {
	var list OAuthClientAppConfigList
	if err := r.c.ReadQuery(BasePath+OAuthClientAppEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single OAuthClientApp configuration identified by id.
func (r *OAuthClientAppResource) Get(id string) (*OAuthClientAppConfig, error) {
	var item OAuthClientAppConfig
	if err := r.c.ReadQuery(BasePath+OAuthClientAppEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new OAuthClientApp configuration.
func (r *OAuthClientAppResource) Create(item OAuthClientAppConfig) (*OAuthClientAppConfig, error) {
	resp, err := r.c.MakeRequest("POST", BasePath+OAuthClientAppEndpoint, item)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respItem OAuthClientAppConfig
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&respItem); err != nil {
		return nil, err
	}
	return &respItem, nil
}

// Edit a OAuthClientApp configuration identified by id.
func (r *OAuthClientAppResource) Edit(id string, item OAuthClientAppConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+OAuthClientAppEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single OAuthClientApp configuration identified by id.
func (r *OAuthClientAppResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+OAuthClientAppEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
