// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type MemoryStatsList struct {
	Entries  map[string]MemoryTopLevelEntries `json:"entries,omitempty"`
	Kind     string                           `json:"kind,omitempty" pretty:"expanded"`
	SelfLink string                           `json:"selfLink,omitempty" pretty:"expanded"`
}

type MemoryTopLevelEntries struct {
	NestedStats MemoryInnerStatsList `json:"nestedStats,omitempty"`
}

type MemoryInnerStatsList struct {
	Entries map[string]MemoryStatsEntries `json:"entries,omitempty"`
}

type MemoryStatsEntries struct {
	NestedStats MemoryStats `json:"nestedStats,omitempty"`
}

type MemoryStats struct {
	Entries struct {
		Allocated struct {
			Value int `json:"value"`
		} `json:"allocated,omitempty"`
		HostId struct {
			Value string `json:"description"`
		} `json:"hostId,omitempty"`
		MaxAllocated struct{
			Value int `json:"value"`
		} `json maxAllocated,omitempty"`
		MemoryFree struct {
			Value int `json:"value"`
		} `json:"memoryFree,omitempty"`
		MemoryTotal struct {
			Value int `json:"value"`
		} `json:"memoryTotal,omitempty"`
		MemoryUsed struct {
			Value int `json:"value"`
		} `json:"memoryUsed,omitempty"`
		OtherMemoryFree struct {
			Value int `json:"value"`
		} `json:"otherMemoryFree,omitempty"`
		OtherMemoryTotal struct {
			Value int `json:"value"`
		} `json:"otherMemoryTotal,omitempty"`
		OtherMemoryUsed struct {
			Value int `json:"value"`
		} `json:"otherMemoryUsed,omitempty"`
		Size struct {
			Value int `json:"value"`
		} `json:"size,omitempty"`
		SwapFree struct {
			Value int `json:"value"`
		} `json:"swapFree,omitempty"`
		SwapTotal struct {
			Value int `json:"value"`
		} `json:"swapTotal,omitempty"`
		SwapUsed struct {
			Value int `json:"value"`
		} `json:"swapUsed,omitempty"`
		TmmId struct {
			Value string `json:"description"`
		} `json:"tmmId,omitempty"`
		TmmMemoryFree struct {
			Value int `json:"value"`
		} `json:"tmmMemoryFree,omitempty"`
		TmmMemoryTotal struct {
			Value int `json:"value"`
		} `json:"tmmMemoryTotal,omitempty"`
		TmmMemoryUsed struct {
			Value int `json:"value"`
		} `json:"tmmMemoryUsed,omitempty"`
		TmName struct {
			Value string `json:"description"`
		} `json:"tmName,omitempty"`
	} `json:"entries,omitempty"`
}

// MemoryStatsEndpoint represents the REST resource for managing MemoryStats.
const MemoryStatsEndpoint = "/memory"

// MemoryStatsResource provides an API to manage MemoryStats entries.
type MemoryStatsResource struct {
	c *f5.Client
}

func (r *MemoryStatsResource) All() (*MemoryStatsList, error) {
	var list MemoryStatsList
	if err := r.c.ReadQuery(BasePath+MemoryStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
