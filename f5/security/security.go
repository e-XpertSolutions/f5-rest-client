// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

const BasePath = "mgmt/tm/security"

type Security struct {
	c *f5.Client

	analytics                                 AnalyticsResource
	antiFraud                                 AntiFraudResource
	antiFraudProfile                          AntiFraudProfileResource
	antiFraudProfileUsers                     AntiFraudProfileUsersResource
	antiFraudProfileUsersModes                AntiFraudProfileUsersModesResource
	antiFraudSignaturesUpdate                 AntiFraudSignaturesUpdateResource
	blacklistPublisher                        BlacklistPublisherResource
	blacklistPublisherBlacklistPublisherStats BlacklistPublisherBlacklistPublisherStatsResource
	blacklistPublisherCategory                BlacklistPublisherCategoryResource
	blacklistPublisherProfile                 BlacklistPublisherProfileResource
	botDefense                                BotDefenseResource
	botDefenseASMProfile                      BotDefenseASMProfileResource
	dNS                                       DNSResource
	dNSProfile                                DNSProfileResource
	dOS                                       DOSResource
	dOSBotSignature                           DOSBotSignatureResource
	dOSBotSignatureCategory                   DOSBotSignatureCategoryResource
	dOSDeviceConfig                           DOSDeviceConfigResource
	dOSNetworkWhitelist                       DOSNetworkWhitelistResource
	dOSProfile                                DOSProfileResource
	dOSProfileApplication                     DOSProfileApplicationResource
	dOSProfileDOSNetwork                      DOSProfileDOSNetworkResource
	dOSProfileProtocolDNS                     DOSProfileProtocolDNSResource
	dOSProfileProtocolSIP                     DOSProfileProtocolSIPResource
	dOSUDPPortlist                            DOSUDPPortlistResource
	datasync                                  DatasyncResource
	datasyncDeviceStats                       DatasyncDeviceStatsResource
	datasyncGlobalProfile                     DatasyncGlobalProfileResource
	datasyncLocalProfile                      DatasyncLocalProfileResource
	datasyncUpdateFile                        DatasyncUpdateFileResource
	device                                    DeviceResource
	deviceDeviceContext                       DeviceDeviceContextResource
	deviceDeviceContextNATRule                DeviceDeviceContextNATRuleResource
	engineUpdate                              EngineUpdateResource
	firewall                                  FirewallResource
	firewallAddressList                       FirewallAddressListResource
	firewallConfigChangeLog                   FirewallConfigChangeLogResource
	firewallFQDNEntity                        FirewallFQDNEntityResource
	firewallFQDNInfo                          FirewallFQDNInfoResource
	firewallGlobalFQDNPolicy                  FirewallGlobalFQDNPolicyResource
	firewallGlobalRules                       FirewallGlobalRulesResource
	firewallGlobalRulesActive                 FirewallGlobalRulesActiveResource
	firewallGlobalRulesEnforcedPolicyRules    FirewallGlobalRulesEnforcedPolicyRulesResource
	firewallGlobalRulesStagedPolicyRules      FirewallGlobalRulesStagedPolicyRulesResource
	firewallManagementIPRules                 FirewallManagementIPRulesResource
	firewallMatchingRule                      FirewallMatchingRuleResource
	firewallOnDemandCompilation               FirewallOnDemandCompilationResource
	firewallOnDemandRuleDeploy                FirewallOnDemandRuleDeployResource
	firewallPolicy                            FirewallPolicyResource
	firewallPolicyRules                       FirewallPolicyRulesResource
	firewallPortList                          FirewallPortListResource
	firewallPortMisusePolicy                  FirewallPortMisusePolicyResource
	firewallRuleList                          FirewallRuleListResource
	firewallRuleListRules                     FirewallRuleListRulesResource
	firewallSchedule                          FirewallScheduleResource
	firewallUserDomain                        FirewallUserDomainResource
	firewallUserGroupEntity                   FirewallUserGroupEntityResource
	firewallUserList                          FirewallUserListResource
	hTTP                                      HTTPResource
	hTTPFileType                              HTTPFileTypeResource
	hTTPMandatoryHeader                       HTTPMandatoryHeaderResource
	hTTPProfile                               HTTPProfileResource
	iPIntelligence                            IPIntelligenceResource
	iPIntelligenceFeedList                    IPIntelligenceFeedListResource
	iPIntelligenceGlobalPolicy                IPIntelligenceGlobalPolicyResource
	iPIntelligenceInfo                        IPIntelligenceInfoResource
	iPIntelligencePolicy                      IPIntelligencePolicyResource
	iPIntelligenseBlacklistCategory           IPIntelligenseBlacklistCategoryResource
	iPIntelligenseCategory                    IPIntelligenseCategoryResource
	log                                       LogResource
	logNetworkStorageField                    LogNetworkStorageFieldResource
	logProfile                                LogProfileResource
	logProfileApplication                     LogProfileApplicationResource
	logProfileNetwork                         LogProfileNetworkResource
	logProfileProtocolDNS                     LogProfileProtocolDNSResource
	logProfileProtocolSIP                     LogProfileProtocolSIPResource
	logProtocolDNSStorageField                LogProtocolDNSStorageFieldResource
	logProtocolSIPStorageField                LogProtocolSIPStorageFieldResource
	logRemoteFormat                           LogRemoteFormatResource
	logStorageField                           LogStorageFieldResource
	managementIPRulesRules                    ManagementIPRulesRulesResource
	nAT                                       NATResource
	nATDestinationTranslation                 NATDestinationTranslationResource
	nATPolicy                                 NATPolicyResource
	nATPolicyRules                            NATPolicyRulesResource
	nATSourceTranslation                      NATSourceTranslationResource
	sSH                                       SSHResource
	sSHProfile                                SSHProfileResource
	sSHProfileAuthInfo                        SSHProfileAuthInfoResource
	sSHProfileRules                           SSHProfileRulesResource
	sSHSSHPluginStats                         SSHSSHPluginStatsResource
	security                                  SecurityResource
	settings                                  SettingsResource
}

