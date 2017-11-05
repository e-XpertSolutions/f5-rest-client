// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GlobalSettingsGeneral holds the configuration of a single GlobalSettingsGeneral.
type GlobalSettingsGeneral struct {
	AutoDiscovery                     string `json:"autoDiscovery,omitempty"`
	AutoDiscoveryInterval             int    `json:"autoDiscoveryInterval,omitempty"`
	AutomaticConfigurationSaveTimeout int    `json:"automaticConfigurationSaveTimeout,omitempty"`
	CacheLdnsServers                  string `json:"cacheLdnsServers,omitempty"`
	DomainNameCheck                   string `json:"domainNameCheck,omitempty"`
	DrainPersistentRequests           string `json:"drainPersistentRequests,omitempty"`
	ForwardStatus                     string `json:"forwardStatus,omitempty"`
	GtmSetsRecursion                  string `json:"gtmSetsRecursion,omitempty"`
	HeartbeatInterval                 int    `json:"heartbeatInterval,omitempty"`
	Kind                              string `json:"kind,omitempty"`
	MonitorDisabledObjects            string `json:"monitorDisabledObjects,omitempty"`
	NethsmTimeout                     int    `json:"nethsmTimeout,omitempty"`
	SelfLink                          string `json:"selfLink,omitempty"`
	SendWildcardRrs                   string `json:"sendWildcardRrs,omitempty"`
	StaticPersistCidrIpv4             int    `json:"staticPersistCidrIpv4,omitempty"`
	StaticPersistCidrIpv6             int    `json:"staticPersistCidrIpv6,omitempty"`
	Synchronization                   string `json:"synchronization,omitempty"`
	SynchronizationGroupName          string `json:"synchronizationGroupName,omitempty"`
	SynchronizationTimeTolerance      int    `json:"synchronizationTimeTolerance,omitempty"`
	SynchronizationTimeout            int    `json:"synchronizationTimeout,omitempty"`
	SynchronizeZoneFiles              string `json:"synchronizeZoneFiles,omitempty"`
	SynchronizeZoneFilesTimeout       int    `json:"synchronizeZoneFilesTimeout,omitempty"`
	VirtualsDependOnServerState       string `json:"virtualsDependOnServerState,omitempty"`
}

// GlobalSettingsGeneralEndpoint represents the REST resource for managing GlobalSettingsGeneral.
const GlobalSettingsGeneralEndpoint = "/global-settings/general"

// GlobalSettingsGeneralResource provides an API to manage GlobalSettingsGeneral configurations.
type GlobalSettingsGeneralResource struct {
	c *f5.Client
}

// List  lists all the GlobalSettingsGeneral configurations.
func (r *GlobalSettingsGeneralResource) List() (*GlobalSettingsGeneral, error) {
	var list GlobalSettingsGeneral
	if err := r.c.ReadQuery(BasePath+GlobalSettingsGeneralEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Modify a GlobalSettingsGeneral configuration.
func (r *GlobalSettingsGeneralResource) Modify(item GlobalSettingsGeneral) error {
	if err := r.c.ModQuery("POST", BasePath+GlobalSettingsGeneralEndpoint, item); err != nil {
		return err
	}
	return nil
}
