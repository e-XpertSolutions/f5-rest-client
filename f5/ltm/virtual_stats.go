// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type VirtualStatsList struct {
	Entries  map[string]VirtualStatsEntries `json:"entries,omitempty"`
	Kind     string                         `json:"kind,omitempty"`
	SelfLink string                         `json:"selflink,omitempty"`
}

type VirtualStatsEntries struct {
	NestedVirtualStats VirtualStats `json:"nestedStats,omitempty"`
}

type VirtualStats struct {
	Entries struct {
		ClientsideBitsIn struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.bitsIn,omitempty"`
		ClientsideBitsOut struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.bitsOut,omitempty"`
		ClientsideCurConns struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.curConns,omitempty"`
		ClientsideEvictedConns struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.evictedConns,omitempty"`
		ClientsideMaxConns struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.maxConns,omitempty"`
		ClientsidePktsIn struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.pktsIn,omitempty"`
		ClientsidePktsOut struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.pktsOut,omitempty"`
		ClientsideSlowKilled struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.slowKilled,omitempty"`
		ClientsideTotConns struct {
			Value int `json:"value,omitempty"`
		} `json:"clientside.totConns,omitempty"`
		CmpEnableMode struct {
			Description string `json:"description,omitempty"`
		} `json:"cmpEnableMode,omitempty"`
		CmpEnabled struct {
			Description string `json:"description,omitempty"`
		} `json:"cmpEnabled,omitempty"`
		CsMaxConnDur struct {
			Value int `json:"value,omitempty"`
		} `json:"csMaxConnDur,omitempty"`
		CsMeanConnDur struct {
			Value int `json:"value,omitempty"`
		} `json:"csMeanConnDur,omitempty"`
		CsMinConnDur struct {
			Value int `json:"value,omitempty"`
		} `json:"csMinConnDur,omitempty"`
		Destination struct {
			Description string `json:"description,omitempty"`
		} `json:"destination,omitempty"`
		EphemeralBitsIn struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.bitsIn,omitempty"`
		EphemeralBitsOut struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.bitsOut,omitempty"`
		EphemeralCurConns struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.curConns,omitempty"`
		EphemeralEvictedConns struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.evictedConns,omitempty"`
		EphemeralMaxConns struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.maxConns,omitempty"`
		EphemeralPktsIn struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.pktsIn,omitempty"`
		EphemeralPktsOut struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.pktsOut,omitempty"`
		EphemeralSlowKilled struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.slowKilled,omitempty"`
		EphemeralTotConns struct {
			Value int `json:"value,omitempty"`
		} `json:"ephemeral.totConns,omitempty"`
		FiveMinAvgUsageRatio struct {
			Value int `json:"value,omitempty"`
		} `json:"fiveMinAvgUsageRatio,omitempty"`
		FiveSecAvgUsageRatio struct {
			Value int `json:"value,omitempty"`
		} `json:"fiveSecAvgUsageRatio,omitempty"`
		OneMinAvgUsageRatio struct {
			Value int `json:"value,omitempty"`
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
			Value int `json:"value,omitempty"`
		} `json:"syncookie.accepts,omitempty"`
		SyncookieHwAccepts struct {
			Value int `json:"value,omitempty"`
		} `json:"syncookie.hwAccepts,omitempty"`
		SyncookieHwSyncookies struct {
			Value int `json:"value,omitempty"`
		} `json:"syncookie.hwSyncookies,omitempty"`
		SyncookieHwsyncookieInstance struct {
			Value int `json:"value,omitempty"`
		} `json:"syncookie.hwsyncookieInstance,omitempty"`
		SyncookieRejects struct {
			Value int `json:"value,omitempty"`
		} `json:"syncookie.rejects,omitempty"`
		SyncookieSwsyncookieInstance struct {
			Value int `json:"value,omitempty"`
		} `json:"syncookie.swsyncookieInstance,omitempty"`
		SyncookieSyncacheCurr struct {
			Value int `json:"value,omitempty"`
		} `json:"syncookie.syncacheCurr,omitempty"`
		SyncookieSyncacheOver struct {
			Value int `json:"value,omitempty"`
		} `json:"syncookie.syncacheOver,omitempty"`
		SyncookieSyncookies struct {
			Value int `json:"value,omitempty"`
		} `json:"syncookie.syncookies,omitempty"`
		SyncookieStatus struct {
			Description string `json:"description,omitempty"`
		} `json:"syncookieStatus,omitempty"`
		TmName struct {
			Description string `json:"description,omitempty"`
		} `json:"tmName,omitempty"`
		TotRequests struct {
			Value int `json:"value,omitempty"`
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