func New(c *f5.Client) Security {
	return Security{
		c: c,

		analytics:                                 AnalyticsResource{c: c},
		antiFraud:                                 AntiFraudResource{c: c},
		antiFraudProfile:                          AntiFraudProfileResource{c: c},
		antiFraudProfileUsers:                     AntiFraudProfileUsersResource{c: c},
		antiFraudProfileUsersModes:                AntiFraudProfileUsersModesResource{c: c},
		antiFraudSignaturesUpdate:                 AntiFraudSignaturesUpdateResource{c: c},
		blacklistPublisher:                        BlacklistPublisherResource{c: c},
		blacklistPublisherBlacklistPublisherStats: BlacklistPublisherBlacklistPublisherStatsResource{c: c},
		blacklistPublisherCategory:                BlacklistPublisherCategoryResource{c: c},
		blacklistPublisherProfile:                 BlacklistPublisherProfileResource{c: c},
		botDefense:                                BotDefenseResource{c: c},
		botDefenseASMProfile:                      BotDefenseASMProfileResource{c: c},
		dNS:                                       DNSResource{c: c},
		dNSProfile:                                DNSProfileResource{c: c},
		dOS:                                       DOSResource{c: c},
		dOSBotSignature:                           DOSBotSignatureResource{c: c},
		dOSBotSignatureCategory:                   DOSBotSignatureCategoryResource{c: c},
		dOSDeviceConfig:                           DOSDeviceConfigResource{c: c},
		dOSNetworkWhitelist:                       DOSNetworkWhitelistResource{c: c},
		dOSProfile:                                DOSProfileResource{c: c},
		dOSProfileApplication:                     DOSProfileApplicationResource{c: c},
		dOSProfileDOSNetwork:                      DOSProfileDOSNetworkResource{c: c},
		dOSProfileProtocolDNS:                     DOSProfileProtocolDNSResource{c: c},
		dOSProfileProtocolSIP:                     DOSProfileProtocolSIPResource{c: c},
		dOSUDPPortlist:                            DOSUDPPortlistResource{c: c},
		datasync:                                  DatasyncResource{c: c},
		datasyncDeviceStats:                       DatasyncDeviceStatsResource{c: c},
		datasyncGlobalProfile:                     DatasyncGlobalProfileResource{c: c},
		datasyncLocalProfile:                      DatasyncLocalProfileResource{c: c},
		datasyncUpdateFile:                        DatasyncUpdateFileResource{c: c},
		device:                                    DeviceResource{c: c},
		deviceDeviceContext:                       DeviceDeviceContextResource{c: c},
		deviceDeviceContextNATRule:                DeviceDeviceContextNATRuleResource{c: c},
		engineUpdate:                              EngineUpdateResource{c: c},
		firewall:                                  FirewallResource{c: c},
		firewallAddressList:                       FirewallAddressListResource{c: c},
		firewallConfigChangeLog:                   FirewallConfigChangeLogResource{c: c},
		firewallFQDNEntity:                        FirewallFQDNEntityResource{c: c},
		firewallFQDNInfo:                          FirewallFQDNInfoResource{c: c},
		firewallGlobalFQDNPolicy:                  FirewallGlobalFQDNPolicyResource{c: c},
		firewallGlobalRules:                       FirewallGlobalRulesResource{c: c},
		firewallGlobalRulesActive:                 FirewallGlobalRulesActiveResource{c: c},
		firewallGlobalRulesEnforcedPolicyRules:    FirewallGlobalRulesEnforcedPolicyRulesResource{c: c},
		firewallGlobalRulesStagedPolicyRules:      FirewallGlobalRulesStagedPolicyRulesResource{c: c},
		firewallManagementIPRules:                 FirewallManagementIPRulesResource{c: c},
		firewallMatchingRule:                      FirewallMatchingRuleResource{c: c},
		firewallOnDemandCompilation:               FirewallOnDemandCompilationResource{c: c},
		firewallOnDemandRuleDeploy:                FirewallOnDemandRuleDeployResource{c: c},
		firewallPolicy:                            FirewallPolicyResource{c: c},
		firewallPolicyRules:                       FirewallPolicyRulesResource{c: c},
		firewallPortList:                          FirewallPortListResource{c: c},
		firewallPortMisusePolicy:                  FirewallPortMisusePolicyResource{c: c},
		firewallRuleList:                          FirewallRuleListResource{c: c},
		firewallRuleListRules:                     FirewallRuleListRulesResource{c: c},
		firewallSchedule:                          FirewallScheduleResource{c: c},
		firewallUserDomain:                        FirewallUserDomainResource{c: c},
		firewallUserGroupEntity:                   FirewallUserGroupEntityResource{c: c},
		firewallUserList:                          FirewallUserListResource{c: c},
		hTTP:                                      HTTPResource{c: c},
		hTTPFileType:                              HTTPFileTypeResource{c: c},
		hTTPMandatoryHeader:                       HTTPMandatoryHeaderResource{c: c},
		hTTPProfile:                               HTTPProfileResource{c: c},
		iPIntelligence:                            IPIntelligenceResource{c: c},
		iPIntelligenceFeedList:                    IPIntelligenceFeedListResource{c: c},
		iPIntelligenceGlobalPolicy:                IPIntelligenceGlobalPolicyResource{c: c},
		iPIntelligenceInfo:                        IPIntelligenceInfoResource{c: c},
		iPIntelligencePolicy:                      IPIntelligencePolicyResource{c: c},
		iPIntelligenseBlacklistCategory:           IPIntelligenseBlacklistCategoryResource{c: c},
		iPIntelligenseCategory:                    IPIntelligenseCategoryResource{c: c},
		log: LogResource{c: c},
		logNetworkStorageField:     LogNetworkStorageFieldResource{c: c},
		logProfile:                 LogProfileResource{c: c},
		logProfileApplication:      LogProfileApplicationResource{c: c},
		logProfileNetwork:          LogProfileNetworkResource{c: c},
		logProfileProtocolDNS:      LogProfileProtocolDNSResource{c: c},
		logProfileProtocolSIP:      LogProfileProtocolSIPResource{c: c},
		logProtocolDNSStorageField: LogProtocolDNSStorageFieldResource{c: c},
		logProtocolSIPStorageField: LogProtocolSIPStorageFieldResource{c: c},
		logRemoteFormat:            LogRemoteFormatResource{c: c},
		logStorageField:            LogStorageFieldResource{c: c},
		managementIPRulesRules:     ManagementIPRulesRulesResource{c: c},
		nAT: NATResource{c: c},
		nATDestinationTranslation: NATDestinationTranslationResource{c: c},
		nATPolicy:                 NATPolicyResource{c: c},
		nATPolicyRules:            NATPolicyRulesResource{c: c},
		nATSourceTranslation:      NATSourceTranslationResource{c: c},
		sSH:                       SSHResource{c: c},
		sSHProfile:                SSHProfileResource{c: c},
		sSHProfileAuthInfo:        SSHProfileAuthInfoResource{c: c},
		sSHProfileRules:           SSHProfileRulesResource{c: c},
		sSHSSHPluginStats:         SSHSSHPluginStatsResource{c: c},
		security:                  SecurityResource{c: c},
		settings:                  SettingsResource{c: c},
	}
}

