package f5

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_SyncStatusDetails_Bug34(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{"kind":"tm:cm:sync-status:sync-statusstats","selfLink":"https://localhost/mgmt/tm/cm/sync-status?ver=14.1.4.6","entries":{"https://localhost/mgmt/tm/cm/sync-status/0":{"nestedStats":{"entries":{"color":{"description":"green"},"https://localhost/mgmt/tm/cm/syncStatus/0/details":{"nestedStats":{"entries":{"https://localhost/mgmt/tm/cm/syncStatus/0/details/0":{"nestedStats":{"entries":{"details":{"description":"i5800b-1.labo.e-xpertsolutions.lan: connected"}}}},"https://localhost/mgmt/tm/cm/syncStatus/0/details/1":{"nestedStats":{"entries":{"details":{"description":" All devices in the device group are in sync"}}}},"https://localhost/mgmt/tm/cm/syncStatus/0/details/2":{"nestedStats":{"entries":{"details":{"description":"In Sync): All devices in the device group are in sync"}}}},"https://localhost/mgmt/tm/cm/syncStatus/0/details/3":{"nestedStats":{"entries":{"details":{"description":" (In Sync): All devices in the device group are in sync"}}}},"https://localhost/mgmt/tm/cm/syncStatus/0/details/4":{"nestedStats":{"entries":{"details":{"description":"datasync-global-dg (In Sync): All devices in the device group are in sync"}}}},"https://localhost/mgmt/tm/cm/syncStatus/0/details/5":{"nestedStats":{"entries":{"details":{"description":"device_trust_group (In Sync): All devices in the device group are in sync"}}}}}}},"mode":{"description":"high-availability"},"status":{"description":"offline"},"summary":{"description":"All devices in the device group are in sync"}}}}}}`)
	}))
	defer ts.Close()

	c := &Client{
		baseURL: ts.URL,
		c:       http.Client{},
		makeAuth: authFunc(func(req *http.Request) error {
			return nil
		}),
	}

	resp, err := c.SyncStatusDetails()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.GroupName != "unknown" {
		t.Errorf("unexpected group name: %v", resp.GroupName)
	}
}
