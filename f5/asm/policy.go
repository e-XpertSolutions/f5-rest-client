// Copyright 2016 e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asm

import (
	//"math/big"
	"time"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
)

// A PolicyList holds a list of Policy.
type PolicyList struct {
	Items      []Policy `json:"items,omitempty"`
	Kind       string   `json:"kind,omitempty"`
	TotalItems int      `json:"totalItems,omitempty"`
	SelfLink   string   `json:"selfLink,omitempty"`
}

// A Policy hold the configuration for a policy.
type Policy struct {
	Name              string    `json:"name"`
	CreatedDatetime   time.Time `json:"createdDatetime"`
	EnforcementMode   string    `json:"enforcementMode"`
	VersionDeviceName string    `json:"versionDeviceName"`
	Active            bool      `json:"active"`
	CreatorName       string    `json:"creatorName"`

	Partition           string    `json:"partition"`
	VirtualServers      []string  `json:"virtualServers"`
	ApplicationLanguage string    `json:"applicationLanguage"`
	ID                  string    `json:"id"`
	ModifierName        string    `json:"modifierName"`
	VersionDatetime     time.Time `json:"versionDatetime"`

	PlainTextProfileReference                   PolicyReference `json:"plainTextProfileReference,omitempty"`
	DataGuardReference                          PolicyReference `json:"dataGuardReference,omitempty"`
	DatabaseProtectionReference                 PolicyReference `json:"databaseProtectionReference,omitempty"`
	CookieSettingsReference                     PolicyReference `json:"cookieSettingsReference,omitempty"`
	CSRFURLReference                            PolicyReference `json:"csrfUrlReference,omitempty"`
	VersionLastChange                           string          `json:"versionLastChange,omitempty"`
	CaseInsensitive                             bool            `json:"caseInsensitive,omitempty"`
	HeaderSettingsReference                     PolicyReference `json:"headerSettingsReference,omitempty"`
	SectionReference                            PolicyReference `json:"sectionReference,omitempty"`
	FlowReference                               PolicyReference `json:"flowReference,omitempty"`
	LoginPageReference                          PolicyReference `json:"loginPageReference,omitempty"`
	Description                                 string          `json:"description,omitempty"`
	FullPath                                    string          `json:"fullPath,omitempty"`
	PolicyBuilderParameterReference             PolicyReference `json:"policyBuilderParameterReference,omitempty"`
	HasParent                                   bool            `json:"hasParent,omitempty"`
	ThreatCampaignReference                     PolicyReference `json:"threatCampaignReference,omitempty"`
	CSRFProtectionReference                     PolicyReference `json:"csrfProtectionReference,omitempty"`
	PolicyAntivirusReference                    PolicyReference `json:"policyAntivirusReference,omitempty"`
	Kind                                        string          `json:"kind,omitempty"`
	PolicyBuilderCookieReference                PolicyReference `json:"policyBuilderCookieReference,omitempty"`
	IPIntelligenceReference                     PolicyReference `json:"ipIntelligenceReference,omitempty"`
	ProtocolIndependent                         bool            `json:"protocolIndependent,omitempty"`
	SessionAwarenessSettingsReference           PolicyReference `json:"sessionAwarenessSettingsReference,omitempty"`
	PolicyBuilderURLReference                   PolicyReference `json:"policyBuilderUrlReference,omitempty"`
	PolicyBuilderServerTechnologiesReference    PolicyReference `json:"policyBuilderServerTechnologiesReference,omitempty"`
	PolicyBuilderFiletypeReference              PolicyReference `json:"policyBuilderFiletypeReference,omitempty"`
	SignatureSetReference                       PolicyReference `json:"signatureSetReference,omitempty"`
	ParameterReference                          PolicyReference `json:"parameterReference,omitempty"`
	LoginEnforcementReference                   PolicyReference `json:"loginEnforcementReference,omitempty"`
	NavigationParameterReference                PolicyReference `json:"navigationParameterReference,omitempty"`
	GWTProfileReference                         PolicyReference `json:"gwtProfileReference,omitempty"`
	WhitelistIPReference                        PolicyReference `json:"whitelistIpReference,omitempty"`
	HistoryRevisionReference                    PolicyReference `json:"historyRevisionReference,omitempty"`
	PolicyBuilderReference                      PolicyReference `json:"policyBuilderReference,omitempty"`
	ResponsePageReference                       PolicyReference `json:"responsePageReference,omitempty"`
	VulnerabilityAssessmentReference            PolicyReference `json:"vulnerabilityAssessmentReference,omitempty"`
	ServerTechnologyReference                   PolicyReference `json:"serverTechnologyReference,omitempty"`
	CookieReference                             PolicyReference `json:"cookieReference,omitempty"`
	BlockingSettingReference                    PolicyReference `json:"blockingSettingReference,omitempty"`
	HostNameReference                           PolicyReference `json:"hostNameReference,omitempty"`
	SelfLink                                    string          `json:"selfLink,omitempty"`
	ThreatCampaignSettingReference              PolicyReference `json:"threatCampaignSettingReference,omitempty"`
	SignatureReference                          PolicyReference `json:"signatureReference,omitempty"`
	PolicyBuilderRedirectionProtectionReference PolicyReference `json:"policyBuilderRedirectionProtectionReferen,omitemptyce"`
	FiletypeReference                           PolicyReference `json:"filetypeReference,omitempty"`
	ManualVirtualServers                        []interface{}   `json:"manualVirtualServers,omitempty"`
	SubPath                                     string          `json:"subPath,omitempty"`
	SessionTrackingStatusReference              PolicyReference `json:"sessionTrackingStatusReference,omitempty"`
	AuditLogReference                           PolicyReference `json:"auditLogReference,omitempty"`
	DisallowedGeolocationReference              PolicyReference `json:"disallowedGeolocationReference,omitempty"`
	RedirectionProtectionDomainReference        PolicyReference `json:"redirectionProtectionDomainReference,omitempty"`
	Type                                        string          `json:"type,omitempty"`
	SignatureSettingReference                   PolicyReference `json:"signatureSettingReference,omitempty"`
	WebsocketURLReference                       PolicyReference `json:"websocketUrlReference,omitempty"`
	XMLProfileReference                         PolicyReference `json:"xmlProfileReference,omitempty"`
	MethodReference                             PolicyReference `json:"methodReference,omitempty"`
	VulnerabilityReference                      PolicyReference `json:"vulnerabilityReference,omitempty"`
	RedirectionProtectionReference              PolicyReference `json:"redirectionProtectionReference,omitempty"`
	PolicyBuilderSessionsAndLoginsReference     PolicyReference `json:"policyBuilderSessionsAndLoginsReference,omitempty"`
	PolicyReference                             PolicyReference `json:"PolicyReference,omitempty"`
	PolicyBuilderHeaderReference                PolicyReference `json:"policyBuilderHeaderReference,omitempty"`
	URLReference                                PolicyReference `json:"urlReference,omitempty"`
	HeaderReference                             PolicyReference `json:"headerReference,omitempty"`
	ActionItemReference                         PolicyReference `json:"actionItemReference,omitempty"`
	MicroserviceReference                       PolicyReference `json:"microserviceReference,omitempty"`
	XMLValidationFileReference                  PolicyReference `json:"xmlValidationFileReference,omitempty"`
	LastUpdateMicros                            float64         `json:"lastUpdateMicros,omitempty"`
	JSONProfileReference                        PolicyReference `json:"jsonProfileReference,omitempty"`
	BruteForceAttackPreventionReference         PolicyReference `json:"bruteForceAttackPreventionReference,omitempty"`
	DisabledActionItemReference                 PolicyReference `json:"disabledActionItemReference,omitempty"`
	ExtractionReference                         PolicyReference `json:"extractionReference,omitempty"`
	CharacterSetReference                       PolicyReference `json:"characterSetReference,omitempty"`
	SuggestionReference                         PolicyReference `json:"suggestionReference,omitempty"`
	IsModified                                  bool            `json:"isModified,omitempty"`
	SensitiveParameterReference                 PolicyReference `json:"sensitiveParameterReference,omitempty"`
	GeneralReference                            PolicyReference `json:"generalReference,omitempty"`
	VersionPolicyName                           string          `json:"versionPolicyName,omitempty"`
	PolicyBuilderCentralConfigurationReference  PolicyReference `json:"policyBuilderCentralConfigurationReference,omitempty"`
}

