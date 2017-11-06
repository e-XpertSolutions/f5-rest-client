// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// PoolList holds a list of Pool configuration.
type PoolList struct {
	Items    []Pool `json:"items,omitempty"`
	Kind     string `json:"kind,omitempty"`
	SelfLink string `json:"selflink,omitempty"`
}

// Pool holds the configuration of a single Pool.
type Pool struct {
	AlternateMode             string `json:"alternateMode,omitempty"`
	DynamicRatio              string `json:"dynamicRatio,omitempty"`
	Enabled                   bool   `json:"enabled,omitempty"`
	FallbackIP                string `json:"fallbackIp,omitempty"`
	FallbackMode              string `json:"fallbackMode,omitempty"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	LimitMaxBps               int    `json:"limitMaxBps,omitempty"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus,omitempty"`
	LimitMaxConnections       int    `json:"limitMaxConnections,omitempty"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus,omitempty"`
	LimitMaxPps               int    `json:"limitMaxPps,omitempty"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus,omitempty"`
	LoadBalancingMode         string `json:"loadBalancingMode,omitempty"`
	ManualResume              string `json:"manualResume,omitempty"`
	MaxAnswersReturned        int    `json:"maxAnswersReturned,omitempty"`
	MembersReference          struct {
		Members         []PoolMembers `json:"items,omitempty"`
		IsSubcollection bool          `json:"isSubcollection,omitempty"`
		Link            string        `json:"link,omitempty"`
	} `json:"membersReference,omitempty"`
	Monitor                  string `json:"monitor,omitempty"`
	Name                     string `json:"name,omitempty"`
	Partition                string `json:"partition,omitempty"`
	QosHitRatio              int    `json:"qosHitRatio,omitempty"`
	QosHops                  int    `json:"qosHops,omitempty"`
	QosKilobytesSecond       int    `json:"qosKilobytesSecond,omitempty"`
	QosLcs                   int    `json:"qosLcs,omitempty"`
	QosPacketRate            int    `json:"qosPacketRate,omitempty"`
	QosRtt                   int    `json:"qosRtt,omitempty"`
	QosTopology              int    `json:"qosTopology,omitempty"`
	QosVsCapacity            int    `json:"qosVsCapacity,omitempty"`
	QosVsScore               int    `json:"qosVsScore,omitempty"`
	SelfLink                 string `json:"selfLink,omitempty"`
	TTL                      int    `json:"ttl,omitempty"`
	VerifyMemberAvailability string `json:"verifyMemberAvailability,omitempty"`
}

type PoolMembersList struct {
	Items    []PoolMembers `json:"items,omitempty"`
	Kind     string        `json:"kind,omitempty"`
	SelfLink string        `json:"selflink,omitempty"`
}

// PoolMembers holds the configuration of a single PoolMembers.
type PoolMembers struct {
	Enabled                   bool   `json:"enabled,omitempty"`
	FullPath                  string `json:"fullPath,omitempty"`
	Generation                int    `json:"generation,omitempty"`
	Kind                      string `json:"kind,omitempty"`
	LimitMaxBps               int    `json:"limitMaxBps,omitempty"`
	LimitMaxBpsStatus         string `json:"limitMaxBpsStatus,omitempty"`
	LimitMaxConnections       int    `json:"limitMaxConnections,omitempty"`
	LimitMaxConnectionsStatus string `json:"limitMaxConnectionsStatus,omitempty"`
	LimitMaxPps               int    `json:"limitMaxPps,omitempty"`
	LimitMaxPpsStatus         string `json:"limitMaxPpsStatus,omitempty"`
	MemberOrder               int    `json:"memberOrder,omitempty"`
	Monitor                   string `json:"monitor,omitempty"`
	Name                      string `json:"name,omitempty"`
	Partition                 string `json:"partition,omitempty"`
	Ratio                     int    `json:"ratio,omitempty"`
	SelfLink                  string `json:"selfLink,omitempty"`
}

type PoolStatsList struct {
	Entries  map[string]PoolStatsEntries `json:"entries,omitempty"`
	Kind     string                      `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                      `json:"selflink,omitempty" pretty:",expanded"`
}

type PoolStatsEntries struct {
	NestedPoolStats PoolStats `json:"nestedStats,omitempty"`
}

type PoolStats struct {
	Entries struct {
		Alternate struct {
			Value int `json:"value"`
		} `json:"alternate"`
		Dropped struct {
			Value int `json:"value"`
		} `json:"dropped"`
		Fallback struct {
			Value int `json:"value"`
		} `json:"fallback"`
		PoolType struct {
			Description string `json:"description"`
		} `json:"poolType"`
		Preferred struct {
			Value int `json:"value"`
		} `json:"preferred"`
		ReturnFromDNS struct {
			Value int `json:"value"`
		} `json:"returnFromDns"`
		ReturnToDNS struct {
			Value int `json:"value"`
		} `json:"returnToDns"`
		Status_availabilityState struct {
			Description string `json:"description"`
		} `json:"status.availabilityState"`
		Status_enabledState struct {
			Description string `json:"description"`
		} `json:"status.enabledState"`
		Status_statusReason struct {
			Description string `json:"description"`
		} `json:"status.statusReason"`
		TmName struct {
			Description string `json:"description"`
		} `json:"tmName"`
	} `json:"entries"`
}

// PoolEndpoint represents the REST resource for managing Pool.
const PoolAEndpoint = "/pool/a"

// PoolAResource provides an API to manage Pool configurations.
type PoolAResource struct {
	c *f5.Client
}

// ListAll  lists all the PoolA configurations.
func (r *PoolAResource) ListAll() (*PoolList, error) {
	var list PoolList
	if err := r.c.ReadQuery(BasePath+PoolAEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// GetAMembers  lists all the PoolMembers configurations.
func (r *PoolAResource) GetAMembers(id string) (*PoolMembersList, error) {
	var list PoolMembersList
	if err := r.c.ReadQuery(BasePath+PoolAEndpoint+"/"+id+"/members", &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Pool configuration identified by id.
func (r *PoolAResource) Get(id string) (*Pool, error) {
	var item Pool
	if err := r.c.ReadQuery(BasePath+PoolAEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new Pool configuration.
func (r *PoolAResource) Create(item Pool) error {
	if err := r.c.ModQuery("POST", BasePath+PoolAEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Pool configuration identified by id.
func (r *PoolAResource) Edit(id string, item Pool) error {
	if err := r.c.ModQuery("PUT", BasePath+PoolAEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Pool configuration identified by id.
func (r *PoolAResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+PoolAEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}

func (pr *PoolAResource) ShowAStats(id string) (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolAEndpoint+"/"+id+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (pr *PoolAResource) ShowAllAStats() (*PoolStatsList, error) {
	var item PoolStatsList
	if err := pr.c.ReadQuery(BasePath+PoolAEndpoint+"/stats", &item); err != nil {
		return nil, err
	}
	return &item, nil
}
