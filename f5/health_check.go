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

// SyncStatusResp contains the values obtained from the sync-status check.
type SyncStatusResp struct {
	Status    string
	Color     string
	Action    string
	GroupName string
}

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

// SyncStatus returns the sync status of the BIG-IP along with the status
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

// SyncStatusDetails returns the sync status and if it is different than
// "In Sync" also the Group Name that is out of sync.
func (c *Client) SyncStatusDetails() (SyncStatusResp, error) {
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
						Details struct {
							NestedStats struct {
								Entries struct {
									Details1 struct {
										NestedStats struct {
											Entries struct {
												Details struct {
													Description string `json:"description"`
												} `json:"details"`
											} `json:"entries"`
										} `json:"nestedStats"`
									} `json:"https://localhost/mgmt/tm/cm/syncStatus/0/details/1"`
									Details2 struct {
										NestedStats struct {
											Entries struct {
												Details struct {
													Description string `json:"description"`
												} `json:"details"`
											} `json:"entries"`
										} `json:"nestedStats"`
									} `json:"https://localhost/mgmt/tm/cm/syncStatus/0/details/2"`
								} `json:"entries"`
							} `json:"nestedStats"`
						} `json:"https://localhost/mgmt/tm/cm/syncStatus/0/details"`
					} `json:"entries"`
				} `json:"nestedStats"`
			} `json:"https://localhost/mgmt/tm/cm/sync-status/0"`
		} `json:"entries"`
	}
	var syncStat SyncStatusResp

	if err := c.ReadQuery(PathSyncStatus, &data); err != nil {
		return syncStat, fmt.Errorf("cannot retrieve sync status: %v", err)
	}

	syncStat.Status = data.Entries.Entry.NestedStats.Entries.Status.Description
	syncStat.Color = data.Entries.Entry.NestedStats.Entries.Color.Description
	if syncStat.Status == "In Sync" {
		return syncStat, nil
	}

	// if the status is different than "In Sync" try to get more details. Normally they should be present in the response details.
	syncStat.Action = data.Entries.Entry.NestedStats.Entries.Details.NestedStats.Entries.Details2.NestedStats.Entries.Details.Description

	description := data.Entries.Entry.NestedStats.Entries.Details.NestedStats.Entries.Details1.NestedStats.Entries.Details.Description
	if description == "" {
		return syncStat, fmt.Errorf("could not retrieve group name")
	}

	syncStat.GroupName = description[:strings.Index(description, "(")-1]

	return syncStat, nil
}
