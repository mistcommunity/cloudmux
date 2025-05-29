package cloudprovider

type SLoadbalancerListenerCreateOptions struct {
	Name                    string
	ListenerType            string
	ListenerPort            int
	BackendGroupType        string
	BackendGroupId          string
	Scheduler               string
	AccessControlListStatus string
	AccessControlListType   string
	AccessControlListId     string
	EnableHTTP2             bool
	CertificateId           string
	EgressMbps              int
	Description             string
	EstablishedTimeout      int

	ClientRequestTimeout  int
	ClientIdleTimeout     int
	BackendConnectTimeout int
	BackendIdleTimeout    int

	ListenerHealthCheckOptions
	ListenerStickySessionOptions

	BackendServerPort int
	XForwardedFor     bool
	Gzip              bool

	TLSCipherPolicy string
}

type ListenerStickySessionOptions struct {
	StickySession              string
	StickySessionCookie        string
	StickySessionType          string
	StickySessionCookieTimeout int
}

type ChangeListenerSchedulerOptions struct {
	Scheduler string
	ListenerStickySessionOptions
}

type ListenerHealthCheckOptions struct {
	HealthCheckReq string
	HealthCheckExp string

	HealthCheck         string
	HealthCheckType     string
	HealthCheckTimeout  int
	HealthCheckDomain   string
	HealthCheckHttpCode string
	HealthCheckURI      string
	HealthCheckInterval int

	HealthCheckRise int
	HealthCheckFail int
}

type ListenerCertificateOptions struct {
	CertificateId string
}

type ListenerAclOptions struct {
	AclId     string
	AclStatus string
	AclType   string
}

type SLoadbalancerListenerRule struct {
	Name             string
	Domain           string
	Path             string
	BackendGroupId   string
	BackendGroupType string

	Condition string // for aws only

	Scheduler           string // for qcloud only
	HealthCheck         string // for qcloud only
	HealthCheckType     string // for qcloud only
	HealthCheckTimeout  int    // for qcloud only
	HealthCheckDomain   string // for qcloud only
	HealthCheckHttpCode string // for qcloud only
	HealthCheckURI      string // for qcloud only
	HealthCheckInterval int    // for qcloud only

	HealthCheckRise int // for qcloud only
	HealthCheckFail int // for qcloud only

	StickySessionCookieTimeout int // for qcloud only
}
