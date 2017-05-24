package elysium

import (
	"encoding/json"
	"fmt"
)

type Backup struct {
	Name         string `json:"-"`
	BuildTag     string `json:"BUILD_TAG"`
	BuildURL     string `json:"BUILD_URL"`
	EndpointUUID string `json:"endpoint_uuid"`
	Folder       string `json:"folder"`
	Size         int64  `json:"size"`
	Timestamp    int64  `json:"timestamp"`
	TotalDirs    int64  `json:"total_dirs"`
	TotalEntries int64  `json:"total_entries"`
	TotalFiles   int64  `json:"total_files"`
	TotalSize    int64  `json:"total_size"`
	TTL          int64  `json:"ttl"`
}

type BackupList struct {
	EnvironmentName string
	SiteID          string
	Backups         map[string]Backup
}

// Path returns the API endpoint which can be used to get a BackupList for the current user.
func (bl BackupList) Path(method string, auth AuthSession) string {
	return fmt.Sprintf("sites/%s/environments", bl.SiteID)
}

// JSON prepares the BackupList for HTTP transport.
func (bl BackupList) JSON() ([]byte, error) {
	return json.Marshal(bl.Backups)
}

// Unmarshal is responsible for converting a HTTP response into a BackupList struct.
func (bl *BackupList) Unmarshal(data []byte) error {
	err := json.Unmarshal(data, &bl.Backups)
	if err != nil {
		return err
	}

	if len(bl.Backups) > 0 {
		for name, env := range bl.Backups {
			env.Name = name
			bl.Backups[name] = env
		}
	}

	return nil
}
