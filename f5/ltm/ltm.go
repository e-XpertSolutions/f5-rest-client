// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ltm provides a REST client for the /tm/ltm F5 BigIP API.
package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// BasePath is the base path of the LTM API.
const BasePath = "mgmt/tm/ltm"

// LTM implements a REST client for the F5 BigIP LTM API.
type LTM struct {
	c *f5.Client

	virtual           VirtualResource
	virtualStats      VirtualStatsResource
	rule              RuleResource
	pool              PoolResource
	poolStats         PoolStatsResource
	poolMembers       PoolMembersResource
	node              NodeResource
	nodeStats         NodeStatsResource
	iFile             IFileResource
	dataGroupInternal DataGroupInternalResource

	profileClientSSL ProfileClientSSLResource

	monitorDiameter         MonitorDiameterResource
	monitorDNS              MonitorDNSResource
	monitorExternal         MonitorExternalResource
	monitorFirepass         MonitorFirepassResource
	monitorFTP              MonitorFTPResource
	monitorGatewayICMP      MonitorGatewayICMPResource
	monitorHTTP             MonitorHTTPResource
	monitorHTTPS            MonitorHTTPSResource
	monitorICMP             MonitorICMPResource
	monitorIMAP             MonitorIMAPResource
	monitorInband           MonitorInbandResource
	monitorLDAP             MonitorLDAPResource
	monitorModuleScore      MonitorModuleScoreResource
	monitorMSSQL            MonitorMSSQLResource
	monitorMySQL            MonitorMySQLResource
	monitorNNTP             MonitorNNTPResource
	monitorOracle           MonitorOracleResource
	monitorPOP3             MonitorPOP3Resource
	monitorPostgreSQL       MonitorPostgreSQLResource
	monitorRadiusAccounting MonitorRadiusAccountingResource
	monitorRadius           MonitorRadiusResource
	monitorRealServer       MonitorRealServerResource
	monitorRPC              MonitorRPCResource
	monitorSASP             MonitorSASPResource
	monitorScripted         MonitorScriptedResource
	monitorSIP              MonitorSIPResource
	monitorSMB              MonitorSMBResource
	monitorSMTP             MonitorSMTPResource
	monitorSNMPDCABase      MonitorSNMPDCABaseResource
	monitorSNMPDCA          MonitorSNMPDCAResource
	monitorSOAP             MonitorSOAPResource
	monitorTCPEcho          MonitorTCPEchoResource
	monitorTCP              MonitorTCPResource
	monitorTCPHalfOpen      MonitorTCPHalfOpenResource
	monitorUDP              MonitorUDPResource
	monitorVirtualLocation  MonitorVirtualLocationResource
	monitorWAP              MonitorWAPResource
	monitorWMI              MonitorWMIResource
}