// Analytics returns a configured AnalyticsResource.
func (sec Security) Analytics() *AnalyticsResource {
	return &sec.analytics
}

// AntiFraud returns a configured AntiFraudResource.
func (sec Security) AntiFraud() *AntiFraudResource {
	return &sec.antiFraud
}

// AntiFraudProfile returns a configured AntiFraudProfileResource.
func (sec Security) AntiFraudProfile() *AntiFraudProfileResource {
	return &sec.antiFraudProfile
}

// AntiFraudProfileUsers returns a configured AntiFraudProfileUsersResource.
func (sec Security) AntiFraudProfileUsers() *AntiFraudProfileUsersResource {
	return &sec.antiFraudProfileUsers
}

// AntiFraudProfileUsersModes returns a configured AntiFraudProfileUsersModesResource.
func (sec Security) AntiFraudProfileUsersModes() *AntiFraudProfileUsersModesResource {
	return &sec.antiFraudProfileUsersModes
}

// AntiFraudSignaturesUpdate returns a configured AntiFraudSignaturesUpdateResource.
func (sec Security) AntiFraudSignaturesUpdate() *AntiFraudSignaturesUpdateResource {
	return &sec.antiFraudSignaturesUpdate
}

// BlacklistPublisher returns a configured BlacklistPublisherResource.
func (sec Security) BlacklistPublisher() *BlacklistPublisherResource {
	return &sec.blacklistPublisher
}

