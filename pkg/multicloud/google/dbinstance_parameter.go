
package google

import "fmt"

type SDBInstanceParameter struct {
	rds *SDBInstance

	Name  string
	Value string
}

func (parameter *SDBInstanceParameter) GetGlobalId() string {
	return fmt.Sprintf("%s/%s", parameter.rds.GetGlobalId(), parameter.Name)
}

func (parameter *SDBInstanceParameter) GetKey() string {
	return parameter.Name
}

func (parameter *SDBInstanceParameter) GetValue() string {
	return parameter.Value
}

func (parameter *SDBInstanceParameter) GetDescription() string {
	return ""
}
