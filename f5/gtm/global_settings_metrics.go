// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gtm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// GlobalSettingsMetrics holds the configuration of a single GlobalSettingsMetrics.
type GlobalSettingsMetrics struct {
	DefaultProbeLimit             int      `json:"defaultProbeLimit,omitempty"`
	HopsPacketLength              int      `json:"hopsPacketLength,omitempty"`
	HopsSampleCount               int      `json:"hopsSampleCount,omitempty"`
	HopsTimeout                   int      `json:"hopsTimeout,omitempty"`
	HopsTTL                       int      `json:"hopsTtl,omitempty"`
	InactiveLdnsTTL               int      `json:"inactiveLdnsTtl,omitempty"`
	InactivePathsTTL              int      `json:"inactivePathsTtl,omitempty"`
	Kind                          string   `json:"kind,omitempty"`
	LdnsUpdateInterval            int      `json:"ldnsUpdateInterval,omitempty"`
	MaxSynchronousMonitorRequests int      `json:"maxSynchronousMonitorRequests,omitempty"`
	MetricsCaching                int      `json:"metricsCaching,omitempty"`
	MetricsCollectionProtocols    []string `json:"metricsCollectionProtocols,omitempty"`
	PathTTL                       int      `json:"pathTtl,omitempty"`
	PathsRetry                    int      `json:"pathsRetry,omitempty"`
	SelfLink                      string   `json:"selfLink,omitempty"`
}

// GlobalSettingsMetricsEndpoint represents the REST resource for managing GlobalSettingsMetrics.
const GlobalSettingsMetricsEndpoint = "/global-settings/metrics"

// GlobalSettingsMetricsResource provides an API to manage GlobalSettingsMetrics configurations.
type GlobalSettingsMetricsResource struct {
	c *f5.Client
}

// List  lists all the GlobalSettingsMetrics configurations.
func (r *GlobalSettingsMetricsResource) List() (*GlobalSettingsMetrics, error) {
	var list GlobalSettingsMetrics
	if err := r.c.ReadQuery(BasePath+GlobalSettingsMetricsEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Modify a new GlobalSettingsMetrics configuration.
func (r *GlobalSettingsMetricsResource) Modify(item GlobalSettingsMetrics) error {
	if err := r.c.ModQuery("POST", BasePath+GlobalSettingsMetricsEndpoint, item); err != nil {
		return err
	}
	return nil
}