// BlacklistPublisherBlacklistPublisherStats returns a configured BlacklistPublisherBlacklistPublisherStatsResource.
func (sec Security) BlacklistPublisherBlacklistPublisherStats() *BlacklistPublisherBlacklistPublisherStatsResource {
	return &sec.blacklistPublisherBlacklistPublisherStats
}

// BlacklistPublisherCategory returns a configured BlacklistPublisherCategoryResource.
func (sec Security) BlacklistPublisherCategory() *BlacklistPublisherCategoryResource {
	return &sec.blacklistPublisherCategory
}

// BlacklistPublisherProfile returns a configured BlacklistPublisherProfileResource.
func (sec Security) BlacklistPublisherProfile() *BlacklistPublisherProfileResource {
	return &sec.blacklistPublisherProfile
}

// BotDefense returns a configured BotDefenseResource.
func (sec Security) BotDefense() *BotDefenseResource {
	return &sec.botDefense
}

// BotDefenseASMProfile returns a configured BotDefenseASMProfileResource.
func (sec Security) BotDefenseASMProfile() *BotDefenseASMProfileResource {
	return &sec.botDefenseASMProfile
}

// DNS returns a configured DNSResource.
func (sec Security) DNS() *DNSResource {
	return &sec.dNS
}

// DNSProfile returns a configured DNSProfileResource.
func (sec Security) DNSProfile() *DNSProfileResource {
	return &sec.dNSProfile
}

