// Copyright e-Xpert Solutions SA. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import "github.com/e-XpertSolutions/f5-rest-client/f5"

// AntiFraudProfileConfigList holds a list of AntiFraudProfile configuration.
type AntiFraudProfileConfigList struct {
	Items    []AntiFraudProfileConfig `json:"items"`
	Kind     string                   `json:"kind"`
	SelfLink string                   `json:"selflink"`
}

// AntiFraudProfileConfig holds the configuration of a single AntiFraudProfile.
type AntiFraudProfileConfig struct {
	AlertPath        string `json:"alertPath"`
	AutoTransactions struct {
		BotScore            int `json:"botScore"`
		ClickScore          int `json:"clickScore"`
		IntegrityFailScore  int `json:"integrityFailScore"`
		MinMouseMoveCount   int `json:"minMouseMoveCount"`
		MinMouseOverCount   int `json:"minMouseOverCount"`
		MinReportScore      int `json:"minReportScore"`
		MinTimeToRequest    int `json:"minTimeToRequest"`
		NotHumanScore       int `json:"notHumanScore"`
		TamperedCookieScore int `json:"tamperedCookieScore"`
		TimeFailScore       int `json:"timeFailScore"`
	} `json:"autoTransactions"`
	BlockingPage struct {
		ResponseBody    string `json:"responseBody"`
		ResponseHeaders string `json:"responseHeaders"`
	} `json:"blockingPage"`
	CaseInsensitive bool   `json:"caseInsensitive"`
	CheckPathInfo   string `json:"checkPathInfo"`
	Cookies         struct {
		ComponentsState              string `json:"componentsState"`
		ComponentsStateLifetime      string `json:"componentsStateLifetime"`
		EncryptionDisabled           string `json:"encryptionDisabled"`
		EncryptionDisabledLifetime   string `json:"encryptionDisabledLifetime"`
		HTMLFieldObfuscation         string `json:"htmlFieldObfuscation"`
		HTMLFieldObfuscationLifetime string `json:"htmlFieldObfuscationLifetime"`
		MalwareCache                 string `json:"malwareCache"`
		MalwareCacheLifetime         string `json:"malwareCacheLifetime"`
		MalwareCounter               string `json:"malwareCounter"`
		MalwareCounterLifetime       string `json:"malwareCounterLifetime"`
		MalwareForensic              string `json:"malwareForensic"`
		MalwareForensicLifetime      string `json:"malwareForensicLifetime"`
		MalwareGUID                  string `json:"malwareGuid"`
		MalwareGUIDLifetime          string `json:"malwareGuidLifetime"`
		PhishingCache                string `json:"phishingCache"`
		PhishingCacheLifetime        string `json:"phishingCacheLifetime"`
		SecureChannel                string `json:"secureChannel"`
		SecureChannelLifetime        string `json:"secureChannelLifetime"`
		TransactionData              string `json:"transactionData"`
		TransactionDataLifetime      string `json:"transactionDataLifetime"`
		UserInspection               string `json:"userInspection"`
		UserName                     string `json:"userName"`
		UserNameLifetime             string `json:"userNameLifetime"`
	} `json:"cookies"`
	Forensic struct {
		AlertPath            string `json:"alertPath"`
		CloudConfigPath      string `json:"cloudConfigPath"`
		CloudForensicsMode   int    `json:"cloudForensicsMode"`
		CloudRemediationMode int    `json:"cloudRemediationMode"`
		ContinueElement      string `json:"continueElement"`
		ExeLocation          string `json:"exeLocation"`
		HTML                 string `json:"html"`
		SelfPostLocation     string `json:"selfPostLocation"`
		SkipElement          string `json:"skipElement"`
		SkipPath             string `json:"skipPath"`
	} `json:"forensic"`
	FullPath           string `json:"fullPath"`
	Generation         int    `json:"generation"`
	JavascriptLocation string `json:"javascriptLocation"`
	Kind               string `json:"kind"`
	Malware            struct {
		BaitLocation            string `json:"baitLocation"`
		FlashCookieLocation     string `json:"flashCookieLocation"`
		FlashCookies            string `json:"flashCookies"`
		SourceIntegrityLocation string `json:"sourceIntegrityLocation"`
	} `json:"malware"`
	Mobilesafe struct {
		AlertThreshold int `json:"alertThreshold"`
		AppIntegrity   struct {
			Android struct {
				Score int `json:"score"`
			} `json:"android"`
			Enabled bool `json:"enabled"`
			Ios     struct {
				Score int `json:"score"`
			} `json:"ios"`
		} `json:"appIntegrity"`
		Malware struct {
			Android           struct{} `json:"android"`
			BehaviourAnalysis struct {
				Run   string `json:"run"`
				Score int    `json:"score"`
			} `json:"behaviourAnalysis"`
			CheckCustom  string   `json:"checkCustom"`
			CheckGeneric string   `json:"checkGeneric"`
			Enabled      bool     `json:"enabled"`
			Ios          struct{} `json:"ios"`
		} `json:"malware"`
		Mitm struct {
			Enabled bool `json:"enabled"`
		} `json:"mitm"`
		OsSecurity struct {
			Android struct {
				UntrustedAppsScore int `json:"untrustedAppsScore"`
			} `json:"android"`
			Enabled bool     `json:"enabled"`
			Ios     struct{} `json:"ios"`
		} `json:"osSecurity"`
		RootingJailbreak struct {
			Enabled        bool `json:"enabled"`
			JailbreakScore int  `json:"jailbreakScore"`
			RootingScore   int  `json:"rootingScore"`
		} `json:"rootingJailbreak"`
	} `json:"mobilesafe"`
	Name      string `json:"name"`
	Partition string `json:"partition"`
	Phishing  struct {
		CSSAttributeName string `json:"cssAttributeName"`
		CSSContent       string `json:"cssContent"`
		CSSLocation      string `json:"cssLocation"`
		ExpirationChecks string `json:"expirationChecks"`
		ImageLocation    string `json:"imageLocation"`
		ReferrerChecks   string `json:"referrerChecks"`
	} `json:"phishing"`
	SelfLink       string `json:"selfLink"`
	TriggerIrule   string `json:"triggerIrule"`
	UsersReference struct {
		IsSubcollection bool   `json:"isSubcollection"`
		Link            string `json:"link"`
	} `json:"usersReference"`
}

