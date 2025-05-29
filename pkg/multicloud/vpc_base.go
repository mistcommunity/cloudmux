
package multicloud

import (
	"yunion.io/x/pkg/errors"

	apis "yunion.io/x/cloudmux/pkg/apis/compute"
	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

type SVpc struct {
	SResourceBase
}

func (self *SVpc) GetINatGateways() ([]cloudprovider.ICloudNatGateway, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetINatGateways")
}

func (self *SVpc) GetICloudVpcPeeringConnections() ([]cloudprovider.ICloudVpcPeeringConnection, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudVpcPeeringConnections")
}

func (self *SVpc) GetICloudAccepterVpcPeeringConnections() ([]cloudprovider.ICloudVpcPeeringConnection, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudVpcPeeringConnections")
}

func (self *SVpc) GetICloudVpcPeeringConnectionById(id string) (cloudprovider.ICloudVpcPeeringConnection, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudVpcPeeringConnectionById")
}

func (self *SVpc) CreateICloudVpcPeeringConnection(opts *cloudprovider.VpcPeeringConnectionCreateOptions) (cloudprovider.ICloudVpcPeeringConnection, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateICloudVpcPeeringConnection")
}

func (self *SVpc) AcceptICloudVpcPeeringConnection(id string) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "AcceptICloudVpcPeeringConnection")
}

func (self *SVpc) GetAuthorityOwnerId() string {
	return ""
}

func (self *SVpc) CreateRouteToVpcPeeringConnection(cidrBlock, peerId string) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateRouteToVpcPeeringConnection")
}

func (self *SVpc) DeleteVpcPeeringConnectionRoute(vpcPeeringConnectionId string) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "DeleteVpcPeeringConnectionRoute")
}

func (self *SVpc) ProposeJoinICloudInterVpcNetwork(opts *cloudprovider.SVpcJointInterVpcNetworkOption) error {
	return errors.Wrapf(cloudprovider.ErrNotImplemented, "ProposeJoinICloudInterVpcNetwork")
}

func (self *SVpc) IsSupportSetExternalAccess() bool {
	return false
}

func (self *SVpc) GetExternalAccessMode() string {
	return apis.VPC_EXTERNAL_ACCESS_MODE_EIP
}

func (self *SVpc) AttachInternetGateway(igwId string) error {
	return errors.Wrap(cloudprovider.ErrNotSupported, "AttachInternetGateway")
}

func (self *SVpc) CreateINatGateway(opts *cloudprovider.NatGatewayCreateOptions) (cloudprovider.ICloudNatGateway, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateINatGateway")
}

func (self *SVpc) CreateIWire(opts *cloudprovider.SWireCreateOptions) (cloudprovider.ICloudWire, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "CreateIWire")
}

func (self *SVpc) GetGlobalVpcId() string {
	return ""
}

func (self *SVpc) GetCidrBlock6() string {
	return ""
}

func (self *SVpc) GetICloudIPv6Gateways() ([]cloudprovider.ICloudIPv6Gateway, error) {
	return nil, errors.Wrapf(cloudprovider.ErrNotImplemented, "GetICloudIPv6Gateways")
}