// DOS returns a configured DOSResource.
func (sec Security) DOS() *DOSResource {
	return &sec.dOS
}

// DOSBotSignature returns a configured DOSBotSignatureResource.
func (sec Security) DOSBotSignature() *DOSBotSignatureResource {
	return &sec.dOSBotSignature
}

// DOSBotSignatureCategory returns a configured DOSBotSignatureCategoryResource.
func (sec Security) DOSBotSignatureCategory() *DOSBotSignatureCategoryResource {
	return &sec.dOSBotSignatureCategory
}

// DOSDeviceConfig returns a configured DOSDeviceConfigResource.
func (sec Security) DOSDeviceConfig() *DOSDeviceConfigResource {
	return &sec.dOSDeviceConfig
}

// DOSNetworkWhitelist returns a configured DOSNetworkWhitelistResource.
func (sec Security) DOSNetworkWhitelist() *DOSNetworkWhitelistResource {
	return &sec.dOSNetworkWhitelist
}

// DOSProfile returns a configured DOSProfileResource.
func (sec Security) DOSProfile() *DOSProfileResource {
	return &sec.dOSProfile
}

// DOSProfileApplication returns a configured DOSProfileApplicationResource.
func (sec Security) DOSProfileApplication() *DOSProfileApplicationResource {
	return &sec.dOSProfileApplication
}

// DOSProfileDOSNetwork returns a configured DOSProfileDOSNetworkResource.
func (sec Security) DOSProfileDOSNetwork() *DOSProfileDOSNetworkResource {
	return &sec.dOSProfileDOSNetwork
}

// DOSProfileProtocolDNS returns a configured DOSProfileProtocolDNSResource.
func (sec Security) DOSProfileProtocolDNS() *DOSProfileProtocolDNSResource {
	return &sec.dOSProfileProtocolDNS
}

// DOSProfileProtocolSIP returns a configured DOSProfileProtocolSIPResource.
func (sec Security) DOSProfileProtocolSIP() *DOSProfileProtocolSIPResource {
	return &sec.dOSProfileProtocolSIP
}

// DOSUDPPortlist returns a configured DOSUDPPortlistResource.
func (sec Security) DOSUDPPortlist() *DOSUDPPortlistResource {
	return &sec.dOSUDPPortlist
}

// Datasync returns a configured DatasyncResource.
func (sec Security) Datasync() *DatasyncResource {
	return &sec.datasync
}

// DatasyncDeviceStats returns a configured DatasyncDeviceStatsResource.
func (sec Security) DatasyncDeviceStats() *DatasyncDeviceStatsResource {
	return &sec.datasyncDeviceStats
}

// DatasyncGlobalProfile returns a configured DatasyncGlobalProfileResource.
func (sec Security) DatasyncGlobalProfile() *DatasyncGlobalProfileResource {
	return &sec.datasyncGlobalProfile
}

// DatasyncLocalProfile returns a configured DatasyncLocalProfileResource.
func (sec Security) DatasyncLocalProfile() *DatasyncLocalProfileResource {
	return &sec.datasyncLocalProfile
}

// DatasyncUpdateFile returns a configured DatasyncUpdateFileResource.
func (sec Security) DatasyncUpdateFile() *DatasyncUpdateFileResource {
	return &sec.datasyncUpdateFile
}

// Device returns a configured DeviceResource.
func (sec Security) Device() *DeviceResource {
	return &sec.device
}

// DeviceDeviceContext returns a configured DeviceDeviceContextResource.
func (sec Security) DeviceDeviceContext() *DeviceDeviceContextResource {
	return &sec.deviceDeviceContext
}

// DeviceDeviceContextNATRule returns a configured DeviceDeviceContextNATRuleResource.
func (sec Security) DeviceDeviceContextNATRule() *DeviceDeviceContextNATRuleResource {
	return &sec.deviceDeviceContextNATRule
}

