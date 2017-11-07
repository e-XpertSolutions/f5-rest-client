// Copyright 2014-2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcmp

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// BasePath is the base path of the VCMP API.
const BasePath = "mgmt/tm/vcmp"

// VCMP implements a REST client for the F5 BigIP VCMP API.
type VCMP struct {
	c *f5.Client

	guest               GuestResource
	health              HealthResource
	trafficProfile      TrafficProfileResource
	virtualDisk         VirtualDiskResource
	virtualDiskTemplate VirtualDiskTemplateResource
}

// New creates a new VCMP client.
func New(c *f5.Client) VCMP {
	return VCMP{
		c: c,

		guest:               GuestResource{c: c},
		health:              HealthResource{c: c},
		trafficProfile:      TrafficProfileResource{c: c},
		virtualDisk:         VirtualDiskResource{c: c},
		virtualDiskTemplate: VirtualDiskTemplateResource{c: c},
	}
}

// Guest returns a configured GuestResource.
func (vcmp VCMP) Guest() *GuestResource {
	return &vcmp.guest
}

// Health returns a configured HealthResource.
func (vcmp VCMP) Health() *HealthResource {
	return &vcmp.health
}

// TrafficProfile returns a configured TrafficProfileResource.
func (vcmp VCMP) TrafficProfile() *TrafficProfileResource {
	return &vcmp.trafficProfile
}

// VirtualDisk returns a configured VirtualDiskResource.
func (vcmp VCMP) VirtualDisk() *VirtualDiskResource {
	return &vcmp.virtualDisk
}

// VirtualDiskTemplate returns a configured VirtualDiskTemplateResource.
func (vcmp VCMP) VirtualDiskTemplate() *VirtualDiskTemplateResource {
	return &vcmp.virtualDiskTemplate
}
