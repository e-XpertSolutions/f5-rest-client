// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type NodeStatsList struct {
	Entries  map[string]NodeStatsEntries `json:"entries"`
	Kind     string                      `json:"kind"`
	SelfLink string                      `json:"selflink"`
}

type NodeStatsEntries struct {
	NestedNodeStats NodeStats `json:"nestedStats"`
}

type NodeStats struct {
	Entries struct {
		Addr struct {
			Description string `json:"description"`
		} `json:"addr"`
		CurSessions struct {
			Value int `json:"value"`
		} `json:"curSessions"`
		MonitorRule struct {
			Description string `json:"description"`
		} `json:"monitorRule"`
		MonitorStatus struct {
			Description string `json:"description"`
		} `json:"monitorStatus"`
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
		SessionStatus struct {
			Description string `json:"description"`
		} `json:"sessionStatus"`
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
