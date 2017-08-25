// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// SyslogConfig holds the configuration of a single Syslog.
type SyslogConfig struct {
	AuthPrivFrom         string         `json:"authPrivFrom,omitempty"`
	AuthPrivTo           string         `json:"authPrivTo,omitempty"`
	ClusteredHostSlot    string         `json:"clusteredHostSlot,omitempty"`
	ClusteredMessageSlot string         `json:"clusteredMessageSlot,omitempty"`
	ConsoleLog           string         `json:"consoleLog,omitempty"`
	CronFrom             string         `json:"cronFrom,omitempty"`
	CronTo               string         `json:"cronTo,omitempty"`
	DaemonFrom           string         `json:"daemonFrom,omitempty"`
	DaemonTo             string         `json:"daemonTo,omitempty"`
	IsoDate              string         `json:"isoDate,omitempty"`
	KernFrom             string         `json:"kernFrom,omitempty"`
	KernTo               string         `json:"kernTo,omitempty"`
	Kind                 string         `json:"kind,omitempty"`
	Local6From           string         `json:"local6From,omitempty"`
	Local6To             string         `json:"local6To,omitempty"`
	MailFrom             string         `json:"mailFrom,omitempty"`
	MailTo               string         `json:"mailTo,omitempty"`
	MessagesFrom         string         `json:"messagesFrom,omitempty"`
	MessagesTo           string         `json:"messagesTo,omitempty"`
	RemoteServers        []RemoteServer `json:"remoteServers,omitempty"`
	SelfLink             string         `json:"selfLink,omitempty"`
	UserLogFrom          string         `json:"userLogFrom,omitempty"`
	UserLogTo            string         `json:"userLogTo,omitempty"`
}

type RemoteServer struct {
	Host       string `json:"host,omitempty"`
	LocalIP    string `json:"localIp,omitempty"`
	Name       string `json:"name,omitempty"`
	RemotePort int    `json:"remotePort,omitempty"`
}

// SyslogEndpoint represents the REST resource for managing Syslog.
const SyslogEndpoint = "/syslog"

// SyslogResource provides an API to manage Syslog configurations.
type SyslogResource struct {
	c *f5.Client
}

// Show the Syslog configuration.
func (r *SyslogResource) Show() (*SyslogConfig, error) {
	var list SyslogConfig
	if err := r.c.ReadQuery(BasePath+SyslogEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Edit a Syslog configuration identified by id.
func (r *SyslogResource) Edit(item SyslogConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+SyslogEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *SyslogResource) AddRemoteServers(rs ...RemoteServer) error {
	if len(rs) == 0 {
		return nil
	}
	var item SyslogConfig
	if err := r.c.ReadQuery(BasePath+SyslogEndpoint, &item); err != nil {
		return err
	}

	item.RemoteServers = append(item.RemoteServers, rs...)

	if err := r.c.ModQuery("PUT", BasePath+SyslogEndpoint, item); err != nil {
		return err
	}
	return nil
}
