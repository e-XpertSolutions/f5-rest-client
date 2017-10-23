// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// OAuthProfileConfigList holds a list of OAuthProfile configuration.
type OAuthProfileConfigList struct {
	Items    []OAuthProfileConfig `json:"items,omitempty"`
	Kind     string               `json:"kind,omitempty"`
	SelfLink string               `json:"selflink,omitempty"`
}

// OAuthProfileConfig holds the configuration of a single OAuthProfile.
type OAuthProfileConfig struct {
	AccessTokenLifetime int    `json:"accessTokenLifetime,omitempty"`
	AuthCodeLifetime    int    `json:"authCodeLifetime,omitempty"`
	AuthURL             string `json:"authUrl,omitempty"`
	ClientAppsReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"clientAppsReference,omitempty"`
	DbInstance          string `json:"dbInstance,omitempty"`
	DbInstanceReference struct {
		Link string `json:"link,omitempty"`
	} `json:"dbInstanceReference,omitempty"`
	FullPath                 string `json:"fullPath,omitempty"`
	GenerateJwtRefreshToken  string `json:"generateJwtRefreshToken,omitempty"`
	GenerateRefreshToken     string `json:"generateRefreshToken,omitempty"`
	Generation               int    `json:"generation,omitempty"`
	IgnoreExpiredCert        string `json:"ignoreExpiredCert,omitempty"`
	Kind                     string `json:"kind,omitempty"`
	Name                     string `json:"name,omitempty"`
	OpaqueToken              string `json:"opaqueToken,omitempty"`
	Partition                string `json:"partition,omitempty"`
	PerUserTokenLimit        int    `json:"perUserTokenLimit,omitempty"`
	RefreshTokenLifetime     int    `json:"refreshTokenLifetime,omitempty"`
	RefreshTokenUsageLimit   int    `json:"refreshTokenUsageLimit,omitempty"`
	ResourceServersReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"resourceServersReference,omitempty"`
	ReuseAccessToken      string `json:"reuseAccessToken,omitempty"`
	ReuseRefreshToken     string `json:"reuseRefreshToken,omitempty"`
	RotationKeysReference struct {
		IsSubcollection bool   `json:"isSubcollection,omitempty"`
		Link            string `json:"link,omitempty"`
	} `json:"rotationKeysReference,omitempty"`
	SelfLink              string `json:"selfLink,omitempty"`
	Subject               string `json:"subject,omitempty"`
	TokenIntrospectionURL string `json:"tokenIntrospectionUrl,omitempty"`
	TokenIssuanceURL      string `json:"tokenIssuanceUrl,omitempty"`
	TokenRevocationURL    string `json:"tokenRevocationUrl,omitempty"`
}

// OAuthProfileEndpoint represents the REST resource for managing OAuthProfile.
const OAuthProfileEndpoint = "/profile/oauth"

// OAuthProfileResource provides an API to manage OAuthProfile configurations.
type OAuthProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the OAuthProfile configurations.
func (r *OAuthProfileResource) ListAll() (*OAuthProfileConfigList, error) {
	var list OAuthProfileConfigList
	if err := r.c.ReadQuery(BasePath+OAuthProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single OAuthProfile configuration identified by id.
func (r *OAuthProfileResource) Get(id string) (*OAuthProfileConfig, error) {
	var item OAuthProfileConfig
	if err := r.c.ReadQuery(BasePath+OAuthProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new OAuthProfile configuration.
func (r *OAuthProfileResource) Create(item OAuthProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+OAuthProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// AppendAppClient appends Client Application to an existing OAuth profile.
func (r *OAuthProfileResource) AppendAppClient(id string, idClientApp string) error {
	data := OAuthProfileConfig{Name: idClientApp}
	if err := r.c.ModQuery("POST", BasePath+OAuthProfileEndpoint+"/"+id+"/client-apps", data); err != nil {
		return err
	}
	return nil
}

// Edit a OAuthProfile configuration identified by id.
func (r *OAuthProfileResource) Edit(id string, item OAuthProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+OAuthProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single OAuthProfile configuration identified by id.
func (r *OAuthProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+OAuthProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
