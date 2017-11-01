// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package util

import (
	"encoding/json"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

type BashCommand struct {
	Command         string `json:"command,omitempty"`
	UtilCommandArgs string `json:"utilCmdArgs,omitempty"`
	CommandResult   string `json:"commandResult,omitempty"`
}

// BashEndpoint represents the REST resource used to execute bash commands.
const BashEndpoint = "/bash"

// A BashResource provides an API to run bash commands
type BashResource struct {
	c *f5.Client
}

func (br *BashResource) Run(item BashCommand) (*BashCommand, error) {
	resp, err := br.c.MakeRequest("POST", BasePath+BashEndpoint, item)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var respItem BashCommand
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&respItem); err != nil {
		return nil, err
	}
	return &respItem, nil
}