// New creates a new LTM client.
func New(c *f5.Client) LTM {
	return LTM{
		c:                 c,
		virtual:           VirtualResource{c: c},
		virtualStats:      VirtualStatsResource{c: c},
		rule:              RuleResource{c: c},
		pool:              PoolResource{c: c},
		poolStats:         PoolStatsResource{c: c},
		poolMembers:       PoolMembersResource{c: c},
		node:              NodeResource{c: c},
		nodeStats:         NodeStatsResource{c: c},
		iFile:             IFileResource{c: c},
		dataGroupInternal: DataGroupInternalResource{c: c},

		profileClientSSL: ProfileClientSSLResource{c: c},

		monitorDiameter:         MonitorDiameterResource{c: c},
		monitorDNS:              MonitorDNSResource{c: c},
		monitorExternal:         MonitorExternalResource{c: c},
		monitorFirepass:         MonitorFirepassResource{c: c},
		monitorFTP:              MonitorFTPResource{c: c},
		monitorGatewayICMP:      MonitorGatewayICMPResource{c: c},
		monitorHTTP:             MonitorHTTPResource{c: c},
		monitorHTTPS:            MonitorHTTPSResource{c: c},
		monitorICMP:             MonitorICMPResource{c: c},
		monitorIMAP:             MonitorIMAPResource{c: c},
		monitorInband:           MonitorInbandResource{c: c},
		monitorLDAP:             MonitorLDAPResource{c: c},
		monitorModuleScore:      MonitorModuleScoreResource{c: c},
		monitorMSSQL:            MonitorMSSQLResource{c: c},
		monitorMySQL:            MonitorMySQLResource{c: c},
		monitorNNTP:             MonitorNNTPResource{c: c},
		monitorOracle:           MonitorOracleResource{c: c},
		monitorPOP3:             MonitorPOP3Resource{c: c},
		monitorPostgreSQL:       MonitorPostgreSQLResource{c: c},
		monitorRadiusAccounting: MonitorRadiusAccountingResource{c: c},
		monitorRadius:           MonitorRadiusResource{c: c},
		monitorRealServer:       MonitorRealServerResource{c: c},
		monitorRPC:              MonitorRPCResource{c: c},
		monitorSASP:             MonitorSASPResource{c: c},
		monitorScripted:         MonitorScriptedResource{c: c},
		monitorSIP:              MonitorSIPResource{c: c},
		monitorSMB:              MonitorSMBResource{c: c},
		monitorSMTP:             MonitorSMTPResource{c: c},
		monitorSNMPDCABase:      MonitorSNMPDCABaseResource{c: c},
		monitorSNMPDCA:          MonitorSNMPDCAResource{c: c},
		monitorSOAP:             MonitorSOAPResource{c: c},
		monitorTCPEcho:          MonitorTCPEchoResource{c: c},
		monitorTCP:              MonitorTCPResource{c: c},
		monitorTCPHalfOpen:      MonitorTCPHalfOpenResource{c: c},
		monitorUDP:              MonitorUDPResource{c: c},
		monitorVirtualLocation:  MonitorVirtualLocationResource{c: c},
		monitorWAP:              MonitorWAPResource{c: c},
		monitorWMI:              MonitorWMIResource{c: c},
	}
}

// Virtual returns a VirtualResource configured to query tm/ltm/virtual API.
func (ltm LTM) Virtual() *VirtualResource {
	return &ltm.virtual
}

func (ltm LTM) VirtualStats() *VirtualStatsResource {
	return &ltm.virtualStats
}

// Rule returns a RuleResource configured to query tm/ltm/rule API.
func (ltm LTM) Rule() *RuleResource {
	return &ltm.rule
}

// Pool returns a PoolResource configured to query /tm/ltm/pool API.
func (ltm LTM) Pool() *PoolResource {
	return &ltm.pool
}

func (ltm LTM) PoolStats() *PoolStatsResource {
	return &ltm.poolStats
}

// PoolMembers returns a PoolMembersResource configured to query /tm/ltm/pool API.
func (ltm LTM) PoolMembers() *PoolMembersResource {
	return &ltm.poolMembers
}

// Node returns a NodeResource configured to query /tm/ltm/node API.
func (ltm LTM) Node() *NodeResource {
	return &ltm.node
}

func (ltm LTM) NodeStats() *NodeStatsResource {
	return &ltm.nodeStats
}

// IFile returns an IFileResource configured to query /tm/ltm/ifile API.
func (ltm LTM) IFile() *IFileResource {
	return &ltm.iFile
}

// DataGroupInternal returns a DataGroupInternalResource configured to query /tm/ltm/data-group/internal API.
func (ltm LTM) DataGroupInternal() *DataGroupInternalResource {
	return &ltm.dataGroupInternal
}

func (ltm LTM) MonitorDiameter() *MonitorDiameterResource {
	return &ltm.monitorDiameter
}

func (ltm LTM) MonitorDNS() *MonitorDNSResource {
	return &ltm.monitorDNS
}

func (ltm LTM) MonitorExternal() *MonitorExternalResource {
	return &ltm.monitorExternal
}

func (ltm LTM) MonitorFirepass() *MonitorFirepassResource {
	return &ltm.monitorFirepass
}

func (ltm LTM) MonitorFTP() *MonitorFTPResource {
	return &ltm.monitorFTP
}

func (ltm LTM) MonitorGatewayICMP() *MonitorGatewayICMPResource {
	return &ltm.monitorGatewayICMP
}

