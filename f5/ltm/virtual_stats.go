// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type VirtualStatsList struct {
	Entries  map[string]VirtualStatsEntries `json:"entries,omitempty"`
	Kind     string                         `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string                         `json:"selflink,omitempty" pretty:",expanded"`
}

type VirtualStatsEntries struct {
	NestedVirtualStats VirtualStats `json:"nestedStats,omitempty"`
}

type VirtualStats struct {
	Entries struct {
		ClientsideBitsIn struct {
			Value int `json:"value"`
		} `json:"clientside.bitsIn,omitempty"`
		ClientsideBitsOut struct {
			Value int `json:"value"`
		} `json:"clientside.bitsOut,omitempty"`
		ClientsideCurConns struct {
			Value int `json:"value"`
		} `json:"clientside.curConns,omitempty"`
		ClientsideEvictedConns struct {
			Value int `json:"value"`
		} `json:"clientside.evictedConns,omitempty"`
		ClientsideMaxConns struct {
			Value int `json:"value"`
		} `json:"clientside.maxConns,omitempty"`
		ClientsidePktsIn struct {
			Value int `json:"value"`
		} `json:"clientside.pktsIn,omitempty"`
		ClientsidePktsOut struct {
			Value int `json:"value"`
		} `json:"clientside.pktsOut,omitempty"`
		ClientsideSlowKilled struct {
			Value int `json:"value"`
		} `json:"clientside.slowKilled,omitempty"`
		ClientsideTotConns struct {
			Value int `json:"value"`
		} `json:"clientside.totConns,omitempty"`
		CmpEnableMode struct {
			Description string `json:"description,omitempty"`
		} `json:"cmpEnableMode,omitempty"`
		CmpEnabled struct {
			Description string `json:"description,omitempty"`
		} `json:"cmpEnabled,omitempty"`
		CsMaxConnDur struct {
			Value int `json:"value"`
		} `json:"csMaxConnDur,omitempty"`
		CsMeanConnDur struct {
			Value int `json:"value"`
		} `json:"csMeanConnDur,omitempty"`
		CsMinConnDur struct {
			Value int `json:"value"`
		} `json:"csMinConnDur,omitempty"`
		Destination struct {
			Description string `json:"description,omitempty"`
		} `json:"destination,omitempty"`
		EphemeralBitsIn struct {
			Value int `json:"value"`
		} `json:"ephemeral.bitsIn,omitempty"`
		EphemeralBitsOut struct {
			Value int `json:"value"`
		} `json:"ephemeral.bitsOut,omitempty"`
		EphemeralCurConns struct {
			Value int `json:"value"`
		} `json:"ephemeral.curConns,omitempty"`
		EphemeralEvictedConns struct {
			Value int `json:"value"`
		} `json:"ephemeral.evictedConns,omitempty"`
		EphemeralMaxConns struct {
			Value int `json:"value"`
		} `json:"ephemeral.maxConns,omitempty"`
		EphemeralPktsIn struct {
			Value int `json:"value"`
		} `json:"ephemeral.pktsIn,omitempty"`
		EphemeralPktsOut struct {
			Value int `json:"value"`
		} `json:"ephemeral.pktsOut,omitempty"`
		EphemeralSlowKilled struct {
			Value int `json:"value"`
		} `json:"ephemeral.slowKilled,omitempty"`
		EphemeralTotConns struct {
			Value int `json:"value"`
		} `json:"ephemeral.totConns,omitempty"`
		FiveMinAvgUsageRatio struct {
			Value int `json:"value"`
		} `json:"fiveMinAvgUsageRatio,omitempty"`
		FiveSecAvgUsageRatio struct {
			Value int `json:"value"`
		} `json:"fiveSecAvgUsageRatio,omitempty"`
		OneMinAvgUsageRatio struct {
			Value int `json:"value"`
		} `json:"oneMinAvgUsageRatio,omitempty"`
		StatusAvailabilityState struct {
			Description string `json:"description,omitempty"`
		} `json:"status.availabilityState,omitempty"`
		StatusEnabledState struct {
			Description string `json:"description,omitempty"`
		} `json:"status.enabledState,omitempty"`
		StatusStatusReason struct {
			Description string `json:"description,omitempty"`
		} `json:"status.statusReason,omitempty"`
		SyncookieAccepts struct {
			Value int `json:"value"`
		} `json:"syncookie.accepts,omitempty"`
		SyncookieHwAccepts struct {
			Value int `json:"value"`
		} `json:"syncookie.hwAccepts,omitempty"`
		SyncookieHwSyncookies struct {
			Value int `json:"value"`
		} `json:"syncookie.hwSyncookies,omitempty"`
		SyncookieHwsyncookieInstance struct {
			Value int `json:"value"`
		} `json:"syncookie.hwsyncookieInstance,omitempty"`
		SyncookieRejects struct {
			Value int `json:"value"`
		} `json:"syncookie.rejects,omitempty"`
		SyncookieSwsyncookieInstance struct {
			Value int `json:"value"`
		} `json:"syncookie.swsyncookieInstance,omitempty"`
		SyncookieSyncacheCurr struct {
			Value int `json:"value"`
		} `json:"syncookie.syncacheCurr,omitempty"`
		SyncookieSyncacheOver struct {
			Value int `json:"value"`
		} `json:"syncookie.syncacheOver,omitempty"`
		SyncookieSyncookies struct {
			Value int `json:"value"`
		} `json:"syncookie.syncookies,omitempty"`
		SyncookieStatus struct {
			Description string `json:"description,omitempty"`
		} `json:"syncookieStatus,omitempty"`
		TmName struct {
			Description string `json:"description,omitempty"`
		} `json:"tmName,omitempty"`
		TotRequests struct {
			Value int `json:"value"`
		} `json:"totRequests,omitempty"`
	} `json:"entries,omitempty"`
}

// VirtualStatsEndpoint represents the REST resource for managing VirtualStats.
const VirtualStatsEndpoint = "/virtual/stats"

// VirtualStatsResource provides an API to manage VirtualStats configurations.
type VirtualStatsResource struct {
	c *f5.Client
}

func (r *VirtualStatsResource) All() (*VirtualStatsList, error) {
	var list VirtualStatsList
	if err := r.c.ReadQuery(BasePath+VirtualStatsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}
