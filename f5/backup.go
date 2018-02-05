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

// IsDone reports whether the status indicates that the action is terminated,
// even if it is an error or that the task has been canceled.
func (resp BackupResponse) IsDone() bool {
	switch resp.Status {
	case "STARTED", "CANCEL_REQUESTED":
		return false
	}
	return true
}

// IsFailure reports whether the status is FAILED.
func (resp BackupResponse) IsFailure() bool {
	return resp.Status == "FAILED"
}

// IsCanceled reports whether the status is CANCELED.
func (resp BackupResponse) IsCanceled() bool {
	return resp.Status == "CANCELED"
}

// IsSuccess reports whether the status is FINISHED.
func (resp BackupResponse) IsSuccess() bool {
	return resp.Status == "FINISHED"
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

// RestoreBackupWithNoLicense works exactly as RestoreBackup but do no check the
// license.
func (c *Client) RestoreBackupWithNoLicense(filename string) (*BackupResponse, error) {
	return c.backupRequest("RESTORE_WITH_NO_LICENSE", filename)
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
