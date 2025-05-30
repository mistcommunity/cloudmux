
package aliyun

import (
	// "time"
	"yunion.io/x/log"
)

// {"CpuCoreCount":1,"EniQuantity":1,"GPUAmount":0,"GPUSpec":"","InstanceTypeFamily":"ecs.t1","InstanceTypeId":"ecs.t1.xsmall","LocalStorageCategory":"","MemorySize":0.500000}
// InstanceBandwidthRx":26214400,"InstanceBandwidthTx":26214400,"InstancePpsRx":4500000,"InstancePpsTx":4500000

type SInstanceType struct {
	BaselineCredit       int
	CpuCoreCount         int
	MemorySize           float32
	EniQuantity          int // 实例规格支持网卡数量
	GPUAmount            int
	GPUSpec              string
	InstanceTypeFamily   string
	InstanceFamilyLevel  string
	InstanceTypeId       string
	LocalStorageCategory string
	LocalStorageAmount   int
	LocalStorageCapacity int64
	InstanceBandwidthRx  int
	InstanceBandwidthTx  int
	InstancePpsRx        int
	InstancePpsTx        int
}

func (self *SRegion) GetInstanceTypes() ([]SInstanceType, error) {
	params := make(map[string]string)
	params["RegionId"] = self.RegionId

	body, err := self.ecsRequest("DescribeInstanceTypes", params)
	if err != nil {
		log.Errorf("GetInstanceTypes fail %s", err)
		return nil, err
	}

	instanceTypes := make([]SInstanceType, 0)
	err = body.Unmarshal(&instanceTypes, "InstanceTypes", "InstanceType")
	if err != nil {
		log.Errorf("Unmarshal instance type details fail %s", err)
		return nil, err
	}
	return instanceTypes, nil
}

func (self *SInstanceType) memoryMB() int {
	return int(self.MemorySize * 1024)
}
