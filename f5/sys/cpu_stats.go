// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type CPUList struct {
	Entries  map[string]CPUEntries  `json:"entries,omitempty"`
	Kind     string                      `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                      `json:"selfLink,omitempty" pretty:",expanded"`
}

type CPUEntries struct {
	NestedStats CPUStatsList `json:"nestedStats,omitempty"`
}

type CPUStatsList struct {
	Entries  map[string]CPUStatsEntries `json:"entries,omitempty"`
	Kind     string                      `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                      `json:"selfLink,omitempty" pretty:",expanded"`
}

type CPUStatsEntries struct {
	NestedStats CPUCoreStatsList `json:"nestedStats,omitempty"`
	Description string `json:"description,omitempty"`
}

type CPUCoreStatsList struct {
	Entries  map[string]CPUCoreStatsEntries `json:"entries,omitempty"`
	Kind     string                      `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                      `json:"selfLink,omitempty" pretty:",expanded"`
}

type CPUCoreStatsEntries struct {
	NestedStats CPUCoreStats `json:"nestedStats,omitempty"`
}

type CPUCoreStats struct {
	Entries struct {
		CpuId struct {
			Value int `json:"value"`
		} `json:"cpuId,omitempty"`
		FiveMinAvgIdle struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgIdle,omitempty"`
		FiveMinAvgIowait struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgIowait,omitempty"`
		FiveMinAvgIrq struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgIrq,omitempty"`
		FiveMinAvgNiced struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgNiced,omitempty"`
		FiveMinAvgSoftirq struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgSoftirq,omitempty"`
		FiveMinAvgStolen struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgStolen,omitempty"`
		FiveMinAvgSystem struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgSystem,omitempty"`
		FiveMinAvgUser struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgUser,omitempty"`
		FiveSecAvgIdle struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgIdle,omitempty"`
		FiveSecAvgIowait struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgIowait,omitempty"`
		FiveSecAvgIrq struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgIrq,omitempty"`
		FiveSecAvgNiced struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgNiced,omitempty"`
		FiveSecAvgSoftirq struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgSoftirq,omitempty"`
		FiveSecAvgStolen struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgStolen,omitempty"`
		FiveSecAvgSystem struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgSystem,omitempty"`
		FiveSecAvgUser struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgUser,omitempty"`
		Idle struct {
			Value int `json:"value"`
		} `json:"idle,omitempty"`
		Iowait struct {
			Value int `json:"value"`
		} `json:"iowait,omitempty"`
		Irq struct {
			Value int `json:"value"`
		} `json:"irq,omitempty"`
		Niced struct {
			Value int `json:"value"`
		} `json:"niced,omitempty"`
		OneMinAvgIdle struct {
			Value int `json:"value"`
		} `json:"oneMinAvgIdle,omitempty"`
		OneMinAvgIowait struct {
			Value int `json:"value"`
		} `json:"oneMinAvgIowait,omitempty"`
		OneMinAvgIrq struct {
			Value int `json:"value"`
		} `json:"oneMinAvgIrq,omitempty"`
		OneMinAvgNiced struct {
			Value int `json:"value"`
		} `json:"oneMinAvgNiced,omitempty"`
		OneMinAvgSoftirq struct {
			Value int `json:"value"`
		} `json:"oneMinAvgSoftirq,omitempty"`
		OneMinAvgStolen struct {
			Value int `json:"value"`
		} `json:"oneMinAvgStolen,omitempty"`
		OneMinAvgSystem struct {
			Value int `json:"value"`
		} `json:"oneMinAvgSystem,omitempty"`
		OneMinAvgUser struct {
			Value int `json:"value"`
		} `json:"oneMinAvgUser,omitempty"`
		Softirq struct {
			Value int `json:"value"`
		} `json:"softirq,omitempty"`
		Stolen struct {
			Value int `json:"value"`
		} `json:"stolen,omitempty"`
		System struct {
			Value int `json:"value"`
		} `json:"system,omitempty"`
		UsageRatio struct {
			Value int `json:"value"`
		} `json:"usageRatio,omitempty"`
		User struct {
			Value int `json:"value"`
		} `json:"user,omitempty"`
	} `json:"entries,omitempty"`
}

// CPUStatsEndpoint represents the REST resource for managing CPUStats.
const CPUStatsEndpoint = "/cpu"

// CPUStatsResource provides an API to manage CPUStats entries.
type CPUStatsResource struct {
	c *f5.Client
}

func (r *CPUStatsResource) All() (*CPUList, error) {
	var list CPUList
	if err := r.c.ReadQuery(BasePath+CPUStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