func (ltm LTM) MonitorHTTP() *MonitorHTTPResource {
	return &ltm.monitorHTTP
}

func (ltm LTM) MonitorHTTPS() *MonitorHTTPSResource {
	return &ltm.monitorHTTPS
}

func (ltm LTM) MonitorICMP() *MonitorICMPResource {
	return &ltm.monitorICMP
}

func (ltm LTM) MonitorIMAP() *MonitorIMAPResource {
	return &ltm.monitorIMAP
}

func (ltm LTM) MonitorInband() *MonitorInbandResource {
	return &ltm.monitorInband
}

func (ltm LTM) MonitorLDAP() *MonitorLDAPResource {
	return &ltm.monitorLDAP
}

func (ltm LTM) MonitorModuleScore() *MonitorModuleScoreResource {
	return &ltm.monitorModuleScore
}

func (ltm LTM) MonitorMSSQL() *MonitorMSSQLResource {
	return &ltm.monitorMSSQL
}

func (ltm LTM) MonitorMySQL() *MonitorMySQLResource {
	return &ltm.monitorMySQL
}

func (ltm LTM) MonitorNNTP() *MonitorNNTPResource {
	return &ltm.monitorNNTP
}

func (ltm LTM) MonitorOracle() *MonitorOracleResource {
	return &ltm.monitorOracle
}

func (ltm LTM) MonitorPOP3() *MonitorPOP3Resource {
	return &ltm.monitorPOP3
}

func (ltm LTM) MonitorPostgreSQL() *MonitorPostgreSQLResource {
	return &ltm.monitorPostgreSQL
}

func (ltm LTM) MonitorRadiusAccounting() *MonitorRadiusAccountingResource {
	return &ltm.monitorRadiusAccounting
}

func (ltm LTM) MonitorRadius() *MonitorRadiusResource {
	return &ltm.monitorRadius
}

func (ltm LTM) MonitorRealServer() *MonitorRealServerResource {
	return &ltm.monitorRealServer
}

func (ltm LTM) MonitorRPC() *MonitorRPCResource {
	return &ltm.monitorRPC
}

func (ltm LTM) MonitorSASP() *MonitorSASPResource {
	return &ltm.monitorSASP
}

func (ltm LTM) MonitorScripted() *MonitorScriptedResource {
	return &ltm.monitorScripted
}

func (ltm LTM) MonitorSIP() *MonitorSIPResource {
	return &ltm.monitorSIP
}

func (ltm LTM) MonitorSMB() *MonitorSMBResource {
	return &ltm.monitorSMB
}

func (ltm LTM) MonitorSMTP() *MonitorSMTPResource {
	return &ltm.monitorSMTP
}

func (ltm LTM) MonitorSNMPDCABase() *MonitorSNMPDCABaseResource {
	return &ltm.monitorSNMPDCABase
}

func (ltm LTM) MonitorSNMPDCA() *MonitorSNMPDCAResource {
	return &ltm.monitorSNMPDCA
}

func (ltm LTM) MonitorSOAP() *MonitorSOAPResource {
	return &ltm.monitorSOAP
}

func (ltm LTM) MonitorTCPEcho() *MonitorTCPEchoResource {
	return &ltm.monitorTCPEcho
}

func (ltm LTM) MonitorTCP() *MonitorTCPResource {
	return &ltm.monitorTCP
}

func (ltm LTM) MonitorTCPHalfOpen() *MonitorTCPHalfOpenResource {
	return &ltm.monitorTCPHalfOpen
}

func (ltm LTM) MonitorUDP() *MonitorUDPResource {
	return &ltm.monitorUDP
}

func (ltm LTM) MonitorVirtualLocation() *MonitorVirtualLocationResource {
	return &ltm.monitorVirtualLocation
}

func (ltm LTM) MonitorWAP() *MonitorWAPResource {
	return &ltm.monitorWAP
}

func (ltm LTM) MonitorWMI() *MonitorWMIResource {
	return &ltm.monitorWMI
}

func (ltm LTM) ProfileClientSSL() *ProfileClientSSLResource {
	return &ltm.profileClientSSL
}
