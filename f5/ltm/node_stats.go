// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type NodeStatsList struct {
	Entries  map[string]NodeStatsEntries `json:"entries,omitempty"`
	Kind     string                      `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                      `json:"selflink,omitempty" pretty:",expanded"`
}

type NodeStatsEntries struct {
	NestedNodeStats NodeStats `json:"nestedStats,omitempty"`
}

type NodeStats struct {
	Entries struct {
		Addr struct {
			Description string `json:"description,omitempty"`
		} `json:"addr,omitempty"`
		CurSessions struct {
			Value int `json:"value"`
		} `json:"curSessions,omitempty"`
		MonitorRule struct {
			Description string `json:"description,omitempty"`
		} `json:"monitorRule,omitempty"`
		MonitorStatus struct {
			Description string `json:"description,omitempty"`
		} `json:"monitorStatus,omitempty"`
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
		SessionStatus struct {
			Description string `json:"description,omitempty"`
		} `json:"sessionStatus,omitempty"`
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
		} `json:"tmName,omitempty" pretty:",expanded"`
		TotRequests struct {
			Value int `json:"value"`
		} `json:"totRequests,omitempty"`
	} `json:"entries,omitempty"`
}

// NodeStatsEndpoint represents the REST resource for managing NodeStats.
const NodeStatsEndpoint = "/node/stats"

// NodeStatsResource provides an API to manage NodeStats entries.
type NodeStatsResource struct {
	c *f5.Client
}

func (r *NodeStatsResource) All() (*NodeStatsList, error) {
	var list NodeStatsList
	if err := r.c.ReadQuery(BasePath+NodeStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
