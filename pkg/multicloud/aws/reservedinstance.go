
package aws

func (self *SRegion) GetReservedInstance() error {
	params := map[string]string{}
	ret := struct{}{}
	return self.ec2Request("DescribeReservedInstances", params, &ret)
}

type SReservedHostOffering struct {
	Duration       int
	HourlyPrice    float64
	InstanceFamily string
	OfferingId     string
	PaymentOption  string
	UpfrontPrice   float64
}

func (self *SRegion) GetReservedHostOfferings() error {
	params := map[string]string{}
	ret := struct{}{}
	return self.ec2Request("DescribeHostReservationOfferings", params, &ret)
}