// EngineUpdate returns a configured EngineUpdateResource.
func (sec Security) EngineUpdate() *EngineUpdateResource {
	return &sec.engineUpdate
}

// Firewall returns a configured FirewallResource.
func (sec Security) Firewall() *FirewallResource {
	return &sec.firewall
}

// FirewallAddressList returns a configured FirewallAddressListResource.
func (sec Security) FirewallAddressList() *FirewallAddressListResource {
	return &sec.firewallAddressList
}

// FirewallConfigChangeLog returns a configured FirewallConfigChangeLogResource.
func (sec Security) FirewallConfigChangeLog() *FirewallConfigChangeLogResource {
	return &sec.firewallConfigChangeLog
}

// FirewallFQDNEntity returns a configured FirewallFQDNEntityResource.
func (sec Security) FirewallFQDNEntity() *FirewallFQDNEntityResource {
	return &sec.firewallFQDNEntity
}

// FirewallFQDNInfo returns a configured FirewallFQDNInfoResource.
func (sec Security) FirewallFQDNInfo() *FirewallFQDNInfoResource {
	return &sec.firewallFQDNInfo
}

// FirewallGlobalFQDNPolicy returns a configured FirewallGlobalFQDNPolicyResource.
func (sec Security) FirewallGlobalFQDNPolicy() *FirewallGlobalFQDNPolicyResource {
	return &sec.firewallGlobalFQDNPolicy
}

// FirewallGlobalRules returns a configured FirewallGlobalRulesResource.
func (sec Security) FirewallGlobalRules() *FirewallGlobalRulesResource {
	return &sec.firewallGlobalRules
}

// FirewallGlobalRulesActive returns a configured FirewallGlobalRulesActiveResource.
func (sec Security) FirewallGlobalRulesActive() *FirewallGlobalRulesActiveResource {
	return &sec.firewallGlobalRulesActive
}

// FirewallGlobalRulesEnforcedPolicyRules returns a configured FirewallGlobalRulesEnforcedPolicyRulesResource.
func (sec Security) FirewallGlobalRulesEnforcedPolicyRules() *FirewallGlobalRulesEnforcedPolicyRulesResource {
	return &sec.firewallGlobalRulesEnforcedPolicyRules
}

// FirewallGlobalRulesStagedPolicyRules returns a configured FirewallGlobalRulesStagedPolicyRulesResource.
func (sec Security) FirewallGlobalRulesStagedPolicyRules() *FirewallGlobalRulesStagedPolicyRulesResource {
	return &sec.firewallGlobalRulesStagedPolicyRules
}

// FirewallManagementIPRules returns a configured FirewallManagementIPRulesResource.
func (sec Security) FirewallManagementIPRules() *FirewallManagementIPRulesResource {
	return &sec.firewallManagementIPRules
}

// FirewallMatchingRule returns a configured FirewallMatchingRuleResource.
func (sec Security) FirewallMatchingRule() *FirewallMatchingRuleResource {
	return &sec.firewallMatchingRule
}

// FirewallOnDemandCompilation returns a configured FirewallOnDemandCompilationResource.
func (sec Security) FirewallOnDemandCompilation() *FirewallOnDemandCompilationResource {
	return &sec.firewallOnDemandCompilation
}

// FirewallOnDemandRuleDeploy returns a configured FirewallOnDemandRuleDeployResource.
func (sec Security) FirewallOnDemandRuleDeploy() *FirewallOnDemandRuleDeployResource {
	return &sec.firewallOnDemandRuleDeploy
}

// FirewallPolicy returns a configured FirewallPolicyResource.
func (sec Security) FirewallPolicy() *FirewallPolicyResource {
	return &sec.firewallPolicy
}

// FirewallPolicyRules returns a configured FirewallPolicyRulesResource.
func (sec Security) FirewallPolicyRules() *FirewallPolicyRulesResource {
	return &sec.firewallPolicyRules
}

// FirewallPortList returns a configured FirewallPortListResource.
func (sec Security) FirewallPortList() *FirewallPortListResource {
	return &sec.firewallPortList
}

