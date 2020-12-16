package f5

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

//
const (
	PathConfigSync = "/mgmt/tm/cm/config-sync"
)

// ConfigSyncOptions represents available parameters for config-sync query.
type ConfigSyncOptions struct {
	FromGroup         string `mapstructure:",omitempty"`
	ToGroup           string `mapstructure:",omitempty"`
	RecoverSync       bool   `mapstructure:",omitempty"`
	ForceFullLoadPush bool   `mapstructure:",omitempty"`
}

// ConfigSyncOption is a function prototype that sets the
type ConfigSyncOption func(*ConfigSyncOptions)

// WithFromGroup sets the name of from-group parameter.
func WithFromGroup(name string) ConfigSyncOption {
	return func(o *ConfigSyncOptions) {
		o.FromGroup = name
	}
}

// WithToGroup sets the name of to-group parameter.
func WithToGroup(name string) ConfigSyncOption {
	return func(o *ConfigSyncOptions) {
		o.ToGroup = name
	}
}

// WithRecoverSync sets recover-sync parameter to true.
func WithRecoverSync() ConfigSyncOption {
	return func(o *ConfigSyncOptions) {
		o.RecoverSync = true
	}
}

// WithForceFullLoadPush sets force-full-load-push parameter to true.
func WithForceFullLoadPush() ConfigSyncOption {
	return func(o *ConfigSyncOptions) {
		o.ForceFullLoadPush = true
	}
}

// ConfigSync performs the config-sync operation. It only starts the sync
// and does not wait for it to complete. The synchronization status must be
// check manually.
func (c *Client) ConfigSync(opts ...ConfigSyncOption) error {

	options := ConfigSyncOptions{}
	for _, o := range opts {
		o(&options)
	}

	var reqOptions = struct {
		Command string
		Options []map[string]string
	}{Command: "run", Options: make([]map[string]string, 0)}

	mapOpts := make(map[string]string)
	err := mapstructure.Decode(options, &mapOpts)
	if err != nil {
		return fmt.Errorf("f5#ConfigSync: failed to transform options: %w", err)
	}
	for k, v := range mapOpts {
		reqOptions.Options = append(reqOptions.Options, map[string]string{k: v})
	}

	resp, err := c.SendRequest("POST", PathConfigSync, reqOptions)
	if err != nil {
		return fmt.Errorf("f5#ConfigSync: failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("f5#ConfigSync: failed to read response body: %w", err)
		}
		return fmt.Errorf("f5#ConfigSync: server returned non 200 response: %s, %s", resp.Status, string(body))
	}

	return nil
}
