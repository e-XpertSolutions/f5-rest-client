// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type InterfaceStatsList struct {
	Entries  map[string]InterfaceStatsEntries `json:"entries,omitempty"`
	Kind     string                           `json:"kind,omitempty"`
	SelfLink string                           `json:"selflink,omitempty"`
}

type InterfaceStatsEntries struct {
	NestedInterfaceStats InterfaceStats `json:"nestedStats,omitempty"`
}

type InterfaceStats struct {
	Entries struct {
		CountersBitsIn struct {
			Value int `json:"value"`
		} `json:"counters.bitsIn,omitempty"`
		CountersBitsOut struct {
			Value int `json:"value"`
		} `json:"counters.bitsOut,omitempty"`
		CountersDropsAll struct {
			Value int `json:"value"`
		} `json:"counters.dropsAll,omitempty"`
		CountersErrorsAll struct {
			Value int `json:"value"`
		} `json:"counters.errorsAll,omitempty"`
		CountersPktsIn struct {
			Value int `json:"value"`
		} `json:"counters.pktsIn,omitempty"`
		CountersPktsOut struct {
			Value int `json:"value"`
		} `json:"counters.pktsOut,omitempty"`
		MediaActive struct {
			Description string `json:"description,omitempty"`
		} `json:"mediaActive,omitempty"`
		Status struct {
			Description string `json:"description,omitempty"`
		} `json:"status,omitempty"`
		TmName struct {
			Description string `json:"description,omitempty"`
		} `json:"tmName,omitempty"`
	} `json:"entries,omitempty"`
}

// InterfaceStatsEndpoint represents the REST resource for managing InterfaceStats.
const InterfaceStatsEndpoint = "/interface/stats"

// InterfaceStatsResource provides an API to manage InterfaceStats configurations.
type InetStatsResource struct {
	c *f5.Client
}

func (r *InetStatsResource) All() (*InterfaceStatsList, error) {
	var list InterfaceStatsList
	if err := r.c.ReadQuery(BasePath+InterfaceStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