// FirewallPortMisusePolicy returns a configured FirewallPortMisusePolicyResource.
func (sec Security) FirewallPortMisusePolicy() *FirewallPortMisusePolicyResource {
	return &sec.firewallPortMisusePolicy
}

// FirewallRuleList returns a configured FirewallRuleListResource.
func (sec Security) FirewallRuleList() *FirewallRuleListResource {
	return &sec.firewallRuleList
}

// FirewallRuleListRules returns a configured FirewallRuleListRulesResource.
func (sec Security) FirewallRuleListRules() *FirewallRuleListRulesResource {
	return &sec.firewallRuleListRules
}

// FirewallSchedule returns a configured FirewallScheduleResource.
func (sec Security) FirewallSchedule() *FirewallScheduleResource {
	return &sec.firewallSchedule
}

// FirewallUserDomain returns a configured FirewallUserDomainResource.
func (sec Security) FirewallUserDomain() *FirewallUserDomainResource {
	return &sec.firewallUserDomain
}

// FirewallUserGroupEntity returns a configured FirewallUserGroupEntityResource.
func (sec Security) FirewallUserGroupEntity() *FirewallUserGroupEntityResource {
	return &sec.firewallUserGroupEntity
}

// FirewallUserList returns a configured FirewallUserListResource.
func (sec Security) FirewallUserList() *FirewallUserListResource {
	return &sec.firewallUserList
}

// HTTP returns a configured HTTPResource.
func (sec Security) HTTP() *HTTPResource {
	return &sec.hTTP
}

// HTTPFileType returns a configured HTTPFileTypeResource.
func (sec Security) HTTPFileType() *HTTPFileTypeResource {
	return &sec.hTTPFileType
}

// HTTPMandatoryHeader returns a configured HTTPMandatoryHeaderResource.
func (sec Security) HTTPMandatoryHeader() *HTTPMandatoryHeaderResource {
	return &sec.hTTPMandatoryHeader
}

// HTTPProfile returns a configured HTTPProfileResource.
func (sec Security) HTTPProfile() *HTTPProfileResource {
	return &sec.hTTPProfile
}

// IPIntelligence returns a configured IPIntelligenceResource.
func (sec Security) IPIntelligence() *IPIntelligenceResource {
	return &sec.iPIntelligence
}

// IPIntelligenceFeedList returns a configured IPIntelligenceFeedListResource.
func (sec Security) IPIntelligenceFeedList() *IPIntelligenceFeedListResource {
	return &sec.iPIntelligenceFeedList
}

// IPIntelligenceGlobalPolicy returns a configured IPIntelligenceGlobalPolicyResource.
func (sec Security) IPIntelligenceGlobalPolicy() *IPIntelligenceGlobalPolicyResource {
	return &sec.iPIntelligenceGlobalPolicy
}

// IPIntelligenceInfo returns a configured IPIntelligenceInfoResource.
func (sec Security) IPIntelligenceInfo() *IPIntelligenceInfoResource {
	return &sec.iPIntelligenceInfo
}

// IPIntelligencePolicy returns a configured IPIntelligencePolicyResource.
func (sec Security) IPIntelligencePolicy() *IPIntelligencePolicyResource {
	return &sec.iPIntelligencePolicy
}

// IPIntelligenseBlacklistCategory returns a configured IPIntelligenseBlacklistCategoryResource.
func (sec Security) IPIntelligenseBlacklistCategory() *IPIntelligenseBlacklistCategoryResource {
	return &sec.iPIntelligenseBlacklistCategory
}

// IPIntelligenseCategory returns a configured IPIntelligenseCategoryResource.
func (sec Security) IPIntelligenseCategory() *IPIntelligenseCategoryResource {
	return &sec.iPIntelligenseCategory
}

// Log returns a configured LogResource.
func (sec Security) Log() *LogResource {
	return &sec.log
}

// LogNetworkStorageField returns a configured LogNetworkStorageFieldResource.
func (sec Security) LogNetworkStorageField() *LogNetworkStorageFieldResource {
	return &sec.logNetworkStorageField
}

