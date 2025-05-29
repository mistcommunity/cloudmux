package cloudprovider

import "yunion.io/x/pkg/util/samlutils"

const (
	SAML_ENTITY_ID_ALIYUN_ROLE  = "urn:alibaba:cloudcomputing"
	SAML_ENTITY_ID_AWS_CN       = "urn:amazon:webservices:cn-north-1"
	SAML_ENTITY_ID_AWS          = "urn:amazon:webservices"
	SAML_ENTITY_ID_QCLOUD       = "cloud.tencent.com"
	SAML_ENTITY_ID_HUAWEI_CLOUD = "https://auth.huaweicloud.com/"
	SAML_ENTITY_ID_GOOGLE       = "google.com"
	SAML_ENTITY_ID_AZURE        = "urn:federation:MicrosoftOnline"
	SAML_ENTITY_ID_VOLC_ENGINE  = "https://www.volcengine.com/"
)

type SAMLProviderCreateOptions struct {
	Name     string
	Metadata samlutils.EntityDescriptor
	Desc     string
}
