package f5

import (
	"encoding/json"
	"fmt"
)

// Backup REST paths
const (
	PathBackup = "/mgmt/tm/shared/sys/backup"
)

// BackupResponse holds attributes returned by requests on the backup API.
type BackupResponse struct {
	// Unique ID to identify uniquely the backup action.
	ID string `json:"id"`

	// Name of the file in which the backup is saved to or restord from.
	File string `json:"file"`

	// Type of action performed. Possible values are:
	//    - BACKUP
	//    - RESTORE
	//    - RESTORE_WITH_NO_LICENSE
	//    - BACKUP_WITH_NO_PRIVATE_KEYS
	//    - BACKUP_WITH_ENCRYPTION
	//    - BACKUP_WITH_NO_PRIVATE_KEYS_WITH_ENCRYPTION
	//    - RESTORE_WITH_ENCRYPTION
	//    - RESTORE_WITH_NO_LICENSE_WITH_ENCRYPTION
	//    - CLEANUP
	Action string `json:"action"`

	// Status of the backup. Possible values are:
	//    - CREATED
	//    - STARTED
	//    - CANCEL_REQUESTED
	//    - CANCELED
	//    - FAILED
	//    - FINISHED
	Status string `json:"status"`
}

// Backup creates a backup remotely saved into a file named according to the
// provided filename.
func (c *Client) Backup(filename string) (*BackupResponse, error) {
	return c.backupRequest("BACKUP", filename)
}

// CheckBackup fetches the status of a backup process.
func (c *Client) CheckBackup(id string) (*BackupResponse, error) {
	var backupResp BackupResponse
	if err := c.ReadQuery(PathBackup+"/"+id, &backupResp); err != nil {
		return nil, fmt.Errorf("cannot read backup status: %v", err)
	}
	return &backupResp, nil
}

// RestoreBackup restores a backup from a file having the provided filename and
// located into /var/local/ucs directory.
func (c *Client) RestoreBackup(filename string) (*BackupResponse, error) {
	return c.backupRequest("RESTORE", filename)
}

func (c *Client) backupRequest(action, filename string) (*BackupResponse, error) {
	data := map[string]interface{}{
		"file":   filename,
		"action": action,
	}

	resp, err := c.SendRequest("POST", PathBackup, data)
	if err != nil {
		return nil, fmt.Errorf("error while requesting ucs backup: %v", err)
	}
	defer resp.Body.Close()

	if err := c.ReadError(resp); err != nil {
		return nil, fmt.Errorf("backup api returned an error: %v", err)
	}

	var backupResp BackupResponse
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&backupResp); err != nil {
		return nil, fmt.Errorf("backup api returned a malformed json response: %v", err)
	}

	return &backupResp, nil
}
