
package aws

import (
	"fmt"
	"net/url"
)

type LambdaFunction struct {
	Description   string `json:"Description"`
	TracingConfig struct {
		Mode string `json:"Mode"`
	} `json:"TracingConfig"`
	VpcConfig     string `json:"VpcConfig"`
	SigningJobArn string `json:"SigningJobArn"`
	SnapStart     struct {
		OptimizationStatus string `json:"OptimizationStatus"`
		ApplyOn            string `json:"ApplyOn"`
	} `json:"SnapStart"`
	RevisionID               string `json:"RevisionId"`
	LastModified             string `json:"LastModified"`
	FileSystemConfigs        string `json:"FileSystemConfigs"`
	FunctionName             string `json:"FunctionName"`
	Runtime                  string `json:"Runtime"`
	Version                  string `json:"Version"`
	PackageType              string `json:"PackageType"`
	LastUpdateStatus         string `json:"LastUpdateStatus"`
	Layers                   string `json:"Layers"`
	FunctionArn              string `json:"FunctionArn"`
	KMSKeyArn                string `json:"KMSKeyArn"`
	MemorySize               int    `json:"MemorySize"`
	ImageConfigResponse      string `json:"ImageConfigResponse"`
	LastUpdateStatusReason   string `json:"LastUpdateStatusReason"`
	DeadLetterConfig         string `json:"DeadLetterConfig"`
	Timeout                  int    `json:"Timeout"`
	Handler                  string `json:"Handler"`
	CodeSha256               string `json:"CodeSha256"`
	Role                     string `json:"Role"`
	SigningProfileVersionArn string `json:"SigningProfileVersionArn"`
	MasterArn                string `json:"MasterArn"`
	RuntimeVersionConfig     string `json:"RuntimeVersionConfig"`
	CodeSize                 int    `json:"CodeSize"`
	State                    string `json:"State"`
	StateReason              string `json:"StateReason"`
	LoggingConfig            struct {
		LogFormat           string `json:"LogFormat"`
		ApplicationLogLevel string `json:"ApplicationLogLevel"`
		LogGroup            string `json:"LogGroup"`
		SystemLogLevel      string `json:"SystemLogLevel"`
	} `json:"LoggingConfig"`
	Environment      string `json:"Environment"`
	EphemeralStorage struct {
		Size int `json:"Size"`
	} `json:"EphemeralStorage"`
	StateReasonCode            string   `json:"StateReasonCode"`
	LastUpdateStatusReasonCode string   `json:"LastUpdateStatusReasonCode"`
	Architectures              []string `json:"Architectures"`
}

func (self *SRegion) ListFunctions() ([]LambdaFunction, error) {
	params := url.Values{}
	params.Set("MaxItems", "10000")
	params.Set("FunctionVersion", "ALL")
	ret, marker := []LambdaFunction{}, ""
	for {
		part := struct {
			Functions  []LambdaFunction
			NextMarker string
		}{}
		if len(marker) > 0 {
			params.Set("Marker", marker)
		}
		path := fmt.Sprintf("/2015-03-31/functions/?%s", params.Encode())
		err := self.lambdaRequest("ListFunctions", path, map[string]interface{}{}, &part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, part.Functions...)
		if len(part.Functions) == 0 || len(part.NextMarker) == 0 {
			break
		}
		marker = part.NextMarker
	}
	return ret, nil
}

type ProvisionedConcurrencyConfig struct {
	RequestedProvisionedConcurrentExecutions int
}

func (self *SRegion) GetProvisionedConcurrencyConfig(funcName, qualifier string) (*ProvisionedConcurrencyConfig, error) {
	params := map[string]interface{}{}
	ret := &ProvisionedConcurrencyConfig{}
	path := fmt.Sprintf("/2019-09-30/functions/%s/provisioned-concurrency?Qualifier=%s", funcName, qualifier)
	err := self.lambdaRequest("GetProvisionedConcurrencyConfig", path, params, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
