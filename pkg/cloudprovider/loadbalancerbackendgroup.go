package cloudprovider

type SLoadbalancerBackendGroup struct {
	Name      string
	GroupType string
	Backends  []SLoadbalancerBackend

	// huawei
	Scheduler string
	Protocol  string

	// aws
	ListenPort int    // 后端端口
	VpcId      string // vpc id
}

type SLoadbalancerHealthCheck struct {
	HealthCheckType string
	HealthCheckReq  string
	HealthCheckExp  string

	HealthCheck         string
	HealthCheckTimeout  int
	HealthCheckDomain   string
	HealthCheckHttpCode string
	HealthCheckURI      string
	HealthCheckInterval int

	HealthCheckRise int
	HealthCheckFail int
}

type SLoadbalancerStickySession struct {
	StickySession              string
	StickySessionCookie        string
	StickySessionType          string
	StickySessionCookieTimeout int
}
