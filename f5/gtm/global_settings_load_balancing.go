// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GlobalSettingsLoadBalancing holds the configuration of a single GlobalSettingsLoadBalancing.
type GlobalSettingsLoadBalancing struct {
	FailureRcode              string `json:"failureRcode,omitempty"`
	FailureRcodeResponse      string `json:"failureRcodeResponse,omitempty"`
	FailureRcodeTTL           int    `json:"failureRcodeTtl,omitempty"`
	IgnorePathTTL             string `json:"ignorePathTtl,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	QosFactorBps              int    `json:"qosFactorBps,omitempty"`
	QosFactorHitRatio         int    `json:"qosFactorHitRatio,omitempty"`
	QosFactorHops             int    `json:"qosFactorHops,omitempty"`
	QosFactorLinkCapacity     int    `json:"qosFactorLinkCapacity,omitempty"`
	QosFactorPacketRate       int    `json:"qosFactorPacketRate,omitempty"`
	QosFactorRtt              int    `json:"qosFactorRtt,omitempty"`
	QosFactorTopology         int    `json:"qosFactorTopology,omitempty"`
	QosFactorVsCapacity       int    `json:"qosFactorVsCapacity,omitempty"`
	QosFactorVsScore          int    `json:"qosFactorVsScore,omitempty"`
	RespectFallbackDependency string `json:"respectFallbackDependency,omitempty"`
	SelfLink                  string `json:"selfLink,omitempty"`
	TopologyAllowZeroScores   string `json:"topologyAllowZeroScores,omitempty"`
	TopologyLongestMatch      string `json:"topologyLongestMatch,omitempty"`
	VerifyVsAvailability      string `json:"verifyVsAvailability,omitempty"`
}

// GlobalSettingsLoadBalancingEndpoint represents the REST resource for managing GlobalSettingsLoadBalancing.
const GlobalSettingsLoadBalancingEndpoint = "/global-settings/load-balancing"

// GlobalSettingsLoadBalancingResource provides an API to manage GlobalSettingsLoadBalancing configurations.
type GlobalSettingsLoadBalancingResource struct {
	c *f5.Client
}

// List  lists all the GlobalSettingsLoadBalancing configurations.
func (r *GlobalSettingsLoadBalancingResource) List() (*GlobalSettingsLoadBalancing, error) {
	var list GlobalSettingsLoadBalancing
	if err := r.c.ReadQuery(BasePath+GlobalSettingsLoadBalancingEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Modify a GlobalSettingsLoadBalancing configuration.
func (r *GlobalSettingsLoadBalancingResource) Modify(item GlobalSettingsLoadBalancing) error {
	if err := r.c.ModQuery("POST", BasePath+GlobalSettingsLoadBalancingEndpoint, item); err != nil {
		return err
	}
	return nil
}
