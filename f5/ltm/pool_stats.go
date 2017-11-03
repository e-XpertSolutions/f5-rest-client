// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

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
		ActiveMemberCnt struct {
			Value int `json:"value"`
		} `json:"activeMemberCnt,omitempty"`
		AvailableMemberCnt struct {
			Value int `json:"value"`
		} `json:"availableMemberCnt,omitempty"`
		ConnqAgeEdm struct {
			Value int `json:"value"`
		} `json:"connq.ageEdm,omitempty"`
		ConnqAgeEma struct {
			Value int `json:"value"`
		} `json:"connq.ageEma,omitempty"`
		ConnqAgeHead struct {
			Value int `json:"value"`
		} `json:"connq.ageHead,omitempty"`
		ConnqAgeMax struct {
			Value int `json:"value"`
		} `json:"connq.ageMax,omitempty"`
		ConnqDepth struct {
			Value int `json:"value"`
		} `json:"connq.depth,omitempty"`
		ConnqServiced struct {
			Value int `json:"value"`
		} `json:"connq.serviced,omitempty"`
		ConnqAllAgeEdm struct {
			Value int `json:"value"`
		} `json:"connqAll.ageEdm,omitempty"`
		ConnqAllAgeEma struct {
			Value int `json:"value"`
		} `json:"connqAll.ageEma,omitempty"`
		ConnqAllAgeHead struct {
			Value int `json:"value"`
		} `json:"connqAll.ageHead,omitempty"`
		ConnqAllAgeMax struct {
			Value int `json:"value"`
		} `json:"connqAll.ageMax,omitempty"`
		ConnqAllDepth struct {
			Value int `json:"value"`
		} `json:"connqAll.depth,omitempty"`
		ConnqAllServiced struct {
			Value int `json:"value"`
		} `json:"connqAll.serviced,omitempty"`
		CurSessions struct {
			Value int `json:"value"`
		} `json:"curSessions,omitempty"`
		MemberCnt struct {
			Value int `json:"value"`
		} `json:"memberCnt,omitempty"`
		MinActiveMembers struct {
			Value int `json:"value"`
		} `json:"minActiveMembers,omitempty"`
		MonitorRule struct {
			Description string `json:"description,omitempty"`
		} `json:"monitorRule,omitempty"`
		ServersideBitsIn struct {
			Value int `json:"value"`
		} `json:"serverside.bitsIn,omitempty"`
		ServersideBitsOut struct {
			Value int `json:"value"`
		} `json:"serverside.bitsOut,omitempty"`
		ServersideCurConns struct {
			Value int `json:"value"`
		} `json:"serverside.curConns,omitempty"`
		ServersideMaxConns struct {
			Value int `json:"value"`
		} `json:"serverside.maxConns,omitempty"`
		ServersidePktsIn struct {
			Value int `json:"value"`
		} `json:"serverside.pktsIn,omitempty"`
		ServersidePktsOut struct {
			Value int `json:"value"`
		} `json:"serverside.pktsOut,omitempty"`
		ServersideTotConns struct {
			Value int `json:"value"`
		} `json:"serverside.totConns,omitempty"`
		StatusAvailabilityState struct {
			Description string `json:"description,omitempty"`
		} `json:"status.availabilityState,omitempty"`
		StatusEnabledState struct {
			Description string `json:"description,omitempty"`
		} `json:"status.enabledState,omitempty"`
		StatusStatusReason struct {
			Description string `json:"description,omitempty"`
		} `json:"status.statusReason,omitempty"`
		TmName struct {
			Description string `json:"description,omitempty"`
		} `json:"tmName,omitempty"`
		TotRequests struct {
			Value int `json:"value"`
		} `json:"totRequests,omitempty"`
	} `json:"entries,omitempty"`
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