// AntiFraudProfileEndpoint represents the REST resource for managing AntiFraudProfile.
const AntiFraudProfileEndpoint = "/anti-fraud/profile"

// AntiFraudProfileResource provides an API to manage AntiFraudProfile configurations.
type AntiFraudProfileResource struct {
	c *f5.Client
}

// ListAll  lists all the AntiFraudProfile configurations.
func (r *AntiFraudProfileResource) ListAll() (*AntiFraudProfileConfigList, error) {
	var list AntiFraudProfileConfigList
	if err := r.c.ReadQuery(BasePath+AntiFraudProfileEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

// Get a single AntiFraudProfile configuration identified by id.
func (r *AntiFraudProfileResource) Get(id string) (*AntiFraudProfileConfig, error) {
	var item AntiFraudProfileConfig
	if err := r.c.ReadQuery(BasePath+AntiFraudProfileEndpoint, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// Create a new AntiFraudProfile configuration.
func (r *AntiFraudProfileResource) Create(item AntiFraudProfileConfig) error {
	if err := r.c.ModQuery("POST", BasePath+AntiFraudProfileEndpoint, item); err != nil {
		return err
	}
	return nil
}

// Edit a AntiFraudProfile configuration identified by id.
func (r *AntiFraudProfileResource) Edit(id string, item AntiFraudProfileConfig) error {
	if err := r.c.ModQuery("PUT", BasePath+AntiFraudProfileEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

// Delete a single AntiFraudProfile configuration identified by id.
func (r *AntiFraudProfileResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+AntiFraudProfileEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
