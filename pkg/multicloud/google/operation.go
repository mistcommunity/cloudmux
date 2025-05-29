package google

import (
	"time"

	"yunion.io/x/log"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
)

const (
	OPERATION_STATUS_RUNNING = "RUNNING"
	OPERATION_STATUS_DONE    = "DONE"
)

type SOperation struct {
	Id            string
	Name          string
	OperationType string
	TargetLink    string
	TargetId      string
	Status        string
	User          string
	Progress      int
	InsertTime    time.Time
	StartTime     time.Time
	EndTime       time.Time
	SelfLink      string
	Region        string
	Kind          string
}

func (self *SGoogleClient) GetOperation(id string) (*SOperation, error) {
	operation := &SOperation{}
	err := self.GetBySelfId(id, &operation)
	if err != nil {
		return nil, err
	}
	return operation, nil
}

func (self *SGoogleClient) WaitOperation(id string, resource, action string) (string, error) {
	targetLink := ""
	err := cloudprovider.Wait(time.Second*5, time.Minute*5, func() (bool, error) {
		operation, err := self.GetOperation(id)
		if err != nil {
			return false, err
		}
		log.Debugf("%s %s operation status: %s expect %s", action, resource, operation.Status, OPERATION_STATUS_DONE)
		if operation.Status == OPERATION_STATUS_DONE {
			targetLink = operation.TargetLink
			return true, nil
		}
		return false, nil
	})
	return targetLink, err
}
