package f5

import (
	"fmt"
	"strings"
)

// Cluster Management REST paths.
const (
	PathDeviceInfo = "/mgmt/tm/cm/device"
	PathSyncStatus = "/mgmt/tm/cm/sync-status"
)

// IsActive returns true whether the BigIP is active and the iControl REST are
// accessible. In case of error, false is returned.
func (c *Client) IsActive(host string) bool {
	var respData struct {
		FailoverState string `json:"failoverState"`
	}
	if err := c.ReadQuery(PathDeviceInfo+"/"+host, &respData); err != nil {
		return false
	}
	return respData.FailoverState == "active"
}

// FailoverState returns the status of the BigIP (active, standby,
// forced-offline, ...).
func (c *Client) FailoverState(host, ip string) (string, error) {
	var respData struct {
		Items []struct {
			FailoverState string `json:"failoverState"`
			Name          string `json:"name"`
			Hostname      string `json:"hostname"`
			IP            string `json:"managementIp"`
		} `json:"items"`
	}
	if err := c.ReadQuery(PathDeviceInfo, &respData); err != nil {
		return "unreachable", fmt.Errorf("cannot retrieve failover state: %v", err)
	}
	for _, item := range respData.Items {
		if strings.HasPrefix(item.Name, host) || strings.HasPrefix(item.Hostname, host) {
			if item.FailoverState == "" {
				return "invalid", nil
			}
			return item.FailoverState, nil
		}
	}
	if ip != "" {
		if pos := strings.Index(ip, "/"); pos != -1 {
			ip = ip[:pos]
		}
		for _, item := range respData.Items {
			if strings.HasPrefix(item.IP, ip) {
				if item.FailoverState == "" {
					return "invalid", nil
				}
				return item.FailoverState, nil
			}
		}
	}
	return "invalid hostname", fmt.Errorf("hostname %q and management ip %q mismatch the ones retrieved", host, ip)
}

// SSyncStatus returns the sync status of the BIG-IP along with the status
// color.
func (c *Client) SyncStatus() (status string, color string, err error) {
	var data struct {
		Entries struct {
			Entry struct {
				NestedStats struct {
					Entries struct {
						Color struct {
							Description string `json:"description"`
						} `json:"color"`
						Status struct {
							Description string `json:"description"`
						} `json:"status"`
					} `json:"entries"`
				} `json:"nestedStats"`
			} `json:"https://localhost/mgmt/tm/cm/sync-status/0"`
		} `json:"entries"`
	}
	if err := c.ReadQuery(PathSyncStatus, &data); err != nil {
		return "", "", fmt.Errorf("cannot retrieve sync status: %v", err)
	}
	status = data.Entries.Entry.NestedStats.Entries.Status.Description
	color = data.Entries.Entry.NestedStats.Entries.Color.Description
	return
}
