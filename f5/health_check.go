package f5

import (
	"errors"
	"fmt"
)

// Cluster Management REST paths.
const (
	PathDeviceInfo = "/mgmt/tm/cm/device"
)

// IsActive returns true whether the BigIP is active and the iControl REST are
// accessible. In case of error, false is returned.
func (c *Client) IsActive() bool {
	var respData struct {
		Items []struct {
			FailoverState string `json:"failoverState"`
		} `json:"items"`
	}
	if err := c.ReadQuery(PathDeviceInfo, &respData); err != nil {
		return false
	}
	if respData.Items == nil || len(respData.Items) == 0 {
		return false
	}
	return respData.Items[0].FailoverState == "active"
}

// FailoverState returns the status of the BigIP (active, standby,
// forced-offline, ...).
func (c *Client) FailoverState() (string, error) {
	var respData struct {
		Items []struct {
			FailoverState string `json:"failoverState"`
		} `json:"items"`
	}
	if err := c.ReadQuery(PathDeviceInfo, &respData); err != nil {
		return "", fmt.Errorf("cannot retrieve failover state: %v", err)
	}
	if respData.Items == nil || len(respData.Items) == 0 {
		return "", errors.New("no failover state found")
	}
	return respData.Items[0].FailoverState, nil
}
