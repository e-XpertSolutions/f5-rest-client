// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package f5 provides a client for using the F5 API.
package f5

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// An authFunc is function responsible for setting necessary headers to
// perform authenticated requests.
type authFunc func(req *http.Request)

// A Client manages communication with the F5 API.
type Client struct {
	c        http.Client
	baseURL  string
	makeAuth authFunc
}

// NewBasicClient creates a new F5 client with HTTP Basic Authentication.
//
// baseURL is the base URL of the F5 API server.
func NewBasicClient(baseURL, user, password string) (*Client, error) {
	return &Client{
		c:       http.Client{},
		baseURL: baseURL,
		makeAuth: authFunc(func(req *http.Request) {
			req.SetBasicAuth(user, password)
		}),
	}, nil
}

// NewTokenClient creates a new F5 client with token based authentication.
//
// baseURL is the base URL of the F5 API server.
func NewTokenClient(baseURL, user, password, loginRef string, sslCheck bool) (*Client, error) {
	c := Client{c: http.Client{}, baseURL: baseURL}

	// Negociate token with a pair of username/password.
	data, err := json.Marshal(map[string]string{"username": user, "password": password, "loginReference": loginRef})
	if err != nil {
		return nil, fmt.Errorf("failed to create token client (cannot marshal user credentials): %v", err)
	}

	req, err := http.NewRequest("POST", c.makeURL("/mgmt/shared/authn/login"), bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("failed to create token client, (cannot create login request): %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(user, password)

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: sslCheck},
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to create token client (token negociation failed): %v", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("failed to create token client (token negociation failed): http status %s", resp.Status)
	}
	defer resp.Body.Close()

	tok := struct {
		Token struct {
			Token string `json:"token"`
		} `json:"token"`
	}{}
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&tok); err != nil {
		return nil, fmt.Errorf("failed to create token client (cannot decode token): %v", err)
	}

	// Create auth function for token based authentication.
	c.makeAuth = authFunc(func(req *http.Request) {
		req.Header.Add("X-F5-Auth-Token", tok.Token.Token)
	})

	return &c, nil
}

// DisableCertCheck disables certificate verification, meaning that insecure
// certificate will not cause any error.
func (c *Client) DisableCertCheck() {
	c.c.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
}

// MakeRequest creates a request with headers appropriately set to make
// authenticated requests. This method must be called for every new request.
func (c Client) MakeRequest(method, restPath string, data interface{}) (*http.Request, error) {
	var buf *bytes.Buffer
	if data != nil {
		bs, err := json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal data into json: %v", err)
		}
		buf = bytes.NewBuffer(bs)
	}
	req, err := http.NewRequest(method, c.makeURL(restPath), buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create F5 authenticated request: %v", err)
	}
	req.Header.Add("Accept", "application/json")
	c.makeAuth(req)
	return req, nil
}

// Do sends an HTTP request and returns an HTTP response. It is just a wrapper
// arround http.Client Do method.
//
// Callers should close resp.Body when done reading from it.
//
// See http package documentation for more information:
//    https://golang.org/pkg/net/http/#Client.Do
func (c Client) Do(req *http.Request) (*http.Response, error) {
	return c.c.Do(req)
}

// makeURL creates an URL from the client base URL and a given REST path. What
// this function actually does is to concatenate the base URL and the REST path
// by handling trailing slashes.
func (c Client) makeURL(restPath string) string {
	return strings.TrimSuffix(c.baseURL, "/") + "/" + strings.TrimPrefix(restPath, "/")
}
