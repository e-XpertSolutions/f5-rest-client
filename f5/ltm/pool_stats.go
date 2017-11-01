// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type PoolStatsList struct {
	Entries  map[string]PoolStatsEntries `json:"entries"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

type PoolStatsEntries struct {
	NestedPoolStats PoolStats `json:"nestedStats"`
}

type PoolStats struct {
	Entries struct {
		ActiveMemberCnt struct {
			Value int `json:"value"`
		} `json:"activeMemberCnt"`
		AvailableMemberCnt struct {
			Value int `json:"value"`
		} `json:"availableMemberCnt"`
		ConnqAgeEdm struct {
			Value int `json:"value"`
		} `json:"connq.ageEdm"`
		ConnqAgeEma struct {
			Value int `json:"value"`
		} `json:"connq.ageEma"`
		ConnqAgeHead struct {
			Value int `json:"value"`
		} `json:"connq.ageHead"`
		ConnqAgeMax struct {
			Value int `json:"value"`
		} `json:"connq.ageMax"`
		ConnqDepth struct {
			Value int `json:"value"`
		} `json:"connq.depth"`
		ConnqServiced struct {
			Value int `json:"value"`
		} `json:"connq.serviced"`
		ConnqAllAgeEdm struct {
			Value int `json:"value"`
		} `json:"connqAll.ageEdm"`
		ConnqAllAgeEma struct {
			Value int `json:"value"`
		} `json:"connqAll.ageEma"`
		ConnqAllAgeHead struct {
			Value int `json:"value"`
		} `json:"connqAll.ageHead"`
		ConnqAllAgeMax struct {
			Value int `json:"value"`
		} `json:"connqAll.ageMax"`
		ConnqAllDepth struct {
			Value int `json:"value"`
		} `json:"connqAll.depth"`
		ConnqAllServiced struct {
			Value int `json:"value"`
		} `json:"connqAll.serviced"`
		CurSessions struct {
			Value int `json:"value"`
		} `json:"curSessions"`
		MemberCnt struct {
			Value int `json:"value"`
		} `json:"memberCnt"`
		MinActiveMembers struct {
			Value int `json:"value"`
		} `json:"minActiveMembers"`
		MonitorRule struct {
			Description string `json:"description"`
		} `json:"monitorRule"`
		ServersideBitsIn struct {
			Value int `json:"value"`
		} `json:"serverside.bitsIn"`
		ServersideBitsOut struct {
			Value int `json:"value"`
		} `json:"serverside.bitsOut"`
		ServersideCurConns struct {
			Value int `json:"value"`
		} `json:"serverside.curConns"`
		ServersideMaxConns struct {
			Value int `json:"value"`
		} `json:"serverside.maxConns"`
		ServersidePktsIn struct {
			Value int `json:"value"`
		} `json:"serverside.pktsIn"`
		ServersidePktsOut struct {
			Value int `json:"value"`
		} `json:"serverside.pktsOut"`
		ServersideTotConns struct {
			Value int `json:"value"`
		} `json:"serverside.totConns"`
		StatusAvailabilityState struct {
			Description string `json:"description"`
		} `json:"status.availabilityState"`
		StatusEnabledState struct {
			Description string `json:"description"`
		} `json:"status.enabledState"`
		StatusStatusReason struct {
			Description string `json:"description"`
		} `json:"status.statusReason"`
		TmName struct {
			Description string `json:"description"`
		} `json:"tmName"`
		TotRequests struct {
			Value int `json:"value"`
		} `json:"totRequests"`
	} `json:"entries"`
}

// PoolStatsEndpoint represents the REST resource for managing PoolStats.
const PoolStatsEndpoint = "/pool/stats"

// PoolStatsResource provides an API to manage PoolStats configurations.
type PoolStatsResource struct {
	c *f5.Client
}

func (r *PoolStatsResource) All() (*PoolStatsList, error) {
	var list PoolStatsList
	if err := r.c.ReadQuery(BasePath+PoolStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
