
package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"yunion.io/x/jsonutils"
	"yunion.io/x/log"
	"yunion.io/x/pkg/errors"
	"yunion.io/x/pkg/util/regutils"

	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/esxi"
)

type SESXiProviderFactory struct {
	cloudprovider.SPremiseBaseProviderFactory
}

func (self *SESXiProviderFactory) GetId() string {
	return esxi.CLOUD_PROVIDER_VMWARE
}

func (self *SESXiProviderFactory) GetName() string {
	return esxi.CLOUD_PROVIDER_VMWARE
}

func (self *SESXiProviderFactory) ValidateChangeBandwidth(instanceId string, bandwidth int64) error {
	return fmt.Errorf("Changing %s bandwidth is not supported", esxi.CLOUD_PROVIDER_VMWARE)
}

func (self *SESXiProviderFactory) ValidateCreateCloudaccountData(ctx context.Context, input cloudprovider.SCloudaccountCredential) (cloudprovider.SCloudaccount, error) {
	output := cloudprovider.SCloudaccount{}
	if len(input.Username) == 0 {
		return output, errors.Wrap(cloudprovider.ErrMissingParameter, "username")
	}
	if len(input.Password) == 0 {
		return output, errors.Wrap(cloudprovider.ErrMissingParameter, "password")
	}
	if len(input.Host) == 0 {
		return output, errors.Wrap(cloudprovider.ErrMissingParameter, "host")
	}
	if !regutils.MatchIPAddr(input.Host) && !regutils.MatchDomainName(input.Host) {
		return output, errors.Wrap(cloudprovider.ErrInputParameter, "host should be ip or domain name")
	}
	output.AccessUrl = fmt.Sprintf("https://%s:%d/sdk", input.Host, input.Port)
	if input.Port == 0 || input.Port == 443 {
		output.AccessUrl = fmt.Sprintf("https://%s/sdk", input.Host)
	}
	output.Account = input.Username
	output.Secret = input.Password
	return output, nil
}

func (self *SESXiProviderFactory) ValidateUpdateCloudaccountCredential(ctx context.Context, input cloudprovider.SCloudaccountCredential, cloudaccount string) (cloudprovider.SCloudaccount, error) {
	output := cloudprovider.SCloudaccount{}
	if len(input.Username) == 0 {
		return output, errors.Wrap(cloudprovider.ErrMissingParameter, "username")
	}
	if len(input.Password) == 0 {
		return output, errors.Wrap(cloudprovider.ErrMissingParameter, "password")
	}
	output = cloudprovider.SCloudaccount{
		Account: input.Username,
		Secret:  input.Password,
	}
	if len(input.Host) > 0 {
		if !regutils.MatchIPAddr(input.Host) && !regutils.MatchDomainName(input.Host) {
			return output, errors.Wrap(cloudprovider.ErrInputParameter, "host should be ip or domain name")
		}
		output.AccessUrl = fmt.Sprintf("https://%s:%d/sdk", input.Host, input.Port)
		if input.Port == 0 || input.Port == 443 {
			output.AccessUrl = fmt.Sprintf("https://%s/sdk", input.Host)
		}
	}
	return output, nil
}

func parseHostPort(host string, defPort int) (string, int, error) {
	colonPos := strings.IndexByte(host, ':')
	if colonPos > 0 {
		h := host[:colonPos]
		p, err := strconv.Atoi(host[colonPos+1:])
		if err != nil {
			log.Errorf("Invalid host %s", host)
			return "", 0, err
		}
		if p == 0 {
			p = defPort
		}
		return h, p, nil
	} else {
		return host, defPort, nil
	}
}

func (self *SESXiProviderFactory) GetProvider(cfg cloudprovider.ProviderConfig) (cloudprovider.ICloudProvider, error) {
	parts, err := url.Parse(cfg.URL)
	if err != nil {
		return nil, err
	}
	host, port, err := parseHostPort(parts.Host, 443)
	if err != nil {
		return nil, err
	}

	client, err := esxi.NewESXiClient(
		esxi.NewESXiClientConfig(
			host, port, cfg.Account, cfg.Secret,
		).CloudproviderConfig(cfg),
	)
	if err != nil {
		return nil, err
	}
	return &SESXiProvider{
		SBaseProvider: cloudprovider.NewBaseProvider(self),
		client:        client,
	}, nil
}

func (self *SESXiProviderFactory) GetClientRC(info cloudprovider.SProviderInfo) (map[string]string, error) {
	parts, err := url.Parse(info.Url)
	if err != nil {
		return nil, err
	}
	host, port, err := parseHostPort(parts.Host, 443)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"VMWARE_HOST":     host,
		"VMWARE_PORT":     fmt.Sprintf("%d", port),
		"VMWARE_ACCOUNT":  info.Account,
		"VMWARE_PASSWORD": info.Secret,
	}, nil
}

func (self *SESXiProviderFactory) GetAccountIdEqualizer() func(origin, now string) bool {
	return func(origin, now string) bool {
		if len(now) == 0 {
			return true
		}
		originUserName, nowUserName := origin, now
		index1 := strings.Index(origin, "@")
		index2 := strings.Index(now, "@")
		if index1 != -1 {
			originUserName = originUserName[:index1]
		}
		if index2 != -1 {
			nowUserName = nowUserName[:index2]
		}
		return originUserName == nowUserName
	}
}

func init() {
	factory := SESXiProviderFactory{}
	cloudprovider.RegisterFactory(&factory)
}

type SESXiProvider struct {
	cloudprovider.SBaseProvider
	client *esxi.SESXiClient
}

func (self *SESXiProvider) GetSysInfo() (jsonutils.JSONObject, error) {
	return self.client.About(), nil
}

func (self *SESXiProvider) GetVersion() string {
	return self.client.GetVersion()
}

func (self *SESXiProvider) GetSubAccounts() ([]cloudprovider.SSubAccount, error) {
	return self.client.GetSubAccounts()
}

func (self *SESXiProvider) GetAccountId() string {
	return self.client.GetAccountId()
}

func (self *SESXiProvider) GetIRegions() ([]cloudprovider.ICloudRegion, error) {
	return nil, cloudprovider.ErrNotSupported
}

func (self *SESXiProvider) GetIRegionById(id string) (cloudprovider.ICloudRegion, error) {
	return nil, cloudprovider.ErrNotSupported
}

func (self *SESXiProvider) GetBalance() (*cloudprovider.SBalanceInfo, error) {
	return &cloudprovider.SBalanceInfo{
		Amount:   0.0,
		Currency: "CNY",
		Status:   api.CLOUD_PROVIDER_HEALTH_NORMAL,
	}, cloudprovider.ErrNotSupported
}

func (self *SESXiProvider) GetOnPremiseIRegion() (cloudprovider.ICloudRegion, error) {
	return self.client, nil
}

func (self *SESXiProvider) GetIProjects() ([]cloudprovider.ICloudProject, error) {
	return self.client.GetIProjects()
}

func (self *SESXiProvider) GetStorageClasses(regionId string) []string {
	return nil
}

func (self *SESXiProvider) GetBucketCannedAcls(regionId string) []string {
	return nil
}

func (self *SESXiProvider) GetObjectCannedAcls(regionId string) []string {
	return nil
}

func (self *SESXiProvider) GetCapabilities() []string {
	return self.client.GetCapabilities()
}

func (self *SESXiProvider) GetMetrics(opts *cloudprovider.MetricListOptions) ([]cloudprovider.MetricValues, error) {
	return self.client.GetMetrics(opts)
}
