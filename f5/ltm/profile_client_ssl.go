package ltm

import "github.com/e-XpertSolutions/f5-rest-client/f5"

type ProfileClientSSLList struct {
	Items    []ProfileClientSSL `json:"items,omitempty"`
	Kind     string             `json:"kind,omitempty" pretty:",expanded"`
	SelfLink string             `json:"selflink,omitempty" pretty:",expanded"`
}

type ProfileClientSSL struct {
	AlertTimeout             string   `json:"alertTimeout,omitempty"`
	AllowDynamicRecordSizing string   `json:"allowDynamicRecordSizing,omitempty"`
	AllowExpiredCRL          string   `json:"allowExpiredCrl,omitempty"`
	AllowNonSSL              string   `json:"allowNonSsl,omitempty"`
	AppService               string   `json:"appService,omitempty"`
	Authenticate             string   `json:"authenticate,omitempty"`
	AuthenticateDepth        int      `json:"authenticateDepth,omitempty"`
	BypassOnClientCertFail   string   `json:"bypassOnClientCertFail,omitempty"`
	BypassOnHandshakeAlert   string   `json:"bypassOnHandshakeAlert,omitempty"`
	CAFile                   string   `json:"caFile,omitempty"`
	CacheSize                int      `json:"cacheSize,omitempty"`
	CacheTimeout             int      `json:"cacheTimeout,omitempty"`
	Cert                     string   `json:"cert,omitempty"`
	CertExtensionIncludes    []string `json:"certExtensionIncludes,omitempty"`
	CertKeyChain             []struct {
		AppService    string `json:"appService,omitempty"`
		Cert          string `json:"cert,omitempty"`
		CertReference struct {
			Link string `json:"link,omitempty"`
		} `json:"certReference,omitempty"`
		Chain        string `json:"chain,omitempty"`
		Key          string `json:"key,omitempty"`
		KeyReference struct {
			Link string `json:"link,omitempty"`
		} `json:"keyReference,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"certKeyChain,omitempty"`
	CertLifespan           int    `json:"certLifespan,omitempty"`
	CertLookupByIpaddrPort string `json:"certLookupByIpaddrPort,omitempty"`
	CertReference          struct {
		Link string `json:"link,omitempty"`
	} `json:"certReference,omitempty"`
	Chain            string `json:"chain,omitempty"`
	CipherGroup      string `json:"cipherGroup,omitempty"`
	Ciphers          string `json:"ciphers,omitempty"`
	ClientCertCA     string `json:"clientCertCa,omitempty"`
	CRLFile          string `json:"crlFile,omitempty"`
	CRLFileReference struct {
		Link string `json:"link,omitempty"`
	} `json:"crlFileReference,omitempty"`
	DefaultsFrom                    string `json:"defaultsFrom,omitempty"`
	Description                     string `json:"description,omitempty"`
	DestinationIPBlacklist          string `json:"destinationIpBlacklist,omitempty"`
	DestinationIPWhitelist          string `json:"destinationIpWhitelist,omitempty"`
	ForwardProxyBypassDefaultAction string `json:"forwardProxyBypassDefaultAction,omitempty"`
	FullPath                        string `json:"fullPath,omitempty"`
	Generation                      int    `json:"generation,omitempty"`
	GenericAlert                    string `json:"genericAlert,omitempty"`
	HandshakeTimeout                string `json:"handshakeTimeout,omitempty"`
	HostnameBlacklist               string `json:"hostnameBlacklist,omitempty"`
	HostnameWhitelist               string `json:"hostnameWhitelist,omitempty"`
	InheritCertkeychain             string `json:"inheritCertkeychain,omitempty"`
	Key                             string `json:"key,omitempty"`
	KeyReference                    struct {
		Link string `json:"link,omitempty"`
	} `json:"keyReference,omitempty"`
	Kind                               string   `json:"kind,omitempty"`
	MaxActiveHandshakes                string   `json:"maxActiveHandshakes,omitempty"`
	MaxAggregateRenegotiationPerMinute string   `json:"maxAggregateRenegotiationPerMinute,omitempty"`
	MaxRenegotiationsPerMinute         int      `json:"maxRenegotiationsPerMinute,omitempty"`
	MaximumRecordSize                  int      `json:"maximumRecordSize,omitempty"`
	ModSslMethods                      string   `json:"modSslMethods,omitempty"`
	Mode                               string   `json:"mode,omitempty"`
	Name                               string   `json:"name,omitempty"`
	NotifyCertStatusToVirtualServer    string   `json:"notifyCertStatusToVirtualServer,omitempty"`
	OCSPStapling                       string   `json:"ocspStapling,omitempty"`
	Partition                          string   `json:"partition,omitempty"`
	PeerCertMode                       string   `json:"peerCertMode,omitempty"`
	PeerNoRenegotiateTimeout           string   `json:"peerNoRenegotiateTimeout,omitempty"`
	ProxyCACert                        string   `json:"proxyCaCert,omitempty"`
	ProxyCAKey                         string   `json:"proxyCaKey,omitempty"`
	ProxySSL                           string   `json:"proxySsl,omitempty"`
	ProxySSLPassthrough                string   `json:"proxySslPassthrough,omitempty"`
	RenegotiateMaxRecordDelay          string   `json:"renegotiateMaxRecordDelay,omitempty"`
	RenegotiatePeriod                  string   `json:"renegotiatePeriod,omitempty"`
	RenegotiateSize                    string   `json:"renegotiateSize,omitempty"`
	Renegotiation                      string   `json:"renegotiation,omitempty"`
	RetainCertificate                  string   `json:"retainCertificate,omitempty"`
	SecureRenegotiation                string   `json:"secureRenegotiation,omitempty"`
	SelfLink                           string   `json:"selfLink,omitempty"`
	ServerName                         string   `json:"serverName,omitempty"`
	SessionMirroring                   string   `json:"sessionMirroring,omitempty"`
	SessionTicket                      string   `json:"sessionTicket,omitempty"`
	SessionTicketTimeout               int      `json:"sessionTicketTimeout,omitempty"`
	SNIDefault                         string   `json:"sniDefault,omitempty"`
	SNIRequire                         string   `json:"sniRequire,omitempty"`
	SourceIPBlacklist                  string   `json:"sourceIpBlacklist,omitempty"`
	SourceIPWhitelist                  string   `json:"sourceIpWhitelist,omitempty"`
	SSLForwardProxy                    string   `json:"sslForwardProxy,omitempty"`
	SSLForwardProxyBypass              string   `json:"sslForwardProxyBypass,omitempty"`
	SSLSignHash                        string   `json:"sslSignHash,omitempty"`
	StrictResume                       string   `json:"strictResume,omitempty"`
	TMOptions                          []string `json:"tmOptions,omitempty"`
	UncleanShutdown                    string   `json:"uncleanShutdown,omitempty"`
}

const ProfileClientSSLEndpoint = "/profile/client-ssl"

type ProfileClientSSLResource struct {
	c *f5.Client
}

func (r *ProfileClientSSLResource) ListAll() (*ProfileClientSSLList, error) {
	var list ProfileClientSSLList
	if err := r.c.ReadQuery(BasePath+ProfileClientSSLEndpoint, &list); err != nil {
		return nil, err
	}
	return &list, nil
}

func (r *ProfileClientSSLResource) Get(id string) (*ProfileClientSSL, error) {
	var item ProfileClientSSL
	if err := r.c.ReadQuery(BasePath+ProfileClientSSLEndpoint+"/"+id, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ProfileClientSSLResource) Create(item ProfileClientSSL) error {
	if err := r.c.ModQuery("POST", BasePath+ProfileClientSSLEndpoint, item); err != nil {
		return err
	}
	return nil
}

func (r *ProfileClientSSLResource) Edit(id string, item ProfileClientSSL) error {
	if err := r.c.ModQuery("PUT", BasePath+ProfileClientSSLEndpoint+"/"+id, item); err != nil {
		return err
	}
	return nil
}

func (r *ProfileClientSSLResource) Delete(id string) error {
	if err := r.c.ModQuery("DELETE", BasePath+ProfileClientSSLEndpoint+"/"+id, nil); err != nil {
		return err
	}
	return nil
}
