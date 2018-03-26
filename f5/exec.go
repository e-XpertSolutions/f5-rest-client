package f5

import (
	"encoding/json"
	"fmt"
)

// Bash util REST path.
const (
	PathBashCmd = "/mgmt/tm/util/bash"
)

// ExecOutput represents the output returned by the API afeter having executed a
// bash command.
type ExecOutput struct {
	Kind          string `json:"kind"`
	Command       string `json:"command"`
	CommandResult string `json:"commandResult"`
	UtilCmdArgs   string `json:"utilCmdArgs"`
}

// Exec executes remotely a shell command on the Big-IP.
func (c *Client) Exec(cmd string) (*ExecOutput, error) {
	data := struct {
		Command     string `json:"command"`
		UtilCmdArgs string `json:"utilCmdArgs"`
	}{
		Command:     "run",
		UtilCmdArgs: fmt.Sprintf("-c %q", cmd),
	}
	resp, err := c.SendRequest("POST", PathBashCmd, &data)
	if err != nil {
		return nil, fmt.Errorf("cannot execute command %q: %v", cmd, err)
	}
	defer resp.Body.Close()

	var output ExecOutput
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&output); err != nil {
		return nil, fmt.Errorf("cannot decode output for command %q: %v", cmd, err)
	}

	return &output, nil
}

// ExecTMSH executes a TMSH command on the Big-IP.
func (c *Client) ExecTMSH(cmd string) (*ExecOutput, error) {
	return c.Exec("tmsh " + cmd)
}
