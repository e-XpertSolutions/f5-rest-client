// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package f5

import (
	"encoding/json"
	"fmt"
	"io"
)

// A RequestError is returned as a HTTP Response by the F5 Big IP server in case
// of error.
type RequestError struct {
	Code     int      `json:"code,omitempty"`
	Message  string   `json:"message,omitempty"`
	ErrStack []string `json:"errorStack,omitempty"`
}

// NewRequestError unmarshal a RequestError from a HTTP response body.
func NewRequestError(body io.Reader) (*RequestError, error) {
	var reqErr RequestError
	dec := json.NewDecoder(body)
	if err := dec.Decode(&reqErr); err != nil {
		return nil, fmt.Errorf("failed to decode request error: %v", err)
	}
	return &reqErr, nil
}
