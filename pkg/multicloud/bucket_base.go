
package multicloud

import (
	api "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SBaseBucket struct {
	SResourceBase
}

func (b *SBaseBucket) MaxPartCount() int {
	return 10000
}

func (b *SBaseBucket) MaxPartSizeBytes() int64 {
	return 5 * 1000 * 1000 * 1000
}

func (b *SBaseBucket) GetId() string {
	return ""
}

func (b *SBaseBucket) GetName() string {
	return ""
}

func (b *SBaseBucket) GetGlobalId() string {
	return ""
}

func (b *SBaseBucket) GetStatus() string {
	return api.BUCKET_STATUS_READY
}

func (b *SBaseBucket) Refresh() error {
	return nil
}

func (b *SBaseBucket) IsEmulated() bool {
	return false
}

func (b *SBaseBucket) LimitSupport() cloudprovider.SBucketStats {
	return cloudprovider.SBucketStats{
		SizeBytes:   -1,
		ObjectCount: -1,
	}
}

func (b *SBaseBucket) GetLimit() cloudprovider.SBucketStats {
	return cloudprovider.SBucketStats{}
}

func (b *SBaseBucket) SetLimit(limit cloudprovider.SBucketStats) error {
	return nil
}

func (b *SBaseBucket) SetWebsite(conf cloudprovider.SBucketWebsiteConf) error {
	return cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) GetWebsiteConf() (cloudprovider.SBucketWebsiteConf, error) {
	return cloudprovider.SBucketWebsiteConf{}, cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) DeleteWebSiteConf() error {
	return cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) SetCORS(rules []cloudprovider.SBucketCORSRule) error {
	return cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) GetCORSRules() ([]cloudprovider.SBucketCORSRule, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) DeleteCORS() error {
	return cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) SetReferer(conf cloudprovider.SBucketRefererConf) error {
	return cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) GetReferer() (cloudprovider.SBucketRefererConf, error) {
	return cloudprovider.SBucketRefererConf{}, cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) GetCdnDomains() ([]cloudprovider.SCdnDomain, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) GetPolicy() ([]cloudprovider.SBucketPolicyStatement, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) SetPolicy(policy cloudprovider.SBucketPolicyStatementInput) error {
	return cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) DeletePolicy(id []string) ([]cloudprovider.SBucketPolicyStatement, error) {
	return nil, cloudprovider.ErrNotImplemented
}

func (b *SBaseBucket) ListMultipartUploads() ([]cloudprovider.SBucketMultipartUploads, error) {
	return nil, cloudprovider.ErrNotImplemented
}
