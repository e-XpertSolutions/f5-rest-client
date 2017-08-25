// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/cm"

type CM struct {
	c *f5.Client

	cert         CertResource
	key          KeyResource
	deviceGroup  DeviceGroupResource
	device       DeviceResource
	trafficGroup TrafficGroupResource
	trustDomain  TrustDomainResource
}

// New creates a new CM client.
func New(c *f5.Client) CM {
	return CM{
		c:            c,
		cert:         CertResource{c: c},
		key:          KeyResource{c: c},
		deviceGroup:  DeviceGroupResource{c: c},
		device:       DeviceResource{c: c},
		trafficGroup: TrafficGroupResource{c: c},
		trustDomain:  TrustDomainResource{c: c},
	}
}

// Cert returns a CertResource configured to query tm/cm/cert API.
func (cm CM) Cert() *CertResource {
	return &cm.cert
}

// Key returns a KeyResource configured to query tm/cm/key API.
func (cm CM) Key() *KeyResource {
	return &cm.key
}

// DeviceGroup returns a DeviceGroupResource configured to query tm/cm/device-group API.
func (cm CM) DeviceGroup() *DeviceGroupResource {
	return &cm.deviceGroup
}

// Device returns a DeviceResource configured to query tm/cm/device API.
func (cm CM) Device() *DeviceResource {
	return &cm.device
}

// TrafficGroup returns a TrafficGroupResource configured to query tm/cm/traffic-group API.
func (cm CM) TrafficGroup() *TrafficGroupResource {
	return &cm.trafficGroup
}

// TrustDomain returns a TrustDomainResource configured to query tm/cm/trust-domain API.
func (cm CM) TrustDomain() *TrustDomainResource {
	return &cm.trustDomain
}