// LogProfile returns a configured LogProfileResource.
func (sec Security) LogProfile() *LogProfileResource {
	return &sec.logProfile
}

// LogProfileApplication returns a configured LogProfileApplicationResource.
func (sec Security) LogProfileApplication() *LogProfileApplicationResource {
	return &sec.logProfileApplication
}

// LogProfileNetwork returns a configured LogProfileNetworkResource.
func (sec Security) LogProfileNetwork() *LogProfileNetworkResource {
	return &sec.logProfileNetwork
}

// LogProfileProtocolDNS returns a configured LogProfileProtocolDNSResource.
func (sec Security) LogProfileProtocolDNS() *LogProfileProtocolDNSResource {
	return &sec.logProfileProtocolDNS
}

// LogProfileProtocolSIP returns a configured LogProfileProtocolSIPResource.
func (sec Security) LogProfileProtocolSIP() *LogProfileProtocolSIPResource {
	return &sec.logProfileProtocolSIP
}

// LogProtocolDNSStorageField returns a configured LogProtocolDNSStorageFieldResource.
func (sec Security) LogProtocolDNSStorageField() *LogProtocolDNSStorageFieldResource {
	return &sec.logProtocolDNSStorageField
}

// LogProtocolSIPStorageField returns a configured LogProtocolSIPStorageFieldResource.
func (sec Security) LogProtocolSIPStorageField() *LogProtocolSIPStorageFieldResource {
	return &sec.logProtocolSIPStorageField
}

// LogRemoteFormat returns a configured LogRemoteFormatResource.
func (sec Security) LogRemoteFormat() *LogRemoteFormatResource {
	return &sec.logRemoteFormat
}

// LogStorageField returns a configured LogStorageFieldResource.
func (sec Security) LogStorageField() *LogStorageFieldResource {
	return &sec.logStorageField
}

// ManagementIPRulesRules returns a configured ManagementIPRulesRulesResource.
func (sec Security) ManagementIPRulesRules() *ManagementIPRulesRulesResource {
	return &sec.managementIPRulesRules
}

// NAT returns a configured NATResource.
func (sec Security) NAT() *NATResource {
	return &sec.nAT
}

// NATDestinationTranslation returns a configured NATDestinationTranslationResource.
func (sec Security) NATDestinationTranslation() *NATDestinationTranslationResource {
	return &sec.nATDestinationTranslation
}

// NATPolicy returns a configured NATPolicyResource.
func (sec Security) NATPolicy() *NATPolicyResource {
	return &sec.nATPolicy
}

// NATPolicyRules returns a configured NATPolicyRulesResource.
func (sec Security) NATPolicyRules() *NATPolicyRulesResource {
	return &sec.nATPolicyRules
}

// NATSourceTranslation returns a configured NATSourceTranslationResource.
func (sec Security) NATSourceTranslation() *NATSourceTranslationResource {
	return &sec.nATSourceTranslation
}

// SSH returns a configured SSHResource.
func (sec Security) SSH() *SSHResource {
	return &sec.sSH
}

// SSHProfile returns a configured SSHProfileResource.
func (sec Security) SSHProfile() *SSHProfileResource {
	return &sec.sSHProfile
}

// SSHProfileAuthInfo returns a configured SSHProfileAuthInfoResource.
func (sec Security) SSHProfileAuthInfo() *SSHProfileAuthInfoResource {
	return &sec.sSHProfileAuthInfo
}

// SSHProfileRules returns a configured SSHProfileRulesResource.
func (sec Security) SSHProfileRules() *SSHProfileRulesResource {
	return &sec.sSHProfileRules
}

// SSHSSHPluginStats returns a configured SSHSSHPluginStatsResource.
func (sec Security) SSHSSHPluginStats() *SSHSSHPluginStatsResource {
	return &sec.sSHSSHPluginStats
}

// Security returns a configured SecurityResource.
func (sec Security) Security() *SecurityResource {
	return &sec.security
}

// Settings returns a configured SettingsResource.
func (sec Security) Settings() *SettingsResource {
	return &sec.settings
}
