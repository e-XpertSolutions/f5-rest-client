package f5

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
