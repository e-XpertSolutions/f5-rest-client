// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package apm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/apm"

type APM struct {
	c *f5.Client

	accessProfile  AccessProfileResource
	oAuthClientApp OAuthClientAppResource
	oAuthProfile   OAuthProfileResource
}

func New(c *f5.Client) *APM {
	return &APM{
		c:              c,
		accessProfile:  AccessProfileResource{c: c},
		oAuthClientApp: OAuthClientAppResource{c: c},
		oAuthProfile:   OAuthProfileResource{c: c},
	}
}

// AccessProfile returns a configured AccessProfileResource.
func (apm APM) AccessProfile() *AccessProfileResource {
	return &apm.accessProfile
}

// OAuthClientApp returns a configured OAuthClientAppResource.
func (apm APM) OAuthClientApp() *OAuthClientAppResource {
	return &apm.oAuthClientApp
}

// OAuthProfile returns a configured OAuthProfileResource.
func (apm APM) OAuthProfile() *OAuthProfileResource {
	return &apm.oAuthProfile
}