type PolicyReference struct {
	Link            string `json:"link,omitempty"`
	IsSubCollection bool   `json:"isSubCollection,omitempty"`
	Title           string `json:"title,omitempty"`
}

// PolicyEndpoint represents the REST resource for managing a policy.
const PolicyEndpoint = "/policies"

// A PolicyResource provides API to manage Policys configuration.
type PolicyResource struct {
	c *f5.Client
}

// ListAll lists all the Policy configurations.
func (pr *PolicyResource) ListAll() (*PolicyList, error) {
	var list PolicyList
	if err := pr.c.ReadQuery(BasePath+PolicyEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single Policy configuration identified by id.
func (pr *PolicyResource) Get(id string) (*Policy, error) {
	var item Policy
	if err := pr.c.ReadQuery(BasePath+PolicyEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

/* // Create a new Policy configuration.
func (pr *PolicyResource) Create(item Policy) error {
	if err := pr.c.ModQuery("POST", BasePath+PolicyEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a Policy configuration identified by id.
func (pr *PolicyResource) Edit(id string, item Policy) error {
	if err := pr.c.ModQuery("PUT", BasePath+PolicyEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single Policy configuration identified by id.
func (pr *PolicyResource) Delete(id string) error {
	if err := pr.c.ModQuery("DELETE", BasePath+PolicyEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
} */
